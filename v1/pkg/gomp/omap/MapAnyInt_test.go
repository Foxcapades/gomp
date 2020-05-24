package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt_Put(t *testing.T) {
	Convey("TestMapAnyInt.Put", t, func() {
		var k interface{} = "1c2b1ea5-8334-4223-9d88-ab4c528814f0"
		var v int = 2097790212

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt_Delete(t *testing.T) {
	Convey("TestMapAnyInt.Delete", t, func() {
		var k interface{} = "b8cee46b-dbb0-4385-9c7c-b3a8e2fa4bab"
		var v int = 1382751806

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt_Has(t *testing.T) {
	Convey("TestMapAnyInt.Has", t, func() {
		var k interface{} = "91141ae5-3b90-42f8-8e73-4451954f12ce"
		var v int = 230152401

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("7083fea9-ff9e-46b2-876d-5047e70237e9"+"71fee516-d4cd-4de4-b8d2-c5fc1ae45ddb"), ShouldBeFalse)
	})
}


func TestMapAnyInt_Get(t *testing.T) {
	Convey("TestMapAnyInt.Get", t, func() {
		var k interface{} = "6fb6a1aa-c7df-4bc3-94be-93a106bc0e38"
		var v int = 982150581

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("cd373ea9-e011-4343-a3c5-48a2924cb9c4" + "71e0fea8-c3ca-4512-9b48-2b2645a9ef0c")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt.GetOpt", t, func() {
		var k interface{} = "cc67295a-4159-49e9-ad04-5eea5f147b82"
		var v int = 1248402033

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("23245a42-8aa6-4a5a-aa19-cf41f7f1282a" + "89e5f696-673f-4375-a180-aad82509a2d5")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt_ForEach(t *testing.T) {
	Convey("TestMapAnyInt.ForEach", t, func() {
		var k interface{} = "cd4370f0-d529-4c10-8ad6-7eda6abca8f7"
		var v int = 1065009319
		hits := 0

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt.MarshalYAML", t, func() {
		var k interface{} = "f9038099-cd96-4735-aac9-69bf2b6fe2bf"
		var v int = 795044753

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt.ToYAML", t, func() {
		var k interface{} = "f885f97d-4198-40ee-bb3f-5bfe273cf792"
		var v int = 10417313

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt.PutIfNotNil", t, func() {
		var k interface{} = "084a34f1-9034-4458-96e3-5454ff993786"
		var v int = 1967172181

		test := omap.NewMapAnyInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("a2915525-f614-4564-9805-a29560112582", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 446803678
		So(test.PutIfNotNil("6b9dacdf-3663-484e-ba71-69bb8edef28b", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceIfExists", t, func() {
		var k interface{} = "9fb18280-617c-4e08-bc46-273cef39c84f"
		var v int = 1281829423
		var x int = 624622812

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ec83faaf-4138-4663-9705-0978b727cb20", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceOrPut", t, func() {
		var k interface{} = "f8d7e7a1-3fbb-4d38-8657-7c46580c9d56"
		var v int = 1330580268
		var x int = 644303810

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ea8d7d6c-e233-4d04-b18b-503ee08c896c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt.MarshalJSON", t, func() {
		var k interface{} = "5b066752-95ba-4afa-a74e-7fa02dafba63"
		var v int = 1653797422

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"5b066752-95ba-4afa-a74e-7fa02dafba63","value":1653797422}]`)
	})
}

