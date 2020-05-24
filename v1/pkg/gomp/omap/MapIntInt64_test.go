package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt64_Put(t *testing.T) {
	Convey("TestMapIntInt64.Put", t, func() {
		var k int = 6105384
		var v int64 = 8446960703956728189

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt64_Delete(t *testing.T) {
	Convey("TestMapIntInt64.Delete", t, func() {
		var k int = 1266659267
		var v int64 = 5159484672389300587

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt64_Has(t *testing.T) {
	Convey("TestMapIntInt64.Has", t, func() {
		var k int = 1751069271
		var v int64 = 8098229102784230399

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
