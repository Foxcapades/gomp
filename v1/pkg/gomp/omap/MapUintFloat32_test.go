package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintFloat32_Put(t *testing.T) {
	Convey("TestMapUintFloat32.Put", t, func() {
		var k uint = 903357796
		var v float32 = 0.724

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat32_Delete(t *testing.T) {
	Convey("TestMapUintFloat32.Delete", t, func() {
		var k uint = 3441488327
		var v float32 = 0.215

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat32_Has(t *testing.T) {
	Convey("TestMapUintFloat32.Has", t, func() {
		var k uint = 1054475533
		var v float32 = 0.956

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1806323087+3649770565), ShouldBeFalse)
	})
}

func TestMapUintFloat32_Get(t *testing.T) {
	Convey("TestMapUintFloat32.Get", t, func() {
		var k uint = 823334118
		var v float32 = 0.742

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3338758220 + 1093257179)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintFloat32_GetOpt(t *testing.T) {
	Convey("TestMapUintFloat32.GetOpt", t, func() {
		var k uint = 932666139
		var v float32 = 0.787

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(908644798 + 1000215441)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintFloat32_ForEach(t *testing.T) {
	Convey("TestMapUintFloat32.ForEach", t, func() {
		var k uint = 2680633861
		var v float32 = 0.642
		hits := 0

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapUintFloat32.MarshalYAML", t, func() {
		var k uint = 2639047127
		var v float32 = 0.740

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintFloat32_ToYAML(t *testing.T) {
	Convey("TestMapUintFloat32.ToYAML", t, func() {
		var k uint = 2293049952
		var v float32 = 0.290

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintFloat32.PutIfNotNil", t, func() {
		var k uint = 1363292244
		var v float32 = 0.614

		test := omap.NewMapUintFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2746224769, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.200
		So(test.PutIfNotNil(975880069, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintFloat32.ReplaceIfExists", t, func() {
		var k uint = 3034320410
		var v float32 = 0.260
		var x float32 = 0.612

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(857237778, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintFloat32.ReplaceOrPut", t, func() {
		var k uint = 520486719
		var v float32 = 0.578
		var x float32 = 0.956

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2337248797, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintFloat32.MarshalJSON", t, func() {
		var k uint = 725890878
		var v float32 = 0.778

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":725890878,"value":0.778}]`)
	})
}
