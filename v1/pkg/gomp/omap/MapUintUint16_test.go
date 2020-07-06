package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint16_Put(t *testing.T) {
	Convey("TestMapUintUint16.Put", t, func() {
		var k uint = 2880962979
		var v uint16 = 40728

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint16_Delete(t *testing.T) {
	Convey("TestMapUintUint16.Delete", t, func() {
		var k uint = 2960034541
		var v uint16 = 12252

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint16_Has(t *testing.T) {
	Convey("TestMapUintUint16.Has", t, func() {
		var k uint = 312319573
		var v uint16 = 63149

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(4213265135+2263615900), ShouldBeFalse)
	})
}

func TestMapUintUint16_Get(t *testing.T) {
	Convey("TestMapUintUint16.Get", t, func() {
		var k uint = 4285777064
		var v uint16 = 55246

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(785578192 + 828383807)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint16_GetOpt(t *testing.T) {
	Convey("TestMapUintUint16.GetOpt", t, func() {
		var k uint = 2461931407
		var v uint16 = 53537

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1578484566 + 76659732)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint16_ForEach(t *testing.T) {
	Convey("TestMapUintUint16.ForEach", t, func() {
		var k uint = 3599421356
		var v uint16 = 9809
		hits := 0

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint16.MarshalYAML", t, func() {
		var k uint = 4056943430
		var v uint16 = 23810

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint16_ToYAML(t *testing.T) {
	Convey("TestMapUintUint16.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k uint = 3675948490
			var v uint16 = 37423

			test := omap.NewMapUintUint16(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k uint = 3012820436
			var v uint16 = 29658

			test := omap.NewMapUintUint16(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapUintUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint16.PutIfNotNil", t, func() {
		var k uint = 1892613954
		var v uint16 = 31060

		test := omap.NewMapUintUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(505937878, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 15685
		So(test.PutIfNotNil(1167157359, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceIfExists", t, func() {
		var k uint = 3741771511
		var v uint16 = 52116
		var x uint16 = 21062

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(4085380251, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceOrPut", t, func() {
		var k uint = 4056207037
		var v uint16 = 34615
		var x uint16 = 21767

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3979829794, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint16.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k uint = 236711011
			var v uint16 = 42671

			test := omap.NewMapUintUint16(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":236711011,"value":42671}]`)
		})

		Convey("Unordered", func() {
			var k uint = 236711011
			var v uint16 = 42671

			test := omap.NewMapUintUint16(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"236711011":42671}`)
		})

	})
}
