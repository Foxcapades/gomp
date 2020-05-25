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
		var k string = "92fe7587-aa7b-4af2-baea-0207bb007b59"
		var v string = "274abd09-8912-4075-a77e-f0fbf441c93a"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringString_Delete(t *testing.T) {
	Convey("TestMapStringString.Delete", t, func() {
		var k string = "0c237ee5-7b15-40cd-ab46-5d7ef1c1984c"
		var v string = "5cc054a2-13bb-4687-8cb8-08fbf91da562"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringString_Has(t *testing.T) {
	Convey("TestMapStringString.Has", t, func() {
		var k string = "f4d804da-46e1-4a8b-8537-a128b249631f"
		var v string = "b8b112f2-3e06-4f9e-900c-63b6224403a2"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("754b2e44-7578-412c-bac3-7cbddc818ddc"+"def210a0-5a20-49a3-a299-eecc1a7cbcef"), ShouldBeFalse)
	})
}

func TestMapStringString_Get(t *testing.T) {
	Convey("TestMapStringString.Get", t, func() {
		var k string = "2c846c89-9d85-4209-8174-9c9ee6b1576f"
		var v string = "cea8e29e-e57b-4fa8-9dc3-4ac60300dfca"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("1e76c7d4-6d02-46d7-a914-c2d8d1bf8d8f" + "14e5fdfa-b1a4-442c-956e-86b69abd8f20")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringString_GetOpt(t *testing.T) {
	Convey("TestMapStringString.GetOpt", t, func() {
		var k string = "7f954ed1-3fb8-452b-9649-5b2b4be654e9"
		var v string = "945b0921-d1ba-416e-8006-630feb622ad9"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("62fd1a84-fce0-4630-bebb-d6e52b1ae9ba" + "49529df9-e5ee-4e89-ab5a-4261e1903272")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringString_ForEach(t *testing.T) {
	Convey("TestMapStringString.ForEach", t, func() {
		var k string = "90dd9f67-7b5b-493a-9409-de62af7bf7c7"
		var v string = "d6254c1c-077f-41b1-ac94-d8d887f00c86"
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
		var k string = "7bd2d6e5-7cf0-48b6-a9c9-ab7fb22c1213"
		var v string = "d94de476-145d-405e-8f61-dd51cd55add9"

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
		var k string = "02f717a5-7a80-43f2-a7f6-b6ed071ccf34"
		var v string = "2d722768-f2f5-416d-80c8-c3a8c98b7b5d"

		test := omap.NewMapStringString(1)

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

func TestMapStringString_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringString.PutIfNotNil", t, func() {
		var k string = "cb67d6d0-0827-4c11-8da2-1b2e603540d3"
		var v string = "0041bec4-cc47-46af-9173-b499081d9b12"

		test := omap.NewMapStringString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("54f5b7b7-dd4f-44d0-b3d1-71e9a8be0dd5", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "edf6cef4-eb12-4bc3-802d-06bc7155cee2"
		So(test.PutIfNotNil("3504f1ae-f4e8-487d-a8c7-2bf14949da36", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringString.ReplaceIfExists", t, func() {
		var k string = "5368e237-f78f-4de4-ab18-c55a07cc1ac5"
		var v string = "d0a470ef-dd54-42a4-9e3c-c60d9801711e"
		var x string = "4bfc40e0-8cdc-4119-a2c3-693b84893380"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("92399409-0c57-4e36-a73d-26fecc212976", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringString.ReplaceOrPut", t, func() {
		var k string = "0efc633d-bcd0-42f8-842a-27394a6fefe2"
		var v string = "04986c45-4955-4b17-8eb5-030408a03474"
		var x string = "e66113c2-585c-4930-842e-a8190f535b45"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("fa78eaf2-6187-4c4d-b93f-902b58a50b83", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_MarshalJSON(t *testing.T) {
	Convey("TestMapStringString.MarshalJSON", t, func() {
		var k string = "8301dc6e-2f7b-4d9e-9a93-c9944096d605"
		var v string = "a44df66a-3d39-4325-bd0a-67d2e99cbbbf"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"8301dc6e-2f7b-4d9e-9a93-c9944096d605","value":"a44df66a-3d39-4325-bd0a-67d2e99cbbbf"}]`)
	})
}
