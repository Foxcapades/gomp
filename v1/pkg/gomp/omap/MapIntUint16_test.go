package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint16_Put(t *testing.T) {
	Convey("TestMapIntUint16.Put", t, func() {
		var k int = 1612398902
		var v uint16 = 61845

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint16_Delete(t *testing.T) {
	Convey("TestMapIntUint16.Delete", t, func() {
		var k int = 1963679008
		var v uint16 = 10304

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint16_Has(t *testing.T) {
	Convey("TestMapIntUint16.Has", t, func() {
		var k int = 1092194227
		var v uint16 = 28052

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(810079496+1666746545), ShouldBeFalse)
	})
}


func TestMapIntUint16_Get(t *testing.T) {
	Convey("TestMapIntUint16.Get", t, func() {
		var k int = 68475621
		var v uint16 = 22068

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(500507583 + 1105722243)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint16_GetOpt(t *testing.T) {
	Convey("TestMapIntUint16.GetOpt", t, func() {
		var k int = 132283283
		var v uint16 = 58696

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1595955861 + 733230335)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint16_ForEach(t *testing.T) {
	Convey("TestMapIntUint16.ForEach", t, func() {
		var k int = 1307260433
		var v uint16 = 46777
		hits := 0

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint16.MarshalYAML", t, func() {
		var k int = 1487876453
		var v uint16 = 59716

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint16_ToYAML(t *testing.T) {
	Convey("TestMapIntUint16.ToYAML", t, func() {
		var k int = 704060435
		var v uint16 = 23226

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint16.PutIfNotNil", t, func() {
		var k int = 34059443
		var v uint16 = 2828

		test := omap.NewMapIntUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2093411012, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 16267
		So(test.PutIfNotNil(781161249, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint16.ReplaceIfExists", t, func() {
		var k int = 687032859
		var v uint16 = 56694
		var x uint16 = 61458

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(539759906, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint16.ReplaceOrPut", t, func() {
		var k int = 110876285
		var v uint16 = 16359
		var x uint16 = 38448

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2115512909, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint16.MarshalJSON", t, func() {
		var k int = 677427665
		var v uint16 = 55472

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":677427665,"value":55472}]`)
	})
}
