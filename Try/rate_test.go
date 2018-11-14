package maintry

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	var (
		i       uint32
		tmpuser User
		ok      bool
	)
	ListUser := NewUserMap()
	limiter := NewDataLimiter(5, 1*time.Second)

	for i = 0; i < 10000000; i++ {
		tmpuser = NewUser()
		ListUser.Store(i, tmpuser)
	}
	tmp := 0

	for i = 0; i < 1000; i++ {

		go func(i uint32) {
			for j := 0; j < 10; j++ {
				if limiter.Act(i) {
					//		fmt.Println("true")
					mutex.Lock()
					tmp++
					mutex.Unlock()
				}

			}
			//			fmt.Println("ID====", i)
		}(i)
	}

	fmt.Println("TMP ", tmp, ok)
}
