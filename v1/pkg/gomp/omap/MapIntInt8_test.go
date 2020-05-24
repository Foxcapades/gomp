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
		var k int = 1326330647
		var v int8 = 113

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt8_Delete(t *testing.T) {
	Convey("TestMapIntInt8.Delete", t, func() {
		var k int = 1677007024
		var v int8 = 4

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt8_Has(t *testing.T) {
	Convey("TestMapIntInt8.Has", t, func() {
		var k int = 1320004027
		var v int8 = 62

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1692681206+1530482732), ShouldBeFalse)
	})
}

func TestMapIntInt8_Get(t *testing.T) {
	Convey("TestMapIntInt8.Get", t, func() {
		var k int = 231709663
		var v int8 = 68

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(985390902 + 348790111)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt8_GetOpt(t *testing.T) {
	Convey("TestMapIntInt8.GetOpt", t, func() {
		var k int = 1606167081
		var v int8 = 79

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(492532713 + 1580182036)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt8_ForEach(t *testing.T) {
	Convey("TestMapIntInt8.ForEach", t, func() {
		var k int = 1415593684
		var v int8 = 16
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
		var k int = 996997293
		var v int8 = 42

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
		var k int = 619335
		var v int8 = 71

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
		var k int = 1561855099
		var v int8 = 79

		test := omap.NewMapIntInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1728598779, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 101
		So(test.PutIfNotNil(215278683, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceIfExists", t, func() {
		var k int = 2076391572
		var v int8 = 113
		var x int8 = 60

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1705735194, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceOrPut", t, func() {
		var k int = 812085538
		var v int8 = 56
		var x int8 = 105

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1714471801, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt8.MarshalJSON", t, func() {
		var k int = 660301768
		var v int8 = 20

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":660301768,"value":20}]`)
	})
}
