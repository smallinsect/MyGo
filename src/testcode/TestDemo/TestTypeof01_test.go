package testdemo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeOfFunc(t *testing.T) {
	// TimeFunc01()
	v := make(map[int]string)
	fmt.Println(reflect.TypeOf(v))
	t.Logf("test Add succ")
}
