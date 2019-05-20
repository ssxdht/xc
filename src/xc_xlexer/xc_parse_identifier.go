package xc_xlexer

import "unicode"

func ParseIdentifier(b []byte, curPos int)(int, string, bool){

    str := ""

    i := curPos+1
    for ; i<len(b); {
        if b[i]=='_' || unicode.IsLetter(rune(b[i])) {
            i++
        } else {
            str = string(b[curPos:i])
            return i, str, true
        }
    }

    return i, "", false

}

