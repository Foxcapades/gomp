package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt16_Put(t *testing.T) {
	Convey("TestMapUintInt16.Put", t, func() {
		var k uint = 3964509523
		var v int16 = 4415

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt16_Delete(t *testing.T) {
	Convey("TestMapUintInt16.Delete", t, func() {
		var k uint = 1333411926
		var v int16 = 28104

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt16_Has(t *testing.T) {
	Convey("TestMapUintInt16.Has", t, func() {
		var k uint = 1219267705
		var v int16 = 9705

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2397630177+2517120771), ShouldBeFalse)
	})
}

func TestMapUintInt16_Get(t *testing.T) {
	Convey("TestMapUintInt16.Get", t, func() {
		var k uint = 2740625941
		var v int16 = 11821

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2200892877 + 1704394534)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt16_GetOpt(t *testing.T) {
	Convey("TestMapUintInt16.GetOpt", t, func() {
		var k uint = 2898978604
		var v int16 = 26348

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1720595936 + 562406816)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt16_ForEach(t *testing.T) {
	Convey("TestMapUintInt16.ForEach", t, func() {
		var k uint = 1818587526
		var v int16 = 5009
		hits := 0

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt16.MarshalYAML", t, func() {
		var k uint = 1783512464
		var v int16 = 2345

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt16_ToYAML(t *testing.T) {
	Convey("TestMapUintInt16.ToYAML", t, func() {
		var k uint = 3835503200
		var v int16 = 25560

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt16.PutIfNotNil", t, func() {
		var k uint = 283341296
		var v int16 = 14836

		test := omap.NewMapUintInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3151206256, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 13630
		So(test.PutIfNotNil(3888261920, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt16.ReplaceIfExists", t, func() {
		var k uint = 1152849271
		var v int16 = 9073
		var x int16 = 19107

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3143494737, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt16.ReplaceOrPut", t, func() {
		var k uint = 323110608
		var v int16 = 10505
		var x int16 = 9603

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3043378496, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt16.MarshalJSON", t, func() {
		var k uint = 306930875
		var v int16 = 25895

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":306930875,"value":25895}]`)
	})
}
