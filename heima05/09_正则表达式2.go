// 09_正则表达式2
package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Hello World!")
	buf := "43.33 456 sgasd 1.23 3. .2 dd "
	//解析正则表达式
	reg := regexp.MustCompile(`\d+\.\d+`)

	//提取关键字
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println(result)
}
