package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint32_Put(t *testing.T) {
	Convey("TestMapStringUint32.Put", t, func() {
		var k string = "4ae07575-d70e-4537-a285-c0008fd14088"
		var v uint32 = 1124594788

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint32_Delete(t *testing.T) {
	Convey("TestMapStringUint32.Delete", t, func() {
		var k string = "5873f473-2705-4530-bde1-4b5936d2bcb8"
		var v uint32 = 497135233

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint32_Has(t *testing.T) {
	Convey("TestMapStringUint32.Has", t, func() {
		var k string = "c8f52d67-e537-4773-8a7e-ab111eee8800"
		var v uint32 = 3186169603

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("610143b2-671a-4002-9207-7015f3d76328"+"7c19affe-83bf-4244-bc0d-5db9ea905989"), ShouldBeFalse)
	})
}

func TestMapStringUint32_Get(t *testing.T) {
	Convey("TestMapStringUint32.Get", t, func() {
		var k string = "d5868cda-a6d0-496a-b767-a2990ac6b8dd"
		var v uint32 = 1498468462

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("ea6dfeab-59bb-4d95-b383-2d2a9d664821" + "233aadc1-712e-4484-9a11-85e754aa7505")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint32_GetOpt(t *testing.T) {
	Convey("TestMapStringUint32.GetOpt", t, func() {
		var k string = "8d2b357f-ba3b-471d-bd64-dc24a50006b2"
		var v uint32 = 3194327087

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("091b9976-0128-48ca-968c-4c607030b598" + "155d23e5-eba9-4c11-8e91-653c7f9b5cb9")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint32_ForEach(t *testing.T) {
	Convey("TestMapStringUint32.ForEach", t, func() {
		var k string = "a7a7d514-bd30-477f-96c6-b14fc84b5d6c"
		var v uint32 = 620627179
		hits := 0

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint32.MarshalYAML", t, func() {
		var k string = "9fb5ca7c-8c26-438f-96ee-6bc4dbacc85e"
		var v uint32 = 3558290114

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint32_ToYAML(t *testing.T) {
	Convey("TestMapStringUint32.ToYAML", t, func() {
		var k string = "32234149-0050-410a-8100-4d5a178f20c7"
		var v uint32 = 657301986

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint32.PutIfNotNil", t, func() {
		var k string = "74bcd73e-ed41-4cf1-ac10-63fe52dd84de"
		var v uint32 = 1894025954

		test := omap.NewMapStringUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("06db1aed-db62-4c80-8db5-02f4bfce7c4e", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 1344769443
		So(test.PutIfNotNil("bc21dac0-17c0-46e8-a44a-a1e1901cd200", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceIfExists", t, func() {
		var k string = "b59d397c-bf92-42e2-9998-b841347a2850"
		var v uint32 = 3844988033
		var x uint32 = 965956212

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("072163b0-346b-4fbf-8a52-2bb5762cbdd4", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceOrPut", t, func() {
		var k string = "560ee033-ad3d-4a51-b539-1f50f41665ac"
		var v uint32 = 559332698
		var x uint32 = 2742015551

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("4ef350b4-c125-45cc-a5ac-7d0de2ecafa6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint32.MarshalJSON", t, func() {
		var k string = "3a20deb8-a312-48af-b8ca-828994d49a17"
		var v uint32 = 1249677249

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"3a20deb8-a312-48af-b8ca-828994d49a17","value":1249677249}]`)
	})
}
