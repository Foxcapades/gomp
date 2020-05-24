package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt32_Put(t *testing.T) {
	Convey("TestMapIntInt32.Put", t, func() {
		var k int = 858958590
		var v int32 = 1069163632

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt32_Delete(t *testing.T) {
	Convey("TestMapIntInt32.Delete", t, func() {
		var k int = 1297033098
		var v int32 = 879648553

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt32_Has(t *testing.T) {
	Convey("TestMapIntInt32.Has", t, func() {
		var k int = 63718591
		var v int32 = 4088582

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
