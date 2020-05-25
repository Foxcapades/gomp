package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt32_Put(t *testing.T) {
	Convey("TestMapAnyInt32.Put", t, func() {
		var k interface{} = "cf764f6f-3751-434e-ad86-3c941d17edd3"
		var v int32 = 421279821

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt32_Delete(t *testing.T) {
	Convey("TestMapAnyInt32.Delete", t, func() {
		var k interface{} = "0e62c8e0-adcc-473a-bdec-e256c55bda65"
		var v int32 = 1250371403

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt32_Has(t *testing.T) {
	Convey("TestMapAnyInt32.Has", t, func() {
		var k interface{} = "c68c0138-e2bd-4484-986f-d8148dee2403"
		var v int32 = 1134587119

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("11f99c15-7c85-4252-9331-cec0382ba592"+"97ff28f6-26a6-4377-a41a-93b188b6fec0"), ShouldBeFalse)
	})
}

func TestMapAnyInt32_Get(t *testing.T) {
	Convey("TestMapAnyInt32.Get", t, func() {
		var k interface{} = "6d615e60-e8d9-4066-af9e-12bd1e6dc99b"
		var v int32 = 1286664729

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("c6207e95-01a6-42f3-aad3-3eeb03c0d2d5" + "71ac438f-abee-48b0-88fe-fdb3f06cd331")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt32_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt32.GetOpt", t, func() {
		var k interface{} = "3737646b-d1bd-40d1-bcce-7e69bc353999"
		var v int32 = 933919034

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("1d2f1b5e-2dc7-470e-9893-6099b6b34621" + "a8613a80-9e8a-45cb-91c4-a3f0b5f21d9f")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt32_ForEach(t *testing.T) {
	Convey("TestMapAnyInt32.ForEach", t, func() {
		var k interface{} = "bb769c67-da58-4dc3-9c22-5b98fa897950"
		var v int32 = 232989227
		hits := 0

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalYAML", t, func() {
		var k interface{} = "7f6e55c4-7829-4707-9f93-9a6aa23ec34e"
		var v int32 = 182095852

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt32_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt32.ToYAML", t, func() {
		var k interface{} = "7942a76a-9baf-4f74-b9b8-d29ab1e3dcbb"
		var v int32 = 548764786

		test := omap.NewMapAnyInt32(1)

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

func TestMapAnyInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt32.PutIfNotNil", t, func() {
		var k interface{} = "95b93cef-bbe8-422a-a5cc-f6038cf72cc3"
		var v int32 = 1299480230

		test := omap.NewMapAnyInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("89c0a7bb-9efd-455e-a8d2-faf0986194f3", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 2128995199
		So(test.PutIfNotNil("ed2a2fd2-4687-43b1-b6bb-571ca49ddee5", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceIfExists", t, func() {
		var k interface{} = "36670514-aa8e-4241-b70e-4754ec7512a4"
		var v int32 = 748320682
		var x int32 = 316342045

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("58b15852-08ec-45c9-bb8b-e3f871c93c30", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceOrPut", t, func() {
		var k interface{} = "fa62e84c-e6f9-4c95-871b-ba3283851999"
		var v int32 = 828222138
		var x int32 = 485435948

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("724983d0-bd53-4e46-b909-348167e7e8f1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalJSON", t, func() {
		var k interface{} = "6849262b-e187-475a-b31b-a26b263b4176"
		var v int32 = 439802428

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"6849262b-e187-475a-b31b-a26b263b4176","value":439802428}]`)
	})
}
