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
		var k uint = 4208850642
		var v uint16 = 50199

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint16_Delete(t *testing.T) {
	Convey("TestMapUintUint16.Delete", t, func() {
		var k uint = 360222266
		var v uint16 = 31292

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint16_Has(t *testing.T) {
	Convey("TestMapUintUint16.Has", t, func() {
		var k uint = 237455984
		var v uint16 = 24428

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(442036270+115584042), ShouldBeFalse)
	})
}


func TestMapUintUint16_Get(t *testing.T) {
	Convey("TestMapUintUint16.Get", t, func() {
		var k uint = 3994205889
		var v uint16 = 36714

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(4206807760 + 3914461686)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint16_GetOpt(t *testing.T) {
	Convey("TestMapUintUint16.GetOpt", t, func() {
		var k uint = 2606583368
		var v uint16 = 3834

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(783461274 + 1916457091)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint16_ForEach(t *testing.T) {
	Convey("TestMapUintUint16.ForEach", t, func() {
		var k uint = 856507642
		var v uint16 = 27137
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
		var k uint = 2865680623
		var v uint16 = 60576

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
		var k uint = 2135877975
		var v uint16 = 16287

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint16.PutIfNotNil", t, func() {
		var k uint = 2462625996
		var v uint16 = 4977

		test := omap.NewMapUintUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3882571250, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 35658
		So(test.PutIfNotNil(3414809547, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceIfExists", t, func() {
		var k uint = 485590556
		var v uint16 = 22281
		var x uint16 = 46008

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(262120147, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceOrPut", t, func() {
		var k uint = 2860090043
		var v uint16 = 13448
		var x uint16 = 50223

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2200093416, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint16.MarshalJSON", t, func() {
		var k uint = 1916517818
		var v uint16 = 27509

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1916517818,"value":27509}]`)
	})
}
