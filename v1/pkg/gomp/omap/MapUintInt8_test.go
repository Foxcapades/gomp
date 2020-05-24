package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt8_Put(t *testing.T) {
	Convey("TestMapUintInt8.Put", t, func() {
		var k uint = 3923172815
		var v int8 = 32

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt8_Delete(t *testing.T) {
	Convey("TestMapUintInt8.Delete", t, func() {
		var k uint = 155740811
		var v int8 = 105

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt8_Has(t *testing.T) {
	Convey("TestMapUintInt8.Has", t, func() {
		var k uint = 2342784360
		var v int8 = 75

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2645428032+1852507967), ShouldBeFalse)
	})
}


func TestMapUintInt8_Get(t *testing.T) {
	Convey("TestMapUintInt8.Get", t, func() {
		var k uint = 2853060997
		var v int8 = 68

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(176774955+2963420261)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt8_GetOpt(t *testing.T) {
	Convey("TestMapUintInt8.GetOpt", t, func() {
		var k uint = 4245983145
		var v int8 = 104

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1853571057+879191977)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt8_ForEach(t *testing.T) {
	Convey("TestMapUintInt8.ForEach", t, func() {
		var k uint = 3215668014
		var v int8 = 21
		hits := 0

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt8.MarshalYAML", t, func() {
		var k uint = 3829942198
		var v int8 = 36

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt8_ToYAML(t *testing.T) {
	Convey("TestMapUintInt8.ToYAML", t, func() {
		var k uint = 3320110889
		var v int8 = 76

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt8.PutIfNotNil", t, func() {
		var k uint = 3405014062
		var v int8 = 24

		test := omap.NewMapUintInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(397598066, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 118
		So(test.PutIfNotNil(3759784197, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceIfExists", t, func() {
		var k uint = 2454664293
		var v int8 = 3
		var x int8 = 47

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2697413914, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceOrPut", t, func() {
		var k uint = 1306737148
		var v int8 = 13
		var x int8 = 101

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1160049129, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt8.MarshalJSON", t, func() {
		var k uint = 1057920571
		var v int8 = 79

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1057920571,"value":79}]`)
	})
}

