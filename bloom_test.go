package bloomfilter

import (
	"testing"
)

var (
	myFilter = NewBloomfilter(Config{HashLoop: 20})
	key      = []byte("aasdasdasdasdas5641tegafDASDASDASD21dasdasdgdfgdfgdfg")
	key2     = []byte("a135as1d35as16c5a61s56df1as35d1as32d1a23s1da2s1d2a121d321")
	key3     = []byte("123")
)

func TestBloomBilter(t *testing.T) {
	myFilter.Insert(key)
	myFilter.Insert(key2)
	// 打印bitmap内存信息
	myFilter.Debug()
	if myFilter.Contains(key) == false {
		t.Fail()
	}
	if myFilter.Contains(key2) == false {
		t.Fail()
	}
	if myFilter.Contains(key3) == true {
		t.Fail()
	}
}

func Benchmark_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myFilter.Insert(key)
	}
}

func Benchmark_Contains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myFilter.Contains(key)
	}
}
