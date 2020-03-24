package testdemo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPoolFunc(t *testing.T) {
	PoolFunc()
	t.Logf("test Add succ")
}

func TestPoolFunc01(t *testing.T) {
	PoolFunc01()
	t.Logf("test Add succ")
}
