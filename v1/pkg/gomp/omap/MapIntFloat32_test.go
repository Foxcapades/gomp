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
		var k int = 622294709
		var v float32 = 0.962

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat32_Delete(t *testing.T) {
	Convey("TestMapIntFloat32.Delete", t, func() {
		var k int = 1016614203
		var v float32 = 0.414

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat32_Has(t *testing.T) {
	Convey("TestMapIntFloat32.Has", t, func() {
		var k int = 790035793
		var v float32 = 0.193

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(767201377+81032348), ShouldBeFalse)
	})
}


func TestMapIntFloat32_Get(t *testing.T) {
	Convey("TestMapIntFloat32.Get", t, func() {
		var k int = 1632874485
		var v float32 = 0.931

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1495214235 + 639309846)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat32_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat32.GetOpt", t, func() {
		var k int = 1956034516
		var v float32 = 0.951

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(703929606 + 1120995363)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat32_ForEach(t *testing.T) {
	Convey("TestMapIntFloat32.ForEach", t, func() {
		var k int = 1818529201
		var v float32 = 0.580
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
		var k int = 2135606273
		var v float32 = 0.076

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
		var k int = 218223629
		var v float32 = 0.671

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
		var k int = 1284550498
		var v float32 = 0.143

		test := omap.NewMapIntFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1775879024, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.649
		So(test.PutIfNotNil(384371416, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceIfExists", t, func() {
		var k int = 590155099
		var v float32 = 0.727
		var x float32 = 0.684

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(227073255, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceOrPut", t, func() {
		var k int = 1893092498
		var v float32 = 0.990
		var x float32 = 0.295

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1328671368, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat32.MarshalJSON", t, func() {
		var k int = 753779671
		var v float32 = 0.964

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":753779671,"value":0.964}]`)
	})
}
