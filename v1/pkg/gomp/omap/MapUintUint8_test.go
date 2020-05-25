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
		var k uint = 2520158894
		var v uint8 = 212

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint8_Delete(t *testing.T) {
	Convey("TestMapUintUint8.Delete", t, func() {
		var k uint = 2693544894
		var v uint8 = 192

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint8_Has(t *testing.T) {
	Convey("TestMapUintUint8.Has", t, func() {
		var k uint = 4112981836
		var v uint8 = 3

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3855130054+2210622721), ShouldBeFalse)
	})
}

func TestMapUintUint8_Get(t *testing.T) {
	Convey("TestMapUintUint8.Get", t, func() {
		var k uint = 2068082243
		var v uint8 = 227

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(3525085790 + 465730310)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint8_GetOpt(t *testing.T) {
	Convey("TestMapUintUint8.GetOpt", t, func() {
		var k uint = 2301785447
		var v uint8 = 30

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(4058714635 + 3005368670)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint8_ForEach(t *testing.T) {
	Convey("TestMapUintUint8.ForEach", t, func() {
		var k uint = 1713204380
		var v uint8 = 77
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
		var k uint = 613663870
		var v uint8 = 141

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
		var k uint = 3895407590
		var v uint8 = 43

		test := omap.NewMapUintUint8(1)

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

func TestMapUintUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint8.PutIfNotNil", t, func() {
		var k uint = 722214668
		var v uint8 = 171

		test := omap.NewMapUintUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1802070701, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 99
		So(test.PutIfNotNil(404432813, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint8.ReplaceIfExists", t, func() {
		var k uint = 4104626112
		var v uint8 = 17
		var x uint8 = 121

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3889934742, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint8.ReplaceOrPut", t, func() {
		var k uint = 126116130
		var v uint8 = 153
		var x uint8 = 121

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2336279656, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint8.MarshalJSON", t, func() {
		var k uint = 1076895806
		var v uint8 = 190

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1076895806,"value":190}]`)
	})
}
