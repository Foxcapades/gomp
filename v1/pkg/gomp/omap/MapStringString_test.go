package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringString_Put(t *testing.T) {
	Convey("TestMapStringString.Put", t, func() {
		var k string = "ac26bc0c-7b1e-42c5-abd1-374a6425c859"
		var v string = "39d03eff-6b6a-43b0-b508-fee207c5a91c"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringString_Delete(t *testing.T) {
	Convey("TestMapStringString.Delete", t, func() {
		var k string = "217ba862-88bd-4fab-a808-06ef5e50bd88"
		var v string = "71d96841-e45a-4907-bcfb-aa663001dc0c"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringString_Has(t *testing.T) {
	Convey("TestMapStringString.Has", t, func() {
		var k string = "8e8e54a6-18db-4f25-9b56-616179b9a610"
		var v string = "9a2163c3-3911-46f5-9480-a0b0fc3acedb"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("a71dbcab-e5e4-4bf9-b293-ad2eae1f19ad"+"fcdae294-e373-4954-af7b-a26ce9d49152"), ShouldBeFalse)
	})
}


func TestMapStringString_Get(t *testing.T) {
	Convey("TestMapStringString.Get", t, func() {
		var k string = "1b24600c-6efc-45ac-8d78-f3c478fbcc79"
		var v string = "4421a7e2-7af4-4ec1-a467-7a6fa5b00037"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("f2b838b6-922a-4e12-8299-7ea3242e5e91" + "62fe0771-3466-405c-b9d6-ff842d15ce3a")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringString_GetOpt(t *testing.T) {
	Convey("TestMapStringString.GetOpt", t, func() {
		var k string = "f06dccc0-d0d4-4d23-9cf9-656ad7cc60fb"
		var v string = "81a3b0bb-6f4f-45cf-b63e-58971c64b62f"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("246976a4-04b0-44e5-91b4-1dca7a73fb0a" + "54ea2fd3-8efe-407e-b253-4162d2537406")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringString_ForEach(t *testing.T) {
	Convey("TestMapStringString.ForEach", t, func() {
		var k string = "55fc0eb4-750e-40df-a73b-4c697c3a47c6"
		var v string = "e50613db-f1d9-4467-8c3f-10daf97c63d0"
		hits := 0

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringString_MarshalYAML(t *testing.T) {
	Convey("TestMapStringString.MarshalYAML", t, func() {
		var k string = "78738781-c55b-4bba-a9ed-bc3c6e704834"
		var v string = "b5771d14-869b-459e-82ff-ab51a26ede0f"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringString_ToYAML(t *testing.T) {
	Convey("TestMapStringString.ToYAML", t, func() {
		var k string = "82bec895-16ba-441b-9827-89955bbef479"
		var v string = "0147280c-e7f8-4fc2-8c17-74fb788e4204"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringString_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringString.PutIfNotNil", t, func() {
		var k string = "99c35d62-a5e2-48d5-bbc8-649d6d087223"
		var v string = "7f7f5fdc-71be-431a-80f8-83aaa826f4be"

		test := omap.NewMapStringString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b0435828-7c2e-4ee4-a5bb-bb9e2491d664", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "6e763a68-aa95-4647-a7ee-b751baa23a26"
		So(test.PutIfNotNil("c4f47ff9-303d-4b7d-a1c9-4898600e88d7", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringString.ReplaceIfExists", t, func() {
		var k string = "41f38ea6-39d5-42cc-97d6-54f25e6ed05c"
		var v string = "218f94cc-76df-4c80-a056-91009077fe7f"
		var x string = "0c277e1d-4248-4043-bd42-c86b78099595"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ccecea8b-1e0f-425f-9022-cc14ffb60a2a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringString.ReplaceOrPut", t, func() {
		var k string = "2e478cad-fa43-41e6-98c4-f305d35f8a80"
		var v string = "964f627d-e229-4e6f-9c2b-51076139bfbb"
		var x string = "f3635ab4-1095-4a3c-a302-6367f0738172"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("87c81ffa-bf65-4178-963f-90107fd4b593", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_MarshalJSON(t *testing.T) {
	Convey("TestMapStringString.MarshalJSON", t, func() {
		var k string = "019de07e-f23b-4411-8b2e-bc7f21ed4c66"
		var v string = "685acf3f-779e-4244-b61f-7309a5406e89"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"019de07e-f23b-4411-8b2e-bc7f21ed4c66","value":"685acf3f-779e-4244-b61f-7309a5406e89"}]`)
	})
}
