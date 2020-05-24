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
		var k int = 283304802
		var v int16 = 30949

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt16_Delete(t *testing.T) {
	Convey("TestMapIntInt16.Delete", t, func() {
		var k int = 446589176
		var v int16 = 30030

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt16_Has(t *testing.T) {
	Convey("TestMapIntInt16.Has", t, func() {
		var k int = 469996939
		var v int16 = 8564

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(50964314+763095343), ShouldBeFalse)
	})
}

func TestMapIntInt16_Get(t *testing.T) {
	Convey("TestMapIntInt16.Get", t, func() {
		var k int = 2051694753
		var v int16 = 21724

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(452884727 + 1125094267)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt16_GetOpt(t *testing.T) {
	Convey("TestMapIntInt16.GetOpt", t, func() {
		var k int = 563083789
		var v int16 = 8574

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1772515909 + 662188284)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt16_ForEach(t *testing.T) {
	Convey("TestMapIntInt16.ForEach", t, func() {
		var k int = 311328554
		var v int16 = 21070
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
		var k int = 1436070916
		var v int16 = 30938

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
		var k int = 1566583535
		var v int16 = 22712

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
		var k int = 1504862312
		var v int16 = 3189

		test := omap.NewMapIntInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(103486885, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 13423
		So(test.PutIfNotNil(1894306153, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceIfExists", t, func() {
		var k int = 1515902100
		var v int16 = 9661
		var x int16 = 5698

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(340813216, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceOrPut", t, func() {
		var k int = 672723115
		var v int16 = 30615
		var x int16 = 18051

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1378688504, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt16.MarshalJSON", t, func() {
		var k int = 1624712377
		var v int16 = 28589

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1624712377,"value":28589}]`)
	})
}
