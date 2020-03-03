// 29_goto的使用
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!1")

	goto End

	fmt.Println("Hello World!2")

End:
	fmt.Println("Hello World!3")

}
