package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint8_Put(t *testing.T) {
	Convey("TestMapIntUint8.Put", t, func() {
		var k int = 701554781
		var v uint8 = 8

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint8_Delete(t *testing.T) {
	Convey("TestMapIntUint8.Delete", t, func() {
		var k int = 1901686722
		var v uint8 = 171

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint8_Has(t *testing.T) {
	Convey("TestMapIntUint8.Has", t, func() {
		var k int = 1763930307
		var v uint8 = 202

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1441362015+336433335), ShouldBeFalse)
	})
}

func TestMapIntUint8_Get(t *testing.T) {
	Convey("TestMapIntUint8.Get", t, func() {
		var k int = 65376489
		var v uint8 = 146

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(805345164 + 984881466)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint8_GetOpt(t *testing.T) {
	Convey("TestMapIntUint8.GetOpt", t, func() {
		var k int = 480153789
		var v uint8 = 210

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1271401318 + 357371359)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint8_ForEach(t *testing.T) {
	Convey("TestMapIntUint8.ForEach", t, func() {
		var k int = 1568753293
		var v uint8 = 82
		hits := 0

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint8.MarshalYAML", t, func() {
		var k int = 1081832579
		var v uint8 = 250

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint8_ToYAML(t *testing.T) {
	Convey("TestMapIntUint8.ToYAML", t, func() {
		var k int = 2070465873
		var v uint8 = 79

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapIntUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint8.PutIfNotNil", t, func() {
		var k int = 1756816105
		var v uint8 = 209

		test := omap.NewMapIntUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1901089895, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 136
		So(test.PutIfNotNil(926113480, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceIfExists", t, func() {
		var k int = 721603191
		var v uint8 = 114
		var x uint8 = 136

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1747657195, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceOrPut", t, func() {
		var k int = 383971961
		var v uint8 = 40
		var x uint8 = 214

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(373745070, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint8.MarshalJSON", t, func() {
		var k int = 1609765710
		var v uint8 = 193

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1609765710,"value":193}]`)
	})
}
