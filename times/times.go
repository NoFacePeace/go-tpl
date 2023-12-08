package times

import "time"

const (
	LayoutDate = "20060102"
)

// 秒级清零
func ZeroOutInSecond(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

// 分钟级相等
func EqualInMin(t1, t2 time.Time) bool {
	return ZeroOutInSecond(t1).Equal(ZeroOutInSecond(t2))
}

// 分钟级早于
func BeforeInMin(t1, t2 time.Time) bool {
	return ZeroOutInSecond(t1).Before(ZeroOutInSecond(t2))
}
