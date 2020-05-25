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
		var k string = "06f5c2bb-0000-4b82-8cce-e2f92a39f643"
		var v interface{} = "2b83359a-68b3-483a-9d5d-5ba828c14989"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringAny_Delete(t *testing.T) {
	Convey("TestMapStringAny.Delete", t, func() {
		var k string = "8628e346-2f82-43e1-9243-db9e2b31a7ea"
		var v interface{} = "d3b84699-a78b-47e8-8fcc-898eec61d843"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringAny_Has(t *testing.T) {
	Convey("TestMapStringAny.Has", t, func() {
		var k string = "9f2ac87d-1d58-440c-bf0b-0f9511aa1281"
		var v interface{} = "44a84628-cf2b-4ae3-b633-57f53b4954b0"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d4d2ca92-1899-405c-90fb-4504fa46bbf3"+"6f758b23-65a5-48c5-b289-420b33c39270"), ShouldBeFalse)
	})
}

func TestMapStringAny_Get(t *testing.T) {
	Convey("TestMapStringAny.Get", t, func() {
		var k string = "f91b5222-cef5-46f2-86dd-e5cb33394f2f"
		var v interface{} = "b115e672-1456-4be1-931e-d8290f402ccb"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("b1eb452e-1696-43c3-9576-af5a456b8169" + "ec94996a-749d-4fb2-b384-25d735ff7818")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringAny_GetOpt(t *testing.T) {
	Convey("TestMapStringAny.GetOpt", t, func() {
		var k string = "0b696293-3661-408d-a5fe-93d5a95e1e07"
		var v interface{} = "16be0872-813f-4524-970f-f4609aa05c9f"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("de2669b4-d6f4-41d1-af7c-fd2ff84fd310" + "d549e674-570e-4bd5-8931-284b269dea96")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringAny_ForEach(t *testing.T) {
	Convey("TestMapStringAny.ForEach", t, func() {
		var k string = "10814d3e-219a-4ddc-86c0-c44587ef617b"
		var v interface{} = "4afb810d-71a4-476d-8b27-3abcde2c498f"
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
		var k string = "5534bb00-3d93-4b04-909d-f8b3994d0e06"
		var v interface{} = "f8716161-ee46-4de4-bfa4-f151766edf0b"

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
		var k string = "a26113af-5e56-4355-992e-7161dcca1d70"
		var v interface{} = "7daa4a91-467b-4984-9e61-944f8cb58bcf"

		test := omap.NewMapStringAny(1)

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

func TestMapStringAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringAny.PutIfNotNil", t, func() {
		var k string = "b24657bb-4eec-425c-908a-38d8f85375d9"
		var v interface{} = "d674cdcc-7fbc-47fc-8b47-cef6d6527356"

		test := omap.NewMapStringAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("1dd5a1d5-c78c-442e-9cc0-d6ece4611173", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "36300f5c-aa94-40b7-9520-a33493a47b13"
		So(test.PutIfNotNil("23fd582b-b1e3-4ff4-8ac8-225f753357c1", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringAny.ReplaceIfExists", t, func() {
		var k string = "78a52ef7-a0a9-4847-a3eb-d49041ca4f69"
		var v interface{} = "34efe575-584e-4e04-a950-dca16071a721"
		var x interface{} = "cd6c6b0b-09bf-4d86-9f8e-a80a546aa919"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("8ec0ea93-776f-4294-bd3a-0d9cc9b90532", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringAny.ReplaceOrPut", t, func() {
		var k string = "647a2b25-18e0-487e-878c-7373add257f6"
		var v interface{} = "bffe9d8b-6a0b-402d-8137-10e76a121a6d"
		var x interface{} = "22f67836-bd57-434b-be4b-e2955461119f"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("65ef7e0d-9f94-4c4c-be5f-9d4114e8e1a4", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_MarshalJSON(t *testing.T) {
	Convey("TestMapStringAny.MarshalJSON", t, func() {
		var k string = "196180f3-0f14-4e6b-9ef7-8032cdccbc1f"
		var v interface{} = "ca2286c5-a390-42f5-bf97-e38ffb55a7fe"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"196180f3-0f14-4e6b-9ef7-8032cdccbc1f","value":"ca2286c5-a390-42f5-bf97-e38ffb55a7fe"}]`)
	})
}
