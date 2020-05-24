package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt16_Put(t *testing.T) {
	Convey("TestMapAnyInt16.Put", t, func() {
		var k interface{} = "9ea4cd0e-26b7-482b-a7b6-2f49cab32b43"
		var v int16 = 25557

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt16_Delete(t *testing.T) {
	Convey("TestMapAnyInt16.Delete", t, func() {
		var k interface{} = "910118c0-5563-4de5-b798-554d68ede56a"
		var v int16 = 24926

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt16_Has(t *testing.T) {
	Convey("TestMapAnyInt16.Has", t, func() {
		var k interface{} = "97f4e20d-8c97-4977-aa32-f4c1331ad94d"
		var v int16 = 3476

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d8d3c819-02c1-465a-81a4-0f3e3f33128f"+"6d04e031-af26-4a45-93ae-d7db90692b92"), ShouldBeFalse)
	})
}

func TestMapAnyInt16_Get(t *testing.T) {
	Convey("TestMapAnyInt16.Get", t, func() {
		var k interface{} = "07c6ab5f-5de4-4e70-bab5-0f416e1d3ea7"
		var v int16 = 1438

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("4742e24b-03d5-499d-9c5f-fad9a968a272" + "1ddf4e94-9db9-4930-840f-ca145eeb2056")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt16_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt16.GetOpt", t, func() {
		var k interface{} = "f41db994-6a0a-4d62-8c6f-3add8b835de6"
		var v int16 = 22823

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("bd195dff-02e6-4f9a-934d-82a6bbbd2d2c" + "cb4d5da8-d3d1-4c85-80e8-9e60dd198f19")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt16_ForEach(t *testing.T) {
	Convey("TestMapAnyInt16.ForEach", t, func() {
		var k interface{} = "4b2bcec9-c9df-41a8-9785-22d6d84a2157"
		var v int16 = 8535
		hits := 0

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt16.MarshalYAML", t, func() {
		var k interface{} = "8b8439b9-f455-4b47-9812-7bd875d99d3b"
		var v int16 = 22280

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt16_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt16.ToYAML", t, func() {
		var k interface{} = "36f0a350-fabf-452b-84d5-f9bbae62f192"
		var v int16 = 27573

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt16.PutIfNotNil", t, func() {
		var k interface{} = "33e4d7c4-b2a8-4c88-a4fb-7ca3736d8da0"
		var v int16 = 2147

		test := omap.NewMapAnyInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("adb03c9c-7476-4a64-a1a0-02eb0c0a81c9", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 22316
		So(test.PutIfNotNil("798d0de2-831b-488f-8fe1-c090d33077cc", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceIfExists", t, func() {
		var k interface{} = "24ab2f9e-2f08-48eb-a794-b159a1b5e683"
		var v int16 = 11889
		var x int16 = 31411

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("d1cf42f1-8f00-4edd-a0b4-80849ddcbbec", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceOrPut", t, func() {
		var k interface{} = "afcbb4a4-feb4-4195-9cd6-f62ba69104dd"
		var v int16 = 21418
		var x int16 = 21726

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("fdd4f7a5-b17f-4641-872a-1389b0a6b0d8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt16.MarshalJSON", t, func() {
		var k interface{} = "e904b85e-ec4e-4a5a-bd4c-8b50915037aa"
		var v int16 = 16496

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e904b85e-ec4e-4a5a-bd4c-8b50915037aa","value":16496}]`)
	})
}
