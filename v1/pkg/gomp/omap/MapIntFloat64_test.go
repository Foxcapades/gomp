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
		var k int = 745844917
		var v float64 = 0.348

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat64_Delete(t *testing.T) {
	Convey("TestMapIntFloat64.Delete", t, func() {
		var k int = 903520924
		var v float64 = 0.779

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat64_Has(t *testing.T) {
	Convey("TestMapIntFloat64.Has", t, func() {
		var k int = 503868956
		var v float64 = 0.976

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2073786340+1910867804), ShouldBeFalse)
	})
}

func TestMapIntFloat64_Get(t *testing.T) {
	Convey("TestMapIntFloat64.Get", t, func() {
		var k int = 714677053
		var v float64 = 0.435

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(241126872 + 1539551237)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat64_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat64.GetOpt", t, func() {
		var k int = 1375669038
		var v float64 = 0.124

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1139263119 + 1874854824)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat64_ForEach(t *testing.T) {
	Convey("TestMapIntFloat64.ForEach", t, func() {
		var k int = 62005266
		var v float64 = 0.655
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
		var k int = 1638918456
		var v float64 = 0.006

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
		var k int = 603493607
		var v float64 = 0.615

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
		var k int = 2142328688
		var v float64 = 0.168

		test := omap.NewMapIntFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1953380391, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.262
		So(test.PutIfNotNil(1487196366, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceIfExists", t, func() {
		var k int = 301296186
		var v float64 = 0.935
		var x float64 = 0.255

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(372579158, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceOrPut", t, func() {
		var k int = 1306553698
		var v float64 = 0.187
		var x float64 = 0.948

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(594106419, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat64.MarshalJSON", t, func() {
		var k int = 674634049
		var v float64 = 0.018

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":674634049,"value":0.018}]`)
	})
}
