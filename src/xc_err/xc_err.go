package xc_err

import "os"
import "fmt"

func CheckErr(err error){
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(-1)
    }
}

