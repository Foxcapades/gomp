package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintByte_Put(t *testing.T) {
	Convey("TestMapUintByte.Put", t, func() {
		var k uint = 3029708689
		var v byte = 50

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintByte_Delete(t *testing.T) {
	Convey("TestMapUintByte.Delete", t, func() {
		var k uint = 940666820
		var v byte = 113

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintByte_Has(t *testing.T) {
	Convey("TestMapUintByte.Has", t, func() {
		var k uint = 2933247987
		var v byte = 154

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3011694699+1957027152), ShouldBeFalse)
	})
}

func TestMapUintByte_Get(t *testing.T) {
	Convey("TestMapUintByte.Get", t, func() {
		var k uint = 1687530262
		var v byte = 66

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2239488029 + 1923469425)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintByte_GetOpt(t *testing.T) {
	Convey("TestMapUintByte.GetOpt", t, func() {
		var k uint = 1335137337
		var v byte = 178

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(794566991 + 1356127052)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintByte_ForEach(t *testing.T) {
	Convey("TestMapUintByte.ForEach", t, func() {
		var k uint = 1119281100
		var v byte = 153
		hits := 0

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintByte_MarshalYAML(t *testing.T) {
	Convey("TestMapUintByte.MarshalYAML", t, func() {
		var k uint = 1894749305
		var v byte = 105

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintByte_ToYAML(t *testing.T) {
	Convey("TestMapUintByte.ToYAML", t, func() {
		var k uint = 3146850437
		var v byte = 197

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapUintByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintByte.PutIfNotNil", t, func() {
		var k uint = 3495815315
		var v byte = 31

		test := omap.NewMapUintByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3784506019, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 187
		So(test.PutIfNotNil(1909245885, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintByte.ReplaceIfExists", t, func() {
		var k uint = 678851252
		var v byte = 26
		var x byte = 158

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(799192880, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintByte.ReplaceOrPut", t, func() {
		var k uint = 1296787354
		var v byte = 162
		var x byte = 20

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1849700531, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_MarshalJSON(t *testing.T) {
	Convey("TestMapUintByte.MarshalJSON", t, func() {
		var k uint = 3333040018
		var v byte = 71

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3333040018,"value":71}]`)
	})
}
