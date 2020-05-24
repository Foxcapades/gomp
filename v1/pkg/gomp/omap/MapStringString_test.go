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
		var k string = "ea45ddf4-691f-4dae-bb0e-127d13f2dcfe"
		var v string = "456873a0-60e8-4ada-ad25-3f21f7180c4e"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringString_Delete(t *testing.T) {
	Convey("TestMapStringString.Delete", t, func() {
		var k string = "cf98c695-7827-40b6-a070-614e0935ba9c"
		var v string = "01e41cd3-2022-4c39-a37b-81866df2e9c8"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringString_Has(t *testing.T) {
	Convey("TestMapStringString.Has", t, func() {
		var k string = "853f3fa7-fd56-43fb-a243-0154f1c0e1c3"
		var v string = "58e9042f-7df8-47db-a14b-50fdee638414"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5a436ace-06cc-4c85-9568-e048847b9175"+"5df90b26-70cc-4fdc-b114-7a4beec2efb6"), ShouldBeFalse)
	})
}


func TestMapStringString_Get(t *testing.T) {
	Convey("TestMapStringString.Get", t, func() {
		var k string = "a5460ff5-9647-404b-95de-7aa51eb76add"
		var v string = "f100c20d-1172-4c70-abb0-70ecda1133fb"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("0203c261-b421-43b6-8f47-1be0f6542f74"+"4395d9c6-5e9a-4901-9ceb-c36074aa326d")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringString_GetOpt(t *testing.T) {
	Convey("TestMapStringString.GetOpt", t, func() {
		var k string = "3d58d645-2e62-456e-8148-afec9e5bfa0d"
		var v string = "be2b2a03-c1b1-4623-bb8a-eb514dfdbf1a"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("1dc8f99b-9691-4f9b-8cf9-f2c62a72ac64"+"7b51c841-cb0b-47ec-8e81-874196844750")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringString_ForEach(t *testing.T) {
	Convey("TestMapStringString.ForEach", t, func() {
		var k string = "85838b7d-add6-424f-8782-83f198d25879"
		var v string = "69bee6cb-1ba3-4ce7-ba94-9d02c8a25ffb"
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
		var k string = "6c3c69cc-c398-4119-b38a-0032abfcd880"
		var v string = "eb473585-4751-4b9b-80c9-e028b52dfddc"

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
		var k string = "e190e432-d41c-4350-b9ce-b132f335a10b"
		var v string = "d80a63f0-461d-414c-892a-4183ec8c6d61"

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
		var k string = "3270571d-0e33-44ab-bf10-220523fc51e3"
		var v string = "d1ff3a48-f639-40d1-95f8-fc3a45580553"

		test := omap.NewMapStringString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("16987896-6d40-41f3-a70b-5a8708155145", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "b55327c5-7636-48e0-8052-83586717f48e"
		So(test.PutIfNotNil("74324b54-7798-4d09-9c7d-a99116dc5e99", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringString.ReplaceIfExists", t, func() {
		var k string = "34aec2f0-f035-4f61-99fd-e69ff8865000"
		var v string = "8c39abce-e6d4-4e68-9c02-32938565c1b5"
		var x string = "1ef41a7c-c62e-448c-85a1-16c69384b06d"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("3193b665-7ef4-40de-a031-3e75b9ac5a12", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringString.ReplaceOrPut", t, func() {
		var k string = "32c9d82b-140d-4b2f-9b2c-64fff5c6fe3b"
		var v string = "651fdd70-ed7a-4c71-bcba-e4611fe7dbd3"
		var x string = "2d962f9d-5ed3-486c-b855-3540847015c0"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d3318646-d068-41ff-86e5-b2a1eadae0f6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_MarshalJSON(t *testing.T) {
	Convey("TestMapStringString.MarshalJSON", t, func() {
		var k string = "24bee725-2380-4a17-94eb-68f314534fbd"
		var v string = "a9a1d34c-5e0f-42d4-96ba-3c6b9ac4c928"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"24bee725-2380-4a17-94eb-68f314534fbd","value":"a9a1d34c-5e0f-42d4-96ba-3c6b9ac4c928"}]`)
	})
}

