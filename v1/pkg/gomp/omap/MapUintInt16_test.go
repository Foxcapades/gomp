package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt16_Put(t *testing.T) {
	Convey("TestMapUintInt16.Put", t, func() {
		var k uint = 386353645
		var v int16 = 25961

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt16_Delete(t *testing.T) {
	Convey("TestMapUintInt16.Delete", t, func() {
		var k uint = 2782196029
		var v int16 = 19234

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt16_Has(t *testing.T) {
	Convey("TestMapUintInt16.Has", t, func() {
		var k uint = 1234773051
		var v int16 = 8378

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3367135600+3552081389), ShouldBeFalse)
	})
}

func TestMapUintInt16_Get(t *testing.T) {
	Convey("TestMapUintInt16.Get", t, func() {
		var k uint = 2060841524
		var v int16 = 1849

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(4239727983 + 1915509406)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt16_GetOpt(t *testing.T) {
	Convey("TestMapUintInt16.GetOpt", t, func() {
		var k uint = 836604421
		var v int16 = 4240

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3923664239 + 2955544792)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt16_ForEach(t *testing.T) {
	Convey("TestMapUintInt16.ForEach", t, func() {
		var k uint = 1675265876
		var v int16 = 16244
		hits := 0

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt16.MarshalYAML", t, func() {
		var k uint = 3984871928
		var v int16 = 18631

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt16_ToYAML(t *testing.T) {
	Convey("TestMapUintInt16.ToYAML", t, func() {
		var k uint = 1659884287
		var v int16 = 28637

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt16.PutIfNotNil", t, func() {
		var k uint = 2999774171
		var v int16 = 10083

		test := omap.NewMapUintInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3435664095, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 32486
		So(test.PutIfNotNil(2896787057, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt16.ReplaceIfExists", t, func() {
		var k uint = 1204165592
		var v int16 = 18793
		var x int16 = 12171

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3730750627, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt16.ReplaceOrPut", t, func() {
		var k uint = 1597152395
		var v int16 = 19338
		var x int16 = 13958

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(96237385, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt16.MarshalJSON", t, func() {
		var k uint = 2752745220
		var v int16 = 3232

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2752745220,"value":3232}]`)
	})
}
