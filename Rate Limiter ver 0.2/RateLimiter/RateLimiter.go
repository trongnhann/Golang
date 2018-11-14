package ratelimiter1

import (
	"errors"
	"sync"
	"time"

	lru "github.com/trongnhann/learning_go/LRU1"
)

// User is user's info
type User struct {
	LastTime uint64
	Counter  uint8
}

// NewUser create new user
func NewUser() User {
	x := time.Now().Unix()
	user := User{LastTime: uint64(x), Counter: 1}
	return user
}

// DataLimiter is data of limiter
type DataLimiter struct {
	Token    uint8
	Timer    uint64
	ListUser *lru.LRUCache
}

// Limiter is interface handle
type Limiter interface {
	Act(ID interface{})
}

// NewLimiter is buffer new limiter
func NewLimiter(token uint8, timer time.Duration, nUser uint32) (DataLimiter, error) {
	LRUCache, err := lru.NewLRUCache(nUser)
	if err != nil {
		return DataLimiter{Token: 0, Timer: 0, ListUser: LRUCache}, errors.New("nUser must be bigger than 0")
	}
	tmp := uint64(timer.Seconds())
	return DataLimiter{Token: token, Timer: tmp, ListUser: LRUCache}, nil
}

// Act return true or false of this request for this user
func (limiter *DataLimiter) Act(key interface{}) bool {
	now := time.Now().Unix()
	mutex.Lock()
	tmp, ok := limiter.ListUser.Get(key)
	mutex.Unlock()
	//	fmt.Println("TMP", tmp)

	// User chua khoi tao
	if ok == false {
		user := NewUser()
		mutex.Lock()
		limiter.ListUser.Set(key, user)
		mutex.Unlock()

		//		fmt.Println("0FIRST NODE LRU CACHE : ", limiter.ListUser.FirstNode.Key)
		return true
	} else {
		user := tmp.(User)
		//		fmt.Println("Node", user)
		// Qua moc thoi gian khac thi reset
		if uint64(user.LastTime+limiter.Timer) < uint64(now) {
			user.Counter = 1
			user.LastTime = uint64(now)
			mutex.Lock()
			limiter.ListUser.AddFirst(key, user)
			mutex.Unlock()
			//			fmt.Println("NEW FIRST NODE LRU CACHE : ", limiter.ListUser.FirstNode.Key)
			return true
		} else {
			// Neu vuot qua so luong token thi reject
			if user.Counter >= limiter.Token {
				//				limiter.ListUser.AddFirst(key, user) //Khong Add thi bi loi
				//				fmt.Println("FALSE FIRST NODE LRU CACHE : ", limiter.ListUser.FirstNode.Key)
				return false
			} else {
				// Nguoc lai accept va tang counter
				user.Counter++
				mutex.Lock()
				limiter.ListUser.AddFirst(key, user)
				mutex.Unlock()
				//				fmt.Println("TRUE FIRST NODE LRU CACHE : ", limiter.ListUser.FirstNode.Key)
				return true
			}
		}
	}
}

var mutex sync.Mutex
