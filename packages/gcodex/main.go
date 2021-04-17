package main

import (
    "github.com/darjun/errcode"
)

func main() {
    code := errcode.ERR_CODE_INVALID_PARAMS
    fmt.Println(code, errcode.GetDescription(errCode))
    
    // output:
    // 输出: 1 无效参数
}
