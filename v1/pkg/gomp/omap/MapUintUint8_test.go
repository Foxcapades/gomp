package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint8_Put(t *testing.T) {
	Convey("TestMapUintUint8.Put", t, func() {
		var k uint = 3537966976
		var v uint8 = 247

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint8_Delete(t *testing.T) {
	Convey("TestMapUintUint8.Delete", t, func() {
		var k uint = 877327507
		var v uint8 = 78

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint8_Has(t *testing.T) {
	Convey("TestMapUintUint8.Has", t, func() {
		var k uint = 1897025958
		var v uint8 = 84

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3740103647+2892154294), ShouldBeFalse)
	})
}

func TestMapUintUint8_Get(t *testing.T) {
	Convey("TestMapUintUint8.Get", t, func() {
		var k uint = 2819211685
		var v uint8 = 32

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(4118406663 + 2543395625)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint8_GetOpt(t *testing.T) {
	Convey("TestMapUintUint8.GetOpt", t, func() {
		var k uint = 871160621
		var v uint8 = 127

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2973521921 + 3745619884)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint8_ForEach(t *testing.T) {
	Convey("TestMapUintUint8.ForEach", t, func() {
		var k uint = 2645389846
		var v uint8 = 221
		hits := 0

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint8.MarshalYAML", t, func() {
		var k uint = 464824329
		var v uint8 = 9

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint8_ToYAML(t *testing.T) {
	Convey("TestMapUintUint8.ToYAML", t, func() {
		var k uint = 1747224352
		var v uint8 = 106

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint8.PutIfNotNil", t, func() {
		var k uint = 583276351
		var v uint8 = 35

		test := omap.NewMapUintUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3067035677, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 47
		So(test.PutIfNotNil(739569449, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint8.ReplaceIfExists", t, func() {
		var k uint = 3385734130
		var v uint8 = 215
		var x uint8 = 79

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3689675594, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint8.ReplaceOrPut", t, func() {
		var k uint = 1117320678
		var v uint8 = 183
		var x uint8 = 2

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1717366330, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint8.MarshalJSON", t, func() {
		var k uint = 3954953698
		var v uint8 = 203

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3954953698,"value":203}]`)
	})
}
