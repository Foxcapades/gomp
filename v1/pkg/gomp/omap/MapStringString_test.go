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
		var k string = "d189f0f3-f92d-418d-bce8-407e20aa4958"
		var v string = "4c1abee5-6b22-40d1-9255-82146abccf58"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringString_Delete(t *testing.T) {
	Convey("TestMapStringString.Delete", t, func() {
		var k string = "21f97b5e-a40e-4a6f-9768-1e69b8f104a3"
		var v string = "55780b17-7f41-4ab7-aa40-3582eb622397"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringString_Has(t *testing.T) {
	Convey("TestMapStringString.Has", t, func() {
		var k string = "80e44dba-60ae-4bc4-9b7d-d71a22e650c3"
		var v string = "52eaf0c2-4631-456d-8f70-a2e02c3e5965"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("af3eb4c0-3949-401b-bd42-32dc0d96e23b"+"333d1ba9-18d7-4d41-8962-f54165097bc8"), ShouldBeFalse)
	})
}

func TestMapStringString_Get(t *testing.T) {
	Convey("TestMapStringString.Get", t, func() {
		var k string = "636a035c-e05c-4494-9178-11c837e8462a"
		var v string = "fb5a64df-9e60-411e-a37b-322083b3fdbf"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("0af18e76-bbd5-4462-9938-5748d1e841ee" + "9a6b5984-06bf-4f3e-bdcb-86261a9a6c86")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringString_GetOpt(t *testing.T) {
	Convey("TestMapStringString.GetOpt", t, func() {
		var k string = "d9fe0be4-e3b1-4304-90f8-627f1d0f43cf"
		var v string = "93233219-a0e1-4d1d-bb45-f369ad81698a"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("6ca0495f-0234-4d28-bfd6-d0c80848ea69" + "53b09292-4a57-4e9c-835d-fe8e639878be")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringString_ForEach(t *testing.T) {
	Convey("TestMapStringString.ForEach", t, func() {
		var k string = "e31787cc-279b-4491-98da-5fb4cca9b6e5"
		var v string = "94518242-119d-4d12-8716-7aabd844e6a0"
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
		var k string = "07cd89a6-23d6-436d-9118-6e9239897ba8"
		var v string = "8d6a8324-d1d2-4d12-a106-2b77f3ec59fa"

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
		Convey("Ordered", func() {
			var k string = "3ef0278b-5b1e-4896-8b83-2729980455b9"
			var v string = "aefad4e0-88d8-4bab-adfd-f69093577f44"

			test := omap.NewMapStringString(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k string = "596ead08-3bf5-4681-8393-241c9d0e5024"
			var v string = "c72b423b-f845-4e7b-b70f-43cb04f764d1"

			test := omap.NewMapStringString(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapStringString_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringString.PutIfNotNil", t, func() {
		var k string = "e5d14f6b-dccd-4fb2-9cbd-4f7488f41022"
		var v string = "8800dc20-b1a2-4a5e-9b52-c603dab559a2"

		test := omap.NewMapStringString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("01b2a16c-dd29-4b49-9ba8-2fbb70d3c93a", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "91972caf-8e06-48fd-8ab7-a5a7310974e0"
		So(test.PutIfNotNil("52095d94-7f72-4c16-8309-a7b2e0320515", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringString.ReplaceIfExists", t, func() {
		var k string = "cc6741e3-86a0-4872-9ab4-a554af19b868"
		var v string = "7dbb8c76-6648-4349-8405-2b938164aa05"
		var x string = "2398dfd8-5c78-4c96-b8b4-7cfe6e039866"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("3a407c3b-90ae-4c2e-b7c3-7d1fd270c4c6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringString.ReplaceOrPut", t, func() {
		var k string = "5b8b2510-b428-4617-8d8b-eea4c84407c5"
		var v string = "6ba62b11-40f2-482f-8bce-319b6f91aea8"
		var x string = "2946abe1-4be9-41a9-bceb-5b5f02438b69"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("8c1d83d1-476a-447b-be1a-1eeb6540ec1a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_MarshalJSON(t *testing.T) {
	Convey("TestMapStringString.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "03454523-451b-4371-a16e-5e885f65f997"
			var v string = "7a47727d-9c34-4ae7-afe3-8151ecf32805"

			test := omap.NewMapStringString(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"03454523-451b-4371-a16e-5e885f65f997","value":"7a47727d-9c34-4ae7-afe3-8151ecf32805"}]`)
		})

		Convey("Unordered", func() {
			var k string = "03454523-451b-4371-a16e-5e885f65f997"
			var v string = "7a47727d-9c34-4ae7-afe3-8151ecf32805"

			test := omap.NewMapStringString(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"03454523-451b-4371-a16e-5e885f65f997":"7a47727d-9c34-4ae7-afe3-8151ecf32805"}`)
		})

	})
}
