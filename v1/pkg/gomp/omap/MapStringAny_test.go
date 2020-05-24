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
		var k string = "4e6f89ea-eaaf-401e-aa02-ded2dac52c45"
		var v interface{} = "40554bf7-b19f-4500-ae4a-7ac93ea1cc01"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringAny_Delete(t *testing.T) {
	Convey("TestMapStringAny.Delete", t, func() {
		var k string = "e0e20a63-18d9-4154-94dd-9d82cfb3fd07"
		var v interface{} = "94fc29a0-5cd0-4c12-a02d-6411cf87dc9b"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringAny_Has(t *testing.T) {
	Convey("TestMapStringAny.Has", t, func() {
		var k string = "6b3ff33e-0aae-4504-8c3e-004fc957fdc5"
		var v interface{} = "eb8e1cdd-ac48-497a-8583-0110c08e193b"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("ebf6bad0-81f7-4e7a-a419-3be0cf78fd0e"+"3faccd11-e5aa-4fb7-b5bd-3aeccca4e55b"), ShouldBeFalse)
	})
}


func TestMapStringAny_Get(t *testing.T) {
	Convey("TestMapStringAny.Get", t, func() {
		var k string = "f0a42645-4e2d-4b20-9d5e-6b52e750fe1f"
		var v interface{} = "e7d6356c-e125-45b8-8ab5-9e0cfe6a56a7"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("6111ed9c-ff8a-4d94-9d44-dfb15ec3a226" + "5674929c-f820-4c4e-87fb-e15c527a36a3")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringAny_GetOpt(t *testing.T) {
	Convey("TestMapStringAny.GetOpt", t, func() {
		var k string = "e0543537-16ac-4cd8-940b-8ce9cc009164"
		var v interface{} = "777e0a0b-ae96-4085-a5f9-643cf012be2d"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("a60850df-a61d-4fb7-afed-8bbdeac53369" + "ac357129-e4aa-4ebc-a840-e34147555f71")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringAny_ForEach(t *testing.T) {
	Convey("TestMapStringAny.ForEach", t, func() {
		var k string = "e8bee082-d93b-4ec6-ba92-ec7706305388"
		var v interface{} = "8091a917-dfe3-4f5e-96c5-611c2ee808ae"
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
		var k string = "c2d0710a-74c1-4380-bedc-902c0ce58054"
		var v interface{} = "d59a9e1a-fd0d-43f8-9b2a-f211d33179d9"

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
		var k string = "bc984cfe-bc8a-4ea6-b6ed-c887d31a8639"
		var v interface{} = "aeae91e2-535b-4918-a1e2-2799242c2f15"

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
		var k string = "098a5d0b-f953-4d37-9521-ec27196c316f"
		var v interface{} = "1cf5da4b-1421-46b2-8ef1-5fcd7d8bc998"

		test := omap.NewMapStringAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("ef5dd7d2-7983-4af0-bbbc-df31e2cda992", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "5ab76c4d-02b0-42c7-b4da-f54f585a3afc"
		So(test.PutIfNotNil("c93bd3d0-18ad-4132-bcb4-d4b511b09709", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringAny.ReplaceIfExists", t, func() {
		var k string = "def2c78b-58b3-4c2b-a069-f78e4676505f"
		var v interface{} = "b370660e-0eff-44d3-8a9b-4c06fdbbbfc9"
		var x interface{} = "f0b5032a-b33e-499a-98c5-4d5fdd959e26"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("dbf90828-46a3-401e-8031-eff02687b3b1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringAny.ReplaceOrPut", t, func() {
		var k string = "5689e129-0ebb-48c3-ad79-1268a88f4c11"
		var v interface{} = "99def607-5b47-4311-9801-4c0c6c64b77b"
		var x interface{} = "c0ce598d-fd8a-48c5-950e-03d7040e50ee"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("9b6470ce-dd99-47d2-93b9-32758eaca6ff", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_MarshalJSON(t *testing.T) {
	Convey("TestMapStringAny.MarshalJSON", t, func() {
		var k string = "5d9187e7-b346-4f28-b741-22997d179e12"
		var v interface{} = "8460ca81-43b3-44cc-8936-9ab85aed7f2e"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"5d9187e7-b346-4f28-b741-22997d179e12","value":"8460ca81-43b3-44cc-8936-9ab85aed7f2e"}]`)
	})
}

