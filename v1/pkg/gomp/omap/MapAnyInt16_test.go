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
		var k interface{} = "ed42fd7d-9e83-4a80-8dd8-7fa8f21b4c4a"
		var v int16 = 32456

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt16_Delete(t *testing.T) {
	Convey("TestMapAnyInt16.Delete", t, func() {
		var k interface{} = "60df61d2-23ff-4e63-b1b4-c6d3a10286c7"
		var v int16 = 18168

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt16_Has(t *testing.T) {
	Convey("TestMapAnyInt16.Has", t, func() {
		var k interface{} = "24200d40-4b29-4e9c-98c2-3c023fc7362f"
		var v int16 = 24247

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("fe745a98-39a8-4d6b-9332-6b68aacd954e"+"942ebc96-48cd-4adc-b2b9-51c0d9b6d35f"), ShouldBeFalse)
	})
}

func TestMapAnyInt16_Get(t *testing.T) {
	Convey("TestMapAnyInt16.Get", t, func() {
		var k interface{} = "6cbfd39d-f999-4a8b-8121-51b74b535a78"
		var v int16 = 468

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("28f79a21-1a4a-43c6-8b9b-876864e499af" + "5cf127fb-4304-490f-a1d6-44e92573a48e")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt16_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt16.GetOpt", t, func() {
		var k interface{} = "578d4d78-62a0-4623-88e6-cdbab3ef9803"
		var v int16 = 15408

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("d3f7d59d-27dd-468f-9105-630ac9d11e67" + "3d7fc3d5-8f15-4b19-a5e6-25c7b34f9901")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt16_ForEach(t *testing.T) {
	Convey("TestMapAnyInt16.ForEach", t, func() {
		var k interface{} = "63df078e-cbe2-4e9b-b42f-19208665b27a"
		var v int16 = 21083
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
		var k interface{} = "884af656-2094-4dc2-a3e2-47a3072b626c"
		var v int16 = 20660

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
		var k interface{} = "8eafbefb-fc5b-468a-a8f2-20bf0b975908"
		var v int16 = 3893

		test := omap.NewMapAnyInt16(1)

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

func TestMapAnyInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt16.PutIfNotNil", t, func() {
		var k interface{} = "244d4921-61b0-49b8-a6e0-2529e306328c"
		var v int16 = 31853

		test := omap.NewMapAnyInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("fc5973cd-a1f1-4399-b6e5-7b78025070f6", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 7721
		So(test.PutIfNotNil("df90a12b-7c3b-4235-a266-1080af5aa54b", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceIfExists", t, func() {
		var k interface{} = "91e839d2-b3a6-46ef-95f0-dea1f05a229d"
		var v int16 = 7829
		var x int16 = 29268

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("8556cd22-e58b-476a-b4c7-2ebe670c95a1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceOrPut", t, func() {
		var k interface{} = "3774e239-fd50-4485-900a-a3defa636e65"
		var v int16 = 5004
		var x int16 = 30114

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("34e3ff5d-fcc0-463c-aa49-263775f43fc6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt16.MarshalJSON", t, func() {
		var k interface{} = "f92e2839-f4ee-4d56-8f13-a06dbe0f04aa"
		var v int16 = 593

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"f92e2839-f4ee-4d56-8f13-a06dbe0f04aa","value":593}]`)
	})
}
