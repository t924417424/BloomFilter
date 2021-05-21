package bloomfilter

const (
	// （1024 * 8）bit
	defaultSize = 1 << 10
	// 10
	defaultHashLoop = 10
)

type Config struct {
	// BitMap大小 = (size * 8) bit
	Size int
	// 数据hash计算次数
	HashLoop int
}
