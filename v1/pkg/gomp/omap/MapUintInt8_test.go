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
		var k uint = 238068175
		var v int8 = 120

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt8_Delete(t *testing.T) {
	Convey("TestMapUintInt8.Delete", t, func() {
		var k uint = 4226026232
		var v int8 = 97

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt8_Has(t *testing.T) {
	Convey("TestMapUintInt8.Has", t, func() {
		var k uint = 2768469349
		var v int8 = 6

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3308101739+2898805397), ShouldBeFalse)
	})
}

func TestMapUintInt8_Get(t *testing.T) {
	Convey("TestMapUintInt8.Get", t, func() {
		var k uint = 2589713331
		var v int8 = 25

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1665095209 + 3797687558)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt8_GetOpt(t *testing.T) {
	Convey("TestMapUintInt8.GetOpt", t, func() {
		var k uint = 555336693
		var v int8 = 14

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1570381499 + 1295823927)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt8_ForEach(t *testing.T) {
	Convey("TestMapUintInt8.ForEach", t, func() {
		var k uint = 3546934274
		var v int8 = 65
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
		var k uint = 14876959
		var v int8 = 35

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
		Convey("Ordered", func() {
			var k uint = 879649854
			var v int8 = 39

			test := omap.NewMapUintInt8(1)

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
			var k uint = 1777923693
			var v int8 = 64

			test := omap.NewMapUintInt8(1)
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

func TestMapUintInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt8.PutIfNotNil", t, func() {
		var k uint = 3202968892
		var v int8 = 15

		test := omap.NewMapUintInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(4099505590, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 91
		So(test.PutIfNotNil(1828223689, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceIfExists", t, func() {
		var k uint = 2954361439
		var v int8 = 88
		var x int8 = 122

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(4150447162, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt8.ReplaceOrPut", t, func() {
		var k uint = 3425867367
		var v int8 = 121
		var x int8 = 28

		test := omap.NewMapUintInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1586900895, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt8.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k uint = 4055342435
			var v int8 = 71

			test := omap.NewMapUintInt8(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":4055342435,"value":71}]`)
		})

		Convey("Unordered", func() {
			var k uint = 4055342435
			var v int8 = 71

			test := omap.NewMapUintInt8(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"4055342435":71}`)
		})

	})
}
