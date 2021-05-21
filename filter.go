package bloomfilter

import (
	"hash/crc32"
	"math/rand"
	"time"
)

type filter struct {
	size     int
	hashLoop int
	seed     []int
	*bitmap
}

// new一个布隆过滤器
func NewBloomfilter(c Config) *filter {
	if c.Size <= 0 {
		c.Size = defaultSize
	}
	if c.HashLoop <= 0 {
		c.HashLoop = defaultHashLoop
	}
	bitmap := new(bitmap)
	bitmap.data = make([]byte, c.Size)
	seed := generateSeed(int(c.HashLoop), int(c.Size))
	return &filter{size: c.Size, hashLoop: c.HashLoop, seed: seed, bitmap: bitmap}
}

// 生成随机种子
func generateSeed(s int, size int) []int {
	result := make([]int, s)
	rand.Seed(time.Now().Unix() / 1000)
	for i := 0; i < s; i++ {
		result[i] = rand.Intn(size * 8)
		//println(result[i])
	}
	return result
}

// 保存一个key到布隆过滤器中
func (f *filter) Insert(key []byte) {
	indexs := f.hash(key)
	f.bitmap.insert(indexs)
}

// 检测key是否已经存在
func (f *filter) Contains(key []byte) bool {
	indexs := f.hash(key)
	return f.bitmap.contains(indexs)
}

// 根据key生成bitmap中到索引
func (f *filter) hash(key []byte) []int {
	indexs := make([]int, f.hashLoop)
	for i := 0; i < f.hashLoop; i++ {
		indexs[i] = (int(crc32.ChecksumIEEE(key)) + f.seed[i]) % f.size
	}
	return indexs
}

// 二进制形式打印布隆过滤器源数据
func (f *filter) Debug() {
	f.bitmap.debug()
}
