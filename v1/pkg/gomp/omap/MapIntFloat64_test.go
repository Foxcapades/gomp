package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntFloat64_Put(t *testing.T) {
	Convey("TestMapIntFloat64.Put", t, func() {
		var k int = 1569284587
		var v float64 = 0.712

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat64_Delete(t *testing.T) {
	Convey("TestMapIntFloat64.Delete", t, func() {
		var k int = 1284594058
		var v float64 = 0.855

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat64_Has(t *testing.T) {
	Convey("TestMapIntFloat64.Has", t, func() {
		var k int = 974055516
		var v float64 = 0.733

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(861791543+1184915417), ShouldBeFalse)
	})
}


func TestMapIntFloat64_Get(t *testing.T) {
	Convey("TestMapIntFloat64.Get", t, func() {
		var k int = 639408493
		var v float64 = 0.594

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(213049044+122198374)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat64_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat64.GetOpt", t, func() {
		var k int = 791908068
		var v float64 = 0.890

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(815430191+174515566)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat64_ForEach(t *testing.T) {
	Convey("TestMapIntFloat64.ForEach", t, func() {
		var k int = 1685987739
		var v float64 = 0.195
		hits := 0

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntFloat64.MarshalYAML", t, func() {
		var k int = 1316005681
		var v float64 = 0.020

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntFloat64_ToYAML(t *testing.T) {
	Convey("TestMapIntFloat64.ToYAML", t, func() {
		var k int = 1123285575
		var v float64 = 0.017

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntFloat64.PutIfNotNil", t, func() {
		var k int = 991395920
		var v float64 = 0.238

		test := omap.NewMapIntFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(406473262, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.458
		So(test.PutIfNotNil(465668083, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceIfExists", t, func() {
		var k int = 1376637601
		var v float64 = 0.206
		var x float64 = 0.904

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(868015437, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceOrPut", t, func() {
		var k int = 302695025
		var v float64 = 0.893
		var x float64 = 0.707

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(956029552, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat64.MarshalJSON", t, func() {
		var k int = 1666098724
		var v float64 = 0.804

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1666098724,"value":0.804}]`)
	})
}

