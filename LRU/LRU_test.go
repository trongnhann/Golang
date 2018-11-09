package lrucache

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	LastTime uint64 //8b
	Counter  uint8  //1b
}

func NewUser() User {
	x := time.Now().Unix()
	user := User{LastTime: uint64(x), Counter: 1}
	return user
}

func TestNewLRUCache(t *testing.T) {
	_, err := NewLRUCache(0)
	if err == nil {
		t.Error("Co loi NewLRU")
	}
}

func TestHandleLRUCache(t *testing.T) {
	lruCache, err := NewLRUCache(4)
	if err != nil {
		t.Error("Co loi NewLRU")
	}
	var user User
	user = NewUser()
	fmt.Println("===============TEST ADDFIRST=================")

	lruCache.AddFirst("user1", user)
	x := lruCache.FirstNode
	y := lruCache.LastNode
	fmt.Println(" TEST ADDFIRST ======= LRU, first, last ", lruCache, x.Key, y.Key)
	fmt.Println()
	if (x.Key != "user1") || (y.Key != "user1") || (lruCache.Map.Len()) != 1 {
		t.Error("Loi Addfirst")
	}

	user = NewUser()
	lruCache.AddFirst("user2", user)
	user = NewUser()
	lruCache.AddFirst("user3", user)
	user = NewUser()
	lruCache.AddFirst("user4", user)
	user = NewUser()
	lruCache.AddFirst("user5", user)
	x = lruCache.FirstNode
	y = lruCache.LastNode

	fmt.Println(" TEST ADDFIRST ======= LRU, first, last ", lruCache, x.Key, y.Key)
	fmt.Println()

	if (x.Key != "user5") || (y.Key != "user2") || (lruCache.Map.Len()) != 4 {
		t.Error("Loi Addfirst")
	}

	fmt.Println("===============TEST GET=================")

	val1, err1 := lruCache.Get("user1")
	val2, err2 := lruCache.Get("user2")

	fmt.Println("TEST GET ======= val user 1, val user 2 ", val1, val2)
	fmt.Println()

	if (err1 == true) || (err2 == false) {
		t.Error("Loi GET khong lay duoc gia tri")
	}

	x = lruCache.FirstNode
	y = lruCache.LastNode
	fmt.Println("TEST GET LRU, first, last ", lruCache, x.Key, y.Key)
	if (x.Key != "user2") || (y.Key != "user3") {
		t.Error("Loi GET khong set lai thanh first node")
	}

	fmt.Println("===============TEST SET=================")
	user = NewUser()
	lruCache.Set("user1", user)
	user = NewUser()
	lruCache.Set("user2", user)
	user = NewUser()
	lruCache.Set("user3", user)
	user = NewUser()
	lruCache.Set("user4", user)
	user = NewUser()
	lruCache.Set("user5", user)

	x = lruCache.FirstNode
	y = lruCache.LastNode
	fmt.Println("TEST SET ======= LRU, first, last ", lruCache, x.Key, y.Key)
	fmt.Println()

	if (x.Key != "user5") || (y.Key != "user2") {
		t.Error("Loi SET")
	}

	fmt.Println("===============TEST REMOVE=================")
	lruCache.Remove("user2")
	lruCache.Remove("user5")

	x = lruCache.FirstNode
	y = lruCache.LastNode

	fmt.Println("TEST SET ======= LRU, first, last ", lruCache, x.Key, y.Key)
	fmt.Println()

	if (x.Key != "user4") || (y.Key != "user3") || (lruCache.Map.Len() != 2) {
		t.Error("Loi REMOVE")
	}

}
