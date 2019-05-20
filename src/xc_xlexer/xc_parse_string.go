package xc_xlexer

func ParseString(b []byte, curPos int)(int, string, bool){

    str := ""

    i := curPos+1
    for ; i<len(b); {
        if b[i] != '"' {
            i++
        } else {
            str = string(b[curPos+1:i])
            i++
            return i, str, true
        }
    }

    return i, "", false

}

