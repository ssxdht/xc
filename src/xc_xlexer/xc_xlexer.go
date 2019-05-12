package xc_xlexer

//import "os"
//import "fmt"
import "io/ioutil"
import "container/list"

//import "xc_err"
import "xc_datadef"

func Xlexer(file_name string)(*list.List){

    ioutil.ReadFile(file_name)

    tokens := list.New()

    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_PLUS, "+"})
    tokens.PushBack(xc_datadef.Token{xc_datadef.KW_VOID, "void"})
    tokens.PushBack(xc_datadef.Token{xc_datadef.KW_CHAR, "char"})
    tokens.PushBack(xc_datadef.Token{xc_datadef.KW_ELSE, "else"})

    return tokens

}
