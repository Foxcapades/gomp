package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint_Put(t *testing.T) {
	Convey("TestMapAnyUint.Put", t, func() {
		var k interface{} = "goodbye"
		var v uint = 4006544902

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint_Delete(t *testing.T) {
	Convey("TestMapAnyUint.Delete", t, func() {
		var k interface{} = "goodbye"
		var v uint = 3186217021

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint_Has(t *testing.T) {
	Convey("TestMapAnyUint.Has", t, func() {
		var k interface{} = "goodbye"
		var v uint = 3440505210

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
