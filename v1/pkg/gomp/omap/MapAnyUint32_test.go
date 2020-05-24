package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint32_Put(t *testing.T) {
	Convey("TestMapAnyUint32.Put", t, func() {
		var k interface{} = "724f3160-1858-441c-8e6d-974bd8a365fd"
		var v uint32 = 224056032

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint32_Delete(t *testing.T) {
	Convey("TestMapAnyUint32.Delete", t, func() {
		var k interface{} = "34f1b7b4-c787-4e54-beff-9edb627ed4ba"
		var v uint32 = 3936405515

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint32_Has(t *testing.T) {
	Convey("TestMapAnyUint32.Has", t, func() {
		var k interface{} = "cb7690ce-d561-4b88-a41c-c5fa309fc701"
		var v uint32 = 2963493258

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("4650449c-db1f-4b6d-87d9-2414e02246ed"+"dbcbc294-c9e9-4276-b7a5-2ab911589c3d"), ShouldBeFalse)
	})
}

func TestMapAnyUint32_Get(t *testing.T) {
	Convey("TestMapAnyUint32.Get", t, func() {
		var k interface{} = "6365abe7-492f-45db-b9ec-4cabdac2fa1d"
		var v uint32 = 133257131

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("4761ed55-19f2-4eaa-b0c3-03bcb9747962" + "d7fd1e9b-0d58-44fe-8ad8-cf4d4995b198")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint32_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint32.GetOpt", t, func() {
		var k interface{} = "3bf1fb15-6036-4827-87c6-97a63e8e6010"
		var v uint32 = 1042279004

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("cf7fe527-109b-4496-9503-bf7b1dd4bbb6" + "f16a1285-b753-453d-ab98-dad29d379360")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint32_ForEach(t *testing.T) {
	Convey("TestMapAnyUint32.ForEach", t, func() {
		var k interface{} = "6dd5e462-952a-4dd5-8dd0-d06f22f566e8"
		var v uint32 = 239129289
		hits := 0

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint32.MarshalYAML", t, func() {
		var k interface{} = "33cbb619-73ac-44f2-b405-827149e28e48"
		var v uint32 = 495820761

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint32_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint32.ToYAML", t, func() {
		var k interface{} = "194204db-19df-4e55-bda2-8b6879ccfbd6"
		var v uint32 = 2610550858

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint32.PutIfNotNil", t, func() {
		var k interface{} = "65e1c22b-e08a-403f-b43f-58e74ea36b93"
		var v uint32 = 4021986664

		test := omap.NewMapAnyUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("16794064-9462-48fc-ab33-0cbc3b8870d4", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 3486718485
		So(test.PutIfNotNil("9bdaf1fa-69ae-4fa3-8a44-2e0e715c7b8a", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceIfExists", t, func() {
		var k interface{} = "e2aba6b4-2286-42de-ae0b-e77d52215fce"
		var v uint32 = 2838665237
		var x uint32 = 372334534

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("d8abf172-a1b5-4268-8de9-d6389f010144", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceOrPut", t, func() {
		var k interface{} = "36018a7c-8009-4c4a-9cdd-b2175e8c48fe"
		var v uint32 = 282252937
		var x uint32 = 1479785255

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("0747b1fc-d21e-4eea-ba8d-145794af16d1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint32.MarshalJSON", t, func() {
		var k interface{} = "fcc88823-86d7-4d11-b21f-c7dfba7abba1"
		var v uint32 = 3537048756

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"fcc88823-86d7-4d11-b21f-c7dfba7abba1","value":3537048756}]`)
	})
}
