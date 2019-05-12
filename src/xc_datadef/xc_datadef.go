package xc_datadef

const(
    TK_PLUS = iota /* +   */
    TK_MINUS       /* -   */
    TK_STAR        /* *   */
    TK_DIVIDE      /* /   */
    TK_MOD         /* %   */
    TK_EQ          /* ==  */
    TK_NEQ         /* !=  */
    TK_LT          /* <   */
    TK_LEQ         /* <=  */
    TK_GT          /* >   */
    TK_GEQ         /* >=  */
    TK_ASSIGN      /* =   */
    TK_POINTSTO    /* ->  */
    TK_DOT         /* .   */
    TK_AND         /* &   */
    TK_OPENPA      /* (   */
    TK_CLOSEPA     /* )   */
    TK_OPENBR      /* [   */
    TK_CLESEBK     /* ]   */
    TK_BEGIN       /* {   */
    TK_END         /* }   */
    TK_SEMICOLON   /* ;   */
    TK_COMMA       /* ,   */
    TK_ELLIPSIS    /* ... */
    TK_EOF         /* END OF FILE  */

    TK_CINT        /* const int    */
    TK_CCHAR       /* const char   */
    TK_CSTR        /* const string */

    KW_CHAR        /* char         */
    KW_SHORT       /* short        */
    KW_INT         /* int          */
    KW_VOID        /* void         */
    KW_STRUCT      /* struct       */
    KW_IF          /* if           */
    KW_ELSE        /* else         */
    KW_FOR         /* for          */
    KW_CONTINUE    /* continue     */
    KW_BREAK       /* break        */
    KW_RETURN      /* return       */
    KW_SIZEOF      /* sizeof       */

    KW_ALIGN       /* __align      */
    KW_CDECL       /* __cdecl      */
    KW_STDCALL     /* __stdcall    */

    TK_IDENT       /* identifier   */
)

type Token struct{
    TokenType  int
    TokenValue string
}

