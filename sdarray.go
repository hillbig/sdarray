package sdarray

import (
	"github.com/hillbig/fixvec"
	"github.com/hillbig/rsdic"
	"sort"
)

type SDArray interface {
	Num() uint64
	Lookup(ind uint64) uint64
}

type sdarrayImpl struct {
	high  rsdic.RSDic
	low   fixvec.FixVec
	width uint8
}

type uint64slice []uint64

func (s uint64slice) Len() int {
	return len(s)
}

func (s uint64slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s uint64slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func New(vals []uint64) SDArray {
	m := uint64(len(vals))
	sort.Sort(uint64slice(vals))
	if m == 0 {
		return &sdarrayImpl{
			high:  initHigh(vals, 0, 0),
			low:   initLow(vals, 0, 0),
			width: 0,
		}
	}
	n := vals[m-1]
	width := log2(n / m)

	return &sdarrayImpl{
		high:  initHigh(vals, m, width),
		low:   initLow(vals, m, width),
		width: width,
	}
}

func (sd sdarrayImpl) Num() uint64 {
	return sd.low.Num()
}

func (sd sdarrayImpl) Lookup(ind uint64) uint64 {
	return ((sd.high.Select(ind, true) - ind) << sd.width) + sd.low.Get(ind)
}

func initHigh(vals []uint64, m uint64, width uint8) rsdic.RSDic {
	high := rsdic.New()
	prevPos := uint64(0)
	for i := uint64(0); i < m; i++ {
		pos := (vals[i] >> width) + i
		for j := prevPos; j < pos; j++ {
			high.PushBack(false)
		}
		high.PushBack(true)
		prevPos = pos + 1
	}
	return high
}

func initLow(vals []uint64, m uint64, width uint8) fixvec.FixVec {
	low := fixvec.New(m, width)
	if width == 0 {
		return low
	}
	mask := (uint64(1) << width) - 1
	for i := uint64(0); i < m; i++ {
		low.Set(i, vals[i]&mask)
	}
	return low
}

func log2(x uint64) uint8 {
	r := uint8(0)
	for (x >> r) > 0 {
		r++
	}
	return r
}
