package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringBool_Put(t *testing.T) {
	Convey("TestMapStringBool.Put", t, func() {
		var k string = "ef0bda50-7cf4-4f7a-ab26-8bf14f0ad9d7"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringBool_Delete(t *testing.T) {
	Convey("TestMapStringBool.Delete", t, func() {
		var k string = "b6dad4ce-de39-41f3-b386-9ec906be6c10"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringBool_Has(t *testing.T) {
	Convey("TestMapStringBool.Has", t, func() {
		var k string = "f940d60e-5a17-4902-9196-faa39d7cd26f"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b87fcd38-e97a-4ece-badf-4a660390e7a3"+"026f5357-9331-4924-b258-69a4a9810c08"), ShouldBeFalse)
	})
}

func TestMapStringBool_Get(t *testing.T) {
	Convey("TestMapStringBool.Get", t, func() {
		var k string = "7ce3bfcf-039d-4088-8296-def7c93e5744"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("1e9ad8de-d025-4b8e-82d1-b07f8f7b1188" + "1ec03ba6-eb5a-4e44-bede-7e742eacf654")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringBool_GetOpt(t *testing.T) {
	Convey("TestMapStringBool.GetOpt", t, func() {
		var k string = "ec114e53-08a8-491e-a590-2a327cf5d3e5"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("d91e92d7-1a99-47ed-b286-811b0d053610" + "44a047a5-8706-4f37-9d6e-8cd2fa163e1a")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringBool_ForEach(t *testing.T) {
	Convey("TestMapStringBool.ForEach", t, func() {
		var k string = "9d48da34-cd5b-4d3c-b1a2-b41771c02e04"
		var v bool = false
		hits := 0

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringBool_MarshalYAML(t *testing.T) {
	Convey("TestMapStringBool.MarshalYAML", t, func() {
		var k string = "69fa6617-6116-452a-8ac2-4c3e77f58c0e"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringBool_ToYAML(t *testing.T) {
	Convey("TestMapStringBool.ToYAML", t, func() {
		var k string = "7256770c-f508-441a-8d9a-f53caeb393c6"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringBool.PutIfNotNil", t, func() {
		var k string = "93b7b792-f1ff-44d7-96d8-65485b313ad0"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("3af38de9-72fc-4342-893a-1bf7bf12bee0", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("d5af1d2c-cea8-4fdc-8a12-937a65726485", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringBool.ReplaceIfExists", t, func() {
		var k string = "a61b4fda-a585-44c0-bf80-b2b473c3c332"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("0f4a2599-2a36-47bd-956e-5a9ea56916e2", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringBool.ReplaceOrPut", t, func() {
		var k string = "cd067a75-e749-445f-b000-26be18d9b859"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("52f7d7b8-e3b8-48ee-9046-eab95e1305b4", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_MarshalJSON(t *testing.T) {
	Convey("TestMapStringBool.MarshalJSON", t, func() {
		var k string = "0f883d19-8ca7-45cb-b9c0-5a3860bcf661"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"0f883d19-8ca7-45cb-b9c0-5a3860bcf661","value":false}]`)
	})
}
