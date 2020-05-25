package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyByte_Put(t *testing.T) {
	Convey("TestMapAnyByte.Put", t, func() {
		var k interface{} = "bd383940-4b45-44c3-98c1-573672693643"
		var v byte = 197

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyByte_Delete(t *testing.T) {
	Convey("TestMapAnyByte.Delete", t, func() {
		var k interface{} = "5467b903-04af-47c0-adc7-cd875eb2eadf"
		var v byte = 171

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyByte_Has(t *testing.T) {
	Convey("TestMapAnyByte.Has", t, func() {
		var k interface{} = "3ca659bc-3e50-400a-b7a9-1e76465d2f64"
		var v byte = 85

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5429fc85-b775-45b5-8533-ad6167ea7b95"+"306d39f8-6481-4352-9447-0258911065db"), ShouldBeFalse)
	})
}

func TestMapAnyByte_Get(t *testing.T) {
	Convey("TestMapAnyByte.Get", t, func() {
		var k interface{} = "f7abeb47-76e2-4e73-93ef-b6fc0eed07d3"
		var v byte = 246

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("5806f8a3-f8f8-4f84-bafe-28b97296b043" + "bd2385f6-1dd4-4944-887a-47896d99f3b6")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyByte_GetOpt(t *testing.T) {
	Convey("TestMapAnyByte.GetOpt", t, func() {
		var k interface{} = "57bfdfe5-a9b0-4b7f-bb59-9b4690ea0c4d"
		var v byte = 7

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("cf38a139-c753-4775-953e-57c89365e55e" + "1a1ea61b-ebdf-4871-be98-cb72bfb6ecc5")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyByte_ForEach(t *testing.T) {
	Convey("TestMapAnyByte.ForEach", t, func() {
		var k interface{} = "9e99d63f-c78a-43f4-8e33-6d1569e63aca"
		var v byte = 10
		hits := 0

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyByte_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyByte.MarshalYAML", t, func() {
		var k interface{} = "f39c63cd-8a63-4859-85c2-838d3fb265a9"
		var v byte = 80

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyByte_ToYAML(t *testing.T) {
	Convey("TestMapAnyByte.ToYAML", t, func() {
		var k interface{} = "ee6d0633-f887-48c3-90fc-44be2436d062"
		var v byte = 29

		test := omap.NewMapAnyByte(1)

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

func TestMapAnyByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyByte.PutIfNotNil", t, func() {
		var k interface{} = "021b1c34-c345-4aa2-a148-8781d339fa06"
		var v byte = 248

		test := omap.NewMapAnyByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("001409d9-2dbb-4aa4-821e-93e569822b37", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 85
		So(test.PutIfNotNil("db572b70-21f1-4822-ae64-a0c846dd9922", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceIfExists", t, func() {
		var k interface{} = "cc171d08-4ebf-46b1-8b7f-70f644beed4a"
		var v byte = 126
		var x byte = 135

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("c76a6214-d8d4-491c-ac7c-074497d428d3", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceOrPut", t, func() {
		var k interface{} = "3885d691-e20f-4285-8683-03064cb1c377"
		var v byte = 68
		var x byte = 128

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("2a11b99f-aded-400e-9a84-db3fde46fa5b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyByte.MarshalJSON", t, func() {
		var k interface{} = "e432e33d-ac4c-48e4-a2dc-345cfdd0df1a"
		var v byte = 176

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e432e33d-ac4c-48e4-a2dc-345cfdd0df1a","value":176}]`)
	})
}
