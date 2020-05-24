package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAny_Put(t *testing.T) {
	Convey("TestMapAny.Put", t, func() {
		var k interface{} = "2ac6c1f3-0225-4247-9905-9ef85ea4ba60"
		var v interface{} = "2e4550c9-7ce8-4637-92ec-8ec8086ffc70"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAny_Delete(t *testing.T) {
	Convey("TestMapAny.Delete", t, func() {
		var k interface{} = "8bc46fd9-61dd-424d-a26a-a600442c1a20"
		var v interface{} = "c8e31197-f376-445b-b1c5-f6a4288c8ccf"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAny_Has(t *testing.T) {
	Convey("TestMapAny.Has", t, func() {
		var k interface{} = "6292e516-17fc-48f1-907c-52dd6f384ce7"
		var v interface{} = "fa104fd5-08e6-4dbf-ac68-07f898c56406"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("6a1dff40-6396-4591-ac6c-9dfc94401004"+"8cd506dd-9e56-47ea-a3c5-22e0b3a243fc"), ShouldBeFalse)
	})
}

func TestMapAny_Get(t *testing.T) {
	Convey("TestMapAny.Get", t, func() {
		var k interface{} = "1427255e-ca1c-41b9-8610-2260e29d593e"
		var v interface{} = "cbd3e982-589b-49d3-8334-d468bcb89185"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("74515c56-9420-4bbc-a057-45fa2c4ecd8f" + "5b6fbcc7-98fa-40ed-894d-6684c0d857ba")
		So(b, ShouldBeFalse)
	})
}

func TestMapAny_GetOpt(t *testing.T) {
	Convey("TestMapAny.GetOpt", t, func() {
		var k interface{} = "50a8e9eb-20f0-4993-9476-033703f4d526"
		var v interface{} = "b555ab58-67ae-4716-93a6-e214459e658a"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("782ed60d-3c40-4618-b537-72ad61afefb6" + "3b6e1f24-53e7-45ae-87d7-1ae85077fb42")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAny_ForEach(t *testing.T) {
	Convey("TestMapAny.ForEach", t, func() {
		var k interface{} = "1ea37bc8-dc47-4de7-996a-e81bc7e16059"
		var v interface{} = "38f51a73-a543-4c31-b129-b84a016b20af"
		hits := 0

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAny_MarshalYAML(t *testing.T) {
	Convey("TestMapAny.MarshalYAML", t, func() {
		var k interface{} = "af40e7cf-03b1-4ffd-8563-68fa1fd87dc1"
		var v interface{} = "2453b615-0906-4a6b-8034-62fa2a00b8b2"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAny_ToYAML(t *testing.T) {
	Convey("TestMapAny.ToYAML", t, func() {
		var k interface{} = "0b78d214-beee-495f-bc5e-09412c0c1eca"
		var v interface{} = "f4ca3d69-aa26-4906-91b1-db7cfa059bc2"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapAny.PutIfNotNil", t, func() {
		var k interface{} = "54792609-91d2-43a6-acba-ae688f9118d7"
		var v interface{} = "a4c155af-e300-4198-b262-525a06c9d3b6"

		test := omap.NewMapAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("22423db6-be6e-471b-814b-9d32b8832652", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "a842a9fa-8fa0-4fe7-879f-0a4f03a15532"
		So(test.PutIfNotNil("6747ee71-87bc-4e85-ab87-cd8cd7b14c8b", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAny.ReplaceIfExists", t, func() {
		var k interface{} = "69138b8f-f438-47a2-b538-ddd66577c323"
		var v interface{} = "0ffea2be-5fad-4de2-8698-0f1b270eab51"
		var x interface{} = "da476aa0-0c72-4450-9e20-e21ad14c6cbc"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("f885edc1-d69c-42a6-9b18-b9cf384400f7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAny.ReplaceOrPut", t, func() {
		var k interface{} = "7fbdc19d-b5f8-4480-8f93-070ca956ac2a"
		var v interface{} = "11dacb67-ab7c-440e-9618-ea71a4e7cc69"
		var x interface{} = "6fd9e013-f7f9-40a2-805d-448e42701dc8"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("7d09e94d-d7a5-4775-85ae-8d1afccbc4af", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_MarshalJSON(t *testing.T) {
	Convey("TestMapAny.MarshalJSON", t, func() {
		var k interface{} = "d2ed7eeb-ad65-45c8-b57d-f1a249de8b1e"
		var v interface{} = "60b5e92e-14b0-49b2-b91d-5e1fe4040a42"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"d2ed7eeb-ad65-45c8-b57d-f1a249de8b1e","value":"60b5e92e-14b0-49b2-b91d-5e1fe4040a42"}]`)
	})
}
