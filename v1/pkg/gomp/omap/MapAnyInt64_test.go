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
		var k interface{} = "bfa07347-38dc-4472-8318-e14a941219b1"
		var v int64 = 1148387158466345697

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt64_Delete(t *testing.T) {
	Convey("TestMapAnyInt64.Delete", t, func() {
		var k interface{} = "e6041d0a-3e0a-4488-b2fc-1cc11f018da7"
		var v int64 = 7065569711014336821

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt64_Has(t *testing.T) {
	Convey("TestMapAnyInt64.Has", t, func() {
		var k interface{} = "66d753dc-3efc-4f3b-9775-fcfed7e577a8"
		var v int64 = 7956931224912962127

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("dc08eaa9-488a-4d2d-86e6-a3cf87353f6c"+"db7b28fd-8c89-4e57-a7d8-ac451273b276"), ShouldBeFalse)
	})
}

func TestMapAnyInt64_Get(t *testing.T) {
	Convey("TestMapAnyInt64.Get", t, func() {
		var k interface{} = "3f353391-2586-466a-af73-10c95c47c5fa"
		var v int64 = 5305857148380495285

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("c8ca4974-0f93-4d7d-906c-269c7e4521ed" + "a7a0f4d0-38ee-4de4-a1a5-28965ac4a6d1")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt64_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt64.GetOpt", t, func() {
		var k interface{} = "7c743ea7-519e-488d-8f19-12313b9f7be0"
		var v int64 = 6710708553367291000

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("f2bc6034-731f-44d7-a6b5-b349b424012f" + "5d021fb5-97ca-498a-b1d9-490ae04ad8a6")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt64_ForEach(t *testing.T) {
	Convey("TestMapAnyInt64.ForEach", t, func() {
		var k interface{} = "24dd0768-a267-4d5b-b702-54db126f2f5b"
		var v int64 = 2947437230677987141
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
		var k interface{} = "eb1db00c-3876-4b44-87b3-4df7c862736d"
		var v int64 = 2419935998552149543

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
		var k interface{} = "3252b6b0-b16b-4413-9848-7c497e68160c"
		var v int64 = 8724750492867484316

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
		var k interface{} = "fdd6a6e2-7291-4627-b143-3449888ddb20"
		var v int64 = 7438111455167137810

		test := omap.NewMapAnyInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("1a7abc5b-8d36-4216-bd2a-fa562f8b1141", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 3351238321930029515
		So(test.PutIfNotNil("a4299ac2-7a2b-4b55-a455-8d226211ca1c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceIfExists", t, func() {
		var k interface{} = "9f3f3728-8597-4073-985b-295d3946d167"
		var v int64 = 7227028476385554843
		var x int64 = 5804525085821645716

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("5ceb2daa-fe07-4203-bf7a-a64be229465c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceOrPut", t, func() {
		var k interface{} = "47096fe0-506b-4cce-957d-e94cce13318a"
		var v int64 = 1120981124233653117
		var x int64 = 4006534572413056105

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("b6ddf86d-02df-41b1-a10b-6ce758b940b5", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt64.MarshalJSON", t, func() {
		var k interface{} = "b124d24a-61fe-4fd3-b083-7922f53ff4f4"
		var v int64 = 63695039602833916

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"b124d24a-61fe-4fd3-b083-7922f53ff4f4","value":63695039602833916}]`)
	})
}
