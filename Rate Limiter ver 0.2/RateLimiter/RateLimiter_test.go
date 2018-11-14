package ratelimiter1

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	var i uint32

	_, ok1 := NewLimiter(5, 1*time.Second, 0)
	if ok1 == nil {
		t.Error("Co loi khoi tao")
	}

	limiter, ok2 := NewLimiter(5, 1*time.Second, 100)
	if ok2 != nil {
		t.Error("Co loi khoi tao")
	}
	for i = 0; i < 101; i++ {
		limiter.Act(i)
		if i == 100 {
			fmt.Println(limiter.ListUser.FirstNode.Key, " : ", limiter.ListUser.LastNode.Key)
		}
	}
	if (limiter.ListUser.FirstNode.Key.(uint32) != 100) || (limiter.ListUser.LastNode.Key.(uint32) != 1) {
		t.Error("Co loi khoi tao them phan tu")
	}

	/*
		ok = limiter.Act(1)
		ok = limiter.Act(2)
		ok = limiter.Act(3)
		ok = limiter.Act(4)
		ok = limiter.Act(5)
		ok = limiter.Act(6)
		/*
			if ok {
				fmt.Println("asdsa")
			} else {
				fmt.Println("FLAEW")
			}
	*/
	/*
		for i := 0; i < 50; i++ {
			time.Sleep(100 * time.Millisecond)
			ok := limiter.Act(0)
			if ok {
				fmt.Println("TRUE ", limiter.ListUser.FirstNode.Key)
			}
		}*/

}

func BenchmarkLimiter(b *testing.B) {
	var i uint32
	var mutex sync.Mutex
	tmp := 0
	var wg sync.WaitGroup

	limiter, _ := NewLimiter(5, 1*time.Second, 20000)

	wg.Add(20001)
	for i = 0; i < 20001; i++ {

		go func(i uint32) {
			for j := 0; j < 10; j++ {
				if limiter.Act(i) {
					mutex.Lock()
					tmp++
					mutex.Unlock()
				}
			}
			wg.Done()
			//			fmt.Println("ID====", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("So luong req accept thanh cong ", tmp)
}
