package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringAny_Put(t *testing.T) {
	Convey("TestMapStringAny.Put", t, func() {
		var k string = "d2a74608-a158-49ff-b83f-ccc969da500f"
		var v interface{} = "6b7cf62c-afbc-4d83-9720-f4d9ee27d4e7"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringAny_Delete(t *testing.T) {
	Convey("TestMapStringAny.Delete", t, func() {
		var k string = "0bcbd204-bdc8-4ef0-b6bc-5e422f3d53b4"
		var v interface{} = "07859a5c-3201-4df3-b1d7-6c2b1bbb82e0"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringAny_Has(t *testing.T) {
	Convey("TestMapStringAny.Has", t, func() {
		var k string = "05c78274-688f-4680-b61d-6e2c69acc2a2"
		var v interface{} = "075dfe00-9cf9-4b8f-a966-515ca6ff53f7"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d5631e28-6023-4d9d-b9dc-01eaf6fc67a3"+"efe2cd9b-dcd2-40b0-bb72-89d70e446e2c"), ShouldBeFalse)
	})
}

func TestMapStringAny_Get(t *testing.T) {
	Convey("TestMapStringAny.Get", t, func() {
		var k string = "08926474-4932-4630-b3e4-cad8c48693a7"
		var v interface{} = "d86fe086-1e39-4a21-a527-91d551397bd8"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("afeb24e1-59da-4903-81cd-5ba7df8d41da" + "4e41e7da-e379-4867-929a-108d11da1d3b")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringAny_GetOpt(t *testing.T) {
	Convey("TestMapStringAny.GetOpt", t, func() {
		var k string = "7267c7f5-a3b5-46a7-b792-f2e88bf6e8d6"
		var v interface{} = "fd56bdc4-e033-4796-a59d-baab0a99b1cc"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("bd8c4ef4-5f63-4ae5-b478-1bbb88854bf7" + "a6517ffa-3c73-40d0-9222-2254d775d30a")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringAny_ForEach(t *testing.T) {
	Convey("TestMapStringAny.ForEach", t, func() {
		var k string = "9e38a5b8-2948-4abc-a27f-7215a0f1d047"
		var v interface{} = "7bffeb60-e407-4fe5-b4de-b4ea9f3bd37d"
		hits := 0

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringAny_MarshalYAML(t *testing.T) {
	Convey("TestMapStringAny.MarshalYAML", t, func() {
		var k string = "c643e047-cc46-4d13-a091-e9f87c9015e0"
		var v interface{} = "6b4c395f-175b-45c8-9538-e67317e37655"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringAny_ToYAML(t *testing.T) {
	Convey("TestMapStringAny.ToYAML", t, func() {
		var k string = "ff9460fa-35da-4fa0-ba8d-b19188889223"
		var v interface{} = "e19ae175-e31e-4cfc-b10e-26fe22448726"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringAny.PutIfNotNil", t, func() {
		var k string = "9b8928d5-6a18-4317-9644-c1a316339af4"
		var v interface{} = "1404ecd9-6db2-43f4-a6af-2650a31fac28"

		test := omap.NewMapStringAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("6d31ffa6-d04c-403d-895b-26829238751f", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "3a051412-e85d-42b0-9d02-ecc80cf181e8"
		So(test.PutIfNotNil("27053664-5333-40a9-953b-c06089cdeeb3", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringAny.ReplaceIfExists", t, func() {
		var k string = "f3509a71-5e78-4534-ae95-7a1a4459d458"
		var v interface{} = "fce26ca0-bc7e-4f99-a07b-1a53cd9891f8"
		var x interface{} = "b94293ae-cf04-4dbd-8a69-d5c0966ca009"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1db4a09a-7883-4006-afdf-912af9237468", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringAny.ReplaceOrPut", t, func() {
		var k string = "d6290baa-212a-4947-8044-a4b0a8d5da32"
		var v interface{} = "20ef2b0f-4694-4607-8f0a-fcffc053eb40"
		var x interface{} = "8fca0a3b-5f35-4edb-83e6-3e160fc9ab53"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("a44f4f8f-6a06-4ed8-a58f-5c1e04851b5d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_MarshalJSON(t *testing.T) {
	Convey("TestMapStringAny.MarshalJSON", t, func() {
		var k string = "425e1bb3-2ef0-4329-8775-78960682d877"
		var v interface{} = "6a615b62-90b9-4c67-8d8b-5268d668847e"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"425e1bb3-2ef0-4329-8775-78960682d877","value":"6a615b62-90b9-4c67-8d8b-5268d668847e"}]`)
	})
}
