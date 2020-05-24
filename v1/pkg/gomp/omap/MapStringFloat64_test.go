package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringFloat64_Put(t *testing.T) {
	Convey("TestMapStringFloat64.Put", t, func() {
		var k string = "hello"
		var v float64 = 0.607

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat64_Delete(t *testing.T) {
	Convey("TestMapStringFloat64.Delete", t, func() {
		var k string = "hello"
		var v float64 = 0.975

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat64_Has(t *testing.T) {
	Convey("TestMapStringFloat64.Has", t, func() {
		var k string = "hello"
		var v float64 = 0.079

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
