package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint8_Put(t *testing.T) {
	Convey("TestMapUintUint8.Put", t, func() {
		var k uint = 3422805541
		var v uint8 = 34

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint8_Delete(t *testing.T) {
	Convey("TestMapUintUint8.Delete", t, func() {
		var k uint = 3343008761
		var v uint8 = 222

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint8_Has(t *testing.T) {
	Convey("TestMapUintUint8.Has", t, func() {
		var k uint = 4058262218
		var v uint8 = 239

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(258237584+212825128), ShouldBeFalse)
	})
}


func TestMapUintUint8_Get(t *testing.T) {
	Convey("TestMapUintUint8.Get", t, func() {
		var k uint = 640705041
		var v uint8 = 37

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3198441173 + 3823363631)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint8_GetOpt(t *testing.T) {
	Convey("TestMapUintUint8.GetOpt", t, func() {
		var k uint = 3741819477
		var v uint8 = 109

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3602338646 + 308905796)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint8_ForEach(t *testing.T) {
	Convey("TestMapUintUint8.ForEach", t, func() {
		var k uint = 550805851
		var v uint8 = 239
		hits := 0

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint8.MarshalYAML", t, func() {
		var k uint = 3293759288
		var v uint8 = 15

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint8_ToYAML(t *testing.T) {
	Convey("TestMapUintUint8.ToYAML", t, func() {
		var k uint = 3596441528
		var v uint8 = 138

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint8.PutIfNotNil", t, func() {
		var k uint = 2145219940
		var v uint8 = 174

		test := omap.NewMapUintUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1855955478, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 246
		So(test.PutIfNotNil(717952182, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint8.ReplaceIfExists", t, func() {
		var k uint = 4017795287
		var v uint8 = 133
		var x uint8 = 89

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2285058660, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint8.ReplaceOrPut", t, func() {
		var k uint = 3782440669
		var v uint8 = 28
		var x uint8 = 106

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2356419863, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint8.MarshalJSON", t, func() {
		var k uint = 2223230202
		var v uint8 = 142

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2223230202,"value":142}]`)
	})
}
