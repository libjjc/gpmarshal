package php

import (
    "fmt"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/yvasiyarov/php_session_decoder/php_serialize"
    "math"
    "testing"
)

func TestSerialize(t *testing.T){
    Convey("Field Testing",t,func() {

        Convey("Int Testing", func() {
            int8MarshalTest()
            int16MarshalTest()
            int32MarshalTest()
            int64MarshalTest()
            intMarshalTest()
            uint8MarshalTest()
            uint16MarshalTest()
            uint32MarshalTest()
            uint64MarshalTest()
            uintMarshalTest()
        })



        Convey("Float Testing",func(){
            float32MarshalTest()
            float64MarshalTest()
        })

        Convey("String Testing",func(){
            stringMarshalTest()
        })

        Convey("Bool Testing",func(){
            boolMarshalTest()
        })

    })

    Convey("Struct Testing",t,func(){
        Convey("Struct Simple Testing",func(){
            structSimpleTest()
            structRecursiveTest()
        })
    })

    Convey("Map Testing",t,func(){
        Convey("Map[int]int Testing",func(){
            mapIntIntTest()
        })
        Convey("Map[bool]float64 Testing",func(){
            mapBoolFloatTest()
        })
        Convey("Map[string]string Testing",func(){
            mapStringStringTest()
        })
        Convey("Map[string]*int Testing",func(){
            mapStringIntPtrTest()
        })
    })

    Convey("Complex Testing",t,func(){
        Convey("Complex Struct Testing",func(){
            complexStructTest()
        })

        Convey("Complex Map Testing",func(){
            complexMapTest()
        })

        Convey("Complex Slice Testing",func(){
            complexSliceTest()
        })
    })
}

func int8MarshalTest(){
    in := []int8{-127,-3,-2,-1,0,1,2,3,4,5,125,126,127}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }

    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func int16MarshalTest(){
    in := []int16{math.MaxInt16,math.MaxInt16-1,-3,-2,-1,0,1,2,3,4,5,125,126,127,127,math.MinInt16,math.MinInt16+1}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func int32MarshalTest(){
    in := []int32{math.MaxInt32,math.MaxInt32-1,-3,-2,-1,0,1,2,3,4,5,125,126,127,127,math.MinInt32,math.MinInt32+1}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func int64MarshalTest(){
    in := []int64{math.MaxInt64,math.MaxInt64-1,-3,-2,-1,0,1,2,3,4,5,125,126,127,127,math.MinInt64,math.MinInt64+1}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func intMarshalTest(){
    in := []int{math.MaxInt64,math.MaxInt64-1,-3,-2,-1,0,1,2,3,4,5,125,126,127,127,math.MinInt64,math.MinInt64+1}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func uint8MarshalTest(){
    in := []uint8{0,1,2,3,4,5,125,126,127}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }

    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func uint16MarshalTest(){
    in := []uint16{math.MaxUint16,math.MaxUint16-2,0,1,2,3,4,5,125,126,127}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func uint32MarshalTest(){
    in := []uint32{math.MaxUint32,math.MaxUint32-1,0,1,2,3,4,5,125,126,127,127}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func uint64MarshalTest(){
    in := []uint64{math.MaxUint64,math.MaxUint64-1,0,1,2,3,4,5,125,126,127,127}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func uintMarshalTest(){
    in := []uint{math.MaxUint64,math.MaxUint64-1,0,1,2,3,4,5,125,126,127,127}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        So(s,ShouldEqual,fmt.Sprintf("i:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        So(sp,ShouldEqual,fmt.Sprintf("i:%d;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func float32MarshalTest(){
    in := []float32{math.MaxFloat32,math.MaxFloat32+0.0000001,math.MaxFloat32+0.1,math.MaxFloat32+1,
        -1.1,-1.0,-0.1,-0.0000001,0,1,0.0000001,0.1,0.11,-1.0,-1.1,-11.11,-222.222,-3333.3333,
    math.MaxFloat32-1,math.MaxFloat32-0.000001,math.MaxFloat32-0.1,math.MaxFloat32 -1}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        //So(s,ShouldEqual,fmt.Sprintf("d:%1.6f;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        //So(sp,ShouldEqual,fmt.Sprintf("d:%1.6f;",i))
        pa = append(pa,i)
    }
    // arrStr , _ := php_serialize.Serialize(pa)
    _ ,err := Marshal(in)
    So(err,ShouldBeNil)
    // So(res,ShouldEqual,arrStr)
    _,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    // So(resp,ShouldEqual,arrStr)
}

func float64MarshalTest(){
    in := []float64{math.MaxFloat64,math.MaxFloat64+0.0000001,math.MaxFloat64+0.1,math.MaxFloat64+1,
        -1.1,-1.0,-0.1,-0.0000001,0,1,0.0000001,0.1,0.11,-1.0,-1.1,-11.11,-222.222,-3333.3333,
        math.MaxFloat64-1,math.MaxFloat64-0.000001,math.MaxFloat64-0.1,math.MaxFloat64 -1}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        //So(s,ShouldEqual,fmt.Sprintf("d:%1.6f;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        //So(sp,ShouldEqual,fmt.Sprintf("d:%1.6f;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func stringMarshalTest(){
    in := []string{"","1","ok","hello world!l","中国","`!@#$$%^&*()[]{}|\\`·2·44，。《》/【】",
        "fs;dlkfjslfjslkfjlskfjpwoiehfdns;lkfj;asdf]wjdfjpsa;fk'asldf]awdsbfn;saflkas'dfash" +
        "pdfksad;fm'aslfia[sjdfj;asfjs;flfsjflksdflkshdfsjdkfksjkfjslkdfjlsjflsdkljlljljllj" +
        "s;fs[dfhspjdfsjf;sfl[shfpsafaaaaaaaaasfsdklfjlshfdsakldfjaaaaaaaaaaaaaaaaaaaaaaaaa" +
        "asjfdas[fsajf;sjf[asdfhspjd;flkjsdfbskhxlvbcaslkfhjdawksjdbflasfa;sdhfblsfasf;'saf" +
        "sfj;a'sfjsadfjas['fdjskjfwyetroiwusblskdkjfiwuy385oe4rfbwnsnas'dfajs[dfhwsjdfjsa'f" +
        "ajsfs;dfjasfoishetfowsuyfgjbnqowlasdifyhcbwjkaskdfcajhosedzlfkhopdhfjas;dflkas'fjk" +
        "afjspfkjasnfksxjghvasidfkhopisadhfjpwosdjkfhjpsalkdhfjopskhfjaspodlfjsal;dfja'fda]" +
        "fjas;dfljapso;dfljsjdfpwjsbdfpsi\rwl;esfjs';flkas;jgfopwaskhdfjas;llfja'f'jas[f;ja" +
        "sfan;sfjas;dfl;asjl;dkfjas;ldfja;sj;aslkdfjasodfkhjasokdfhaops;dlkfjas;dfjasd;fjas" +
        "dffsodfkhawosdfhawlsdkhfawiosdhfwoasdhf\n\tkosjfawhasdjfpsahfjaslkfhashfh;ksfkhjjl" +
        "ajfasfjaw;olsdfkjals;jfdals;jfsa;ljfl;asjfl;asjdfolsjdflwjdsfo;lkajw;olsdfjwalskdj" +
        "folawjsdflawjsdfo;lkjaws;ds;fjsladfjkbawsjdhfbalsdflanskfjnalskjf;laljfnkjfljfja]s" +
        "fdja;sf;alfaslkfhnaslkdfjas;kldfh;alsjf;ajsdl;fk;asjdf;nasjdf;j;ljf;oskandfjsdflsf" +
        "中华人民共和国，人民有信仰，国家有力量"}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        //So(s,ShouldEqual,fmt.Sprintf("d:%1.6f;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        //So(sp,ShouldEqual,fmt.Sprintf("d:%1.6f;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

func boolMarshalTest(){
    in := []bool{true,false,true,true,true,false,false,false,false}
    pa := make(php_serialize.PhpSlice,0)
    for _,i := range in{
        s,e := Marshal(i)
        fmt.Println("Marshal(",i,")=",s," | err=",e)
        So(e,ShouldBeNil)
        //So(s,ShouldEqual,fmt.Sprintf("b:%d;",i))
        sp,ep :=Marshal(&i)
        fmt.Println("Marshal(",i,")=",sp," | err=",ep)
        So(ep,ShouldBeNil)
        //So(sp,ShouldEqual,fmt.Sprintf("d:%1.6f;",i))
        pa = append(pa,i)
    }
    arrStr , _ := php_serialize.Serialize(pa)
    res ,err := Marshal(in)
    So(err,ShouldBeNil)
    So(res,ShouldEqual,arrStr)
    resp,errp := Marshal(&in)
    So(errp,ShouldBeNil)
    So(resp,ShouldEqual,arrStr)
}

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
func structSimpleTest(){
    sns := &SimpleNormalStruct{
        I8:    -8,
        I16:   -16,
        I32:   -32,
        I64:   -64,
        I:     -1,
        U8:    8,
        U16:   16,
        U32:   32,
        U64:   64,
        U:     1,
        F32:   1.1,
        F64:   -234.567,
        Empty: "",
        Hello: "hello world!",
        True:  true,
        False: false,
    }
    s , e := Marshal(sns)
    fmt.Println(s)
    So(e,ShouldBeNil)
    s,e = Marshal(*sns)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
}

func structRecursiveTest(){
    type RecursiveStruct struct{
        Name string `php:"name"`
        Recursive SimpleNormalStruct `php:"recursive"`
    }

    rs := &RecursiveStruct{
        Name:      "test name",
        Recursive: SimpleNormalStruct{
            I8:    -8,
            I16:   -16,
            I32:   -32,
            I64:   -64,
            I:     -1,
            U8:    8,
            U16:   16,
            U32:   32,
            U64:   64,
            U:     1,
            F32:   1.1,
            F64:   -234.567,
            Empty: "",
            Hello: "hello world!",
            True:  true,
            False: false,
        },
    }
    s , e := Marshal(rs)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    s,e = Marshal(*rs)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
}

func structUnexportedTest(){
    type UnexportedStruct struct{
        Name string `php:"name"`
        url string `php:"url"`
    }

    rs := &UnexportedStruct{
        Name:      "NNNNNNName",
        url:"http://balabalabala.com",
    }
    s , e := Marshal(rs)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    s,e = Marshal(*rs)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
}

func mapIntIntTest(){
    in := map[int]int{-3:-33,-2:22,-1:-11,0:0,1:-111,2:222,3:-333}
    s,e := Marshal(in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    in = map[int]int{}
    s,e = Marshal(in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
}

func mapBoolFloatTest(){
    in := map[bool]float32{true:-33,false:11}
    s,e := Marshal(in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)

    in = map[bool]float32{}
    s,e = Marshal(in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
}

func mapStringStringTest(){
    in := map[string]string{"key1":"value1","":""}
    s,e := Marshal(in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)

    in = map[string]string{}
    s,e = Marshal(in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
}

func mapStringIntPtrTest(){
    a ,b  := -33,0
    in := map[string]*int{"key1":&a,"":&b}
    s,e := Marshal(in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)

    in = map[string]*int{}
    s,e = Marshal(in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
}

type structSimpleSample struct{
    N int `php:"n"`
    S string `php:"s"`
    B bool `php:"b"`
    F float32 `php:"f"`
}

type structComplexSample struct{
    Stru *structSimpleSample `php:"stru,omitempty"`
    Maps map[string]*structSimpleSample `php:"leet,omitempty"`
    Slice []*structSimpleSample `php:"slice,omitempty"`
}

func complexMapTest(){
    in1 := make(map[string]*structSimpleSample)
    s,e := Marshal(in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)

    in1["first"] = &structSimpleSample{
        N: 1,
        S: "2",
        B: false,
        F: 0.0,
    }
    in1["last"] = &structSimpleSample{
        N: 2,
        S: "",
        B: true,
        F: 0,
    }
    s,e = Marshal(in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)

    in2 := make(map[string][]*structSimpleSample)
    s,e = Marshal(in2)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)

    in2["first"] = append(in2["first"],&structSimpleSample{
        N: 1,
        S: "2",
        B: false,
        F: 0.0,
    })
    in2["first"] = append(in2["first"],&structSimpleSample{
        N: 2,
        S: "",
        B: true,
        F: 0,
    })
    s,e = Marshal(in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
}

func complexSliceTest(){
    in1 := []*structSimpleSample{
        {
            N: 1,
            S: "2",
            B: false,
            F: 0.0,
        },
        {
            N: 2,
            S: "",
            B: true,
            F: 0,
        },
    }
    s,e := Marshal(in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)

    in2 := []structSimpleSample{}
    s,e = Marshal(in2)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)

    in2 = append(in2,structSimpleSample{
        N: 1,
        S: "2",
        B: false,
        F: 0.0,
    },structSimpleSample{
        N: 2,
        S: "",
        B: true,
        F: 0,
    })
    s,e = Marshal(in2)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&in1)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
}

func complexStructTest(){
    cst := &structComplexSample{}
    s,e := Marshal(cst)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&cst)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    cst.Stru = &structSimpleSample{
        N: 0,
        S: "",
        B: false,
        F: 0,
    }
    cst.Slice = []*structSimpleSample{
        {
            N: 0,
            S: "",
            B: false,
            F: 0,
        },
        {
            N: 2,
            S: "hello world!",
            B: true,
            F: 0,
        },
    }
    cst.Maps=map[string]*structSimpleSample{
        "test":{
            N: 2,
            S: "hello world!",
            B: true,
            F: 0,
        },
    }
    s,e = Marshal(cst)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
    s,e = Marshal(&cst)
    fmt.Println(s,e)
    So(e,ShouldBeNil)
    So(s,ShouldNotBeEmpty)
}
