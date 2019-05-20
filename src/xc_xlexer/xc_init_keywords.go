package xc_xlexer

import "xc_datadef"

func InitKeywordsMap(m map[string]int) {

    m["char"] = xc_datadef.KW_CHAR
    m["short"] = xc_datadef.KW_SHORT
    m["int"] = xc_datadef.KW_INT
    m["void"] = xc_datadef.KW_VOID
    m["struct"] = xc_datadef.KW_STRUCT
    m["if"] = xc_datadef.KW_IF
    m["else"] = xc_datadef.KW_ELSE
    m["for"] = xc_datadef.KW_FOR
    m["continue"] = xc_datadef.KW_CONTINUE
    m["break"] = xc_datadef.KW_BREAK
    m["return"] = xc_datadef.KW_RETURN
    m["sizeof"] = xc_datadef.KW_SIZEOF
    m["__align"] = xc_datadef.KW_ALIGN
    m["__cdecl"] = xc_datadef.KW_CDECL
    m["__stdcall"] = xc_datadef.KW_STDCALL

}

