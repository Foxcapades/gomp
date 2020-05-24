package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyBool_Put(t *testing.T) {
	Convey("TestMapAnyBool.Put", t, func() {
		var k interface{} = "8bb88bde-2b72-4f09-9b4f-33217df3b614"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyBool_Delete(t *testing.T) {
	Convey("TestMapAnyBool.Delete", t, func() {
		var k interface{} = "6ed34a50-a670-430f-9342-ba21ab8adab8"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyBool_Has(t *testing.T) {
	Convey("TestMapAnyBool.Has", t, func() {
		var k interface{} = "946a373d-4158-4c92-aca4-88136936bfd7"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("94ae4db2-1751-4b8d-a25e-473839cc4bcd"+"fb3f3e45-3c56-449e-bb23-bae86778349f"), ShouldBeFalse)
	})
}


func TestMapAnyBool_Get(t *testing.T) {
	Convey("TestMapAnyBool.Get", t, func() {
		var k interface{} = "5a2281c2-7de6-4c75-ab00-906a64e6b1e9"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("cb4c9133-6f3d-4598-8bd1-ee4f0f53c42f"+"1612194b-66c4-454a-94da-e5102c1ab13e")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyBool_GetOpt(t *testing.T) {
	Convey("TestMapAnyBool.GetOpt", t, func() {
		var k interface{} = "8c6bd4be-021b-419f-8e93-1bcae1c3ce31"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("74ccedf4-6572-477e-846d-bb37031e458a"+"a2c8d61a-f5c4-4a50-af55-fcd3b5a745fb")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyBool_ForEach(t *testing.T) {
	Convey("TestMapAnyBool.ForEach", t, func() {
		var k interface{} = "bc5ea744-8e94-4b81-8a33-a71ce1058e1e"
		var v bool = false
		hits := 0

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyBool_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyBool.MarshalYAML", t, func() {
		var k interface{} = "bd14120f-275a-497f-9006-4a5c267fde7d"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyBool_ToYAML(t *testing.T) {
	Convey("TestMapAnyBool.ToYAML", t, func() {
		var k interface{} = "2bbceca7-b7a4-40e4-b57a-4be7c56f87c7"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyBool.PutIfNotNil", t, func() {
		var k interface{} = "1ac944be-66be-46ba-a71f-4f0085eb04ec"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("6dd7b198-c553-4aa0-b014-4c4fd2be0de8", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("0b21a2e8-4f41-4fa8-a224-4e7801cdcf51", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceIfExists", t, func() {
		var k interface{} = "6ead8f9b-3a1d-42f5-8d16-a71a30f6ce5a"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("45d1c555-17e3-4b9e-8ea4-7099b60d1893", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceOrPut", t, func() {
		var k interface{} = "12e5a352-77dc-4431-ad8d-a5259c816824"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("6cd70149-23e9-440f-a5aa-3b71b6515fbb", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyBool.MarshalJSON", t, func() {
		var k interface{} = "b8eb653c-cace-4f97-a26f-213be6eb34cf"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"b8eb653c-cace-4f97-a26f-213be6eb34cf","value":false}]`)
	})
}

