package xc_xlexer

import "unicode"

func SkipSpace(b []byte, curPos int) int {

    i := curPos
    for ; i<len(b); {
        if unicode.IsSpace(rune(b[i])) {
            i++
        }else{
            break
        }
    }

    return i

}

