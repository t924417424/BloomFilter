package bloomfilter

import (
	"fmt"
	"sync"
)

type bitmap struct {
	sync.RWMutex
	data []byte
}

// 二进制打印bitmap
func (b *bitmap) debug() {
	b.RLock()
	defer b.RUnlock()
	for i := range b.data {
		fmt.Printf("%08b", b.data[i])
	}
}

// 将数据插入bitmap
func (b *bitmap) insert(indexs []int) {
	b.Lock()
	for i := range indexs {
		b.data[indexs[i]] |= 1 << (indexs[i] % 8)
	}
	b.Unlock()
}

func (b *bitmap) contains(indexs []int) bool {
	b.RLock()
	for i := range indexs {
		match := 1 << (indexs[i] % 8)
		if b.data[indexs[i]]&(1<<(indexs[i]%8)) != byte(match) {
			return false
		}
	}
	b.RUnlock()
	return true
}
