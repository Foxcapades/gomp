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
		var k uint = 2958031440
		var v float32 = 0.179

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat32_Delete(t *testing.T) {
	Convey("TestMapUintFloat32.Delete", t, func() {
		var k uint = 1034705583
		var v float32 = 0.795

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat32_Has(t *testing.T) {
	Convey("TestMapUintFloat32.Has", t, func() {
		var k uint = 472764210
		var v float32 = 0.369

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2099602752+1131021293), ShouldBeFalse)
	})
}

func TestMapUintFloat32_Get(t *testing.T) {
	Convey("TestMapUintFloat32.Get", t, func() {
		var k uint = 3705571361
		var v float32 = 0.305

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(3252377266 + 4189864222)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintFloat32_GetOpt(t *testing.T) {
	Convey("TestMapUintFloat32.GetOpt", t, func() {
		var k uint = 774271115
		var v float32 = 0.685

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2842419977 + 2766611043)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintFloat32_ForEach(t *testing.T) {
	Convey("TestMapUintFloat32.ForEach", t, func() {
		var k uint = 3892076669
		var v float32 = 0.124
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
		var k uint = 2824108777
		var v float32 = 0.074

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
		var k uint = 3466272198
		var v float32 = 0.534

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
		var k uint = 435341864
		var v float32 = 0.301

		test := omap.NewMapUintFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2337850665, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.060
		So(test.PutIfNotNil(1251293605, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintFloat32.ReplaceIfExists", t, func() {
		var k uint = 3755106079
		var v float32 = 0.720
		var x float32 = 0.535

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3474440321, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintFloat32.ReplaceOrPut", t, func() {
		var k uint = 1268185137
		var v float32 = 0.214
		var x float32 = 0.497

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(644525643, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintFloat32.MarshalJSON", t, func() {
		var k uint = 1182859297
		var v float32 = 0.424

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1182859297,"value":0.424}]`)
	})
}
