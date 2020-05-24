package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint16_Put(t *testing.T) {
	Convey("TestMapIntUint16.Put", t, func() {
		var k int = 709131444
		var v uint16 = 47108

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint16_Delete(t *testing.T) {
	Convey("TestMapIntUint16.Delete", t, func() {
		var k int = 811772145
		var v uint16 = 58381

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint16_Has(t *testing.T) {
	Convey("TestMapIntUint16.Has", t, func() {
		var k int = 1997086137
		var v uint16 = 27647

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(380532107+1834250747), ShouldBeFalse)
	})
}

func TestMapIntUint16_Get(t *testing.T) {
	Convey("TestMapIntUint16.Get", t, func() {
		var k int = 925933748
		var v uint16 = 64628

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(640711528 + 1273120493)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint16_GetOpt(t *testing.T) {
	Convey("TestMapIntUint16.GetOpt", t, func() {
		var k int = 1582848175
		var v uint16 = 9213

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1248808281 + 377495901)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint16_ForEach(t *testing.T) {
	Convey("TestMapIntUint16.ForEach", t, func() {
		var k int = 2093703734
		var v uint16 = 28279
		hits := 0

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint16.MarshalYAML", t, func() {
		var k int = 1926513334
		var v uint16 = 27828

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint16_ToYAML(t *testing.T) {
	Convey("TestMapIntUint16.ToYAML", t, func() {
		var k int = 1774450654
		var v uint16 = 58533

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint16.PutIfNotNil", t, func() {
		var k int = 829049014
		var v uint16 = 62305

		test := omap.NewMapIntUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(862809618, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 51284
		So(test.PutIfNotNil(637165491, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint16.ReplaceIfExists", t, func() {
		var k int = 1932322832
		var v uint16 = 6939
		var x uint16 = 2501

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1890267502, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint16.ReplaceOrPut", t, func() {
		var k int = 1440545499
		var v uint16 = 37824
		var x uint16 = 15895

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(962802312, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint16.MarshalJSON", t, func() {
		var k int = 755998387
		var v uint16 = 50462

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":755998387,"value":50462}]`)
	})
}
