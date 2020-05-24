package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringAny_Put(t *testing.T) {
	Convey("TestMapStringAny.Put", t, func() {
		var k string = "6ab01828-fcad-45e6-b437-98695479289e"
		var v interface{} = "c2710ff3-71b2-44ba-984e-2fad2f7534e1"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringAny_Delete(t *testing.T) {
	Convey("TestMapStringAny.Delete", t, func() {
		var k string = "b852f2d0-a3dc-4acc-b747-e114bce2c6fc"
		var v interface{} = "04c7a73b-e49b-463d-a175-8e6f3fe6098b"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringAny_Has(t *testing.T) {
	Convey("TestMapStringAny.Has", t, func() {
		var k string = "3845590e-cbbc-49a0-8d04-8cb1a2bfdb9b"
		var v interface{} = "9e508bb9-42aa-4f7f-ae89-7b4f235e6b56"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5d8c3157-4688-4221-93c5-66f1fb926079"+"a52ebb1f-2b7e-4d27-9a14-92ca3b8e200e"), ShouldBeFalse)
	})
}


func TestMapStringAny_Get(t *testing.T) {
	Convey("TestMapStringAny.Get", t, func() {
		var k string = "086a1ad2-27c9-497d-b761-c6d6986af35e"
		var v interface{} = "2816097b-841c-420b-a741-55421ecce94f"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("a7c78f68-5bb4-416b-9265-923d59d98765"+"a520808f-885b-4233-917b-0faf767e14c6")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringAny_GetOpt(t *testing.T) {
	Convey("TestMapStringAny.GetOpt", t, func() {
		var k string = "479c45fa-b4af-4896-bcf1-d14d365d110b"
		var v interface{} = "7e5cdbf4-92e7-4230-9e4b-724287146a91"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("f7424079-aa01-4b80-a528-eb84c8709dfd"+"3cb6f492-f886-43c6-a78d-da5932578117")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringAny_ForEach(t *testing.T) {
	Convey("TestMapStringAny.ForEach", t, func() {
		var k string = "fb79836a-bdaf-48cf-9f59-dbd6c9e21c7c"
		var v interface{} = "c7f5ed75-20e5-420c-8847-bcb59530872d"
		hits := 0

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringAny_MarshalYAML(t *testing.T) {
	Convey("TestMapStringAny.MarshalYAML", t, func() {
		var k string = "8de0dcad-f1af-4faa-b014-570aa7f44ae9"
		var v interface{} = "4ede366a-d309-4280-b37a-eb152ed403b9"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringAny_ToYAML(t *testing.T) {
	Convey("TestMapStringAny.ToYAML", t, func() {
		var k string = "9ef7e9af-32df-4f09-a472-05ffccf8ab20"
		var v interface{} = "ba4b5ccd-7c3b-44cd-89b3-9ec90a5d3b68"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringAny.PutIfNotNil", t, func() {
		var k string = "63aa77f5-4889-4968-9f51-dca3e81da611"
		var v interface{} = "33267447-a497-40c5-9683-de2080a52ba3"

		test := omap.NewMapStringAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("8333f44a-3a2c-421a-b7e4-d4cc803ec947", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "54b57165-68f9-4f80-b9e0-5480aa045634"
		So(test.PutIfNotNil("a88cee4a-7db8-4b61-bb77-386abbeee296", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringAny.ReplaceIfExists", t, func() {
		var k string = "7e08de1f-c141-4548-b9e0-d0bf7419844a"
		var v interface{} = "987de181-e7b4-419b-8268-9c66080f66d5"
		var x interface{} = "a03d3f33-d2f1-4dae-829a-403b69cdb4a4"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1a680fd2-7a0f-4980-8dea-f4a6d76c0239", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringAny.ReplaceOrPut", t, func() {
		var k string = "819d2ac7-d6ed-46bf-8822-82bf405dff40"
		var v interface{} = "a0793d7a-08ad-4145-9744-cc1b9605f88f"
		var x interface{} = "9184538a-1c6c-442a-bbfc-8d2230fdd0c2"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("57bf2472-b488-4ca3-affb-bed950ad4ee8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_MarshalJSON(t *testing.T) {
	Convey("TestMapStringAny.MarshalJSON", t, func() {
		var k string = "22bab69f-23b1-428e-a64f-54985fdaac44"
		var v interface{} = "8ab38935-9b1b-46d1-a9ee-f500e9525be6"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"22bab69f-23b1-428e-a64f-54985fdaac44","value":"8ab38935-9b1b-46d1-a9ee-f500e9525be6"}]`)
	})
}

