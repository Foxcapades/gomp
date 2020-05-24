package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt_Put(t *testing.T) {
	Convey("TestMapAnyInt.Put", t, func() {
		var k interface{} = "goodbye"
		var v int = 1139424147

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt_Delete(t *testing.T) {
	Convey("TestMapAnyInt.Delete", t, func() {
		var k interface{} = "goodbye"
		var v int = 544474078

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt_Has(t *testing.T) {
	Convey("TestMapAnyInt.Has", t, func() {
		var k interface{} = "goodbye"
		var v int = 605764324

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
