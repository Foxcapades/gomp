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
		var k uint = 3422360947
		var v int32 = 968531173

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt32_Delete(t *testing.T) {
	Convey("TestMapUintInt32.Delete", t, func() {
		var k uint = 2196009564
		var v int32 = 332700128

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt32_Has(t *testing.T) {
	Convey("TestMapUintInt32.Has", t, func() {
		var k uint = 678448446
		var v int32 = 84163968

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3451361772+1846701015), ShouldBeFalse)
	})
}

func TestMapUintInt32_Get(t *testing.T) {
	Convey("TestMapUintInt32.Get", t, func() {
		var k uint = 1812600105
		var v int32 = 1288341685

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(3952392295 + 2004397187)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt32_GetOpt(t *testing.T) {
	Convey("TestMapUintInt32.GetOpt", t, func() {
		var k uint = 3856952138
		var v int32 = 1873173870

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1303648152 + 2823804108)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt32_ForEach(t *testing.T) {
	Convey("TestMapUintInt32.ForEach", t, func() {
		var k uint = 2295918914
		var v int32 = 1336819822
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
		var k uint = 538425502
		var v int32 = 429063278

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
		var k uint = 824371698
		var v int32 = 1409478930

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
		var k uint = 1792649524
		var v int32 = 680199525

		test := omap.NewMapUintInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1055458051, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 1076044941
		So(test.PutIfNotNil(3047932358, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt32.ReplaceIfExists", t, func() {
		var k uint = 2155080514
		var v int32 = 2000856957
		var x int32 = 1584747199

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(4249371994, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt32.ReplaceOrPut", t, func() {
		var k uint = 1211436626
		var v int32 = 1389386095
		var x int32 = 405847837

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(394044116, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt32.MarshalJSON", t, func() {
		var k uint = 2526212365
		var v int32 = 1501227685

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2526212365,"value":1501227685}]`)
	})
}
