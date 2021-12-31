### install
```
go get sgithub.com/libjjc/gpmarshal
```

### Example
```go

package main

import(
    "fmt"
    ."github.com/jjchen/gpmarshal"
)

func main(){

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

    in := `a:16:{s:4:"true";b:1;s:3:"u16";i:16;s:3:"i16";i:-16;s:3:"i32";i:-32;s:1:"i";i:-1;s:3:"u32";i:32;s:1:"u";i:1;s:5:"empty";s:0:"";s:5:"false";b:0;s:2:"i8";i:8;s:5:"hello";s:5:"hello";s:2:"u8";i:8;s:3:"u64";i:64;s:3:"f32";d:32.32;s:3:"f64";d:64.640000000000001;s:3:"i64";i:-64;}`

    out := &SimpleNormalStruct{}

    e := gpmarshal.Unmarshal(in,out)

    fmt.Println(in,out,e)

}

```


### todo
- error type recognizable
- error context details
- apiMarshal&apiUnmarshal for type of the member of struct is interface
- more tag support
- more type support
