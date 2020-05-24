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
		var k uint = 3784765547
		var v int8 = 49

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt8_Delete(t *testing.T) {
	Convey("TestMapUintInt8.Delete", t, func() {
		var k uint = 478272728
		var v int8 = 72

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt8_Has(t *testing.T) {
	Convey("TestMapUintInt8.Has", t, func() {
		var k uint = 663854003
		var v int8 = 14

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2388234787+2997961142), ShouldBeFalse)
	})
}


func TestMapUintInt8_Get(t *testing.T) {
	Convey("TestMapUintInt8.Get", t, func() {
		var k uint = 4091205349
		var v int8 = 114

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1217201980 + 290802064)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt8_GetOpt(t *testing.T) {
	Convey("TestMapUintInt8.GetOpt", t, func() {
		var k uint = 3753451917
		var v int8 = 98

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3482238448 + 3725251835)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt8_ForEach(t *testing.T) {
	Convey("TestMapUintInt8.ForEach", t, func() {
		var k uint = 3335657915
		var v int8 = 22
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
		var k uint = 391141975
		var v int8 = 8

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
		var k uint = 353369192
		var v int8 = 59

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
		var k uint = 1463565755
		var v int8 = 30

		test := omap.NewMapUintInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1969729704, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 36
		So(test.PutIfNotNil(1653635238, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceIfExists", t, func() {
		var k uint = 3167503760
		var v int8 = 38
		var x int8 = 54

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2920835335, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceOrPut", t, func() {
		var k uint = 3603479087
		var v int8 = 93
		var x int8 = 15

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3419097528, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt8.MarshalJSON", t, func() {
		var k uint = 3406546966
		var v int8 = 110

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3406546966,"value":110}]`)
	})
}

