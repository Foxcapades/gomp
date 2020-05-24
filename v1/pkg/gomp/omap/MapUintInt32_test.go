package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt32_Put(t *testing.T) {
	Convey("TestMapUintInt32.Put", t, func() {
		var k uint = 1174440642
		var v int32 = 1673730718

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt32_Delete(t *testing.T) {
	Convey("TestMapUintInt32.Delete", t, func() {
		var k uint = 1418362930
		var v int32 = 1428037979

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt32_Has(t *testing.T) {
	Convey("TestMapUintInt32.Has", t, func() {
		var k uint = 2962167556
		var v int32 = 1442290160

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3909056191+4085585944), ShouldBeFalse)
	})
}


func TestMapUintInt32_Get(t *testing.T) {
	Convey("TestMapUintInt32.Get", t, func() {
		var k uint = 1234463336
		var v int32 = 2081217512

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3904169855+1506998977)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt32_GetOpt(t *testing.T) {
	Convey("TestMapUintInt32.GetOpt", t, func() {
		var k uint = 3989810194
		var v int32 = 526677444

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1016336107+4215759344)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt32_ForEach(t *testing.T) {
	Convey("TestMapUintInt32.ForEach", t, func() {
		var k uint = 825926133
		var v int32 = 76459544
		hits := 0

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt32.MarshalYAML", t, func() {
		var k uint = 3869306180
		var v int32 = 407792047

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt32_ToYAML(t *testing.T) {
	Convey("TestMapUintInt32.ToYAML", t, func() {
		var k uint = 1098877849
		var v int32 = 2092369976

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt32.PutIfNotNil", t, func() {
		var k uint = 1037472946
		var v int32 = 1283479932

		test := omap.NewMapUintInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1018858535, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 175061601
		So(test.PutIfNotNil(2905678306, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt32.ReplaceIfExists", t, func() {
		var k uint = 35560408
		var v int32 = 1032837294
		var x int32 = 1908829512

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1454838632, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt32.ReplaceOrPut", t, func() {
		var k uint = 3681292270
		var v int32 = 157381299
		var x int32 = 582204601

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2914871163, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt32.MarshalJSON", t, func() {
		var k uint = 1933457187
		var v int32 = 570936218

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1933457187,"value":570936218}]`)
	})
}

