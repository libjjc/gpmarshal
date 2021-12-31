package php

import (
    "bytes"
    "fmt"
    "reflect"
    "strconv"
    "strings"
    "sync"
)

var(
    textEncodePool sync.Pool
)

func Marshal(value interface{})(string,error){
    enc := newTextEncoder()
    if err := enc.encode(value);err!=nil{
        mse , ok := err.(*MarshalError)
        if !ok{
            mse = newMarshalError(err).withExtend(extends.encoding())
        }
        return "",mse
    }
    textEncodePool.Put(enc)
    return enc.String(),nil
}

type textEncoder struct{
    bytes.Buffer
    err error
}

func newTextEncoder()*textEncoder{
    if v := textEncodePool.Get(); v != nil {
        e := v.(*textEncoder)
        e.reset()
        return e
    }
    return &textEncoder{}
}

func(te *textEncoder)reset(){
    te.Reset()
    te.err = nil
}

func(te *textEncoder)writeString(s string)error{
    if te.err == nil{
        _, te.err = te.WriteString(s)
    }
    return te.err
}

func(te *textEncoder)write(enc *textEncoder)error{
    if te.err == nil{
        _,te.err = te.Write(enc.Bytes())
    }
    return te.err
}

func(te *textEncoder)writeBytes(bytes[]byte)error{
    if te.err == nil{
        _,te.err = te.Write(bytes)
    }
    return te.err
}

func(te *textEncoder)writeRune(r rune)error{
    if te.err == nil{
        _,te.err = te.WriteRune(r)
    }
    return te.err
}

func(te *textEncoder)readString()(string,error){
    if te.err == nil {
        return te.String(), te.err
    }
    return "",te.err
}

func(te *textEncoder)read()([]byte,error){
    if te.err == nil{
        return te.Bytes(),te.err
    }
    return nil,te.err
}

type tagOpt string

func(opt tagOpt)omitEmpty()bool{
    return strings.Contains(string(opt),",omitempty")
}

func parseTag(s string)(string,tagOpt){
    if idx := strings.Index(s,",");idx!=-1 {
        return s[:idx], tagOpt(s[idx+1:])
    }
    return s,""
}

func (te *textEncoder)encode(in interface{})error{
    if te.err != nil{
        return te.err
    }
    iv := reflect.Indirect(reflect.ValueOf(in))
    return te.value(iv)
}

func(te *textEncoder)value(in reflect.Value)error{
    if te.err != nil{
        return te.err
    }
    switch in.Kind(){
    case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64,
        reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
        return te.numeric(in)
    case reflect.Float32,reflect.Float64:
        return te.float(in)
    case reflect.String:
        return te.string(in)
    case reflect.Bool:
        return te.bool(in)
    case reflect.Slice,reflect.Array:
        return te.slice(in)
    case reflect.Map:
        return te.dict(in)
    case reflect.Struct:
        return te.object(in)
    case reflect.Ptr:
        return te.value(in.Elem())
    case reflect.Interface:
    default:
    }
    return nil
}

func(te *textEncoder)numeric(in reflect.Value)error{
    _ = te.writeString(fmt.Sprintf("i:%d;",in.Interface()))
    return te.err
}

func(te *textEncoder)float(in reflect.Value)error{
    _ = te.writeString(fmt.Sprintf("d:%s;",strconv.FormatFloat(in.Float(), 'g', 17, 64)))
    return te.err
}

func(te *textEncoder)string(in reflect.Value)error{
    te.err =te.writeString(fmt.Sprintf("s:%d:\"%s\";",in.Len(),in.String()))
    return te.err
}

func(te *textEncoder)bool(in reflect.Value)error{
    if in.Bool() {
        _ = te.writeString("b:1;")
    }else{
        _ = te.writeString("b:0;")
    }
    return te.err
}

func(te *textEncoder)ptr(in reflect.Value)error{
    if in.Kind() == reflect.Ptr && in.IsNil(){
        return te.nil(in)
    }
    return nil
}
func(te *textEncoder)nil(in reflect.Value)error{
    _ = te.writeString("N;")
    return te.err
}

func(te *textEncoder)slice(in reflect.Value)error{
    rv := reflect.Indirect(in)
    enc := newTextEncoder()
    defer func(){
        if te.err == nil && enc.err == nil{
            _ = te.writeString(fmt.Sprintf("a:%d:{", rv.Len()))
            _ = te.write(enc)
            _ = te.writeString("}")
        }
    }()
    len := rv.Len()
    for i := 0 ;i < len ;i++{
        if err := enc.encode(i);err!=nil{
            return err
        }
        if err := enc.value(rv.Index(i));err!=nil{
            return err
        }
    }
    return nil
}

func(te *textEncoder)dict(in reflect.Value)error{
    rv := reflect.Indirect(in)
    enc := newTextEncoder()
    defer func(){
        if te.err == nil {
            _ = te.writeString(fmt.Sprintf("a:%d:{", rv.Len()))
            _ = te.write(enc)
            _ = te.writeString("}")
        }
    }()
    iter := rv.MapRange()
    for ; iter.Next();{
        if err := enc.value(iter.Key());err!=nil{
            return err
        }
        if err := enc.value(iter.Value());err!=nil{
            return err
        }
    }
    return nil
}

func isEmptyValue(v reflect.Value) bool {
    switch v.Kind() {
    case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
        return v.Len() == 0
    case reflect.Bool:
        return !v.Bool()
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return v.Int() == 0
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        return v.Uint() == 0
    case reflect.Float32, reflect.Float64:
        return v.Float() == 0
    case reflect.Interface, reflect.Ptr:
        return v.IsNil()
    }
    return false
}

func(te *textEncoder)object(in reflect.Value)error{
    rv := reflect.Indirect(in)
    enc := newTextEncoder()
    numFields := 0
    defer func(){
        if te.err == nil{
            _ = te.writeString(fmt.Sprintf("a:%d:{",numFields))
            _ = te.write(enc)
            _ = te.writeString(fmt.Sprintf("}"))
        }
    }()

    for i := 0 ;i < rv.NumField(); i++{
        field := rv.Type().Field(i)
        stag := field.Tag
        tag,ok := stag.Lookup("php")
        if ok && field.PkgPath==""{
            t,opt := parseTag(tag)
            if (!isEmptyValue(in) || opt.omitEmpty()) && t!="" {
                if err := enc.encode(t); err != nil {
                    return err
                }
                if err := enc.value(rv.Field(i)); err != nil {
                    return err
                }
                numFields++
            }
        }
    }
    return nil
}

