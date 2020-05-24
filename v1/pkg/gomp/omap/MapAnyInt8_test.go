package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt8_Put(t *testing.T) {
	Convey("TestMapAnyInt8.Put", t, func() {
		var k interface{} = "192e0e67-d896-4ed8-b6eb-e2b5c9a4b1fd"
		var v int8 = 76

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt8_Delete(t *testing.T) {
	Convey("TestMapAnyInt8.Delete", t, func() {
		var k interface{} = "23bcd30c-51b0-4008-965e-9a0374d6c685"
		var v int8 = 68

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt8_Has(t *testing.T) {
	Convey("TestMapAnyInt8.Has", t, func() {
		var k interface{} = "7f100c85-3706-4f4a-8834-575e3f40325a"
		var v int8 = 64

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("aa4a79de-d85e-4095-954e-7fe448c57d3a"+"09060730-b146-4734-95f9-8ecec64a4c21"), ShouldBeFalse)
	})
}


func TestMapAnyInt8_Get(t *testing.T) {
	Convey("TestMapAnyInt8.Get", t, func() {
		var k interface{} = "4aec2cd8-4a03-4e69-8cca-88d5a29cb624"
		var v int8 = 125

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("71b12bc1-f81f-4e7a-b1f6-e7711cecd273"+"e9ef538d-7b41-4db0-92e6-ae3c348e93b4")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt8_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt8.GetOpt", t, func() {
		var k interface{} = "cd024ba8-4273-4c65-bece-cd5596262805"
		var v int8 = 62

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("4001bf07-7b9a-49f6-b3db-222bf05fcb70"+"001c81d4-8de5-466b-8bac-187a3bc20d76")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt8_ForEach(t *testing.T) {
	Convey("TestMapAnyInt8.ForEach", t, func() {
		var k interface{} = "1b470f58-0a3f-407b-8f8e-cac2d4d0baae"
		var v int8 = 21
		hits := 0

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt8.MarshalYAML", t, func() {
		var k interface{} = "39ce149e-a101-4519-933b-62dc816e17b7"
		var v int8 = 88

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt8_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt8.ToYAML", t, func() {
		var k interface{} = "9e002402-8e16-4d00-a8ab-d71cfad018e7"
		var v int8 = 53

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt8.PutIfNotNil", t, func() {
		var k interface{} = "29e1b7d3-f233-4edc-ad7d-146c406d2280"
		var v int8 = 35

		test := omap.NewMapAnyInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("393dc88f-1454-49b7-a550-8fca7031580e", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 90
		So(test.PutIfNotNil("216d78de-7cc5-42da-abbf-6f1ba4ba9983", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceIfExists", t, func() {
		var k interface{} = "36dfad35-9bb2-4f10-8f3b-cade6ede39da"
		var v int8 = 60
		var x int8 = 25

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("858d13f4-2d11-4277-a8d9-90785c3870f4", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceOrPut", t, func() {
		var k interface{} = "0ac967c4-a8e1-49f2-a91b-beaa5c6c0850"
		var v int8 = 49
		var x int8 = 123

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("81f58de2-f15d-42b1-93b0-c2e89e66076a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt8.MarshalJSON", t, func() {
		var k interface{} = "9802df92-85ca-4bcd-afcf-dea8288a4af2"
		var v int8 = 88

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"9802df92-85ca-4bcd-afcf-dea8288a4af2","value":88}]`)
	})
}

