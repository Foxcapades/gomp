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
		var k int = 804758148
		var v int16 = 6095

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt16_Delete(t *testing.T) {
	Convey("TestMapIntInt16.Delete", t, func() {
		var k int = 2106541834
		var v int16 = 2800

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt16_Has(t *testing.T) {
	Convey("TestMapIntInt16.Has", t, func() {
		var k int = 2037977833
		var v int16 = 23184

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1452909505+1542942284), ShouldBeFalse)
	})
}

func TestMapIntInt16_Get(t *testing.T) {
	Convey("TestMapIntInt16.Get", t, func() {
		var k int = 324433600
		var v int16 = 6830

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(658847679 + 1112736094)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt16_GetOpt(t *testing.T) {
	Convey("TestMapIntInt16.GetOpt", t, func() {
		var k int = 2009168723
		var v int16 = 12158

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1342283685 + 236381482)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt16_ForEach(t *testing.T) {
	Convey("TestMapIntInt16.ForEach", t, func() {
		var k int = 846204110
		var v int16 = 24401
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
		var k int = 1175485282
		var v int16 = 6685

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
		var k int = 430535800
		var v int16 = 799

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapIntInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt16.PutIfNotNil", t, func() {
		var k int = 24465011
		var v int16 = 21001

		test := omap.NewMapIntInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1345344045, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 29973
		So(test.PutIfNotNil(911621298, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceIfExists", t, func() {
		var k int = 63635244
		var v int16 = 4554
		var x int16 = 3935

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1282417285, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceOrPut", t, func() {
		var k int = 1589612270
		var v int16 = 20983
		var x int16 = 27298

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(215993755, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt16.MarshalJSON", t, func() {
		var k int = 879468120
		var v int16 = 18591

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":879468120,"value":18591}]`)
	})
}
