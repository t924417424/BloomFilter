# Go语言BloomFilter包（布隆过滤器）

## 安装使用
`go get github.com/t924417424/BloomFilter`
```go
package main

import (
	"github.com/t924417424/BloomFilter"
)

func main() {
	filter := bloomfilter.NewBloomfilter(bloomfilter.Config{HashLoop: 30})
	filter.Insert([]byte("123"))
	println(filter.Contains([]byte("123")))
	filter.Debug()
}
```


## 测试用例
```go
myFilter := NewBloomfilter(Config{HashLoop: 20})
key      := []byte("123")
key2     := []byte("456")
key3     := []byte("999")
myFilter.Insert(key)
myFilter.Insert(key2)
myFilter.Contains(key)
myFilter.Contains(key2)
```

## 基准测试
### Insert:
```
goos: darwin
goarch: amd64
pkg: github.com/t924417424/BloomFilter
cpu: Intel(R) Core(TM) i3-8100B CPU @ 3.60GHz
Benchmark_Insert-4   	  976336	      1075 ns/op	     160 B/op	       1 allocs/op
PASS
ok  	github.com/t924417424/BloomFilter	1.505s
```
### Contains:
```
goos: darwin
goarch: amd64
pkg: github.com/t924417424/BloomFilter
cpu: Intel(R) Core(TM) i3-8100B CPU @ 3.60GHz
Benchmark_Contains-4   	 1000000	      1054 ns/op	     160 B/op	       1 allocs/op
PASS
ok  	github.com/t924417424/BloomFilter	1.169s
```