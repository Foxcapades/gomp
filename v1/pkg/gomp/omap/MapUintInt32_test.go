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
		var k uint = 2075076989
		var v int32 = 1705642582

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt32_Delete(t *testing.T) {
	Convey("TestMapUintInt32.Delete", t, func() {
		var k uint = 3725175772
		var v int32 = 733414525

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt32_Has(t *testing.T) {
	Convey("TestMapUintInt32.Has", t, func() {
		var k uint = 4063428554
		var v int32 = 661113878

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1524041172+2091203368), ShouldBeFalse)
	})
}

func TestMapUintInt32_Get(t *testing.T) {
	Convey("TestMapUintInt32.Get", t, func() {
		var k uint = 384118754
		var v int32 = 2125490148

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(3778289650 + 1081784426)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt32_GetOpt(t *testing.T) {
	Convey("TestMapUintInt32.GetOpt", t, func() {
		var k uint = 1029451243
		var v int32 = 1757297738

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(184930592 + 4149145279)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt32_ForEach(t *testing.T) {
	Convey("TestMapUintInt32.ForEach", t, func() {
		var k uint = 2766565063
		var v int32 = 377557257
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
		var k uint = 1478136265
		var v int32 = 1691975758

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
		var k uint = 517964121
		var v int32 = 745450751

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
		var k uint = 3567764741
		var v int32 = 1361667357

		test := omap.NewMapUintInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3597410598, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 570441516
		So(test.PutIfNotNil(2206149893, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt32.ReplaceIfExists", t, func() {
		var k uint = 370381998
		var v int32 = 442865996
		var x int32 = 1403012277

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3074298120, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt32.ReplaceOrPut", t, func() {
		var k uint = 3323418385
		var v int32 = 857616593
		var x int32 = 1073939200

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3867160385, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt32.MarshalJSON", t, func() {
		var k uint = 21671156
		var v int32 = 505923399

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":21671156,"value":505923399}]`)
	})
}
