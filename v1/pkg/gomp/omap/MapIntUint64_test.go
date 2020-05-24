package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint64_Put(t *testing.T) {
	Convey("TestMapIntUint64.Put", t, func() {
		var k int = 339192614
		var v uint64 = 17327455766902076695

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint64_Delete(t *testing.T) {
	Convey("TestMapIntUint64.Delete", t, func() {
		var k int = 1893396575
		var v uint64 = 2591333051498151763

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint64_Has(t *testing.T) {
	Convey("TestMapIntUint64.Has", t, func() {
		var k int = 737495749
		var v uint64 = 9321497739288818975

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(265028545+849469618), ShouldBeFalse)
	})
}

func TestMapIntUint64_Get(t *testing.T) {
	Convey("TestMapIntUint64.Get", t, func() {
		var k int = 1051213018
		var v uint64 = 2246698631846754885

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(45019723 + 1591754424)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint64_GetOpt(t *testing.T) {
	Convey("TestMapIntUint64.GetOpt", t, func() {
		var k int = 1657052989
		var v uint64 = 10665255293779040983

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1825003448 + 192165430)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint64_ForEach(t *testing.T) {
	Convey("TestMapIntUint64.ForEach", t, func() {
		var k int = 1533635739
		var v uint64 = 17622120820660090199
		hits := 0

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint64.MarshalYAML", t, func() {
		var k int = 1910876269
		var v uint64 = 1212251217294974302

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint64_ToYAML(t *testing.T) {
	Convey("TestMapIntUint64.ToYAML", t, func() {
		var k int = 811694090
		var v uint64 = 7171403890970407086

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint64.PutIfNotNil", t, func() {
		var k int = 19800933
		var v uint64 = 13854555070203296529

		test := omap.NewMapIntUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1577721323, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 13472552346844892458
		So(test.PutIfNotNil(16489266, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceIfExists", t, func() {
		var k int = 1561561361
		var v uint64 = 12406551565055682004
		var x uint64 = 8052930516617577000

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1984533390, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceOrPut", t, func() {
		var k int = 1364665455
		var v uint64 = 4985806638274162277
		var x uint64 = 8901915081728749730

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1313201528, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint64.MarshalJSON", t, func() {
		var k int = 1042966162
		var v uint64 = 9049468678114231877

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1042966162,"value":9049468678114231877}]`)
	})
}
