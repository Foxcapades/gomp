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
		var k string = "abd4042c-e59c-4df5-aaa7-e448f49ad704"
		var v interface{} = "61b27346-28d0-4f11-86c3-c3b7044fc473"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringAny_Delete(t *testing.T) {
	Convey("TestMapStringAny.Delete", t, func() {
		var k string = "5fa8c93e-742f-4087-bd87-0260ecb412cd"
		var v interface{} = "299a509d-7d33-49a9-81b2-44bee1274fa3"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringAny_Has(t *testing.T) {
	Convey("TestMapStringAny.Has", t, func() {
		var k string = "0f79f989-601d-4036-b756-b09afa01c113"
		var v interface{} = "82aec66e-ab7e-440e-a878-271c5de2ba74"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b12a0bd0-7456-402f-bca6-7244674a32aa"+"78f99677-92a0-436f-a8e8-98440f0e551f"), ShouldBeFalse)
	})
}


func TestMapStringAny_Get(t *testing.T) {
	Convey("TestMapStringAny.Get", t, func() {
		var k string = "1bf8c95b-5b53-4669-bce7-ae3db53366fa"
		var v interface{} = "a2cd2b94-f9dc-4ecb-8a5d-6b043c775691"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("bbce99e7-637d-439e-9bbf-afc11939125e" + "b7445569-2e09-478a-9ce2-ad642a0eeb76")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringAny_GetOpt(t *testing.T) {
	Convey("TestMapStringAny.GetOpt", t, func() {
		var k string = "2298bc3a-d1f3-4250-8ae9-2d2883490edd"
		var v interface{} = "d09aa896-69dc-497b-b6d7-cb20bd0d8959"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("185f5ddc-30cb-4bb1-a740-094d97b3276a" + "0aed3506-eb93-4959-909a-263ee0e4a4c4")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringAny_ForEach(t *testing.T) {
	Convey("TestMapStringAny.ForEach", t, func() {
		var k string = "fbaa9151-fe34-4938-908d-9562198bb62e"
		var v interface{} = "d5c70737-d6fd-4e5f-966e-1fdb7ccaea5f"
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
		var k string = "c1633eca-a830-41f6-ad14-84c10b853584"
		var v interface{} = "efb55ef4-aabf-4746-b973-6e2a3c20ed9b"

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
		var k string = "9799ceef-f011-4526-b616-68ed4f4e8159"
		var v interface{} = "9fd3de90-44ed-424f-a788-c46707c009f9"

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
		var k string = "f3d1e963-2bcf-4b9e-9701-96d26ed2aa2c"
		var v interface{} = "3e650278-5070-4b2a-8d3f-4499e2984239"

		test := omap.NewMapStringAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("116021d0-0319-4ad3-8999-28384f896c0b", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "fa5985fa-2500-4669-a6de-0ed285d0f446"
		So(test.PutIfNotNil("261efaae-cb17-4e51-8879-6bbd4ac84ce6", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringAny.ReplaceIfExists", t, func() {
		var k string = "177f7146-a32d-4a9f-ad09-1e33cf156057"
		var v interface{} = "c0611c4e-6d6e-4bd9-8cde-fc6f8bfcaf6b"
		var x interface{} = "5283ca84-1ad3-4180-829d-5ecc16afafbc"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("3772a26f-7248-43d1-8b3d-59b8becf6bfd", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringAny.ReplaceOrPut", t, func() {
		var k string = "4a6bb820-be01-4d75-83a5-65aa7da90dd2"
		var v interface{} = "3e8be07f-7763-43d4-abb4-bb654693355f"
		var x interface{} = "6d1f6b5b-684e-4d5c-9fbc-79930f08a54c"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("961be341-c466-4879-bb7e-4284a44dda06", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_MarshalJSON(t *testing.T) {
	Convey("TestMapStringAny.MarshalJSON", t, func() {
		var k string = "9f129334-0d35-481c-aa02-982f3c960eaa"
		var v interface{} = "95ad66e1-afa5-4cc7-b0d7-1c95fcb7a630"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"9f129334-0d35-481c-aa02-982f3c960eaa","value":"95ad66e1-afa5-4cc7-b0d7-1c95fcb7a630"}]`)
	})
}
