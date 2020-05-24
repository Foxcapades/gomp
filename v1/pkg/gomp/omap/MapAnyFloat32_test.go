package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyFloat32_Put(t *testing.T) {
	Convey("TestMapAnyFloat32.Put", t, func() {
		var k interface{} = "500b093e-5465-499d-82c4-92233f622c04"
		var v float32 = 0.835

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat32_Delete(t *testing.T) {
	Convey("TestMapAnyFloat32.Delete", t, func() {
		var k interface{} = "5833578b-7e0d-4e26-b457-ff1b06a403e9"
		var v float32 = 0.868

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat32_Has(t *testing.T) {
	Convey("TestMapAnyFloat32.Has", t, func() {
		var k interface{} = "3fef6fc3-e767-4543-8e9d-91bbbd07c656"
		var v float32 = 0.403

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b5255814-20f5-4100-8209-909a73cd1668"+"031f66c3-6854-496c-a582-8e781b9dc9b8"), ShouldBeFalse)
	})
}


func TestMapAnyFloat32_Get(t *testing.T) {
	Convey("TestMapAnyFloat32.Get", t, func() {
		var k interface{} = "ac6e0be0-90c5-437a-bff3-843d83ca0331"
		var v float32 = 0.786

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("2c252ea4-3707-4fcf-8238-126a69e37dba" + "30b7c9ba-e831-4f59-b875-4a5a4c601162")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat32_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat32.GetOpt", t, func() {
		var k interface{} = "57538774-b84f-44ee-94ad-6e02e532f64a"
		var v float32 = 0.242

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("68e1a866-a887-4e68-b192-c04e11f35d17" + "60b0a18c-a5f4-4d4d-84bf-147b0beda256")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat32_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat32.ForEach", t, func() {
		var k interface{} = "95649200-03eb-4415-b54c-c26e9ee0caa9"
		var v float32 = 0.633
		hits := 0

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyFloat32.MarshalYAML", t, func() {
		var k interface{} = "4495d369-a155-449b-84c4-8cb674aad689"
		var v float32 = 0.086

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyFloat32_ToYAML(t *testing.T) {
	Convey("TestMapAnyFloat32.ToYAML", t, func() {
		var k interface{} = "5cfaa429-b09f-4444-9e8b-886e228c0ca1"
		var v float32 = 0.524

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyFloat32.PutIfNotNil", t, func() {
		var k interface{} = "1249d750-40ac-4742-8a40-aad46b91ceb3"
		var v float32 = 0.474

		test := omap.NewMapAnyFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("e7819ff8-9747-478f-bb28-3558ca06bc00", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.258
		So(test.PutIfNotNil("261fde7f-ebd4-414e-b22a-9b9380c3a8b0", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceIfExists", t, func() {
		var k interface{} = "21dc9b0f-3266-475b-8395-889c93e54aca"
		var v float32 = 0.748
		var x float32 = 0.373

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ccbaf3a7-cdf7-452a-9501-3fc9c0dc9c8f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceOrPut", t, func() {
		var k interface{} = "9ee69e89-ee4f-4a36-bc41-612ced685a91"
		var v float32 = 0.262
		var x float32 = 0.029

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("95114d97-306e-4737-b3ab-1830e06d3c25", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat32.MarshalJSON", t, func() {
		var k interface{} = "b40d6c11-0f45-4b9f-89bf-ad7eda9444ff"
		var v float32 = 0.320

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"b40d6c11-0f45-4b9f-89bf-ad7eda9444ff","value":0.32}]`)
	})
}

