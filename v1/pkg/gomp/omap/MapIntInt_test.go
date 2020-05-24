package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt_Put(t *testing.T) {
	Convey("TestMapIntInt.Put", t, func() {
		var k int = 1337298878
		var v int = 793909336

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt_Delete(t *testing.T) {
	Convey("TestMapIntInt.Delete", t, func() {
		var k int = 508572546
		var v int = 1149509107

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt_Has(t *testing.T) {
	Convey("TestMapIntInt.Has", t, func() {
		var k int = 402107940
		var v int = 512906503

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
