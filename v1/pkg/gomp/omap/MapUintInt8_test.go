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
		var k uint = 3166792360
		var v int8 = 46

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt8_Delete(t *testing.T) {
	Convey("TestMapUintInt8.Delete", t, func() {
		var k uint = 3466878658
		var v int8 = 117

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt8_Has(t *testing.T) {
	Convey("TestMapUintInt8.Has", t, func() {
		var k uint = 1709173990
		var v int8 = 97

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(796826557+1296718743), ShouldBeFalse)
	})
}

func TestMapUintInt8_Get(t *testing.T) {
	Convey("TestMapUintInt8.Get", t, func() {
		var k uint = 400842875
		var v int8 = 112

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2216669141 + 1205402529)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt8_GetOpt(t *testing.T) {
	Convey("TestMapUintInt8.GetOpt", t, func() {
		var k uint = 223662977
		var v int8 = 44

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3745991869 + 1561923919)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt8_ForEach(t *testing.T) {
	Convey("TestMapUintInt8.ForEach", t, func() {
		var k uint = 1003441277
		var v int8 = 84
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
		var k uint = 490272333
		var v int8 = 111

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
		var k uint = 695791395
		var v int8 = 20

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
		var k uint = 640854813
		var v int8 = 43

		test := omap.NewMapUintInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3897517346, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 56
		So(test.PutIfNotNil(879687379, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceIfExists", t, func() {
		var k uint = 2452733404
		var v int8 = 96
		var x int8 = 98

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1922223497, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceOrPut", t, func() {
		var k uint = 1251523211
		var v int8 = 98
		var x int8 = 96

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(841986125, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt8.MarshalJSON", t, func() {
		var k uint = 4193632654
		var v int8 = 5

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":4193632654,"value":5}]`)
	})
}
