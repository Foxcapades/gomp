package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt64_Put(t *testing.T) {
	Convey("TestMapUintInt64.Put", t, func() {
		var k uint = 2885484039
		var v int64 = 3691085478146778751

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt64_Delete(t *testing.T) {
	Convey("TestMapUintInt64.Delete", t, func() {
		var k uint = 3866652314
		var v int64 = 8761126201432260190

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt64_Has(t *testing.T) {
	Convey("TestMapUintInt64.Has", t, func() {
		var k uint = 1371517350
		var v int64 = 4606018198686923411

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
