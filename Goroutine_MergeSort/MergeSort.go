//this code copy in https://medium.com/@_orcaman/when-too-much-concurrency-slows-you-down-golang-9c144ca305a

package msort

import (
	"sync"
)

func merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

// MergeSort sorts the slice s using Merge Sort Algorithm
func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	n := len(s) / 2

	var l []int
	var r []int

	l = MergeSort(s[:n])
	r = MergeSort(s[n:])

	return merge(l, r)
}

// MergeSortMulti sorts the slice s using Merge Sort Algorithm with firing goroutines like crazy
func MergeSortMulti(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	n := len(s) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var l []int
	var r []int

	go func() {
		l = MergeSortMulti(s[:n])
		wg.Done()
	}()

	go func() {
		r = MergeSortMulti(s[n:])
		wg.Done()
	}()

	wg.Wait()
	return merge(l, r)
}

var sem = make(chan struct{}, 32)

// MergeSortMultiChan sorts the slice s using Merge Sort Algorithm with a number goroutine is limited by channel
func MergeSortMultiChan(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	n := len(s) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var l []int
	var r []int

	select {
	case sem <- struct{}{}:
		go func() {
			l = MergeSortMultiChan(s[:n])
			<-sem
			wg.Done()
		}()
	default:
		l = MergeSort(s[:n])
		wg.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			r = MergeSortMultiChan(s[n:])
			<-sem
			wg.Done()
		}()
	default:
		r = MergeSort(s[n:])
		wg.Done()
	}

	wg.Wait()
	return merge(l, r)
}

func x() {

}
