package maintry

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
func TestLimiter(t *testing.T) {
	var (
		i       uint32
		tmpuser User
		ok      bool
	)
	ListUser := NewUserMap()
	limiter := NewDataLimiter(5, 1*time.Second)

	for i = 0; i < 1000; i++ {
		tmpuser = NewUser()
		ListUser.Store(i, tmpuser)
	}

}
*/
func BenchmarkRateLimiter(b *testing.B) {
	var (
		i       uint32
		tmpuser User
	)
	ListUser := NewUserMap()
	limiter := NewDataLimiter(5, 1*time.Second)

	for i = 0; i < 1000000; i++ {
		tmpuser = NewUser()
		ListUser.Store(i, tmpuser)
	}

	tmp := 0

	var wg sync.WaitGroup
	wg.Add(1000000)
	for i = 0; i < 1000000; i++ {

		go func(i uint32) {
			for j := 0; j < 10; j++ {
				if limiter.Act(i) {
					//		fmt.Println("true")
					mutex.Lock()
					tmp++
					//				fmt.Println(i)
					mutex.Unlock()
				}
			}
			wg.Done()
			//			fmt.Println("ID====", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("So luong accept cua 1m user * 10 req ", tmp)
}
