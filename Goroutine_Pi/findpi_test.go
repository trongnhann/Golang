package pi

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkPiNormal(b *testing.B) {
	runtime.GOMAXPROCS(4)
	x := 0.0
	n := 1562500
	h := 1.0 / float64(n)
	for nb := 0; nb < b.N; nb++ {
		x = chunknormal(0, int64(n), h)
	}
	fmt.Println(x)
}

func BenchmarkPiGoroutine(b *testing.B) {
	x := 0.0
	np := 64
	n := 1000000000
	pi := 0.0
	runtime.GOMAXPROCS(4)
	h := 1.0 / float64(n)
	c := make(chan float64, np)
	for nb := 0; nb < b.N; nb++ {
		for i := 0; i < np; i++ {
			go chunk(int64(i)*int64(n)/int64(np), (int64(i)+1)*int64(n)/int64(np), c, h)
		}
		for i := 0; i < np; i++ {
			pi += <-c
		}
		x = pi
	}
	fmt.Println(x)
}
