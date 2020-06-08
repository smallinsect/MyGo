/*
 * @Author: your name
 * @Date: 2020-04-29 18:09:57
 * @LastEditTime: 2020-06-01 16:52:59
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \MyGo\testcode\TestDemo\TestMap01_test.go
 */
package testdemo

import (
	"sync"
	"testing"
)
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

//计算Map中的元素个数
func TestMapFunc02(t *testing.T) {

	var m sync.Map
	m.Store("1111", &Person{Name: "aaaaa"})
	m.Store("2222", &Person{Name: "bbbbb"})
	m.Store("3333", &Person{Name: "ccccc"})

	p, _ := m.Load("1111")
	// fmt.Println(m)
	fmt.Println(p)
	m.Delete("1111")
	// fmt.Println(m)
	fmt.Println(p)

	// m := make(map[string]*Person)
	// m["1111"] = &Person{Name: "aaaaa"}
	// m["2222"] = &Person{Name: "bbbbb"}
	// m["3333"] = &Person{Name: "ccccc"}

	// p := m["1111"]
	// fmt.Println(p)
	// fmt.Println(m)
	// delete(m, "1111")
	// fmt.Println(p)
	// fmt.Println(m)
}

func TestMapFunc03(t *testing.T) {
	m := make(map[string]string, 10)
	m["1111"] = "aaaa"
	fmt.Println(len(m))
	m["1111"] = "aaaa"
	fmt.Println(len(m))
	m["2222"] = "bbbb"
	fmt.Println(len(m))
	m["3333"] = "cccc"
	fmt.Println(len(m))
}
