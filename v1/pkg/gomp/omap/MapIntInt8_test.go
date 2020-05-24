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
		var k int = 439238869
		var v int8 = 111

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt8_Delete(t *testing.T) {
	Convey("TestMapIntInt8.Delete", t, func() {
		var k int = 543170956
		var v int8 = 67

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt8_Has(t *testing.T) {
	Convey("TestMapIntInt8.Has", t, func() {
		var k int = 357618560
		var v int8 = 109

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1595950482+162284895), ShouldBeFalse)
	})
}


func TestMapIntInt8_Get(t *testing.T) {
	Convey("TestMapIntInt8.Get", t, func() {
		var k int = 957869102
		var v int8 = 13

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(298056345+886914602)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt8_GetOpt(t *testing.T) {
	Convey("TestMapIntInt8.GetOpt", t, func() {
		var k int = 1820596441
		var v int8 = 95

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(275094810+1458692340)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt8_ForEach(t *testing.T) {
	Convey("TestMapIntInt8.ForEach", t, func() {
		var k int = 65346164
		var v int8 = 79
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
		var k int = 1320310713
		var v int8 = 66

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
		var k int = 356651768
		var v int8 = 54

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
		var k int = 2093352113
		var v int8 = 105

		test := omap.NewMapIntInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(801882089, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 101
		So(test.PutIfNotNil(1623791288, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceIfExists", t, func() {
		var k int = 1327069992
		var v int8 = 39
		var x int8 = 87

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(809446767, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceOrPut", t, func() {
		var k int = 201537138
		var v int8 = 55
		var x int8 = 78

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1535337916, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt8.MarshalJSON", t, func() {
		var k int = 1995922760
		var v int8 = 31

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1995922760,"value":31}]`)
	})
}

