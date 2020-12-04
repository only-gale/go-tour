package main

import (
	"fmt"
	"math"
)

type ByteSize struct {
	size   uint64
	amount float64
	unit   string
	scale  uint64
}

const (
	// 通过赋予空白标识符来忽略第一个值
	_       = iota // ignore first value by assigning to blank identifier
	KbScale = 1 << (10 * iota)
	MbScale
	GbScale
	TbScale
	PbScale
	EbScale
	ZbScale
	YbScale
)

var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision uint8) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func (u *ByteSize) convert(scale float64) *ByteSize {
	size := u.size
	index := math.Floor(math.Log(float64(size)) / math.Log(scale))
	u.unit = units[uint8(index)]
	u.amount = float64(size) / math.Pow(scale, index)
	return u
}

func (u *ByteSize) ToKB() *ByteSize {
	return u.convert(KbScale)
}

func (u *ByteSize) ToMB() *ByteSize {
	return u.convert(MbScale)
}

func (u *ByteSize) ToGB() *ByteSize {
	return u.convert(GbScale)
}

func (u *ByteSize) ToTB() *ByteSize {
	return u.convert(TbScale)
}

func (u *ByteSize) ToPB() *ByteSize {
	return u.convert(PbScale)
}

func (u *ByteSize) ToEB() *ByteSize {
	return u.convert(EbScale)
}

func (u *ByteSize) ToZB() *ByteSize {
	return u.convert(ZbScale)
}

func (u *ByteSize) ToYB() *ByteSize {
	return u.convert(YbScale)
}

func (u *ByteSize) Format(precision uint8) *ByteSize {
	u.amount = toFixed(u.amount, precision)
	return u
}

func (u ByteSize) String() string {
	return fmt.Sprintf("%v %v", u.amount, u.unit)
}

func main() {
	u1 := ByteSize{size: 999999999}
	fmt.Println(u1.ToKB().Format(6))
}
