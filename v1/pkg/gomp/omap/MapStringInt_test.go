package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt_Put(t *testing.T) {
	Convey("TestMapStringInt.Put", t, func() {
		var k string = "795f89b4-a2bf-46dc-b6a7-e6a098ac1e1d"
		var v int = 585741278

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt_Delete(t *testing.T) {
	Convey("TestMapStringInt.Delete", t, func() {
		var k string = "bbe692de-9389-40cb-9f0c-4e687034698e"
		var v int = 1044186164

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt_Has(t *testing.T) {
	Convey("TestMapStringInt.Has", t, func() {
		var k string = "4d1866ea-bd68-49ee-b245-3d8167c335ad"
		var v int = 416009941

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("99de9359-2684-47fb-9f0e-5b96304056c9"+"069dfa53-2279-4011-a588-be1af6953f4f"), ShouldBeFalse)
	})
}

func TestMapStringInt_Get(t *testing.T) {
	Convey("TestMapStringInt.Get", t, func() {
		var k string = "5331e27e-5feb-4cf1-a3c6-43b9808b3d7e"
		var v int = 526551930

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("bb48c2a4-76d3-413d-accb-d7f132fbcb8a" + "8aae68bc-8631-48f0-ac9e-3bbe8d0ead20")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt_GetOpt(t *testing.T) {
	Convey("TestMapStringInt.GetOpt", t, func() {
		var k string = "d1cad8b5-159c-4724-8316-fa711cb7b80f"
		var v int = 859884144

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b8f05f42-3a33-43c8-bc88-6af89fe83b97" + "03a40376-6833-42c7-8274-e26d090e61cf")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt_ForEach(t *testing.T) {
	Convey("TestMapStringInt.ForEach", t, func() {
		var k string = "195c6869-335d-4d72-8877-8b7c1decb5ab"
		var v int = 1497991022
		hits := 0

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt.MarshalYAML", t, func() {
		var k string = "eac87fce-2eb1-49b4-8eac-b5aebea076e0"
		var v int = 622055052

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt_ToYAML(t *testing.T) {
	Convey("TestMapStringInt.ToYAML", t, func() {
		var k string = "05931162-352c-401b-852d-d7de4bebfa51"
		var v int = 2113613686

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt.PutIfNotNil", t, func() {
		var k string = "1dde5608-f53c-457e-abbd-56e86e6c391f"
		var v int = 1234769798

		test := omap.NewMapStringInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("d7b2983e-372e-461f-8da0-2afbeff22799", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 1226740187
		So(test.PutIfNotNil("0d65e436-fc0c-4057-8b0d-2396cd59b61e", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt.ReplaceIfExists", t, func() {
		var k string = "25b09acc-b068-4832-957f-a078fd9ff3d9"
		var v int = 2126572789
		var x int = 2009336853

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("3574324c-a4a6-40fd-8f6e-63d6b4b76894", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt.ReplaceOrPut", t, func() {
		var k string = "bbdef0f5-4191-45db-badb-6a63373a2e69"
		var v int = 1381323023
		var x int = 1796452043

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("4038edd0-ec8c-4e01-bc47-4c5916857edd", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt.MarshalJSON", t, func() {
		var k string = "e4ded5db-f264-4e62-8c28-382a16b99ca5"
		var v int = 359221359

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e4ded5db-f264-4e62-8c28-382a16b99ca5","value":359221359}]`)
	})
}
