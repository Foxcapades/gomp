package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint8_Put(t *testing.T) {
	Convey("TestMapUintUint8.Put", t, func() {
		var k uint = 3138467168
		var v uint8 = 221

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint8_Delete(t *testing.T) {
	Convey("TestMapUintUint8.Delete", t, func() {
		var k uint = 364018690
		var v uint8 = 84

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint8_Has(t *testing.T) {
	Convey("TestMapUintUint8.Has", t, func() {
		var k uint = 358237628
		var v uint8 = 142

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(847163087+1184163975), ShouldBeFalse)
	})
}


func TestMapUintUint8_Get(t *testing.T) {
	Convey("TestMapUintUint8.Get", t, func() {
		var k uint = 3036021530
		var v uint8 = 8

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1947837191 + 554662460)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint8_GetOpt(t *testing.T) {
	Convey("TestMapUintUint8.GetOpt", t, func() {
		var k uint = 1529762871
		var v uint8 = 253

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(4100490433 + 514750194)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint8_ForEach(t *testing.T) {
	Convey("TestMapUintUint8.ForEach", t, func() {
		var k uint = 3415704247
		var v uint8 = 226
		hits := 0

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint8.MarshalYAML", t, func() {
		var k uint = 2731062935
		var v uint8 = 37

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint8_ToYAML(t *testing.T) {
	Convey("TestMapUintUint8.ToYAML", t, func() {
		var k uint = 3034992532
		var v uint8 = 249

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint8.PutIfNotNil", t, func() {
		var k uint = 292849231
		var v uint8 = 119

		test := omap.NewMapUintUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(447788849, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 97
		So(test.PutIfNotNil(2770243690, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint8.ReplaceIfExists", t, func() {
		var k uint = 2716142505
		var v uint8 = 203
		var x uint8 = 209

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2182688439, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint8.ReplaceOrPut", t, func() {
		var k uint = 1026703038
		var v uint8 = 38
		var x uint8 = 249

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(888247065, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint8.MarshalJSON", t, func() {
		var k uint = 2062197015
		var v uint8 = 118

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2062197015,"value":118}]`)
	})
}

