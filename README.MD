### install
```
go get sgithub.com/libjjc/gpmarshal
```

### Example
```
    type SimpleNormalStruct struct {
        I8 int8 `php:"i8"`
        I16 int16 `php:"i16"`
        I32 int32 `php:"i32"`
        I64 int64 `php:"i64"`
        I int `php:"i"`
        U8 uint8 `php:"u8"`
        U16 uint16 `php:"u16"`
        U32 uint32 `php:"u32"`
        U64 uint64 `php:"u64"`
        U uint `php:"u"`
        F32 float32 `php:"f32"`
        F64 float64 `php:"f64"`
        Empty string `php:"empty"`
        Hello string `php:"hello"`
        True bool `php:"true"`
        False bool `php:"false"`
    }
    out := &SimpleNormalStruct{}
    in := php_serialize.PhpArray{
        "i8":    8,
        "i16":   -16,
        "i32":   -32,
        "i64":   -64,
        "i":     -1,
        "u8":    8,
        "u16":   16,
        "u32":   32,
        "u64":   64,
        "u":     1,
        "f32":   32.32,
        "f64":   64.64,
        "empty": "",
        "hello": "hello",
        "true":  true,
        "false": false,
    }
    s,e := php_serialize.Serialize(in)
    e = php.Unmarshal(s,out)
    fmt.Println(s,e)
```
