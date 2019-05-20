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

    map_keywords := make(map[string]int)
    InitKeywordsMap(map_keywords)

    tokens := list.New()
    str := ""
    ok  := false

    i := 0
    for ; i<len(file_data); {
        i = SkipSpace(file_data, i)
        if i>=len(file_data) {
            break
        }
        switch file_data[i] {
            case '_': fallthrough
            case 'a': fallthrough
            case 'b': fallthrough
            case 'c': fallthrough
            case 'd': fallthrough
            case 'e': fallthrough
            case 'f': fallthrough
            case 'g': fallthrough
            case 'h': fallthrough
            case 'i': fallthrough
            case 'j': fallthrough
            case 'k': fallthrough
            case 'l': fallthrough
            case 'm': fallthrough
            case 'n': fallthrough
            case 'o': fallthrough
            case 'p': fallthrough
            case 'q': fallthrough
            case 'r': fallthrough
            case 's': fallthrough
            case 't': fallthrough
            case 'u': fallthrough
            case 'v': fallthrough
            case 'w': fallthrough
            case 'x': fallthrough
            case 'y': fallthrough
            case 'z': fallthrough
            case 'A': fallthrough
            case 'B': fallthrough
            case 'C': fallthrough
            case 'D': fallthrough
            case 'E': fallthrough
            case 'F': fallthrough
            case 'G': fallthrough
            case 'H': fallthrough
            case 'I': fallthrough
            case 'J': fallthrough
            case 'K': fallthrough
            case 'L': fallthrough
            case 'M': fallthrough
            case 'N': fallthrough
            case 'O': fallthrough
            case 'P': fallthrough
            case 'Q': fallthrough
            case 'R': fallthrough
            case 'S': fallthrough
            case 'T': fallthrough
            case 'U': fallthrough
            case 'V': fallthrough
            case 'W': fallthrough
            case 'X': fallthrough
            case 'Y': fallthrough
            case 'Z':
                i, str, ok = ParseIdentifier(file_data, i)
                if ok {
                    if value, ok := map_keywords[str]; ok{
                        tokens.PushBack(xc_datadef.Token{value, str})
                    }else{
                        tokens.PushBack(xc_datadef.Token{xc_datadef.TK_IDENT, str})
                    }
                    str = ""
                }else {
                    fmt.Println("incorrect number!")
                }
            case '0': fallthrough
            case '1': fallthrough
            case '2': fallthrough
            case '3': fallthrough
            case '4': fallthrough
            case '5': fallthrough
            case '6': fallthrough
            case '7': fallthrough
            case '8': fallthrough
            case '9':
                i, str, ok = ParseNumber(file_data, i)
                if ok {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_CINT, str})
                    str = ""
                }else {
                    fmt.Println("incorrect number!")
                }
            case '"':
                i, str, ok = ParseString(file_data, i)
                if ok {
                    tokens.PushBack(xc_datadef.Token{xc_datadef.TK_CSTR, str})
                    str = ""
                }else {
                    fmt.Println("incorrect string!")
                }
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
                fmt.Println("incorrect character!", i)
                break
        }
    }

    return tokens

}

