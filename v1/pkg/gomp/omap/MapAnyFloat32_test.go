package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyFloat32_Put(t *testing.T) {
	Convey("TestMapAnyFloat32.Put", t, func() {
		var k interface{} = "c6311984-c61c-4627-b6d9-b53c57f11cae"
		var v float32 = 0.331

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat32_Delete(t *testing.T) {
	Convey("TestMapAnyFloat32.Delete", t, func() {
		var k interface{} = "08ac802e-af8c-470e-a1f8-c86da5d7817c"
		var v float32 = 0.900

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat32_Has(t *testing.T) {
	Convey("TestMapAnyFloat32.Has", t, func() {
		var k interface{} = "d855d3f3-4957-4ad5-80c6-b373a40f04d4"
		var v float32 = 0.902

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("4c485aeb-b79d-4f3d-8338-c45c53530e3d"+"a9db3ea1-5034-4e41-9558-fb1bce975482"), ShouldBeFalse)
	})
}


func TestMapAnyFloat32_Get(t *testing.T) {
	Convey("TestMapAnyFloat32.Get", t, func() {
		var k interface{} = "aab09bbc-c9a2-4d84-8855-ce16cdfbedeb"
		var v float32 = 0.473

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("05773dea-583d-4851-8176-dc4c7c89a702" + "b4c03ec1-dfe0-4e5c-ad97-52ca8eec893e")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat32_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat32.GetOpt", t, func() {
		var k interface{} = "9be88f2a-6739-41c4-b47e-da303b9b27a8"
		var v float32 = 0.603

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("a3299007-e367-4afc-8aaf-cf853d16128a" + "066f00b1-890a-48bf-9d8c-2da84b5fe095")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat32_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat32.ForEach", t, func() {
		var k interface{} = "cfb7e6b3-88c8-4e0f-8265-f7ff40abe34f"
		var v float32 = 0.897
		hits := 0

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyFloat32.MarshalYAML", t, func() {
		var k interface{} = "654c482c-4e2e-4db9-a54c-3729f9a08b5c"
		var v float32 = 0.101

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyFloat32_ToYAML(t *testing.T) {
	Convey("TestMapAnyFloat32.ToYAML", t, func() {
		var k interface{} = "fcc1ea4c-0d1b-41bf-9518-23e3d2dadb0b"
		var v float32 = 0.820

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyFloat32.PutIfNotNil", t, func() {
		var k interface{} = "0d36a6bd-15c4-4adb-a2c4-abbcb332cd7a"
		var v float32 = 0.500

		test := omap.NewMapAnyFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("f1c0c92f-f624-44c4-b45d-0e91d467062a", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.741
		So(test.PutIfNotNil("ff2fe0be-10f7-476c-935d-e7404b71284e", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceIfExists", t, func() {
		var k interface{} = "a9b65b2b-53f3-464f-8059-a7e45b68337c"
		var v float32 = 0.122
		var x float32 = 0.918

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("055ca8d3-0936-47f6-bbea-bfc81a43d4b4", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceOrPut", t, func() {
		var k interface{} = "aba97bc0-3acc-4bdd-b006-ff821946c335"
		var v float32 = 0.850
		var x float32 = 0.719

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ac3b60b4-a4b7-4c34-bfb7-89e8dfeb4e6c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat32.MarshalJSON", t, func() {
		var k interface{} = "b1b46c66-8578-44ca-85ec-0a153030075a"
		var v float32 = 0.680

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"b1b46c66-8578-44ca-85ec-0a153030075a","value":0.68}]`)
	})
}
