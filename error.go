package php

import (
    "errors"
    "fmt"
    "reflect"
)

var(
    nullPhpValue = newMarshalError(errors.New("null php value"))
    extends = &errorExtends{}
)
type SerializeError struct{}


type MarshalError struct{
    err error
    stage string
    reason string
    context string
}

func newMarshalError(err error)*MarshalError{
    return &MarshalError{
        err :err,
    }
}


type extendFunc func(*MarshalError)*MarshalError

func(error *MarshalError) withExtend(f extendFunc)*MarshalError{
    return f(error)
}

func (error *MarshalError) Error()string{
    es := error.err.Error()
    if error.stage != ""{
        es = fmt.Sprintf("%s | reason : %s",es,error.stage)
    }
    if error.reason != ""{
        es = fmt.Sprintf("%s | reason : %s",es,error.reason)
    }
    if error.context != ""{
        es = fmt.Sprintf("%s | context : %s",es,error.context)
    }
    return es
}

type errorExtends struct{}

func (ext *errorExtends)encoding() extendFunc{
    return func(err *MarshalError)*MarshalError{
        err.stage = "Encode"
        return err
    }
}

func (ext *errorExtends)decoding() extendFunc{
    return func(err *MarshalError)*MarshalError{
        err.stage = "Decode"
        return err
    }
}

func(ext *errorExtends)unexpected(actually,expected interface{})extendFunc{
    return func(err *MarshalError)*MarshalError{
        err.reason = fmt.Sprintf("unexpected token,actually=%v,expect=%v",actually,expected)
        return err
    }
}

func(ext *errorExtends)invalidConversation(v interface{},to reflect.Type)extendFunc{
    return func(err *MarshalError)*MarshalError{
        err.reason = fmt.Sprintf("invalid conversation,val : %v from=%s,to=%s",v,reflect.TypeOf(v).Name(),to.Name())
        return err
    }
}

func(ext *errorExtends)unRecognize(v interface{})extendFunc{
    return func(err *MarshalError)*MarshalError{
        err.reason = fmt.Sprintf("unexpected token : %v",v)
        return err
    }
}

func(ext *errorExtends)outOfRange(v,max int)extendFunc{
    return func(err *MarshalError)*MarshalError{
        err.reason = fmt.Sprintf("out of range : index : %d , max : %d",v,max)
        return err
    }
}





