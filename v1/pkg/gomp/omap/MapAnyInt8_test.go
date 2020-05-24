package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt8_Put(t *testing.T) {
	Convey("TestMapAnyInt8.Put", t, func() {
		var k interface{} = "1f2debbd-dbb3-472f-9d47-635759e89a15"
		var v int8 = 98

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt8_Delete(t *testing.T) {
	Convey("TestMapAnyInt8.Delete", t, func() {
		var k interface{} = "9678779e-dbae-4e90-8f36-aa2c51473255"
		var v int8 = 27

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt8_Has(t *testing.T) {
	Convey("TestMapAnyInt8.Has", t, func() {
		var k interface{} = "3b3fb26e-7881-4f33-b020-8c9179bab333"
		var v int8 = 11

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("fba5d963-35e4-4c03-8472-3d8f84833f75"+"3e901144-9fe4-4040-bb5d-cd41c46fccbe"), ShouldBeFalse)
	})
}

func TestMapAnyInt8_Get(t *testing.T) {
	Convey("TestMapAnyInt8.Get", t, func() {
		var k interface{} = "a79ef813-cc2f-47d1-ab4b-ea32b136d39b"
		var v int8 = 56

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("aac93a8b-39ed-4916-993d-0ca2b12767c0" + "aa09ec80-9ca6-4175-8891-29db78380799")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt8_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt8.GetOpt", t, func() {
		var k interface{} = "506a702f-e629-40d7-bd78-9389a8598878"
		var v int8 = 106

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("1c884930-5377-4308-ae93-38e33088a8c9" + "fdc4e023-c9c5-4bee-9e32-e1c5b0c5b1cc")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt8_ForEach(t *testing.T) {
	Convey("TestMapAnyInt8.ForEach", t, func() {
		var k interface{} = "4b418105-c603-4b43-8b55-d10233d4d617"
		var v int8 = 23
		hits := 0

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt8.MarshalYAML", t, func() {
		var k interface{} = "2748c97a-d137-4a04-b304-1b9cc430ce3f"
		var v int8 = 113

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt8_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt8.ToYAML", t, func() {
		var k interface{} = "e4b4ed78-08f0-4b40-a93c-d9c92b74fe92"
		var v int8 = 76

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt8.PutIfNotNil", t, func() {
		var k interface{} = "578fa9f8-8a10-4a2d-8671-38b9258ffbf3"
		var v int8 = 49

		test := omap.NewMapAnyInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("80f69252-ea7f-439b-adb6-440b9f1742c0", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 55
		So(test.PutIfNotNil("b4812ffa-551c-409e-8198-3912bc05549c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceIfExists", t, func() {
		var k interface{} = "3991645f-41ba-41c9-af95-2734248797f0"
		var v int8 = 103
		var x int8 = 23

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("fee10fa5-7586-4827-94b3-6b76b401b422", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceOrPut", t, func() {
		var k interface{} = "7dc22db1-f777-42b0-9cbb-f14ca16a2b1a"
		var v int8 = 73
		var x int8 = 54

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("9fb05449-2598-46cb-8cc3-e8adf0cda0e4", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt8.MarshalJSON", t, func() {
		var k interface{} = "7335942a-60c9-457c-8f95-b95e5c0a597a"
		var v int8 = 82

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"7335942a-60c9-457c-8f95-b95e5c0a597a","value":82}]`)
	})
}
