package ratelimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	_, ok1 := NewLimiter(5, 1*time.Second, 0)
	if ok1 == nil {
		t.Error("Co loi khoi tao")
	}

	limiter, ok2 := NewLimiter(5, 1*time.Second, 100)
	if ok2 != nil {
		t.Error("Co loi khoi tao")
	}
	var i uint32
	//	var mutex sync.Mutex
	//	tmp := 0
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
		for i = 0; i < 5500000; i++ {

			go func(x uint32) {
				for j := 0; j < 10; j++ {
					if limiter.Act(x) {
						mutex.Lock()
						tmp++
						mutex.Unlock()
					}
				}
				//			fmt.Println("ID====", i)
			}(i)
		}
		fmt.Println("TMP ", tmp)

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
	/*	for i := 0; i <= 5; i++ {
			ok = limiter.Act(i)
			if ok {
				fmt.Println(limiter.ListUser.Get(i))
			} else {
				fmt.Println("FALSE")
			}
		}
	*/
	/*	for i := 0; i < 100; i++ {
			time.Sleep(100 * time.Millisecond)
			ok = limiter.Act(0)
			if ok {
			}
		}
	*/
}
