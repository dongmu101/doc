package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	count := 10
	num := 100
	wg := sync.WaitGroup{}
	c := make(chan struct{}, count)
	for i := 0; i < num; i++ {
		wg.Add(1)
		c <- struct{}{}
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
			<-c
			time.Sleep(time.Second * 10)
		}(i)
	}
	wg.Wait()
}
