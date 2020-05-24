package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyString_Put(t *testing.T) {
	Convey("TestMapAnyString.Put", t, func() {
		var k interface{} = "8f1d60d3-afa7-4191-a2c3-dca143029fc0"
		var v string = "efa3cdbb-640d-44e8-94a1-63f90d76c6a5"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyString_Delete(t *testing.T) {
	Convey("TestMapAnyString.Delete", t, func() {
		var k interface{} = "38a3f34a-ca12-420e-b66e-b00a4915ca00"
		var v string = "7a7c88e7-6332-44f7-a671-ff22095d932b"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyString_Has(t *testing.T) {
	Convey("TestMapAnyString.Has", t, func() {
		var k interface{} = "676b36db-2981-4f75-8543-3a27f14020b5"
		var v string = "02f0998c-b252-44d8-9318-80947c83c2e5"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("bd701a8e-e545-4492-93fe-82f02327e961"+"3b455d18-b1af-4595-abac-d364bc6943f2"), ShouldBeFalse)
	})
}


func TestMapAnyString_Get(t *testing.T) {
	Convey("TestMapAnyString.Get", t, func() {
		var k interface{} = "2ef7888e-1a24-4a52-92c6-9ccca9277eae"
		var v string = "53b812dd-e46b-43bd-a3e2-b90a835b6c7e"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("51f7c97f-0a25-4a59-88b9-94c53ce04871" + "d904e768-bc07-44f2-a007-a0ab5a816ae5")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyString_GetOpt(t *testing.T) {
	Convey("TestMapAnyString.GetOpt", t, func() {
		var k interface{} = "39faa175-2e98-404d-a2a8-e1613f706c42"
		var v string = "e231dfae-86fa-472d-8ad8-6603c95e8e76"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("18ea3d8e-6e4e-4582-a363-765d98cd468e" + "278db2fe-f849-4d9d-b0f9-1ea924b5462a")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyString_ForEach(t *testing.T) {
	Convey("TestMapAnyString.ForEach", t, func() {
		var k interface{} = "f2b1114b-6b08-4894-89e2-1aa4fb8055fa"
		var v string = "cf4eb83a-1434-43d5-9f69-5c490125cfe7"
		hits := 0

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyString_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyString.MarshalYAML", t, func() {
		var k interface{} = "a5e0a87e-a4e8-4314-99c4-692ce168475c"
		var v string = "da68297e-af61-4e84-a74a-91f75e7e296f"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyString_ToYAML(t *testing.T) {
	Convey("TestMapAnyString.ToYAML", t, func() {
		var k interface{} = "07b65163-b9af-4977-b19e-7bd045e923f6"
		var v string = "fbcfade8-042c-4c93-bb2e-4251a4ece393"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyString_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyString.PutIfNotNil", t, func() {
		var k interface{} = "3d6434e2-acbc-434e-9c51-dca1f3752068"
		var v string = "3024ea59-f17e-474f-bcfb-2f10041520c8"

		test := omap.NewMapAnyString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("e3eb8a98-c56e-4b86-8970-67d78ec992fb", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "03bace78-960e-4c0a-aeeb-e352a9149657"
		So(test.PutIfNotNil("3bda9738-368f-499c-816a-7fb87d95234c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyString.ReplaceIfExists", t, func() {
		var k interface{} = "d02acd4a-37bd-4f3c-adb4-7dd6abd47198"
		var v string = "c5a4635e-ffba-481f-adbb-038f2a014540"
		var x string = "6c37f2b6-7caf-47b9-8ea8-8066344f71a8"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("429dbf50-caf8-42da-bfdf-29071118185a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyString.ReplaceOrPut", t, func() {
		var k interface{} = "e9125a83-fbdd-4aee-b4a4-64d78dd1c74c"
		var v string = "ce513027-3ec0-422c-be9a-518976ba3ff3"
		var x string = "de679955-d75b-462b-a136-a53203f58132"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("f08633c4-78bd-44b8-8475-5e8c642f2204", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyString.MarshalJSON", t, func() {
		var k interface{} = "ef4051ee-bad1-44d1-b3e4-f0a2a9d45e1d"
		var v string = "9aebba41-697b-4906-bbe8-0d52fc0ca50f"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"ef4051ee-bad1-44d1-b3e4-f0a2a9d45e1d","value":"9aebba41-697b-4906-bbe8-0d52fc0ca50f"}]`)
	})
}

