package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint16_Put(t *testing.T) {
	Convey("TestMapUintUint16.Put", t, func() {
		var k uint = 2057242093
		var v uint16 = 37389

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint16_Delete(t *testing.T) {
	Convey("TestMapUintUint16.Delete", t, func() {
		var k uint = 1291926752
		var v uint16 = 61987

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint16_Has(t *testing.T) {
	Convey("TestMapUintUint16.Has", t, func() {
		var k uint = 1171012672
		var v uint16 = 36115

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(4171856508+3752517822), ShouldBeFalse)
	})
}


func TestMapUintUint16_Get(t *testing.T) {
	Convey("TestMapUintUint16.Get", t, func() {
		var k uint = 2547351115
		var v uint16 = 46184

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1825690514+349563666)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint16_GetOpt(t *testing.T) {
	Convey("TestMapUintUint16.GetOpt", t, func() {
		var k uint = 1713142978
		var v uint16 = 10556

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1744416891+3426882900)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint16_ForEach(t *testing.T) {
	Convey("TestMapUintUint16.ForEach", t, func() {
		var k uint = 3680372035
		var v uint16 = 21587
		hits := 0

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint16.MarshalYAML", t, func() {
		var k uint = 181372392
		var v uint16 = 18594

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint16_ToYAML(t *testing.T) {
	Convey("TestMapUintUint16.ToYAML", t, func() {
		var k uint = 3302313464
		var v uint16 = 6956

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint16.PutIfNotNil", t, func() {
		var k uint = 3792960031
		var v uint16 = 7984

		test := omap.NewMapUintUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2444233459, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 50805
		So(test.PutIfNotNil(3099994918, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceIfExists", t, func() {
		var k uint = 1845700059
		var v uint16 = 25859
		var x uint16 = 37363

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3764101637, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceOrPut", t, func() {
		var k uint = 75725442
		var v uint16 = 58218
		var x uint16 = 27645

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(4285786386, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint16.MarshalJSON", t, func() {
		var k uint = 975452225
		var v uint16 = 2617

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":975452225,"value":2617}]`)
	})
}

