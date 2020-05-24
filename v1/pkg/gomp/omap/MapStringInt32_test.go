package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt32_Put(t *testing.T) {
	Convey("TestMapStringInt32.Put", t, func() {
		var k string = "hello"
		var v int32 = 646203300

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt32_Delete(t *testing.T) {
	Convey("TestMapStringInt32.Delete", t, func() {
		var k string = "hello"
		var v int32 = 1106410694

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt32_Has(t *testing.T) {
	Convey("TestMapStringInt32.Has", t, func() {
		var k string = "hello"
		var v int32 = 1747278511

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
