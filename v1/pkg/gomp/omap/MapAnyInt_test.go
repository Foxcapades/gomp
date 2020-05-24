package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt_Put(t *testing.T) {
	Convey("TestMapAnyInt.Put", t, func() {
		var k interface{} = "adc852a5-d89b-41a3-bb4c-0908db30f6a8"
		var v int = 1057570495

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt_Delete(t *testing.T) {
	Convey("TestMapAnyInt.Delete", t, func() {
		var k interface{} = "7a4d72e4-6964-4685-80ef-5f65136ff08b"
		var v int = 1579516850

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt_Has(t *testing.T) {
	Convey("TestMapAnyInt.Has", t, func() {
		var k interface{} = "611c2fda-2f0c-405c-827f-c6bd0a835469"
		var v int = 1202915671

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3a983d60-1a35-4bf1-8a84-468baa660dc7"+"b28c1975-433c-4081-9e38-10347fbce997"), ShouldBeFalse)
	})
}

func TestMapAnyInt_Get(t *testing.T) {
	Convey("TestMapAnyInt.Get", t, func() {
		var k interface{} = "17fba1ae-1dad-41c0-aa5f-7c7b49400063"
		var v int = 318880867

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("631eed65-4ac5-4e71-8d06-46eee2af4dde" + "50b01796-fedf-4c81-a4cc-b36d28c0f490")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt.GetOpt", t, func() {
		var k interface{} = "5d5c75b8-1130-4eaf-8a55-4f8c59913066"
		var v int = 1538154506

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("ba5fde13-94f0-48af-878a-76dc74dcba59" + "86489da0-06ed-46f8-aea0-bf61d928e5e7")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt_ForEach(t *testing.T) {
	Convey("TestMapAnyInt.ForEach", t, func() {
		var k interface{} = "c4e14af0-a4e7-4a84-b976-0ec30e75c083"
		var v int = 534765927
		hits := 0

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt.MarshalYAML", t, func() {
		var k interface{} = "4e2859ff-7322-47ce-b12e-148d249816a3"
		var v int = 724519061

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt.ToYAML", t, func() {
		var k interface{} = "aa1fda43-f05d-4173-8558-51f63a7bb344"
		var v int = 612380518

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt.PutIfNotNil", t, func() {
		var k interface{} = "bda92905-75d3-492c-b347-9afa2de9b0b8"
		var v int = 2003548252

		test := omap.NewMapAnyInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("406b2568-6d99-445a-b6df-bc74d6bc0d35", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 1479190083
		So(test.PutIfNotNil("ba6187f2-6776-4a3c-8662-d98165c270c4", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceIfExists", t, func() {
		var k interface{} = "2556cd7b-cd16-46ef-8267-4f81051db903"
		var v int = 203417402
		var x int = 1792998722

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("413e751d-6d0f-4a26-a693-2c78d6b31704", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceOrPut", t, func() {
		var k interface{} = "84633ebd-6b55-4aed-ab68-63c1b217b376"
		var v int = 645896543
		var x int = 1638706026

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("2093ec6c-23bd-4c69-b0e4-c590cd4f7754", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt.MarshalJSON", t, func() {
		var k interface{} = "75e5de8b-50eb-43d8-a0c7-893efb0f59bc"
		var v int = 218127091

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"75e5de8b-50eb-43d8-a0c7-893efb0f59bc","value":218127091}]`)
	})
}
