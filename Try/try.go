package maintry

import (
	"sync"
	"time"
)

//======================================================
//======================================================

//User's info
//======================================================
type User struct {
	LastTime uint64 //8b
	Counter  uint8  //1b
}

func NewUser() User {
	x := time.Now().Unix()
	user := User{LastTime: uint64(x), Counter: 1}
	return user
}

//Dinh nghia 1 map[uint32]User cho phep Concurrent
//======================================================
type UserMap struct {
	sync.RWMutex
	List map[uint32]User
}

func NewUserMap() *UserMap {
	return &UserMap{
		List: make(map[uint32]User),
	}
}

func (rm *UserMap) Load(key uint32) (value User, ok bool) {
	rm.RLock()
	defer rm.RUnlock()
	value, ok = rm.List[key]
	return
}

func (rm *UserMap) Delete(key uint32) {
	rm.Lock()
	defer rm.Unlock()
	delete(rm.List, key)
}

func (rm *UserMap) Store(key uint32, value User) {
	rm.Lock()
	defer rm.Unlock()
	rm.List[key] = value
}

//======================================================
//Data 1 limiter
//======================================================
type DataLimiter struct {
	Token    uint8
	Timer    uint64
	ListUser *UserMap
}

type Limiter interface {
	Act(ID uint32)
}

func NewDataLimiter(token uint8, timer time.Duration) DataLimiter {
	tmp := uint64(timer.Seconds())
	listUser := NewUserMap()
	return DataLimiter{Token: token, Timer: tmp, ListUser: listUser}
}

func (limiter *DataLimiter) Act(key uint32) bool {
	now := time.Now().Unix()
	user, ok := limiter.ListUser.Load(key)
	//	fmt.Println(user)
	if ok {
		if uint64(user.LastTime+limiter.Timer) < uint64(now) {
			user.Counter = 1
			user.LastTime = uint64(now)
			//			fmt.Println("new")
			limiter.ListUser.Store(key, user)
			return true
		} else {
			if user.Counter >= limiter.Token {
				//				fmt.Println("FALSE")
				return false
			} else {
				user.Counter++
				limiter.ListUser.Store(key, user)
				return true
			}
		}
	} else {
		user = NewUser()
		limiter.ListUser.Store(key, user)
		return true
	}
}

//======================================================

//======================================================

var (
	mutex sync.Mutex
)

/*
func main() {
	//Khoi tao 1 ListUser
	var (
		i       uint32
		tmpuser User
		ok      bool
	) /*
		limiter := NewDataLimiter(5, 1*time.Second)

			for i = 0; i < 10000000; i++ {
				ok = limiter.Act(i)
						if ok {
							fmt.Println(limiter.ListUser.Load(i))
						}
			}

				for i = 0; i < 100; i++ {
					time.Sleep(100 * time.Millisecond)
					tmpuser, ok = limiter.ListUser.Load(1)
					if ok {
						if tmpuser.Counter == 1 {
							fmt.Println("RESET time: ", tmpuser.LastTime)
						}
						if limiter.Act(1) {
							fmt.Println("ACCEPT REQ : ", tmpuser.Counter)
						} else {
							fmt.Println("REJECT REQ : ", tmpuser.Counter)
						}
					} else {
						fmt.Println("MAIN ERROR READ DATA")
					}
				}
*/
/*
	ListUser := NewUserMap()
	limiter := NewDataLimiter(5, 1*time.Second)

	for i = 0; i < 10000000; i++ {
		tmpuser = NewUser()
		ListUser.Store(i, tmpuser)
	}
	/*
		x(0, ListUser)
		val, ok := ListUser.Load(0)
		fmt.Println("==============", val)

		for i = 0; i < 100; i++ {
			time.Sleep(100 * time.Millisecond)
			tmp_user, ok = ListUser.Load(1)
			if ok {
				if tmp_user.Counter == 1 {
					fmt.Println("===time: ", tmp_user.LastTime)
				}
				if Act(ListUser, 1, limiter) {
					fmt.Println("TRUE : ", tmp_user.Counter)
				}
			} else {
				fmt.Println("MAIN ERROR READ DATA")
			}
		}
*/
/*	tmp := 0

	for i = 0; i < 5000000; i++ {

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
*/
