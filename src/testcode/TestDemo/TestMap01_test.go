package testdemo

import "testing"
import "reflect"
import "fmt"

func TestMapFunc(t *testing.T) {
	MapFunc()
	t.Logf("test Add succ")
}

//计算Map中的元素个数
func TestMapFunc01(t *testing.T) {
	// for i := 0; i < 10000; i++ {
	// 	MapFunc04()
	// }
	m := make(map[string]string)
	m["11111"] = "小昆虫"
	m["22222"] = "小白菜"
	m["33333"] = "小东风"
	fmt.Println("m=", m)
	fmt.Println("reflect.TypeOf(len(m))=", reflect.TypeOf(len(m)))
	t.Logf("test Add succ")
}
