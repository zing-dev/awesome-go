package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"io"
)
//Golang计算MD5

func main() {
	str := "hello world"
	/**************************/
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Println(hex.EncodeToString(cipherStr))

	/****************************/
	hash := md5.New()
	io.WriteString(hash,str)
	fmt.Println(hash.Sum(nil))
	fmt.Printf("%X\n",hash.Sum(nil))


}
