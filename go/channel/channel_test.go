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

func Test2(t *testing.T) {
	var c = make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v) //这里会依次打印123，456
		}
	}()
	c <- 123
	c <- 456
	close(c)
	time.Sleep(time.Second * 2)
}

// 一个 goroutine 顺序发送 0,1,2,3,4 个数字
func Test3(t *testing.T) {

	var data = []int{0, 1, 2, 3, 4}
	var c = make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		for {
			select {
			case v, ok := <-c:
				if !ok {
					fmt.Println("接收结束")
					return
				}
				fmt.Println(v)

			}
		}

	}(c)

	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		for index := range data {
			c <- data[index]
		}
		close(c)
		fmt.Println("发送结束")
	}(c)

	wg.Wait()

}
