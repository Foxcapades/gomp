package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt16_Put(t *testing.T) {
	Convey("TestMapIntInt16.Put", t, func() {
		var k int = 40908538
		var v int16 = 3616

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt16_Delete(t *testing.T) {
	Convey("TestMapIntInt16.Delete", t, func() {
		var k int = 1501283279
		var v int16 = 7071

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt16_Has(t *testing.T) {
	Convey("TestMapIntInt16.Has", t, func() {
		var k int = 1424244789
		var v int16 = 20441

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1821110289+937458360), ShouldBeFalse)
	})
}


func TestMapIntInt16_Get(t *testing.T) {
	Convey("TestMapIntInt16.Get", t, func() {
		var k int = 321350860
		var v int16 = 2627

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(60725891+1234882434)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt16_GetOpt(t *testing.T) {
	Convey("TestMapIntInt16.GetOpt", t, func() {
		var k int = 551455651
		var v int16 = 17012

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1588983018+866580491)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt16_ForEach(t *testing.T) {
	Convey("TestMapIntInt16.ForEach", t, func() {
		var k int = 907128143
		var v int16 = 1581
		hits := 0

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt16.MarshalYAML", t, func() {
		var k int = 634915258
		var v int16 = 5508

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt16_ToYAML(t *testing.T) {
	Convey("TestMapIntInt16.ToYAML", t, func() {
		var k int = 309548818
		var v int16 = 19891

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt16.PutIfNotNil", t, func() {
		var k int = 1397749752
		var v int16 = 28493

		test := omap.NewMapIntInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2045215648, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 3776
		So(test.PutIfNotNil(1870697503, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceIfExists", t, func() {
		var k int = 1957289535
		var v int16 = 23357
		var x int16 = 18075

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1371035693, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceOrPut", t, func() {
		var k int = 937999704
		var v int16 = 16301
		var x int16 = 28584

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1923675083, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt16.MarshalJSON", t, func() {
		var k int = 227925803
		var v int16 = 11695

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":227925803,"value":11695}]`)
	})
}

