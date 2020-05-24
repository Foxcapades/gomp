package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint8_Put(t *testing.T) {
	Convey("TestMapIntUint8.Put", t, func() {
		var k int = 911678976
		var v uint8 = 233

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint8_Delete(t *testing.T) {
	Convey("TestMapIntUint8.Delete", t, func() {
		var k int = 1089549519
		var v uint8 = 56

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint8_Has(t *testing.T) {
	Convey("TestMapIntUint8.Has", t, func() {
		var k int = 444586586
		var v uint8 = 113

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1226244108+1891656010), ShouldBeFalse)
	})
}


func TestMapIntUint8_Get(t *testing.T) {
	Convey("TestMapIntUint8.Get", t, func() {
		var k int = 796965517
		var v uint8 = 236

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2037285864 + 1637778570)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint8_GetOpt(t *testing.T) {
	Convey("TestMapIntUint8.GetOpt", t, func() {
		var k int = 613269676
		var v uint8 = 244

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(839058477 + 1408391285)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint8_ForEach(t *testing.T) {
	Convey("TestMapIntUint8.ForEach", t, func() {
		var k int = 484915695
		var v uint8 = 101
		hits := 0

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint8.MarshalYAML", t, func() {
		var k int = 894043082
		var v uint8 = 77

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint8_ToYAML(t *testing.T) {
	Convey("TestMapIntUint8.ToYAML", t, func() {
		var k int = 818949594
		var v uint8 = 96

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint8.PutIfNotNil", t, func() {
		var k int = 211888096
		var v uint8 = 30

		test := omap.NewMapIntUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1956840573, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 118
		So(test.PutIfNotNil(252777591, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceIfExists", t, func() {
		var k int = 1633992816
		var v uint8 = 184
		var x uint8 = 68

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(165098511, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceOrPut", t, func() {
		var k int = 1545664020
		var v uint8 = 245
		var x uint8 = 164

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(707474755, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint8.MarshalJSON", t, func() {
		var k int = 27012023
		var v uint8 = 23

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":27012023,"value":23}]`)
	})
}

