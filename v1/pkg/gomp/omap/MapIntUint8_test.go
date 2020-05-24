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
		var k int = 146253649
		var v uint8 = 58

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint8_Delete(t *testing.T) {
	Convey("TestMapIntUint8.Delete", t, func() {
		var k int = 795306952
		var v uint8 = 211

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint8_Has(t *testing.T) {
	Convey("TestMapIntUint8.Has", t, func() {
		var k int = 1623780801
		var v uint8 = 160

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2112084220+1078971926), ShouldBeFalse)
	})
}


func TestMapIntUint8_Get(t *testing.T) {
	Convey("TestMapIntUint8.Get", t, func() {
		var k int = 1338136152
		var v uint8 = 192

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1608483996+875010197)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint8_GetOpt(t *testing.T) {
	Convey("TestMapIntUint8.GetOpt", t, func() {
		var k int = 1902539118
		var v uint8 = 197

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(764714240+1562264348)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint8_ForEach(t *testing.T) {
	Convey("TestMapIntUint8.ForEach", t, func() {
		var k int = 2025791331
		var v uint8 = 211
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
		var k int = 445365630
		var v uint8 = 183

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
		var k int = 500757473
		var v uint8 = 91

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
		var k int = 847179391
		var v uint8 = 203

		test := omap.NewMapIntUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(111926015, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 190
		So(test.PutIfNotNil(1658000569, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceIfExists", t, func() {
		var k int = 551896985
		var v uint8 = 29
		var x uint8 = 19

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1792851424, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceOrPut", t, func() {
		var k int = 1537378819
		var v uint8 = 168
		var x uint8 = 37

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1033450102, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint8.MarshalJSON", t, func() {
		var k int = 892836190
		var v uint8 = 198

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":892836190,"value":198}]`)
	})
}

