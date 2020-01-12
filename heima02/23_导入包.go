// 23_导入包
package main

//. 操作
import . "fmt" //调用函数，无需通过包名
import io "os" //给包取别名
import _ "bufio"

func main() {
	Println("Hello World!")
	Println(io.Args)
}

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("Hello World!")
// 	fmt.Println(os.Args)
// }
