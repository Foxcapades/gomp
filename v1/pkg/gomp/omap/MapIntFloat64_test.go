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
		var k int = 1710658177
		var v float64 = 0.223

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat64_Delete(t *testing.T) {
	Convey("TestMapIntFloat64.Delete", t, func() {
		var k int = 1635341794
		var v float64 = 0.481

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat64_Has(t *testing.T) {
	Convey("TestMapIntFloat64.Has", t, func() {
		var k int = 424897798
		var v float64 = 0.541

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(863363699+408149431), ShouldBeFalse)
	})
}


func TestMapIntFloat64_Get(t *testing.T) {
	Convey("TestMapIntFloat64.Get", t, func() {
		var k int = 2066858068
		var v float64 = 0.118

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(817314971 + 924217424)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat64_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat64.GetOpt", t, func() {
		var k int = 1221113308
		var v float64 = 0.693

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(438404915 + 1174976602)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat64_ForEach(t *testing.T) {
	Convey("TestMapIntFloat64.ForEach", t, func() {
		var k int = 141346512
		var v float64 = 0.852
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
		var k int = 1069536821
		var v float64 = 0.032

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
		var k int = 777791794
		var v float64 = 0.240

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
		var k int = 1408184868
		var v float64 = 0.015

		test := omap.NewMapIntFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(440428690, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.647
		So(test.PutIfNotNil(1415989427, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceIfExists", t, func() {
		var k int = 1794059728
		var v float64 = 0.018
		var x float64 = 0.216

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(43754473, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceOrPut", t, func() {
		var k int = 1836030512
		var v float64 = 0.308
		var x float64 = 0.834

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1497715235, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat64.MarshalJSON", t, func() {
		var k int = 1736161627
		var v float64 = 0.496

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1736161627,"value":0.496}]`)
	})
}

