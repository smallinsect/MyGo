package testdemo

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFunc(t *testing.T) {
	// TimeFunc01()
	fmt.Println(time.Unix(1588176000, 0).Format("2006-01-02 15:04:05"))
	t.Logf("test Add succ")
}
