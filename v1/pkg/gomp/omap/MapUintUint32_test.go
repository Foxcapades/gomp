package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint32_Put(t *testing.T) {
	Convey("TestMapUintUint32.Put", t, func() {
		var k uint = 4258006228
		var v uint32 = 2246381294

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint32_Delete(t *testing.T) {
	Convey("TestMapUintUint32.Delete", t, func() {
		var k uint = 2285927582
		var v uint32 = 2613104789

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint32_Has(t *testing.T) {
	Convey("TestMapUintUint32.Has", t, func() {
		var k uint = 1827533599
		var v uint32 = 1109529568

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2458777819+3311033988), ShouldBeFalse)
	})
}


func TestMapUintUint32_Get(t *testing.T) {
	Convey("TestMapUintUint32.Get", t, func() {
		var k uint = 542955213
		var v uint32 = 2984479386

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1865163621+1071597199)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint32_GetOpt(t *testing.T) {
	Convey("TestMapUintUint32.GetOpt", t, func() {
		var k uint = 927884378
		var v uint32 = 1589552023

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(975154344+3360670363)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint32_ForEach(t *testing.T) {
	Convey("TestMapUintUint32.ForEach", t, func() {
		var k uint = 3094756036
		var v uint32 = 643395765
		hits := 0

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint32.MarshalYAML", t, func() {
		var k uint = 2268381463
		var v uint32 = 2451022227

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint32_ToYAML(t *testing.T) {
	Convey("TestMapUintUint32.ToYAML", t, func() {
		var k uint = 1710566630
		var v uint32 = 776004607

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint32.PutIfNotNil", t, func() {
		var k uint = 2249812207
		var v uint32 = 2284189760

		test := omap.NewMapUintUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3738155211, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 3659309840
		So(test.PutIfNotNil(2022699524, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceIfExists", t, func() {
		var k uint = 3766732915
		var v uint32 = 353324031
		var x uint32 = 4214386944

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2712380844, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceOrPut", t, func() {
		var k uint = 1023794939
		var v uint32 = 1898330603
		var x uint32 = 3236507819

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3214239955, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint32.MarshalJSON", t, func() {
		var k uint = 3665757374
		var v uint32 = 2352073420

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3665757374,"value":2352073420}]`)
	})
}

