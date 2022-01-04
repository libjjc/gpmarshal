package gpmarshal

import (
    "Github/go/src/pkg/bytes"
    "errors"
    "fmt"
    "reflect"
    "runtime"
    "strconv"
)

// Unmarshal parses the php-json-encoded data and stores the result
// in the value pointed to by v.
//
// it supported all type unmarshal from php-json-encode,but except the
// type chan,interface,function
//
//  the parameter `out` must be a pointer,if the parameter can not be addr
//  Unmarshal will be panic
//
// it use a scanner for scan the php-json-encode data. and a decoder
// for decode the result of scanner into go types.
//
// if the type of parameter `out` is slice or array.the buf must encode
// from php-sequence-array
//
// if the type of parameter `out` is map or struct,the buf must encode from the
// same type of key and value of php-array
//
// golang is a strong type language,before we marshal and unmarshal the data of php
// we should exactly process the type,make sure golang can recognize it. Unmarshal
// would not do any cast for value .
//
// if parameter `out` is a struct,and you want to decode the field of struct.you
// should make sure the field is exportable,and with `php` tag .
//
// Unmarshal will return MarshalError when somethings error(except runtime.error)
// in decoding.
//
// for example:
//  decoding int variable below
//  source := "i:3;"
//  var target int
//  Unmarshal([]byte(source),&target)
//  go run this code ,the target will be 3
//
//  decoding string variable below
//  source := "s:11:\"hello world!\"";
//  var target string
//  Unmarshal(source,&target)
//
//  decoding slice
//  source :="a:2{i:1;i:3;i:2;i:7;}"
//  target := make([]int,0)
//  Unmarshal(source,&target)
//
//  more usage see README.md
func Unmarshal(buf []byte,out interface{})(err error){
    defer func()(){
        r := recover()
        if r != nil {
            switch x := r.(type) {
            case runtime.Error:
                panic(r)
            case string:
                err = newMarshalError(errors.New(x)).withExtend(extends.decoding())
            case error:
                err = newMarshalError(x).withExtend(extends.decoding())
            default:
                err = newMarshalError(errors.New("unknown panic")).withExtend(extends.decoding())
            }
        }
        return
    }()
    dec := &textDecoder{
        err:  nil,
        scan: newScanner(buf),
    }
    err = dec.decode(out)
    if err != nil {
        _, ok := err.(*MarshalError)
        if !ok {
            err = newMarshalError(err).withExtend(extends.decoding())
        }
    }
    return
}

type apiUnmarshal interface {
    Unmarshal([]byte,interface{})error
}

type textDecoder struct{
    err error
    scan *scanner
}

type scanner struct{
    bytes.Reader
}

func newScanner(buf []byte)*scanner{
    return &scanner{
        *bytes.NewReader(buf),
    }
}

// get the current cursor of scanner
func (scan *scanner)cursor()int64{
    cursor,err := scan.Seek(0,1)
    if err!=nil{
        panic(err)
    }
    return cursor
}

func(scan *scanner)skipSpace(){
    scan.skip(' ')
}

// skip sequenced byte `tk`. and return the prev index of the first byte not `tk`.
// if the byte which is not `tk` is not found, it will panic
func(scan *scanner)skip(tk byte){
    for{
        ch,err := scan.ReadByte()
        if err!=nil{
            panic(err)
        }
        if ch == tk{
            continue
        }
        if err := scan.UnreadByte();err != nil {
            panic(err)
        }
        break
    }
}

// skip all byte except `tk`. if `tk` not found,it will panic
// if `tk` found,the cursor will point to the prev byte index of `tk`
func(scan *scanner)until(tk uint8)int64{
    for {
        ch, err := scan.ReadByte()
        if err != nil {
            panic(err)
        }
        if ch == tk {
            break
        }
    }
    return scan.cursor()-1
}


func(scan *scanner)nextByte()(byte,error){
    ch,err := scan.ReadByte()
    if err == nil {
        err = scan.UnreadByte()
        if err!=nil{
            return 0,err
        }
    }
    return ch,err
}

func (scan *scanner)readByte()byte{
    ch,err:=scan.ReadByte()
    if err!=nil{
        panic(err)
    }
    return ch
}

func (scan *scanner)read(len int64)[]byte{
    buf :=make([]byte,len)
    n,err := scan.Read(buf)
    if err != nil{
        panic(err)
    }
    if int64(n)!=len{
        panic(fmt.Sprintf("unexpected lentgh , actually=%d,expect =%d",n,len))
    }
    return buf
}

func(scan *scanner)tag(expected byte)error{
    tag:=[2]byte{}
    len,err := scan.Read(tag[:])
    if len < 2{
        return err
    }
    if tag[0]=='N'&&tag[1]==';'{
        return nullPhpValue
    }
    if tag[0] != expected || tag[1] != ':'{
        return newMarshalError(errors.New("unexpected tag")).withExtend(extends.unexpected(tag[1],':'))
    }
    return nil
}

//
func(scan *scanner)expect(expected byte){
    scan.skipSpace()
    actually,err := scan.ReadByte()
    if err == nil && actually!=expected{
        panic(fmt.Sprintf("unexpected token , actually=%d,expect =%d",actually,expected))
    }
}

func(scan *scanner)val(){
    scan.until(';')
}

func(scan *scanner)int()(int64,error){
    scan.skipSpace()
    if err := scan.tag('i');err!=nil{
        return 0,err
    }
    prev := scan.cursor()
    cur := scan.until(';')
    size := cur-prev
    buf := make([]byte,size)
    n,err := scan.ReadAt(buf,prev)
    if err!=nil{
        return 0,err
    }
    if int64(n)!=size{
        return 0,newMarshalError(errors.New("unexpected length")).withExtend(extends.unexpected(n,size))
    }
    return strconv.ParseInt(string(buf),10,64)
}

func(scan *scanner)uint()(uint64,error){
    scan.skipSpace()
    if err := scan.tag('i');err!=nil{
        return 0,err
    }
    prev := scan.cursor()
    cur := scan.until(';')
    size := cur-prev
    buf := make([]byte,size)
    n,err := scan.ReadAt(buf,prev)
    if err!=nil{
        return 0,err
    }
    if int64(n)!=size{
        return 0,newMarshalError(errors.New("unexpected length")).withExtend(extends.unexpected(n,size))
    }
    return strconv.ParseUint(string(buf),10,64)
}

func(scan *scanner)float()(float64,error){
    scan.skipSpace()
    if err := scan.tag('d');err!=nil{
        return 0,err
    }
    prev := scan.cursor()
    cur := scan.until(';')
    size := cur-prev
    buf := make([]byte,size)
    n,err := scan.ReadAt(buf,prev)
    if err!=nil{
        return 0,err
    }
    if int64(n)!=size{
        return 0,newMarshalError(errors.New("unexpected length")).withExtend(extends.unexpected(n,size))
    }
    return strconv.ParseFloat(string(buf),64)
}

func(scan *scanner)bool()(bool,error){
    scan.skipSpace()
    if err := scan.tag('b');err!=nil{
        return false,err
    }
    b := scan.readByte()
    scan.expect(';')
    val := true
    if b == '0'{
        val = false
    }

    return val,nil
}

func(scan *scanner)length()(int64,error){
    prev := scan.cursor()
    cur := scan.until(':')
    size := cur-prev
    buf := make([]byte,size)
    n,err := scan.ReadAt(buf,prev)
    if err!=nil{
        return 0,err
    }
    if int64(n) != size{
        return 0,newMarshalError(
            errors.New("php array length parse failed")).
            withExtend(extends.unexpected(n,size))
    }
    return strconv.ParseInt(string(buf),10,64)
}

func(scan *scanner)string()(string,error){
    scan.skipSpace()
    if err:=scan.tag('s');err!=nil{
        return "",err
    }
    len, err := scan.length()
    if err!=nil{
        return "",err
    }
    scan.expect('"')
    s := scan.read(len)
    scan.expect('"')
    scan.expect(';')
    return string(s),nil
}

func(scan *scanner)array()error{
    if err:= scan.tag('a');err!=nil{
        return err
    }
    _,err := scan.length()
    if err!= nil{
        return err
    }
    scan.expect('{')
    if err:= scan.ignore();err!=nil{
        return err
    }
    scan.expect('}')
    return nil
}

func(scan *scanner)object()error{
    if _,err := scan.string();err!=nil{
        return err
    }
    return scan.array()
}

func(scan *scanner)nil()error{
    scan.expect(';')
    return nil
}

func(scan *scanner)objectSerialized()error{
    if _,err := scan.string();err!=nil{
        return err
    }
    return scan.ignore()
}

func(scan *scanner)reference()error{
    scan.until(';')
    return nil
}

func(scan *scanner)splArray()error{
    _,err := scan.int()
    if err!=nil{
        return err
    }
    if err:= scan.ignore();err!=nil{
        return err
    }
    scan.until(';')
    if err := scan.tag('m');err!=nil{
        return err
    }
    return scan.ignore()
}

func(scan *scanner)ignore()error{
    tag,err := scan.nextByte()
    if err != nil{
        return err
    }
    switch tag {
    case 'i':
        _,err := scan.int()
        return err
    case 'd':
        _,err := scan.float()
        return err
    case 'b':
        _,err := scan.bool()
        return err
    case 's':
        _,err := scan.string()
        return err
    case 'a':
        return scan.array()
    case 'O':
        return scan.object()
    case 'C':
        return scan.objectSerialized()
    case 'R','r':
        return scan.reference()
    case 'x':
        return scan.splArray()
    case 'N':
        return scan.nil()
    case 'm':
    default:
        return newMarshalError(
            errors.New("token parse failed")).
            withExtend(extends.unRecognize(tag))
    }
    return err
}


func(dec *textDecoder)decode(out interface{})error{
    if dec.err!=nil{
        if mse , ok := dec.err.(*MarshalError);!ok{
            dec.err = mse
        }
        return dec.err
    }
    pv := reflect.ValueOf(out)
    if pv.Kind() != reflect.Ptr || pv.IsNil() {
        return newMarshalError(errors.New("type invalid")).withExtend(extends.decoding()).withExtend(extends.canNotSet())
    }
    return dec.value(reflect.ValueOf(out))
}

func(dec *textDecoder)value(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    switch out.Kind() {
    case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
        dec.err = dec.int(out)
    case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
        dec.err = dec.uint(out)
    case reflect.Float32,reflect.Float64:
        dec.err = dec.float(out)
    case reflect.String:
        dec.err = dec.string(out)
    case reflect.Bool:
        dec.err = dec.bool(out)
    case reflect.Slice:
        dec.err = dec.slice(out)
    case reflect.Array:
        dec.err = dec.array(out)
    case reflect.Map:
        dec.err = dec.dict(out)
    case reflect.Struct:
        dec.err = dec.object(out)
    case reflect.Interface:
    case reflect.Ptr:
        dec.err = dec.ptr(out)
    default:

    }
    if dec.err == nullPhpValue{
        dec.err = nil
    }
    return dec.err
}

func(dec *textDecoder)int(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    n,err := dec.scan.int()
    if err !=nil{
        dec.err = err
        return dec.err
    }

    nv := reflect.ValueOf(n)
    if reflect.Zero(out.Type()).OverflowInt(n){
        return newMarshalError(errors.New("int decode failed")).
            withExtend(extends.invalidConversation(n,out.Type()))
    }
    out.Set(nv.Convert(out.Type()))
    return nil
}

func(dec *textDecoder)uint(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    n,err := dec.scan.uint()
    if err !=nil{
        dec.err = err
        return dec.err
    }

    nv := reflect.ValueOf(n)
    if reflect.Zero(out.Type()).OverflowUint(n){
        return newMarshalError(errors.New("uint decode failed")).
            withExtend(extends.invalidConversation(n,out.Type()))
    }
    out.Set(nv.Convert(out.Type()))
    return nil
}

func(dec *textDecoder)float(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    n,err := dec.scan.float()
    if err!=nil{
        dec.err = err
        return dec.err
    }

    if out.OverflowFloat(n) {
        dec.err = newMarshalError(errors.New("float decode failed")).
            withExtend(extends.invalidConversation(n,out.Type()))
        return dec.err
    }
    out.SetFloat(n)
    return nil
}

func(dec *textDecoder)string(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    s,err := dec.scan.string()
    if err!=nil{
        dec.err = err
        return dec.err
    }

    out.SetString(s)
    return nil
}

func(dec *textDecoder)bool(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    v,err := dec.scan.bool()
    if err!=nil{
        dec.err = err
        return dec.err
    }

    out.SetBool(v)
    return nil
}

func(dec *textDecoder)slice(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    scan := dec.scan
    if err := scan.tag('a');err!=nil{
        return err
    }
    l,err :=scan.length()
    if err!=nil{
        return err
    }
    scan.expect('{')
    ot := out.Type()
    newV := reflect.MakeSlice(ot,int(l),int(l))
    idx := 0
    for i:= 0;i < int(l);i++{
        if err := dec.decode(&idx);err!=nil{
            return nil
        }
        if idx <0 || idx >= int(l){
            dec.err = newMarshalError(errors.New("int decode failed")).
                withExtend(extends.outOfRange(idx,int(l)))
            return dec.err
        }
        ele := newV.Index(i)
        if err :=dec.value(ele);err!=nil{
            return err
        }
    }
    out.Set(newV)
    scan.until('}')
    return nil
}

func(dec *textDecoder)array(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    scan := dec.scan
    if err := scan.tag('a');err!=nil{
        return err
    }

    l,err :=scan.length()
    if err!=nil{
        return err
    }
    scan.until('{')
    idx := 0
    for i:= 0;i < int(l);i++{
        if err := dec.decode(&idx);err!=nil{
            return nil
        }
        if idx <0 || idx >=int(l){
            dec.err = newMarshalError(errors.New("int decode failed")).
                withExtend(extends.outOfRange(idx,int(l)))
            return dec.err
        }
        ele := out.Index(i)
        if err :=dec.value(ele);err!=nil{
            return dec.err
        }
    }
    scan.expect('}')
    return nil
}

func(dec *textDecoder)dict(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    scan := dec.scan
    if err := scan.tag('a');err!=nil{
        return err
    }
    l,err :=scan.length()
    if err!=nil{
        return err
    }
    scan.expect('{')

    ekt := out.Type().Key()
    evt := out.Type().Elem()
    newMap := out
    if out.IsNil() {
        // Make a new map to hold our result
        mapType := reflect.MapOf(ekt, evt)
        newMap = reflect.MakeMap(mapType)
        out.Set(newMap)
    }

    for i:= 0;i < int(l);i++{
        fk,fv := reflect.New(ekt),reflect.New(evt)
        if err := dec.value(fk);err!=nil{
            return nil
        }
        if err := dec.value(fv);err!=nil{
            return nil
        }
        newMap.SetMapIndex(fk.Elem(),fv.Elem())
    }
    scan.until('}')

    return nil
}

func(dec *textDecoder)object(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    if out.Kind() != reflect.Struct{
        return fmt.Errorf("type error")
    }
    scan := dec.scan
    if err := scan.tag('a');err!=nil{
        return err
    }
    l,err :=scan.length()
    if err!=nil{
        return err
    }
    scan.expect('{')

    vmap,err := dec.fieldMap(out)
    if err != nil{
        return err
    }

    for i:= 0 ;i < int(l) ;i++{
        tag :=""
        if err:= dec.decode(&tag);err!=nil{
            return err
        }
        if pev ,ok := vmap[tag];ok{
            if err := dec.value(*pev);err!=nil{
                return err
            }
        }else{
            err:= dec.scan.ignore()
            if err != nil{
                return err
            }
        }
    }
    scan.until('}')
    return nil
}

func(dec *textDecoder)fieldMap(out reflect.Value)(map[string]*reflect.Value,error){
    if dec.err!=nil{
        return nil,dec.err
    }
    fieldMap := make(map[string]*reflect.Value)
    ot := out.Type()
    for i := 0 ; i < out.NumField();i++{
        ev := out.Field(i)
        et := ot.Field(i)
        tag,ok := et.Tag.Lookup("php")
        if ok&& tag !=""{
            t,_ :=parseTag(tag)
            if t != ""{
                fieldMap[t] = &ev
            }
        }
    }
    return fieldMap,nil
}

func(dec *textDecoder)ptr(out reflect.Value)error{
    if dec.err!=nil{
        return dec.err
    }
    rout := out
    if out.IsNil(){
        rout = reflect.New(out.Type().Elem())
        out.Set(rout)
    }
    rout = reflect.Indirect(rout)
    return dec.value(rout)
}


