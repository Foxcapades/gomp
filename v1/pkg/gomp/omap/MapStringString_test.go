package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringString_Put(t *testing.T) {
	Convey("TestMapStringString.Put", t, func() {
		var k string = "2f989a64-0da8-4fc0-9972-3eb8e2259fb3"
		var v string = "93ee9661-e6fb-4f97-affa-a754b4ce3241"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringString_Delete(t *testing.T) {
	Convey("TestMapStringString.Delete", t, func() {
		var k string = "9a9776ca-15aa-4ee4-9a00-298fc87e7710"
		var v string = "1a8d75b5-2e4e-4307-9369-e12cfc1c8bdf"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringString_Has(t *testing.T) {
	Convey("TestMapStringString.Has", t, func() {
		var k string = "60a23b2d-e2fb-4a7d-bfc5-eab0beff8172"
		var v string = "9782698e-b1b2-4d85-9747-4c8e9055e16e"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("bbe34857-b767-434c-9cde-a499bfcbad09"+"35dbe013-f222-4a8b-b291-3a3bd76aa73a"), ShouldBeFalse)
	})
}

func TestMapStringString_Get(t *testing.T) {
	Convey("TestMapStringString.Get", t, func() {
		var k string = "9888c855-e449-401d-986b-a90bedf67ceb"
		var v string = "e49cccc8-9362-4622-9f5d-33b985415525"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("23276dc4-8e8f-47c9-951c-719ec14ab6e3" + "b739065c-ff7f-4587-88a3-26da72daf8bb")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringString_GetOpt(t *testing.T) {
	Convey("TestMapStringString.GetOpt", t, func() {
		var k string = "525b11d5-ca07-44be-9609-d18bbbdc1f8d"
		var v string = "a5ab9b39-8aa1-49ca-ae46-bf1f2c3d15f4"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("e9718c78-968b-428b-af9d-ebb24173ebae" + "cf8abbb7-a7d8-4a73-8702-eed47735ee00")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringString_ForEach(t *testing.T) {
	Convey("TestMapStringString.ForEach", t, func() {
		var k string = "ff29fe60-5264-4825-81c1-49a0496a2c26"
		var v string = "00dd5cab-57be-4e44-9ea4-327f92013a25"
		hits := 0

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringString_MarshalYAML(t *testing.T) {
	Convey("TestMapStringString.MarshalYAML", t, func() {
		var k string = "4e6dd4b7-fdbe-488c-a844-ad7f60181084"
		var v string = "d419b340-0805-4f62-869f-f19967cd8131"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringString_ToYAML(t *testing.T) {
	Convey("TestMapStringString.ToYAML", t, func() {
		var k string = "55c5f3c7-4a64-4434-8254-ac852c8711ca"
		var v string = "104c5f67-ed53-428e-8c5d-983188cb656d"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringString_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringString.PutIfNotNil", t, func() {
		var k string = "fdd515cb-e8a6-40fe-9d36-6a31e5bdfd40"
		var v string = "4aeed5f7-18cb-477e-b6af-47732bc5556e"

		test := omap.NewMapStringString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("4de27734-ded3-40a4-aa1b-1c653aaf9885", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "3f924d69-8250-4647-bebe-b3495bc005e5"
		So(test.PutIfNotNil("b1b4cc95-166e-46b1-a365-ae72ba4cbf9e", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringString.ReplaceIfExists", t, func() {
		var k string = "9ad21258-be72-4aa9-ba1d-49a26ca4ef74"
		var v string = "16ea1875-318d-4c90-8e65-2aa93b615d75"
		var x string = "50b81379-ce3c-4685-8cf5-60652802780a"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e1bd5429-211a-4d51-b807-ddee785b1daf", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringString.ReplaceOrPut", t, func() {
		var k string = "225e2d4e-aca9-4f4c-9e34-2510629a9f79"
		var v string = "66c4f26b-9b6b-43eb-91d8-9bb65e78c43b"
		var x string = "143b0021-6a61-4d8e-9cef-1a7d64a302f4"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("53713f70-6df7-4a45-ae5d-f39cd12634c6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_MarshalJSON(t *testing.T) {
	Convey("TestMapStringString.MarshalJSON", t, func() {
		var k string = "db6d585d-2dca-4daf-b6d5-2b60aa21d60c"
		var v string = "66b74394-52c7-4a81-a290-63c35bf10a55"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"db6d585d-2dca-4daf-b6d5-2b60aa21d60c","value":"66b74394-52c7-4a81-a290-63c35bf10a55"}]`)
	})
}
