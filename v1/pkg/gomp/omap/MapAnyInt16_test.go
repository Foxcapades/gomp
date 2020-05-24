package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt16_Put(t *testing.T) {
	Convey("TestMapAnyInt16.Put", t, func() {
		var k interface{} = "3bb4b15b-3784-469e-9ce1-dfd742016684"
		var v int16 = 29058

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt16_Delete(t *testing.T) {
	Convey("TestMapAnyInt16.Delete", t, func() {
		var k interface{} = "ce217cc3-66f6-4346-ad36-97983691a0a9"
		var v int16 = 21275

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt16_Has(t *testing.T) {
	Convey("TestMapAnyInt16.Has", t, func() {
		var k interface{} = "7775bfaf-eaae-4ebe-b03c-0be6c6258966"
		var v int16 = 14057

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("272a6d71-db2c-4c15-9bd4-18fc2ea400f8"+"8e858ece-f90e-4a00-bdf2-fc5352a59c9b"), ShouldBeFalse)
	})
}


func TestMapAnyInt16_Get(t *testing.T) {
	Convey("TestMapAnyInt16.Get", t, func() {
		var k interface{} = "37eb609f-44b6-413b-a644-4f1d0fed6406"
		var v int16 = 17024

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("270b15bd-12a6-4d96-878c-c2200d2711ec" + "11484053-c83d-4580-bf38-47d6d240a27f")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt16_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt16.GetOpt", t, func() {
		var k interface{} = "fc71ead6-89e8-4db0-a451-ca3b201ebcb3"
		var v int16 = 20361

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("fb2864f2-6624-4cbb-bbcb-51e5ffe233f4" + "f231e812-8947-4388-9c95-c4d910e56b9e")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt16_ForEach(t *testing.T) {
	Convey("TestMapAnyInt16.ForEach", t, func() {
		var k interface{} = "fe8f13bd-f9a0-4b4e-b8a5-faf9bcf306ec"
		var v int16 = 2771
		hits := 0

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt16.MarshalYAML", t, func() {
		var k interface{} = "b0612ef3-dddc-4635-abb5-b32383a64d4a"
		var v int16 = 5642

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt16_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt16.ToYAML", t, func() {
		var k interface{} = "38380ae2-86b7-4fc4-9420-146ff7c0ff81"
		var v int16 = 18463

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt16.PutIfNotNil", t, func() {
		var k interface{} = "56bb1e79-a79b-41eb-804c-dc21d5fda4fc"
		var v int16 = 6136

		test := omap.NewMapAnyInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b460aa34-0527-493b-ac6f-52d0cf103a4e", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 14256
		So(test.PutIfNotNil("a6b7622f-9a41-4d7b-8baa-9a8d0cd539da", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceIfExists", t, func() {
		var k interface{} = "cb5b425f-16bc-4356-a505-7c16d491ff97"
		var v int16 = 7657
		var x int16 = 9636

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("810a20be-3b9c-4beb-873b-59b201e4e5cc", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceOrPut", t, func() {
		var k interface{} = "f4889590-397e-4b28-9330-2abc9ded84e8"
		var v int16 = 17376
		var x int16 = 17223

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("fc485190-a0d3-446d-ae91-505d11894769", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt16.MarshalJSON", t, func() {
		var k interface{} = "9d983adc-5cd7-4ae1-b4f0-9e748b412525"
		var v int16 = 29574

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"9d983adc-5cd7-4ae1-b4f0-9e748b412525","value":29574}]`)
	})
}
