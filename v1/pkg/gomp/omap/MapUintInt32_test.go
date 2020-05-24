package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt32_Put(t *testing.T) {
	Convey("TestMapUintInt32.Put", t, func() {
		var k uint = 2361431427
		var v int32 = 855940061

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt32_Delete(t *testing.T) {
	Convey("TestMapUintInt32.Delete", t, func() {
		var k uint = 626367047
		var v int32 = 1526013465

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt32_Has(t *testing.T) {
	Convey("TestMapUintInt32.Has", t, func() {
		var k uint = 3150738808
		var v int32 = 916564930

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1854238361+3232279415), ShouldBeFalse)
	})
}


func TestMapUintInt32_Get(t *testing.T) {
	Convey("TestMapUintInt32.Get", t, func() {
		var k uint = 331967242
		var v int32 = 679474790

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2772286771 + 11524710)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt32_GetOpt(t *testing.T) {
	Convey("TestMapUintInt32.GetOpt", t, func() {
		var k uint = 1882492335
		var v int32 = 992339693

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(73993385 + 3949212131)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt32_ForEach(t *testing.T) {
	Convey("TestMapUintInt32.ForEach", t, func() {
		var k uint = 2406167280
		var v int32 = 1872515142
		hits := 0

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt32.MarshalYAML", t, func() {
		var k uint = 2329792529
		var v int32 = 444743161

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt32_ToYAML(t *testing.T) {
	Convey("TestMapUintInt32.ToYAML", t, func() {
		var k uint = 2686232082
		var v int32 = 1166379017

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt32.PutIfNotNil", t, func() {
		var k uint = 58718886
		var v int32 = 1436998770

		test := omap.NewMapUintInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(4000956613, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 1016811295
		So(test.PutIfNotNil(524117566, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt32.ReplaceIfExists", t, func() {
		var k uint = 1299999861
		var v int32 = 1447490546
		var x int32 = 1487721019

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1478168974, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt32.ReplaceOrPut", t, func() {
		var k uint = 1110096378
		var v int32 = 740156180
		var x int32 = 1367916044

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3786801845, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt32.MarshalJSON", t, func() {
		var k uint = 3675310966
		var v int32 = 91381534

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3675310966,"value":91381534}]`)
	})
}
