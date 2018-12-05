package msort

import (
	"fmt"
	"runtime"
	"testing"
)

var a []int

// Khởi tạo các giá trị cần để thực hiện thuật toán
func init() {
	for i := 0; i < 10000000; i++ {
		a = append(a, i)
	}
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
	fmt.Println(numcpu)
}

func BenchmarkMergeSortMultiChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortMultiChan(a)
	}
}

func BenchmarkMergeSortMulti(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortMulti(a)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(a)
	}
}

// Tính benchmark của Go khi: Gọi 1 hàm, gọi 1 goroutine, các phép tính cơ bản
func BenchmarkX(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x()
	}
}

func BenchmarkGor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go func() {
		}()
	}
}
func BenchmarkGan(b *testing.B) {
	a := 1
	for n := 0; n < b.N; n++ {
		a = 2 + 3*3 - 4/1 + 3*3
	}
	fmt.Println(a)
}
func BenchmarkEp(b *testing.B) {
	a := 1
	c := 0.0
	for n := 0; n < b.N; n++ {
		c = float64(a)
	}
	fmt.Println(a, c)
}
func BenchmarkSosanh(b *testing.B) {
	a := 1
	for n := 0; n < b.N; n++ {
		if a < 2 {
		}
	}
	fmt.Println(a)
}
func BenchmarkCong(b *testing.B) {
	a := 1
	for n := 0; n < b.N; n++ {
		a = a + 1
	}
	fmt.Println(a)
}
func BenchmarkTru(b *testing.B) {
	a := 1
	for n := 0; n < b.N; n++ {
		a = a - 0
	}
	fmt.Println(a)
}
func BenchmarkNhan(b *testing.B) {
	a := 1
	for n := 0; n < b.N; n++ {
		a = a * 1
	}
	fmt.Println(a)
}
func BenchmarkChia(b *testing.B) {
	a := 1
	for n := 0; n < b.N; n++ {
		a = a / 1
	}
	fmt.Println(a)
}
func BenchmarkSwap(b *testing.B) {
	a1 := 1
	b1 := 2
	c1 := 3
	for n := 0; n < b.N; n++ {
		a1 = b1
		b1 = c1
		c1 = a1
	}
}
