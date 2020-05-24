package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint64_Put(t *testing.T) {
	Convey("TestMapAnyUint64.Put", t, func() {
		var k interface{} = "goodbye"
		var v uint64 = 8549944162621642512

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint64_Delete(t *testing.T) {
	Convey("TestMapAnyUint64.Delete", t, func() {
		var k interface{} = "goodbye"
		var v uint64 = 8807817071862113702

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint64_Has(t *testing.T) {
	Convey("TestMapAnyUint64.Has", t, func() {
		var k interface{} = "goodbye"
		var v uint64 = 12432680895096110463

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
