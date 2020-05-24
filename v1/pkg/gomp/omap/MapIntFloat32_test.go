package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntFloat32_Put(t *testing.T) {
	Convey("TestMapIntFloat32.Put", t, func() {
		var k int = 328926302
		var v float32 = 0.460

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat32_Delete(t *testing.T) {
	Convey("TestMapIntFloat32.Delete", t, func() {
		var k int = 1984384489
		var v float32 = 0.788

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat32_Has(t *testing.T) {
	Convey("TestMapIntFloat32.Has", t, func() {
		var k int = 1827536786
		var v float32 = 0.822

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(527654514+1478318694), ShouldBeFalse)
	})
}


func TestMapIntFloat32_Get(t *testing.T) {
	Convey("TestMapIntFloat32.Get", t, func() {
		var k int = 277311092
		var v float32 = 0.303

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1821629080+1616914754)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat32_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat32.GetOpt", t, func() {
		var k int = 1081742914
		var v float32 = 0.925

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(7927701+1427412220)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat32_ForEach(t *testing.T) {
	Convey("TestMapIntFloat32.ForEach", t, func() {
		var k int = 419294402
		var v float32 = 0.062
		hits := 0

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapIntFloat32.MarshalYAML", t, func() {
		var k int = 712071766
		var v float32 = 0.261

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntFloat32_ToYAML(t *testing.T) {
	Convey("TestMapIntFloat32.ToYAML", t, func() {
		var k int = 386416439
		var v float32 = 0.739

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntFloat32.PutIfNotNil", t, func() {
		var k int = 269777965
		var v float32 = 0.653

		test := omap.NewMapIntFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1928970828, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.361
		So(test.PutIfNotNil(908959327, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceIfExists", t, func() {
		var k int = 1103716923
		var v float32 = 0.628
		var x float32 = 0.162

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(213820140, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceOrPut", t, func() {
		var k int = 1104558341
		var v float32 = 0.156
		var x float32 = 0.272

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(631237710, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat32.MarshalJSON", t, func() {
		var k int = 40486686
		var v float32 = 0.476

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":40486686,"value":0.476}]`)
	})
}

