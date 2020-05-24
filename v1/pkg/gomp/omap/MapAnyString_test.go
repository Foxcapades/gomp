package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyString_Put(t *testing.T) {
	Convey("TestMapAnyString.Put", t, func() {
		var k interface{} = "ad8d52d4-c4ae-461c-a7b0-c2fe9ed57192"
		var v string = "49b4e05d-1ef0-49ea-9904-cb95c46e9489"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyString_Delete(t *testing.T) {
	Convey("TestMapAnyString.Delete", t, func() {
		var k interface{} = "e357fd83-24dd-4e7f-9769-b7d8e3ea9ffd"
		var v string = "ad2fcdde-fa05-46b6-9df4-09420c762636"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyString_Has(t *testing.T) {
	Convey("TestMapAnyString.Has", t, func() {
		var k interface{} = "ba0eb37a-b613-4aec-9bf6-1af20f1f55ad"
		var v string = "113384c3-5f92-4ed7-a872-65175b082dfc"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("349e5ee6-bf47-464f-9466-bdcd0531b3af"+"a1a3ed7c-481f-4ecd-a3c1-d433fac22035"), ShouldBeFalse)
	})
}


func TestMapAnyString_Get(t *testing.T) {
	Convey("TestMapAnyString.Get", t, func() {
		var k interface{} = "26b01466-4ac5-49a5-948c-dc170d7bad65"
		var v string = "1763dfc4-8d26-4f16-b5ab-b6f38ac452fd"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("1a1ba204-f1f5-4c30-a167-8f2e5be595b2"+"d7c60802-871c-4250-ad26-2c5b81db12c4")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyString_GetOpt(t *testing.T) {
	Convey("TestMapAnyString.GetOpt", t, func() {
		var k interface{} = "b288259b-ce67-42bd-9ca5-36580f78b434"
		var v string = "d359ab1f-2409-4677-8706-23e499fa05bb"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("86e12516-22b4-4c4c-abb2-444590a9e511"+"52bc095b-c69d-4941-8598-892672421cfe")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyString_ForEach(t *testing.T) {
	Convey("TestMapAnyString.ForEach", t, func() {
		var k interface{} = "0c440020-ab32-42af-a990-f8d65764f6f0"
		var v string = "21cb0f0a-d825-42bd-9210-a61e0bee1878"
		hits := 0

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyString_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyString.MarshalYAML", t, func() {
		var k interface{} = "8d96970f-2c1f-4c47-8f59-972fd33f3964"
		var v string = "7423de95-f1ae-4f9a-8743-bc8c52db54cd"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyString_ToYAML(t *testing.T) {
	Convey("TestMapAnyString.ToYAML", t, func() {
		var k interface{} = "dd835901-5121-43ae-a1c5-a9b225fa4d9b"
		var v string = "f77af3d0-05f3-4389-8cf2-ea136ee77e7e"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyString_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyString.PutIfNotNil", t, func() {
		var k interface{} = "83fbea50-1eae-45db-91c2-2fd4dd158878"
		var v string = "a374b6d6-67c0-47ab-8288-2df50f46fdcf"

		test := omap.NewMapAnyString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("9e03d49d-d21e-4398-ae3f-6cd81d09a639", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "8b841add-2d24-4d02-894e-ad0ffda63834"
		So(test.PutIfNotNil("903cf9ee-4d32-4988-8167-5fb08f44d104", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyString.ReplaceIfExists", t, func() {
		var k interface{} = "eff0d552-3b69-4491-8198-15d70f0d103b"
		var v string = "3ae28533-0c36-493c-a2f5-40df44ae0101"
		var x string = "9b94a131-3673-4ee2-8eb9-35cf8f5ad1f3"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("20c18bce-3b3a-4795-838e-df52485a0621", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyString.ReplaceOrPut", t, func() {
		var k interface{} = "5402c164-c056-426a-bf35-2849276ce829"
		var v string = "1a03d8bc-528e-4d51-a7c1-d3545ee61df8"
		var x string = "8020ca9a-e0c5-4153-ba89-6ab0a0bd8d96"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("f76cd1d4-79d2-47b3-88ad-a34ed359fc80", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyString.MarshalJSON", t, func() {
		var k interface{} = "2376b598-c222-4eec-b797-b0b0829b3675"
		var v string = "7f12eaa6-c5bc-47b0-abf6-25d7509dee1b"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"2376b598-c222-4eec-b797-b0b0829b3675","value":"7f12eaa6-c5bc-47b0-abf6-25d7509dee1b"}]`)
	})
}

