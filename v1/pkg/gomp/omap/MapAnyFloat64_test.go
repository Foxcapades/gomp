package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyFloat64_Put(t *testing.T) {
	Convey("TestMapAnyFloat64.Put", t, func() {
		var k interface{} = "657c614c-642b-4212-9d91-13ea55cdcda3"
		var v float64 = 0.263

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat64_Delete(t *testing.T) {
	Convey("TestMapAnyFloat64.Delete", t, func() {
		var k interface{} = "878e697a-d766-4668-814d-5d044bdab13f"
		var v float64 = 0.627

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat64_Has(t *testing.T) {
	Convey("TestMapAnyFloat64.Has", t, func() {
		var k interface{} = "b87bd0c2-308f-4864-a58e-be67c67ac72a"
		var v float64 = 0.430

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("0d1b9b34-df71-45e4-a6a9-d0887ccd0f3c"+"8ea8b268-f777-4630-af5d-9aed4cdaccd8"), ShouldBeFalse)
	})
}


func TestMapAnyFloat64_Get(t *testing.T) {
	Convey("TestMapAnyFloat64.Get", t, func() {
		var k interface{} = "23cbcc57-0cb9-41cb-b604-ceba1ccb7262"
		var v float64 = 0.124

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("f0827300-402e-4967-bc96-838cb3db4b25"+"a2811c6c-0ebd-48a2-b8a5-0295f8e46c70")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat64_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat64.GetOpt", t, func() {
		var k interface{} = "493f3a55-0938-4d99-9520-5e497d24d95c"
		var v float64 = 0.215

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("66573346-f50c-4954-8ea4-6f815ddebda7"+"0ca717df-8c62-4d83-b918-770096fa66dc")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat64_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat64.ForEach", t, func() {
		var k interface{} = "ec6c07bd-0259-4e12-97ec-bc44d42eee11"
		var v float64 = 0.149
		hits := 0

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyFloat64.MarshalYAML", t, func() {
		var k interface{} = "950a2a6d-7dce-4300-8fd9-e9f0d3bffced"
		var v float64 = 0.672

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyFloat64_ToYAML(t *testing.T) {
	Convey("TestMapAnyFloat64.ToYAML", t, func() {
		var k interface{} = "05adad46-f68c-465b-ab2b-548402a27395"
		var v float64 = 0.989

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyFloat64.PutIfNotNil", t, func() {
		var k interface{} = "251921f9-c2b8-4d88-b0f3-25b1c6baf6ef"
		var v float64 = 0.967

		test := omap.NewMapAnyFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("42f5235b-e100-4347-bd11-47300aaa6bf0", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.320
		So(test.PutIfNotNil("6e5c05ac-9ef5-4ea0-aee7-28232a04a8dd", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceIfExists", t, func() {
		var k interface{} = "af01dc17-7d12-4cfa-9cb7-b3af9d462e93"
		var v float64 = 0.318
		var x float64 = 0.422

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("c8df2054-bc6b-4d50-8ec4-bedc5b85b95c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceOrPut", t, func() {
		var k interface{} = "52e8f54c-ea32-4f17-ba36-370bd0de1080"
		var v float64 = 0.251
		var x float64 = 0.641

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("2fcdb386-87c0-4baa-b98b-7df8600a9d1f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat64.MarshalJSON", t, func() {
		var k interface{} = "0d3de874-9e98-445e-ab48-060a5db182c5"
		var v float64 = 0.237

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"0d3de874-9e98-445e-ab48-060a5db182c5","value":0.237}]`)
	})
}

