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
		var k interface{} = "f5c1457c-eb11-4669-a234-990c75e00b64"
		var v interface{} = "8465b91d-5b4b-4e07-8983-8391098603c2"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAny_Delete(t *testing.T) {
	Convey("TestMapAny.Delete", t, func() {
		var k interface{} = "fb955262-b0b5-4b04-b664-ca704c9d49b9"
		var v interface{} = "398d4b20-3d7d-4920-84e7-4c2ff3045df1"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAny_Has(t *testing.T) {
	Convey("TestMapAny.Has", t, func() {
		var k interface{} = "82d0e5c4-927e-4a53-846f-20e280e93ab1"
		var v interface{} = "1dd0c172-87c2-40b4-a72d-2791804e28d1"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3d0b3f94-ba34-4d39-95e5-f43c05de58f7"+"2576b746-6d12-4b34-820a-692d0abde76a"), ShouldBeFalse)
	})
}

func TestMapAny_Get(t *testing.T) {
	Convey("TestMapAny.Get", t, func() {
		var k interface{} = "44e8be36-36b9-4ae9-b869-51af0339415d"
		var v interface{} = "2b47b89e-611e-4c52-b571-cd4fffeb57ab"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("48825d19-e02d-447c-b8d8-bf9c93ba12bf" + "f568e0d0-bd87-4445-a3e9-125dc2f987dc")
		So(b, ShouldBeFalse)
	})
}

func TestMapAny_GetOpt(t *testing.T) {
	Convey("TestMapAny.GetOpt", t, func() {
		var k interface{} = "843f8012-54f8-44fc-bd76-2f1c186e931c"
		var v interface{} = "0d93bc22-be1c-43bc-a2cd-1c223ff59698"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("fbc54c93-e615-44bc-b41b-2feb64ad1713" + "5ab4feaf-b6c0-4064-b1a2-7a12caa938f5")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAny_ForEach(t *testing.T) {
	Convey("TestMapAny.ForEach", t, func() {
		var k interface{} = "045023fc-e89a-48cb-b737-4845e3dc16aa"
		var v interface{} = "aa24fc90-8cc6-4e33-acd6-be456b044bee"
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
		var k interface{} = "2897a2f4-3b53-4127-b18a-dde32c85b71a"
		var v interface{} = "2bcf9bf0-49ea-4f20-a68c-b71aaa7919f4"

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
		var k interface{} = "c77d1933-5f72-4472-ad07-a8978a989bd2"
		var v interface{} = "583726df-47a8-41f3-b194-954389de1609"

		test := omap.NewMapAny(1)

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

func TestMapAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapAny.PutIfNotNil", t, func() {
		var k interface{} = "8f811de9-1923-46b5-9087-d58c6d1ba60f"
		var v interface{} = "46bd83d4-92cd-4fba-9251-aff362268509"

		test := omap.NewMapAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("7c36fcc9-98bc-4664-932c-d19610e0e8d0", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "81c1a742-a750-46e1-bfc7-cb8aba00c523"
		So(test.PutIfNotNil("34fc575a-bee6-47b5-9008-6d4919c36763", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAny.ReplaceIfExists", t, func() {
		var k interface{} = "ef2fcca1-f619-43bc-8af1-44d5e43b8e63"
		var v interface{} = "2ef6ff62-e66d-4378-addf-738d060e6a6e"
		var x interface{} = "09510c2f-2ce7-469c-837b-a08391f4d6d9"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ef8adad5-bbe3-483f-93f0-51a5f0d873f5", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAny.ReplaceOrPut", t, func() {
		var k interface{} = "2af447ce-e636-4aaa-9bd7-0a5d569797c0"
		var v interface{} = "5c123bf4-2d67-44e7-b719-1e97e9ceb50e"
		var x interface{} = "d860c152-6502-44dd-b67d-e0706eac461a"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("16017ab1-f7aa-4442-8964-2c66fd69d823", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_MarshalJSON(t *testing.T) {
	Convey("TestMapAny.MarshalJSON", t, func() {
		var k interface{} = "921833bc-7d8e-4698-8117-97dbf2fe6be8"
		var v interface{} = "d1a29a65-e224-43f6-b77d-c7cdf0a4edbe"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"921833bc-7d8e-4698-8117-97dbf2fe6be8","value":"d1a29a65-e224-43f6-b77d-c7cdf0a4edbe"}]`)
	})
}
