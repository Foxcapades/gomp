package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt8_Put(t *testing.T) {
	Convey("TestMapStringInt8.Put", t, func() {
		var k string = "aed53990-35a8-49e1-b5e0-1d962c09a61a"
		var v int8 = 79

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt8_Delete(t *testing.T) {
	Convey("TestMapStringInt8.Delete", t, func() {
		var k string = "d151caf3-f949-4c89-a7c7-e272812fe96b"
		var v int8 = 92

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt8_Has(t *testing.T) {
	Convey("TestMapStringInt8.Has", t, func() {
		var k string = "aba0ca4f-88a7-449b-9c7a-69edbdea495e"
		var v int8 = 19

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5e00e46a-4516-4c59-a6b1-8cff952d9dcc"+"d0c23f49-adcb-40ea-830b-a4ef5dbd01ee"), ShouldBeFalse)
	})
}

func TestMapStringInt8_Get(t *testing.T) {
	Convey("TestMapStringInt8.Get", t, func() {
		var k string = "834a575e-e9d8-4758-a290-c751a0fe08fa"
		var v int8 = 40

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("8563b587-4f3a-4e7d-8aed-79c502c3fb10" + "4bbf20ff-1090-4e2b-89e4-3be69cbe76d0")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt8_GetOpt(t *testing.T) {
	Convey("TestMapStringInt8.GetOpt", t, func() {
		var k string = "e850ebed-43e8-4c28-967b-cb23413bbb39"
		var v int8 = 39

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("69eeaa92-ade6-4737-9e37-25bc733a3ce6" + "f6e34605-a6e3-43b6-8192-2cdcacc5a05b")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt8_ForEach(t *testing.T) {
	Convey("TestMapStringInt8.ForEach", t, func() {
		var k string = "1fa25ca5-7247-4d6d-b6b7-fb255a3e5a86"
		var v int8 = 59
		hits := 0

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt8.MarshalYAML", t, func() {
		var k string = "e2c21347-5712-4f92-bab9-d7a144a9d136"
		var v int8 = 126

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt8_ToYAML(t *testing.T) {
	Convey("TestMapStringInt8.ToYAML", t, func() {
		var k string = "28bb09c5-72d2-490f-9bbb-639c9cc9902f"
		var v int8 = 89

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt8.PutIfNotNil", t, func() {
		var k string = "7e97a9b4-5c0d-40fb-84b3-af6ff276fa17"
		var v int8 = 21

		test := omap.NewMapStringInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("050f07de-f6c0-4c3f-8a25-ee14f3bc368c", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 113
		So(test.PutIfNotNil("9cf43421-b9dc-48d3-877e-b78e14e38876", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceIfExists", t, func() {
		var k string = "82dc7870-122c-401d-9c72-c09e96b78ddd"
		var v int8 = 80
		var x int8 = 57

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("31d809f3-d74f-4b46-bc56-1ca95ae8e833", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceOrPut", t, func() {
		var k string = "0d8a158d-c3ff-47f5-9581-1caf4623ed8a"
		var v int8 = 3
		var x int8 = 40

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("36846879-68c5-4e21-a5bd-bdf1de6432d3", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt8.MarshalJSON", t, func() {
		var k string = "1b07f286-0392-4a95-a12e-94ad82746df7"
		var v int8 = 36

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"1b07f286-0392-4a95-a12e-94ad82746df7","value":36}]`)
	})
}
