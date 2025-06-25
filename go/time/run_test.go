package time

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	count := 1
	for _ = range ticker.C {
		fmt.Println("执行了：", count)
		count++
		if count >= 10 {
			ticker.Stop()
			break
		}
	}
}

func Test2(t *testing.T) {

	var s = max(1, 2)
	fmt.Println("s:", s)
}
