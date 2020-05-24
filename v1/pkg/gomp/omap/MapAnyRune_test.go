package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyRune_Put(t *testing.T) {
	Convey("TestMapAnyRune.Put", t, func() {
		var k interface{} = "33b9742e-6ea0-4f5c-92fc-41af10d62cfe"
		var v rune = 1260517990

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyRune_Delete(t *testing.T) {
	Convey("TestMapAnyRune.Delete", t, func() {
		var k interface{} = "265255ea-3d2f-4eb1-ab04-41cff5b5d680"
		var v rune = 772408207

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyRune_Has(t *testing.T) {
	Convey("TestMapAnyRune.Has", t, func() {
		var k interface{} = "fe0fc97e-ad33-405d-a52c-6267326f123e"
		var v rune = 159415927

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("8256990c-e852-4e69-9ffd-7dd4e205e54c"+"e36a9d7b-6191-4b38-a119-df29cce070a7"), ShouldBeFalse)
	})
}

func TestMapAnyRune_Get(t *testing.T) {
	Convey("TestMapAnyRune.Get", t, func() {
		var k interface{} = "65ef5edb-e06f-489c-86c8-99af08e8eb69"
		var v rune = 743097955

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("c45c19e3-ed90-4a9f-aa08-53875269d7dc" + "b802e826-7c82-44dd-a7e0-fb5341e05274")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyRune_GetOpt(t *testing.T) {
	Convey("TestMapAnyRune.GetOpt", t, func() {
		var k interface{} = "eb4b7a98-7012-425a-b257-f4070f54a5a7"
		var v rune = 1697524104

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("df45cf1c-e8ee-4f93-9303-14c1e44f8b95" + "f27fca73-4e6c-4ed3-8135-09475a4ed99c")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyRune_ForEach(t *testing.T) {
	Convey("TestMapAnyRune.ForEach", t, func() {
		var k interface{} = "1865f6d7-173f-48e8-9742-c0b33460ebf4"
		var v rune = 683840234
		hits := 0

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv rune) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyRune_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyRune.MarshalYAML", t, func() {
		var k interface{} = "712335c7-cb2a-4e66-9663-71d2824548da"
		var v rune = 1220934261

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyRune_ToYAML(t *testing.T) {
	Convey("TestMapAnyRune.ToYAML", t, func() {
		var k interface{} = "944bcbe8-5145-43f8-bb0e-eee86b29a1f6"
		var v rune = 432754424

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyRune.PutIfNotNil", t, func() {
		var k interface{} = "f4e1f084-7c4b-4eea-85b5-a6459439eac2"
		var v rune = 2057230812

		test := omap.NewMapAnyRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("bcf77c65-c24f-40a2-b085-700241a61bbe", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1217418772
		So(test.PutIfNotNil("93c52a2b-366d-4bc9-bd1e-7f495be0c1bd", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceIfExists", t, func() {
		var k interface{} = "3673ddc5-bc8c-43a1-ac3a-7cb1f36855a1"
		var v rune = 1523857781
		var x rune = 1916083252

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("cfd3ddb0-e4b6-40c4-afd7-5deffc681005", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceOrPut", t, func() {
		var k interface{} = "664f0044-c8ea-4504-8804-01149877fe4b"
		var v rune = 1301715308
		var x rune = 1840657048

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("756f23fe-fd6d-4b7d-a375-1c217c435015", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyRune.MarshalJSON", t, func() {
		var k interface{} = "c12b8216-4221-4917-a7cc-07790a64d05d"
		var v rune = 2084212532

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"c12b8216-4221-4917-a7cc-07790a64d05d","value":2084212532}]`)
	})
}
