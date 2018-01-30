package main

import (
	"fmt"
	"path"
	"strings"
	"bytes"
	"strconv"
)

var p = "/Users/zhangrongxiang/WorkSpace/goProjects/src/awesomeGo/basic/string.go"

func MyBase(p string) string {
	for index := len(p) - 1; index > 0; index-- {
		if p[index] == '/' {
			return p[index+1:]
		}
	}
	return ""
}
func MyBase2(p string) string {
	return p[strings.LastIndex(p, "/")+1:]
}
func Comma(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	return Comma(str[:n-3]) + "," + str[n-3:]
}
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(path.Base(p))
	fmt.Println(MyBase(p))
	fmt.Println(MyBase2(p))
	fmt.Println(Comma("1234567890"))
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"
	fmt.Printf("%+9s\n",strconv.FormatInt(int64(x), 2)) // "1111011"
	fmt.Printf("%+8s\n",strconv.FormatInt(int64(-x), 2)) // "-1111011"


}
