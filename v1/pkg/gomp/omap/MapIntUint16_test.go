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
		var k int = 577578678
		var v uint16 = 14830

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint16_Delete(t *testing.T) {
	Convey("TestMapIntUint16.Delete", t, func() {
		var k int = 143908775
		var v uint16 = 1323

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint16_Has(t *testing.T) {
	Convey("TestMapIntUint16.Has", t, func() {
		var k int = 1589369104
		var v uint16 = 23044

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(321589550+1402578678), ShouldBeFalse)
	})
}


func TestMapIntUint16_Get(t *testing.T) {
	Convey("TestMapIntUint16.Get", t, func() {
		var k int = 1289451052
		var v uint16 = 62363

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(747860227 + 267796506)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint16_GetOpt(t *testing.T) {
	Convey("TestMapIntUint16.GetOpt", t, func() {
		var k int = 156075449
		var v uint16 = 4183

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(604555718 + 1290293983)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint16_ForEach(t *testing.T) {
	Convey("TestMapIntUint16.ForEach", t, func() {
		var k int = 2073342938
		var v uint16 = 16687
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
		var k int = 1931456836
		var v uint16 = 56156

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
		var k int = 143279512
		var v uint16 = 37556

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
		var k int = 590363170
		var v uint16 = 55216

		test := omap.NewMapIntUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1518206718, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 50717
		So(test.PutIfNotNil(781436982, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint16.ReplaceIfExists", t, func() {
		var k int = 45542709
		var v uint16 = 26628
		var x uint16 = 30141

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2022922008, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint16.ReplaceOrPut", t, func() {
		var k int = 1385185235
		var v uint16 = 25182
		var x uint16 = 12920

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1994359522, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint16.MarshalJSON", t, func() {
		var k int = 613361228
		var v uint16 = 25104

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":613361228,"value":25104}]`)
	})
}

