package testdemo

import (
	"fmt"
	"testing"
)

func TestMathFunc(t *testing.T) {
	f := int64(5800000.0 * 30.0 / 100.0)
	fmt.Println(f)
	f += 50
	fmt.Println(f)
	f /= 100
	fmt.Println(f)
	f *= 100
	fmt.Println(f)
	t.Logf("test ChannelFunc succ")
}
