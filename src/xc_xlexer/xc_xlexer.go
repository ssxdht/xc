package xc_xlexer

//import "os"
import "fmt"
import "io/ioutil"
import "container/list"

import "xc_err"
import "xc_datadef"

func Xlexer(file_name string)(*list.List){

    file_data, err := ioutil.ReadFile(file_name)
    xc_err.CheckErr(err)

    tokens := list.New()

    i := 0
    for ; i<len(file_data); {
        switch file_data[i] {
            case '+':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_PLUS, "+"})
                i++
            case '-':
                i++
                if file_data[i] == '>' {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_POINTSTO, "->"})
                    i++
                } else {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_MINUS, "-"})
                }
            case '*':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_STAR, "*"})
                i++
            case '/':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_DIVIDE, "/"})
                i++
            case '%':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_MOD, "%"})
                i++
            case '=':
                i++
                if file_data[i] == '=' {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_EQ, "=="})
                    i++
                } else {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_ASSIGN, "="})
                }
            case '<':
                i++
                if file_data[i] == '=' {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_LEQ, "<="})
                    i++
                } else {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_LT, "<"})
                }
            case '>':
                i++
                if file_data[i] == '=' {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_GEQ, ">="})
                    i++
                } else {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_GT, ">"})
                }
            case '.':
                i++
                if file_data[i] == '.' {
                    i++
                    if file_data[i] == '.' {
                        tokens.PushBack(xc_datadef.Token{xc_datadef.TK_ELLIPSIS, "..."})
                        i++
                    } else {
                        i--
                    }
                } else {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_DOT, "."})
                }
            case '&':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_AND, "&"})
                i++
            case '(':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_OPENPA, "("})
                i++
            case ')':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_CLOSEPA, ")"})
                i++
            case '[':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_OPENBR, "["})
                i++
            case ']':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_CLOSEBR, "]"})
                i++
            case '{':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_BEGIN, "{"})
                i++
            case '}':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_END, "}"})
                i++
            case ';':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_SEMICOLON, ";"})
                i++
            case ',':
                tokens.PushBack(xc_datadef.Token{xc_datadef.TK_COMMA, ","})
                i++
            default:
                i++
                fmt.Println("incorrect character!")
                break
        }
    }

    return tokens

}

