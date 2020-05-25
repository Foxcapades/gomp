package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint8_Put(t *testing.T) {
	Convey("TestMapAnyUint8.Put", t, func() {
		var k interface{} = "94d3c241-51b2-423d-9a53-5537afe3ca79"
		var v uint8 = 71

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint8_Delete(t *testing.T) {
	Convey("TestMapAnyUint8.Delete", t, func() {
		var k interface{} = "17728d35-9a47-49b5-902e-f2830825f394"
		var v uint8 = 145

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint8_Has(t *testing.T) {
	Convey("TestMapAnyUint8.Has", t, func() {
		var k interface{} = "a52108b4-cdb1-4752-a34d-066a08736731"
		var v uint8 = 30

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("9aace314-77e5-48b7-beed-263554a9f257"+"4220eaf3-b823-4ca2-9e43-f589f34a75a3"), ShouldBeFalse)
	})
}

func TestMapAnyUint8_Get(t *testing.T) {
	Convey("TestMapAnyUint8.Get", t, func() {
		var k interface{} = "216862f3-4c36-4b64-88f4-2cd1c1ff7f1e"
		var v uint8 = 113

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("9ef4a0c2-2b91-466a-b2a2-253f764b0b01" + "b18267f7-62b9-4c54-9f44-3cc0f63d5432")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint8_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint8.GetOpt", t, func() {
		var k interface{} = "0b4b0ab4-4d93-4329-9507-ad046b79998d"
		var v uint8 = 83

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("ce3a60f0-c19b-4f2d-a8fd-435ea712ad05" + "6b50a4ab-b87f-4336-a606-4dbe905c06f0")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint8_ForEach(t *testing.T) {
	Convey("TestMapAnyUint8.ForEach", t, func() {
		var k interface{} = "1c437c1a-f315-49bb-962d-97d7a0164c1c"
		var v uint8 = 221
		hits := 0

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint8.MarshalYAML", t, func() {
		var k interface{} = "41b93a66-fb8c-4e00-a0eb-77dc68007098"
		var v uint8 = 207

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint8_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint8.ToYAML", t, func() {
		var k interface{} = "883f502c-0b0d-40d0-8b01-23290ae39119"
		var v uint8 = 106

		test := omap.NewMapAnyUint8(1)

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

func TestMapAnyUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint8.PutIfNotNil", t, func() {
		var k interface{} = "fec2b252-7b70-4ad2-9843-cf74afdbc27e"
		var v uint8 = 225

		test := omap.NewMapAnyUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("805803c6-3e80-4f80-9f8c-8b9c9a5b7eb8", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 33
		So(test.PutIfNotNil("81898f7e-9d63-4477-bc71-84124735f8e5", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceIfExists", t, func() {
		var k interface{} = "66fbd939-b32d-4711-a333-e17c74d21c16"
		var v uint8 = 77
		var x uint8 = 22

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("74ec969d-ae35-4da7-b09e-d94f41d29a2b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceOrPut", t, func() {
		var k interface{} = "c9c45567-24b8-489f-80a5-f3033a74c6c3"
		var v uint8 = 198
		var x uint8 = 124

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("6b209558-8289-418e-a2cf-317307aa5e42", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint8.MarshalJSON", t, func() {
		var k interface{} = "e0ddf359-c7a1-49b8-a2e9-50ab87490f85"
		var v uint8 = 99

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e0ddf359-c7a1-49b8-a2e9-50ab87490f85","value":99}]`)
	})
}
