package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt32_Put(t *testing.T) {
	Convey("TestMapStringInt32.Put", t, func() {
		var k string = "1823a57b-5ec8-49bf-b619-b845aa1e6f57"
		var v int32 = 817125795

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt32_Delete(t *testing.T) {
	Convey("TestMapStringInt32.Delete", t, func() {
		var k string = "f40c52f7-98b9-43e2-8e4a-6970cbaf0320"
		var v int32 = 16479037

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt32_Has(t *testing.T) {
	Convey("TestMapStringInt32.Has", t, func() {
		var k string = "c6fe3504-3404-418b-8550-680eff058bbf"
		var v int32 = 264368154

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d6d9a9f0-8895-4471-ad29-5de3137dd47a"+"d186890f-a5a9-4661-95d1-6ec51706be5f"), ShouldBeFalse)
	})
}


func TestMapStringInt32_Get(t *testing.T) {
	Convey("TestMapStringInt32.Get", t, func() {
		var k string = "83a08648-9128-4eba-b6f8-5e22705d96da"
		var v int32 = 1523971306

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("dd864155-8796-4dd6-86bc-eb17bee5cea6"+"33f6d5b9-b95f-44fb-8c99-290ac39fdde6")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt32_GetOpt(t *testing.T) {
	Convey("TestMapStringInt32.GetOpt", t, func() {
		var k string = "bec616fc-92f0-4595-93e6-9389d0758947"
		var v int32 = 1921574758

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("42ba564a-e045-4795-8f52-32eac8c9f0dc"+"8941600a-d672-49ad-856f-cf19b8b4be15")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt32_ForEach(t *testing.T) {
	Convey("TestMapStringInt32.ForEach", t, func() {
		var k string = "2b5a75e6-d784-453c-aa62-e6abd93489c3"
		var v int32 = 2606282
		hits := 0

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt32.MarshalYAML", t, func() {
		var k string = "ec1b36ed-2a78-4723-9b6d-a085d18272af"
		var v int32 = 1271456055

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt32_ToYAML(t *testing.T) {
	Convey("TestMapStringInt32.ToYAML", t, func() {
		var k string = "c4a4eed3-936d-4156-b40a-a7af446c1d6e"
		var v int32 = 1019368667

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt32.PutIfNotNil", t, func() {
		var k string = "1b1a02cb-2235-46b1-8289-439a75f411bb"
		var v int32 = 1908650463

		test := omap.NewMapStringInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("c927e050-d190-422b-bec0-dd1ca81d3b6e", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 819147274
		So(test.PutIfNotNil("7d7e3a60-8462-4a74-8000-923c99e55d77", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceIfExists", t, func() {
		var k string = "50f7df1d-40ed-4f2a-a32b-0ae99cc7745b"
		var v int32 = 2022447581
		var x int32 = 1538443132

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("a7118dcf-8e4f-4cef-9337-58b3a0e178f0", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceOrPut", t, func() {
		var k string = "fe07f2e5-d76a-452a-b036-d2e2addf551e"
		var v int32 = 292988578
		var x int32 = 12155796

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("28693fdd-0d36-44e6-98ba-be202ff62bb1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt32.MarshalJSON", t, func() {
		var k string = "a96befc0-dba1-45ce-8c6a-3b21f354b8f7"
		var v int32 = 981957051

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a96befc0-dba1-45ce-8c6a-3b21f354b8f7","value":981957051}]`)
	})
}

