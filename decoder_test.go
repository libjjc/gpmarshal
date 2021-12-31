package php

import (
    "fmt"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/yvasiyarov/php_session_decoder/php_serialize"
    "math"
    "testing"
)

func TestUnmarshal(t *testing.T){

    Convey("Base Type Test",t,func(){

        Convey("int type test",func(){
            intUnmarshalTest()
        })

        Convey("uint type test",func(){
            uintUnmarshalTest()
        })

        Convey("float type test",func(){
            floatUnmarshalTest()
        })


        Convey("string type test",func(){
            stringUnmarshalTest()
        })

        Convey("bool type test",func(){
            boolUnmarshalTest()
        })


    })



    Convey("Struct Test",t,func(){

        Convey("simple struct test",func(){
            simpleStructTest()
        })
        Convey("inherited struct test",func(){
            inheritedStructTest()
        })
        Convey("combination struct test",func(){
            combinationStructTest()
        })
    })



    Convey("Slice Test",t,func(){
        sliceUnmarshalTest()
    })

    Convey("Array Test",t,func(){
        arrayUnmarshalTest()
    })
    Convey("Map Test",t,func(){
        mapUnmarshalTest()
    })
    Convey("Complex Type Test",t,func(){
        complexFieldTest()
    })


}

func intUnmarshalTest(){
    i8s := []int8{/*math.MinInt8,math.MinInt8+1,-1,0,1,math.MaxInt8-1,math.MaxInt8*/ }
    i8test := func(in int8){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        fmt.Println(s,e)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r int8
        e = Unmarshal(s,&r)
        fmt.Println(s,e)
        So(r,ShouldEqual,in)
        So(e,ShouldBeNil)
    }
    for _,i := range i8s{
        i8test(i)
    }


    i16s := []int16{math.MinInt16,math.MinInt16+1,-1,0,1,math.MaxInt16-1,math.MaxInt16 }
    i16test := func(in int16){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r int16
        e = Unmarshal(s,&r)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }
    for _,i := range i16s{
        i16test(i)
    }

    i32s := []int32{math.MinInt32,math.MinInt32+1,-1,0,1,math.MaxInt32-1,math.MaxInt32 }
    i32test := func(in int32){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r int32
        e = Unmarshal(s,&r)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }
    for _,i := range i32s{
        i32test(i)
    }

    i64s := []int64{math.MinInt64,math.MinInt64+1,-1,0,1,math.MaxInt64-1,math.MaxInt64 }
    i64test := func(in int64){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r int64
        e = Unmarshal(s,&r)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }
    for _,i := range i64s{
        i64test(i)
    }

    is := []int{math.MinInt64,math.MinInt64+1,-1,0,1,math.MaxInt64-1,math.MaxInt64 }
    itest := func(in int){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r int
        e = Unmarshal(s,&r)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }
    for _,i := range is{
        itest(i)
    }
}

func uintUnmarshalTest(){
    i8s := []uint8{0,1,math.MaxUint8-1,math.MaxUint8 }
    i8test := func(in uint8){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r uint8
        e = Unmarshal(s,&r)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }
    for _,i := range i8s{
        i8test(i)
    }

    i16s := []uint16{0,1,math.MaxUint16-1,math.MaxUint16 }
    i16test := func(in uint16){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r uint16
        e = Unmarshal(s,&r)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }
    for _,i := range i16s{
        i16test(i)
    }

    i32s := []uint32{0,1,math.MaxUint32-1,math.MaxUint32 }
    i32test := func(in uint32){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r uint32
        e = Unmarshal(s,&r)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }
    for _,i := range i32s{
        i32test(i)
    }


    i64s := []uint64{0,1,math.MaxUint64-1,math.MaxUint64 }
    i64test := func(in uint64){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        fmt.Println(s,e)
        // So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r uint64
        e = Unmarshal(s,&r)
        fmt.Println(s,e)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }
    for _,i := range i64s{
        i64test(i)
    }


    is := []uint{0,1,math.MaxUint64-1,math.MaxUint64 }
    itest := func(in uint){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        fmt.Println(s,e)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",in))
        So(e,ShouldBeNil)
        var r uint
        e = Unmarshal(s,&r)
        fmt.Println(s,e)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }
    for _,i := range is{
        itest(i)
    }

}

func floatUnmarshalTest(){
    f32s := []float32{-1 * math.MaxFloat32,-1.0,0.0,1.0,math.MaxFloat32}
    f64s := []float64{-1 * math.MaxFloat64,-1.0,0.0,1.0,math.MaxFloat64}

    f32test := func(in float32){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(e,ShouldBeNil)
        var r float32
        exp := 0.000001
        e = Unmarshal(s,&r)
        So(math.Abs(float64(in-r)),ShouldBeLessThanOrEqualTo,exp)
        So(e,ShouldBeNil)
    }
    f64test := func(in float64){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(e,ShouldBeNil)
        var r float64
        exp := 0.000001
        e = Unmarshal(s,&r)
        So(math.Abs(in-r),ShouldBeLessThanOrEqualTo,exp)
        So(e,ShouldBeNil)
    }

    for _,f := range f32s{
        f32test(f)
    }

    for _,f := range f64s{
        f64test(f)
    }
}

func stringUnmarshalTest(){
    strs := []string{"","hello world!","zhong国","!@#@%$#^$%^$@#&!@【】，。/你吃不··。，/'；！#"}

    strTest := func(in string){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        fmt.Println(s,e)
        So(e,ShouldBeNil)
        r :=""
        e = Unmarshal(s,&r)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }

    for _,s := range strs{
        strTest(s)
    }
}

func boolUnmarshalTest(){
    bs := []bool{true,false,false,false,true,true,false}
    strTest := func(in bool){
        pv := php_serialize.PhpValue(in)
        s,e := php_serialize.Serialize(pv)
        So(e,ShouldBeNil)
        var r bool
        e = Unmarshal(s,&r)
        So(in,ShouldEqual,r)
        So(e,ShouldBeNil)
    }

    for _,s := range bs{
        strTest(s)
    }
}

type InheritedStruct struct{
    SimpleNormalStruct
    Name string
}
func simpleStructTest(){
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
    So(e,ShouldBeNil)
    e = Unmarshal(s,out)
    So(e,ShouldBeNil)
    So(out.I8,ShouldEqual,8)
    So(out.I16,ShouldEqual,-16)
    So(out.I32,ShouldEqual,-32)
    So(out.I64,ShouldEqual,-64)
    So(out.I,ShouldEqual,-1)
    So(out.I8,ShouldEqual,8)
    So(out.U16,ShouldEqual,16)
    So(out.U32,ShouldEqual,32)
    So(out.U64,ShouldEqual,64)
    So(out.U,ShouldEqual,1)
    So(math.Abs(float64(out.F32-32.32)),ShouldBeLessThanOrEqualTo,0.000001)
    So(math.Abs(out.F64-64.64),ShouldBeLessThanOrEqualTo,0.000001)
    So(out.True,ShouldBeTrue)
    So(out.False,ShouldBeFalse)
}
type Struct1 struct{
    Name string `php:"name"`
}

type Struct2 struct{
    Struct1 `php:"base"`
    Age int `php:"age"`
}

type Struct3 struct{
    S1 Struct1 `php:"s1"`
    S2 Struct2 `php:"s2"`
    Sex int `php:"sex"`
}
func inheritedStructTest(){
    in := php_serialize.PhpArray{
        "base":php_serialize.PhpArray{
            "name":"jjchen",
        },
        "age":33,
    }
    out := &Struct2{}
    s,e := php_serialize.Serialize(in)
    So(e,ShouldBeNil)
    e = Unmarshal(s,out)
    So(e,ShouldBeNil)
    So(out.Name,ShouldEqual,"jjchen")
    So(out.Age,ShouldEqual,33)
}
func combinationStructTest(){
    in := php_serialize.PhpArray{
        "s1":php_serialize.PhpArray{
            "name":"jjchen",
        },
        "s2":php_serialize.PhpArray{
            "age":33,
        },
        "sex":1,
    }
    out := &Struct3{}
    s,e := php_serialize.Serialize(in)
    So(e,ShouldBeNil)
    e = Unmarshal(s,out)
    So(e,ShouldBeNil)
    So(out.S1.Name,ShouldEqual,"jjchen")
    So(out.S2.Age,ShouldEqual,33)
    So(out.Sex,ShouldEqual,1)
}

func sliceUnmarshalTest(){
    in := php_serialize.PhpSlice{
        "","hello","中国","#@$#%#$【】'；，。",
    }
    s,e := php_serialize.Serialize(in)
    So(e,ShouldBeNil)
    rs:=make( []string,0)
    e = Unmarshal(s,&rs)
    fmt.Println(rs)
}

func arrayUnmarshalTest(){
    in := php_serialize.PhpSlice{
        "","hello","中国","#@$#%#$【】'；，。",
    }
    s,e := php_serialize.Serialize(in)
    So(e,ShouldBeNil)
    rs:=[4]string{}
    e = Unmarshal(s,&rs)
    fmt.Println(rs)
}

func mapUnmarshalTest(){
    in := php_serialize.PhpSlice{
        "","hello","中国","#@$#%#$【】'；，。",
    }
    s,e := php_serialize.Serialize(in)
    So(e,ShouldBeNil)
    rs:=map[int]string{}
    e = Unmarshal(s,&rs)
    fmt.Println(rs)
}


type ComplexPointerStruct struct{
    P1 *int `php:"p1"`
}
type ComplexStruct struct{
    S3 Struct3 `php:"s3"`
    As2 []*Struct2 `php:"as2"`
    Ms1 map[string]*Struct1 `php:"ms1"`
    Ps3 *Struct3 `php:"ps3"`
    Pas2 *[]*Struct2 `php:"pas2"`
    Pms1 map[string]*Struct1 `php:"pms1"`
    Pcp *ComplexPointerStruct `php:"pcp"`
}

func complexFieldTest(){
    in := php_serialize.PhpArray{
        "s3":php_serialize.PhpArray{
            "s1":php_serialize.PhpArray{
                "name":"jjchen",
            },
            "s2":php_serialize.PhpArray{
                "age":33,
            },
            "sex":1,
        },
        "ps3":php_serialize.PhpArray{
            "s1":php_serialize.PhpArray{
                "name":"jjchen",
            },
            "s2":php_serialize.PhpArray{
                "age":33,
            },
            "sex":1,
        },
        "as2":php_serialize.PhpSlice{
            php_serialize.PhpArray{
                "age": 18,
            },
            php_serialize.PhpArray{
                "age": 880,
            },
        },
        "pas2":php_serialize.PhpSlice{
        php_serialize.PhpArray{
        "age": 18,
    },
        php_serialize.PhpArray{
        "age": 880,
    },
    },
        "ms1":php_serialize.PhpArray{
            "jjchen":php_serialize.PhpArray{
                "name":"chen",
            },
            "wenwen":php_serialize.PhpArray{
                "name":"ivan",
            },
        },
        "pms1":php_serialize.PhpArray{
            "jjchen":php_serialize.PhpArray{
                "name":"chen",
            },
            "wenwen":php_serialize.PhpArray{
                "name":"ivan",
            },
        },
        "pcp":php_serialize.PhpArray{
            "p1":5,
        },
    }
    s,e := php_serialize.Serialize(in)
    So(e,ShouldBeNil)
    rs:= &ComplexStruct{}
    e = Unmarshal(s,&rs)
    fmt.Println(rs)
    So(rs.S3.Sex,ShouldEqual,1)
    So(rs.S3.S1.Name,ShouldEqual,"jjchen")
    So(rs.S3.S2.Age,ShouldEqual,33)
    So(rs.Ps3.Sex,ShouldEqual,1)
    So(rs.Ps3.S1.Name,ShouldEqual,"jjchen")
    So(rs.Ps3.S2.Age,ShouldEqual,33)
    So(rs.As2[0].Age,ShouldEqual,18)
    So(rs.As2[1].Age,ShouldEqual,880)
    So((*rs.Pas2)[0].Age,ShouldEqual,18)
    So((*rs.Pas2)[1].Age,ShouldEqual,880)
    So(rs.Ms1["jjchen"].Name,ShouldEqual,"chen")
    So(rs.Ms1["wenwen"].Name,ShouldEqual,"ivan")
    So(rs.Pms1["jjchen"].Name,ShouldEqual,"chen")
    So(rs.Pms1["wenwen"].Name,ShouldEqual,"ivan")
    So(*rs.Pcp.P1,ShouldEqual,5)
}

