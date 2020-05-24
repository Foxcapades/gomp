package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringRune_Put(t *testing.T) {
	Convey("TestMapStringRune.Put", t, func() {
		var k string = "hello"
		var v rune = 1168565194

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringRune_Delete(t *testing.T) {
	Convey("TestMapStringRune.Delete", t, func() {
		var k string = "hello"
		var v rune = 598090563

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringRune_Has(t *testing.T) {
	Convey("TestMapStringRune.Has", t, func() {
		var k string = "hello"
		var v rune = 908712433

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
