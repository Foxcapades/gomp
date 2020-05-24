package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt8_Put(t *testing.T) {
	Convey("TestMapIntInt8.Put", t, func() {
		var k int = 2066813994
		var v int8 = 0

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt8_Delete(t *testing.T) {
	Convey("TestMapIntInt8.Delete", t, func() {
		var k int = 771169720
		var v int8 = 90

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt8_Has(t *testing.T) {
	Convey("TestMapIntInt8.Has", t, func() {
		var k int = 726939707
		var v int8 = 89

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1122554006+1633835935), ShouldBeFalse)
	})
}

func TestMapIntInt8_Get(t *testing.T) {
	Convey("TestMapIntInt8.Get", t, func() {
		var k int = 1259113172
		var v int8 = 8

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2053926295 + 1544600825)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt8_GetOpt(t *testing.T) {
	Convey("TestMapIntInt8.GetOpt", t, func() {
		var k int = 924598401
		var v int8 = 108

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1819269014 + 1954856589)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt8_ForEach(t *testing.T) {
	Convey("TestMapIntInt8.ForEach", t, func() {
		var k int = 1802058896
		var v int8 = 115
		hits := 0

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt8.MarshalYAML", t, func() {
		var k int = 882906858
		var v int8 = 22

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt8_ToYAML(t *testing.T) {
	Convey("TestMapIntInt8.ToYAML", t, func() {
		var k int = 36532334
		var v int8 = 56

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt8.PutIfNotNil", t, func() {
		var k int = 1415004831
		var v int8 = 14

		test := omap.NewMapIntInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1216099991, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 35
		So(test.PutIfNotNil(1010369058, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceIfExists", t, func() {
		var k int = 2112560385
		var v int8 = 98
		var x int8 = 94

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(759113880, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceOrPut", t, func() {
		var k int = 1919152720
		var v int8 = 68
		var x int8 = 126

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(786570971, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt8.MarshalJSON", t, func() {
		var k int = 1076333509
		var v int8 = 113

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1076333509,"value":113}]`)
	})
}
