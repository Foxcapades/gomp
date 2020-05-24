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
		var k int = 35100906
		var v uint8 = 63

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint8_Delete(t *testing.T) {
	Convey("TestMapIntUint8.Delete", t, func() {
		var k int = 2065261912
		var v uint8 = 168

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint8_Has(t *testing.T) {
	Convey("TestMapIntUint8.Has", t, func() {
		var k int = 485715269
		var v uint8 = 157

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1597736675+837261649), ShouldBeFalse)
	})
}

func TestMapIntUint8_Get(t *testing.T) {
	Convey("TestMapIntUint8.Get", t, func() {
		var k int = 367227154
		var v uint8 = 172

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1110324175 + 717971317)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint8_GetOpt(t *testing.T) {
	Convey("TestMapIntUint8.GetOpt", t, func() {
		var k int = 400747850
		var v uint8 = 16

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(574455513 + 1416549268)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint8_ForEach(t *testing.T) {
	Convey("TestMapIntUint8.ForEach", t, func() {
		var k int = 1799160874
		var v uint8 = 159
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
		var k int = 938564748
		var v uint8 = 251

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
		var k int = 574481831
		var v uint8 = 216

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
		var k int = 499427741
		var v uint8 = 159

		test := omap.NewMapIntUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(908153154, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 239
		So(test.PutIfNotNil(1082010658, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceIfExists", t, func() {
		var k int = 1667856260
		var v uint8 = 7
		var x uint8 = 236

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(708453304, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceOrPut", t, func() {
		var k int = 2034448723
		var v uint8 = 230
		var x uint8 = 165

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(60901775, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint8.MarshalJSON", t, func() {
		var k int = 994532441
		var v uint8 = 160

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":994532441,"value":160}]`)
	})
}
