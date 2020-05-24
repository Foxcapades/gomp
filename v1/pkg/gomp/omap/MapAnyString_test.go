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
		var k interface{} = "6aa1284d-e869-4261-b51e-6fc814349796"
		var v string = "f6a10c8f-cc89-49c1-8c79-0a77df986156"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyString_Delete(t *testing.T) {
	Convey("TestMapAnyString.Delete", t, func() {
		var k interface{} = "39f8a145-4261-49e0-aadd-275e968bd813"
		var v string = "a78ea64f-5a25-4756-92ab-623fcd1aa7d8"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyString_Has(t *testing.T) {
	Convey("TestMapAnyString.Has", t, func() {
		var k interface{} = "054104b3-5f20-430e-bb93-fca58eceac89"
		var v string = "c84be320-38e8-4d5b-beac-ea18be0a6bcd"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("2e2d3296-6d61-4f6c-b768-2ba038efdec0"+"abde3044-bb21-4994-872c-8f73c2a29b6e"), ShouldBeFalse)
	})
}

func TestMapAnyString_Get(t *testing.T) {
	Convey("TestMapAnyString.Get", t, func() {
		var k interface{} = "c63cdd88-35cd-4f85-b2b4-34879cde09f9"
		var v string = "cbfd4e2d-f259-4c56-997d-8cf8fddd9002"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("e5d147f3-9d01-4b1f-afb2-c0965bdea919" + "1729a170-a52c-4b9d-9d02-24e8e0040f6f")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyString_GetOpt(t *testing.T) {
	Convey("TestMapAnyString.GetOpt", t, func() {
		var k interface{} = "a5cef51c-ac7c-4965-8f32-1ed0b5f6762a"
		var v string = "4048cb47-1697-4c77-a950-603418d564c6"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("73ccbbec-23ce-4393-acfc-c42adf98f718" + "96222986-b0a4-4ba2-b292-0dd4f9ca4765")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyString_ForEach(t *testing.T) {
	Convey("TestMapAnyString.ForEach", t, func() {
		var k interface{} = "43eb776f-6d6e-4452-9f87-1f9656dfc3d8"
		var v string = "e93b611c-6574-46f0-855c-2bddcdb3d1ed"
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
		var k interface{} = "3748aa2c-87b3-4262-bd89-dccc8511ed78"
		var v string = "4ffa9304-3ac2-4bb1-bb4b-e00d936551c2"

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
		var k interface{} = "b81b0627-d985-4955-a89e-1d34423f242a"
		var v string = "3ce607d8-1159-4be0-bdf0-52523bf2b833"

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
		var k interface{} = "246778eb-ddf6-484f-b7bb-36e99e1922e7"
		var v string = "75aab102-8428-4737-b304-d23c3818de42"

		test := omap.NewMapAnyString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("042b578e-a138-4d9d-b68c-489d936e02c8", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "d313dfa1-08ce-486c-a54a-e55352eb1a87"
		So(test.PutIfNotNil("f596fe7f-38df-4882-be13-49a4cd2a8fff", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyString.ReplaceIfExists", t, func() {
		var k interface{} = "4de19a10-c5a7-4a7c-bafd-f93af9416934"
		var v string = "ca137472-8ec2-4955-bcf6-f8ecf5c8564e"
		var x string = "fcae56fb-b041-43cf-a206-cb990addff97"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("07fd6f2a-2306-4fe8-b405-2f6a0eebd04b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyString.ReplaceOrPut", t, func() {
		var k interface{} = "bfe4ea07-23e7-4e77-b43c-8e70d0154cee"
		var v string = "d57a034b-0855-4e27-8731-528201025154"
		var x string = "4c2231a2-f403-49ed-bb82-240d1689578e"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("1ddf211b-1c06-4c56-a96d-5c8e495066f2", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyString.MarshalJSON", t, func() {
		var k interface{} = "4ce1a902-4a36-4291-a1aa-d55d9b10a6e2"
		var v string = "4b142f85-1171-4f78-b77d-699313425d16"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"4ce1a902-4a36-4291-a1aa-d55d9b10a6e2","value":"4b142f85-1171-4f78-b77d-699313425d16"}]`)
	})
}
