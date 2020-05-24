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
		var k int = 732537725
		var v int8 = 65

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt8_Delete(t *testing.T) {
	Convey("TestMapIntInt8.Delete", t, func() {
		var k int = 354604105
		var v int8 = 84

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt8_Has(t *testing.T) {
	Convey("TestMapIntInt8.Has", t, func() {
		var k int = 564637407
		var v int8 = 0

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1466677+1606477867), ShouldBeFalse)
	})
}


func TestMapIntInt8_Get(t *testing.T) {
	Convey("TestMapIntInt8.Get", t, func() {
		var k int = 1217421432
		var v int8 = 89

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(115840630 + 1770981521)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt8_GetOpt(t *testing.T) {
	Convey("TestMapIntInt8.GetOpt", t, func() {
		var k int = 768492538
		var v int8 = 115

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(181296477 + 24936847)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt8_ForEach(t *testing.T) {
	Convey("TestMapIntInt8.ForEach", t, func() {
		var k int = 971747726
		var v int8 = 62
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
		var k int = 1350852020
		var v int8 = 52

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
		var k int = 1845811913
		var v int8 = 67

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
		var k int = 1428942780
		var v int8 = 93

		test := omap.NewMapIntInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(634960754, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 42
		So(test.PutIfNotNil(2134542272, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceIfExists", t, func() {
		var k int = 429068036
		var v int8 = 41
		var x int8 = 23

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1486206123, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceOrPut", t, func() {
		var k int = 1699689262
		var v int8 = 80
		var x int8 = 51

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1642754674, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt8.MarshalJSON", t, func() {
		var k int = 1414776090
		var v int8 = 24

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1414776090,"value":24}]`)
	})
}

