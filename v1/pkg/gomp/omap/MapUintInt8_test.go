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
		var k uint = 3327590197
		var v int8 = 25

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt8_Delete(t *testing.T) {
	Convey("TestMapUintInt8.Delete", t, func() {
		var k uint = 1677302396
		var v int8 = 31

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt8_Has(t *testing.T) {
	Convey("TestMapUintInt8.Has", t, func() {
		var k uint = 4194880703
		var v int8 = 18

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3736608938+3853141998), ShouldBeFalse)
	})
}

func TestMapUintInt8_Get(t *testing.T) {
	Convey("TestMapUintInt8.Get", t, func() {
		var k uint = 2745627088
		var v int8 = 54

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2688825166 + 294257440)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt8_GetOpt(t *testing.T) {
	Convey("TestMapUintInt8.GetOpt", t, func() {
		var k uint = 736358662
		var v int8 = 68

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3074656710 + 668771680)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt8_ForEach(t *testing.T) {
	Convey("TestMapUintInt8.ForEach", t, func() {
		var k uint = 2383764296
		var v int8 = 78
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
		var k uint = 4257083314
		var v int8 = 87

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
		var k uint = 2297996831
		var v int8 = 39

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
		var k uint = 2356482491
		var v int8 = 58

		test := omap.NewMapUintInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1858150799, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 103
		So(test.PutIfNotNil(1774572056, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceIfExists", t, func() {
		var k uint = 978661266
		var v int8 = 75
		var x int8 = 63

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1424555154, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceOrPut", t, func() {
		var k uint = 2681765627
		var v int8 = 24
		var x int8 = 120

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(384639553, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt8.MarshalJSON", t, func() {
		var k uint = 4253412326
		var v int8 = 13

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":4253412326,"value":13}]`)
	})
}
