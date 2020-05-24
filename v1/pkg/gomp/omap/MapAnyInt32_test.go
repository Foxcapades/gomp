package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt32_Put(t *testing.T) {
	Convey("TestMapAnyInt32.Put", t, func() {
		var k interface{} = "9ac2fc8d-ef65-4854-b423-cb498ccf79bf"
		var v int32 = 467365531

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt32_Delete(t *testing.T) {
	Convey("TestMapAnyInt32.Delete", t, func() {
		var k interface{} = "347de7f3-9505-4549-8322-96a5492f373c"
		var v int32 = 1639737541

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt32_Has(t *testing.T) {
	Convey("TestMapAnyInt32.Has", t, func() {
		var k interface{} = "222d0166-aed1-4cea-ae63-f6e3d5ff7e64"
		var v int32 = 358779560

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d84d3868-713d-4406-863b-9721e5210f05"+"852d0bb6-8cb0-4ae6-8346-79d7c05fc2d1"), ShouldBeFalse)
	})
}

func TestMapAnyInt32_Get(t *testing.T) {
	Convey("TestMapAnyInt32.Get", t, func() {
		var k interface{} = "b718d0f3-7fab-46f3-8dd9-d7dd649711f4"
		var v int32 = 157113616

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("4ab6b26e-687e-4d34-8c3b-e7c39c0b85fd" + "137dc181-23ec-4542-8a12-c461ee0d4d9e")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt32_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt32.GetOpt", t, func() {
		var k interface{} = "062055af-cc23-4bdc-a8dc-64a8b12d828d"
		var v int32 = 283212872

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("60975fa8-0a04-4ab3-b90c-cc17fca45fb6" + "894be8ec-6d62-4c4e-b86c-9fe249665b75")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt32_ForEach(t *testing.T) {
	Convey("TestMapAnyInt32.ForEach", t, func() {
		var k interface{} = "ecafba0d-0fed-42c2-aba8-4e460b088ddf"
		var v int32 = 1094776019
		hits := 0

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalYAML", t, func() {
		var k interface{} = "20634f67-e119-419f-a3b1-4d3c37cc0765"
		var v int32 = 515792164

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt32_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt32.ToYAML", t, func() {
		var k interface{} = "79e82c12-2d7e-4192-9ce4-fa49edeb9eef"
		var v int32 = 1847734385

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt32.PutIfNotNil", t, func() {
		var k interface{} = "be6fa9b8-970a-4e9b-aa0a-60178ac6219e"
		var v int32 = 320676112

		test := omap.NewMapAnyInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("480f33f3-a38b-424c-935e-cf4904ee35cb", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 268632063
		So(test.PutIfNotNil("84bfd21e-fd05-41b3-89aa-bd71fb1a0caf", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceIfExists", t, func() {
		var k interface{} = "36727f4b-66d3-4252-b519-68255706c728"
		var v int32 = 1489667727
		var x int32 = 647258554

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("bd7c7a38-1b44-4632-be3f-02d255b16568", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceOrPut", t, func() {
		var k interface{} = "dbb5c114-6092-4360-94b1-33212f19106f"
		var v int32 = 1751334151
		var x int32 = 2067251546

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("1718f385-7aab-4b10-9669-3bed56b2bc11", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalJSON", t, func() {
		var k interface{} = "37ed1266-9eaf-4406-aab5-7d7684d76d9a"
		var v int32 = 451168577

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"37ed1266-9eaf-4406-aab5-7d7684d76d9a","value":451168577}]`)
	})
}
