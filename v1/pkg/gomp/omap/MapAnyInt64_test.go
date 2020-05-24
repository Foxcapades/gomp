package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt64_Put(t *testing.T) {
	Convey("TestMapAnyInt64.Put", t, func() {
		var k interface{} = "76548f43-9831-4e0b-94a0-9b689bff143f"
		var v int64 = 6537455606374654930

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt64_Delete(t *testing.T) {
	Convey("TestMapAnyInt64.Delete", t, func() {
		var k interface{} = "8c4f3b9b-a27e-4525-b814-cee7471e1416"
		var v int64 = 602396646982864962

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt64_Has(t *testing.T) {
	Convey("TestMapAnyInt64.Has", t, func() {
		var k interface{} = "201021da-f876-4599-a35f-e0fc59a80e6c"
		var v int64 = 3079506066043696322

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b4f4b919-954c-4993-9a90-d3375d9a8731"+"739fa0e5-aa71-4b3d-94b2-0017f332dc28"), ShouldBeFalse)
	})
}


func TestMapAnyInt64_Get(t *testing.T) {
	Convey("TestMapAnyInt64.Get", t, func() {
		var k interface{} = "41f6cf56-216a-4136-9ec1-339390a38c2e"
		var v int64 = 7624397298430798911

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("0cf7c532-a3e9-43b4-baf3-6015e4b78c41"+"fce2c5eb-fa99-4053-ad9e-fec65d9332ab")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt64_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt64.GetOpt", t, func() {
		var k interface{} = "7f5ca029-575b-4348-a8dd-5a661e072578"
		var v int64 = 3205773528237374867

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("9577d7fb-598b-4b25-9cf0-f4a01bed6df6"+"0317b27c-1538-4f0f-8a14-60536e14377a")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt64_ForEach(t *testing.T) {
	Convey("TestMapAnyInt64.ForEach", t, func() {
		var k interface{} = "018afb32-0c5e-4b5f-9f92-b633a6abfdf0"
		var v int64 = 4395877090088224857
		hits := 0

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt64.MarshalYAML", t, func() {
		var k interface{} = "d29313f4-e332-4e50-9035-3f1a6268077b"
		var v int64 = 8160842172668730632

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt64_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt64.ToYAML", t, func() {
		var k interface{} = "a772ff63-a0bc-48b0-8320-9aee2953f707"
		var v int64 = 5648666986208480712

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt64.PutIfNotNil", t, func() {
		var k interface{} = "bcc70694-32f5-477d-b741-dea402ed9eb0"
		var v int64 = 8758960434135172522

		test := omap.NewMapAnyInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("559ccfdd-2d2d-44b6-bf05-88be3454fb4a", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 5976884318317583331
		So(test.PutIfNotNil("2174f205-145b-4e89-a340-ded212954814", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceIfExists", t, func() {
		var k interface{} = "a6b6d0cd-b19c-4c52-8d14-16d7e19000e9"
		var v int64 = 8806419837506100943
		var x int64 = 4076855892784972369

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("6a0fd162-c63e-481f-8edb-4c6cbe344930", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceOrPut", t, func() {
		var k interface{} = "071bd084-4508-4886-b77f-5596e108bd5d"
		var v int64 = 6879589710393220218
		var x int64 = 6734455400291046233

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e64f0cef-3ff9-4aa4-8558-ce674f4569ec", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt64.MarshalJSON", t, func() {
		var k interface{} = "c7e0c127-fbc2-4a2b-a1eb-feded42f3e81"
		var v int64 = 5373790647116372079

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"c7e0c127-fbc2-4a2b-a1eb-feded42f3e81","value":5373790647116372079}]`)
	})
}

