package sdarray

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"sort"
	"testing"
)

func TestSDArray(t *testing.T) {
	Convey("When an empty sdarary is created", t, func() {
		vals := make([]uint64, 0)
		sdarray := New(vals)
		So(sdarray.Num(), ShouldEqual, 0)
	})
	Convey("When a sdarray is created", t, func() {
		num := uint64(1000)
		vals := make([]uint64, num)
		for i := 0; i < len(vals); i++ {
			vals[i] = uint64(rand.Int63n(10000))
		}
		sdarray := New(vals)
		So(sdarray.Num(), ShouldEqual, num)
		sort.Sort(uint64slice(vals))
		for i := uint64(0); i < num; i++ {
			So(sdarray.Lookup(i), ShouldEqual, vals[i])
		}
	})
}
