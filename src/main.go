package main

import "os"
import "fmt"
//import "reflect"
//import "io/ioutil"

//import "xc_err"
import "xc_xlexer"
import "xc_datadef"

func main() {
    if len(os.Args)<2{
        fmt.Printf("用法：%s [source file]\n", os.Args[0])
        os.Exit(-1)
    }

    //fmt.Printf("xc_datadef.KW_CHAR=[%d]\n", xc_datadef.KW_CHAR)
    //fmt.Println("xc_datadef.KW_CHAR's type:", reflect.TypeOf(xc_datadef.KW_CHAR))

    tokens := xc_xlexer.Xlexer(os.Args[1])
    //xc_err.CheckErr(err)
    for e:=tokens.Front(); e!=nil; e=e.Next(){
        //fmt.Println("e.Value's type:", reflect.TypeOf(e.Value))
        if token, ok := e.Value.(xc_datadef.Token); ok{
            fmt.Printf("TokenType:[%02d]-->TokeValue[%s]\n", token.TokenType, token.TokenValue)
        }
    }

}
