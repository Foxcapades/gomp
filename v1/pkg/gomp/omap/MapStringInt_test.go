package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt_Put(t *testing.T) {
	Convey("TestMapStringInt.Put", t, func() {
		var k string = "b0ecc686-1775-48ec-bb68-98da0f5e50f8"
		var v int = 1638738368

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt_Delete(t *testing.T) {
	Convey("TestMapStringInt.Delete", t, func() {
		var k string = "1ed0a894-c842-440b-bb71-d675581d3cc6"
		var v int = 1825361514

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt_Has(t *testing.T) {
	Convey("TestMapStringInt.Has", t, func() {
		var k string = "9f2699e4-14f5-467c-b993-e8dcadf80139"
		var v int = 395706472

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("e5bdce7a-b492-4b89-a423-e6cbf902cb86"+"5226716d-59ff-4aff-a018-8105afbe85be"), ShouldBeFalse)
	})
}

func TestMapStringInt_Get(t *testing.T) {
	Convey("TestMapStringInt.Get", t, func() {
		var k string = "297c5058-9c1a-4e85-aaf0-cc30ccb22f38"
		var v int = 1230407880

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("8cfb2448-6441-405d-995f-652f09cbf518" + "5b0eda71-b883-441e-8b94-9db024f45ac5")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt_GetOpt(t *testing.T) {
	Convey("TestMapStringInt.GetOpt", t, func() {
		var k string = "4fcfe8fe-85ab-415b-ab2a-109e300b319d"
		var v int = 1828886640

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("e03a9f20-30e8-40a5-93d1-b4bdd70fed29" + "cc460247-0229-4fe0-9d25-c9365d5027de")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt_ForEach(t *testing.T) {
	Convey("TestMapStringInt.ForEach", t, func() {
		var k string = "d847aefb-e3ac-48f7-acd0-cc141e3424a6"
		var v int = 1753771023
		hits := 0

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt.MarshalYAML", t, func() {
		var k string = "dd054ab5-afeb-4fbb-8878-5a9307d5571b"
		var v int = 1233459068

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt_ToYAML(t *testing.T) {
	Convey("TestMapStringInt.ToYAML", t, func() {
		var k string = "45b5e1bd-4eec-4c75-90ae-57c3a3d48dbd"
		var v int = 2058511913

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt.PutIfNotNil", t, func() {
		var k string = "b21ae9fd-6faf-4839-8c6a-0bca2a97e38c"
		var v int = 455968802

		test := omap.NewMapStringInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("1561a32d-664b-4263-be85-fd187212cd88", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 1288282352
		So(test.PutIfNotNil("09f545eb-f7f8-42d8-afc9-cd05174e9b39", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt.ReplaceIfExists", t, func() {
		var k string = "64e8865d-87da-4759-a018-0ecc8a9ea1bc"
		var v int = 901660005
		var x int = 306364613

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("0788f17b-5f0b-400a-ac4c-970437278144", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt.ReplaceOrPut", t, func() {
		var k string = "189c79ef-b83a-4143-84bb-2e1a043e1a3f"
		var v int = 1432316907
		var x int = 1434397136

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ca3f98b3-fbb6-4ff3-94fa-8a596631ad63", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt.MarshalJSON", t, func() {
		var k string = "18110435-0693-4a30-bc8b-7435062abb92"
		var v int = 115497890

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"18110435-0693-4a30-bc8b-7435062abb92","value":115497890}]`)
	})
}
