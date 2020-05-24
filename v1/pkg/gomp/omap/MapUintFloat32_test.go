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
		var k uint = 2021242674
		var v float32 = 0.596

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat32_Delete(t *testing.T) {
	Convey("TestMapUintFloat32.Delete", t, func() {
		var k uint = 170421826
		var v float32 = 0.264

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat32_Has(t *testing.T) {
	Convey("TestMapUintFloat32.Has", t, func() {
		var k uint = 3928891343
		var v float32 = 0.970

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2754961201+364749334), ShouldBeFalse)
	})
}


func TestMapUintFloat32_Get(t *testing.T) {
	Convey("TestMapUintFloat32.Get", t, func() {
		var k uint = 2668562572
		var v float32 = 0.863

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1873405014 + 2265209970)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintFloat32_GetOpt(t *testing.T) {
	Convey("TestMapUintFloat32.GetOpt", t, func() {
		var k uint = 645999263
		var v float32 = 0.443

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3723348132 + 3127340248)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintFloat32_ForEach(t *testing.T) {
	Convey("TestMapUintFloat32.ForEach", t, func() {
		var k uint = 2751400305
		var v float32 = 0.790
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
		var k uint = 2827551431
		var v float32 = 0.215

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
		var k uint = 1936192282
		var v float32 = 0.932

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
		var k uint = 1820669625
		var v float32 = 0.575

		test := omap.NewMapUintFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1510700718, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.263
		So(test.PutIfNotNil(71349498, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintFloat32.ReplaceIfExists", t, func() {
		var k uint = 1234881748
		var v float32 = 0.389
		var x float32 = 0.807

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(462842692, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintFloat32.ReplaceOrPut", t, func() {
		var k uint = 196411113
		var v float32 = 0.699
		var x float32 = 0.603

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1996850961, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintFloat32.MarshalJSON", t, func() {
		var k uint = 1306144646
		var v float32 = 0.380

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1306144646,"value":0.38}]`)
	})
}

