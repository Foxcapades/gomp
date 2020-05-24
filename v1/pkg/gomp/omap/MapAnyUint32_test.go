package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint32_Put(t *testing.T) {
	Convey("TestMapAnyUint32.Put", t, func() {
		var k interface{} = "goodbye"
		var v uint32 = 3960871492

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint32_Delete(t *testing.T) {
	Convey("TestMapAnyUint32.Delete", t, func() {
		var k interface{} = "goodbye"
		var v uint32 = 390143126

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint32_Has(t *testing.T) {
	Convey("TestMapAnyUint32.Has", t, func() {
		var k interface{} = "goodbye"
		var v uint32 = 2118028752

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
