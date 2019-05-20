package xc_xlexer

import "unicode"

func ParseNumber(b []byte, curPos int)(int, string, bool){

    str := ""

    i := curPos+1
    for ; i<len(b); {
        if unicode.IsNumber(rune(b[i])) {
            i++
        } else {
            if b[i]==';' || unicode.IsSpace(rune(b[i])) {
                str = string(b[curPos:i])
                return i, str, true
            }else{
                return i, "", false
            }
        }
    }

    return i, "", false

}

