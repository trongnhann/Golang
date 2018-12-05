# Golang
src Golang

-------------------------------------------------------
Rate_Limiter_ver_0.1 consist:  
 - LRU consist:
    + LRU.go is package LRUCache for type MapGoroutine.
    + LRU_test is testing file.
 - MapGoroutine is package as map in go but it can handle cocurrent.
 - RateLimiting consist:
    + RateLimiting.go is package rate limiter which is used to control the rate of traffic. It is based on "fixed window counter algorithm".
    + RateLimiting_test.go is testing.

-------------------------------------------------------
Rate_Limiter_ver_0.2 consist:  New version, faster than ver_0.1 because it use normal map
 - LRU consist:
    + LRU.go is package LRUCache for type map.
    + LRU_test is testing file.
 - RateLimiting consist:
    + RateLimiting.go is package rate limiter which is used to control the rate of traffic. It is based on "fixed window counter algorithm".
    + RateLimiting_test.go is testing.

-------------------------------------------------------
MapGoroutine: Package as map in go but it can handle cocurrent.

-------------------------------------------------------
DevideK: Package devide N to K group

-------------------------------------------------------
Goroutine_Pi: Package demonstrate the ease of multi-core parallel programming
in Go and the efficiency of parallel Go code.

-------------------------------------------------------
Goroutine_MergeSort: Package demonstrate that when too much goroutine, it will slow your code

-------------------------------------------------------
