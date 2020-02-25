package TestMongoDB

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestConfLoad(t *testing.T) {
	b, _ := json.Marshal(Conf)
	fmt.Println("conf=>", string(b))
	t.Log(*Conf)
}
