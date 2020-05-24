package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint64_Put(t *testing.T) {
	Convey("TestMapStringUint64.Put", t, func() {
		var k string = "hello"
		var v uint64 = 1905388747193831650

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint64_Delete(t *testing.T) {
	Convey("TestMapStringUint64.Delete", t, func() {
		var k string = "hello"
		var v uint64 = 17204678798284737396

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint64_Has(t *testing.T) {
	Convey("TestMapStringUint64.Has", t, func() {
		var k string = "hello"
		var v uint64 = 15649472107743074779

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
