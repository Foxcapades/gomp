package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintAny_Put(t *testing.T) {
	Convey("TestMapUintAny.Put", t, func() {
		var k uint = 3306648463
		var v interface{} = "goodbye"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintAny_Delete(t *testing.T) {
	Convey("TestMapUintAny.Delete", t, func() {
		var k uint = 3397857443
		var v interface{} = "goodbye"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintAny_Has(t *testing.T) {
	Convey("TestMapUintAny.Has", t, func() {
		var k uint = 1126921711
		var v interface{} = "goodbye"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
