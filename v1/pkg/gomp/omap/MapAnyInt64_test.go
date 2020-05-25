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
		var k interface{} = "8b082eb7-e29b-4f35-9db8-812543dfc325"
		var v int64 = 10058104035308446

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt64_Delete(t *testing.T) {
	Convey("TestMapAnyInt64.Delete", t, func() {
		var k interface{} = "410efc22-acb4-4116-a329-28ef967a306b"
		var v int64 = 779788048327252753

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt64_Has(t *testing.T) {
	Convey("TestMapAnyInt64.Has", t, func() {
		var k interface{} = "1779cd98-fcfa-4558-8974-9e3219be6f40"
		var v int64 = 5826046052943173683

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5dac43dd-2290-43e2-aa22-04ef1ceba5a9"+"7be247db-10ed-4eed-a84c-b1fe3330d7bf"), ShouldBeFalse)
	})
}

func TestMapAnyInt64_Get(t *testing.T) {
	Convey("TestMapAnyInt64.Get", t, func() {
		var k interface{} = "9b2053fd-ff98-405b-a69f-15244fcfc6bd"
		var v int64 = 8524366549505365042

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("39c2fa10-ff98-4c08-8188-30b9fb3a5822" + "df95f56e-d25d-4f67-80d5-53241d263abf")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt64_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt64.GetOpt", t, func() {
		var k interface{} = "075f0f25-8ddb-4743-8408-e23d0324aed1"
		var v int64 = 5911882964768353185

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("98359170-5723-4dba-b2d8-744a6e3d8d6e" + "54aac36f-d061-4b7f-937f-e2117ac1c796")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt64_ForEach(t *testing.T) {
	Convey("TestMapAnyInt64.ForEach", t, func() {
		var k interface{} = "8e0493d1-f804-4fba-b3eb-6452a640ce7e"
		var v int64 = 9205154844213039955
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
		var k interface{} = "c3cb5541-5483-4383-90a0-d2d13f51dc6c"
		var v int64 = 5391069475297244590

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
		var k interface{} = "6662b96f-948e-4aba-ba71-b0225ce8cb12"
		var v int64 = 4603280481667078437

		test := omap.NewMapAnyInt64(1)

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

func TestMapAnyInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt64.PutIfNotNil", t, func() {
		var k interface{} = "f9aa6cc7-4fa9-403e-b90c-06a884f17f21"
		var v int64 = 853105043853304108

		test := omap.NewMapAnyInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("9b7ef5a9-0a1b-439a-bc7c-378b44031abf", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 5310589983849718567
		So(test.PutIfNotNil("e9fd8504-4f3c-4f64-b87d-33af4c262349", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceIfExists", t, func() {
		var k interface{} = "09c1620b-3f1f-483e-945b-5594dff0a508"
		var v int64 = 3374451856395199873
		var x int64 = 2910658844100545251

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("17d17c05-04fa-4cba-9afd-1299a2fddc9f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceOrPut", t, func() {
		var k interface{} = "a2daa3e2-1dca-42e0-a307-a6ef82ee764d"
		var v int64 = 8019628854235970123
		var x int64 = 7607616167710178603

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("3a977522-0b2f-4e90-b1d1-12aeaf871843", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt64.MarshalJSON", t, func() {
		var k interface{} = "189eed68-75c8-424e-aaeb-bab12cb7ed9a"
		var v int64 = 5601147218404810302

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"189eed68-75c8-424e-aaeb-bab12cb7ed9a","value":5601147218404810302}]`)
	})
}
