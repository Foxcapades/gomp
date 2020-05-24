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
		var k interface{} = "73273ebd-ac08-4f5a-82c6-94b2a3e816ed"
		var v string = "8a3305b2-29ef-40c0-acf5-11f6a065400e"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyString_Delete(t *testing.T) {
	Convey("TestMapAnyString.Delete", t, func() {
		var k interface{} = "e66503cd-6a68-4495-9626-69ac8c4ec80f"
		var v string = "ff832212-065a-4ed7-811e-1571d89d714f"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyString_Has(t *testing.T) {
	Convey("TestMapAnyString.Has", t, func() {
		var k interface{} = "7b4eaf15-10ec-43ce-8a95-65a54211276f"
		var v string = "27d1c367-1c7f-4cb1-bc11-ef10b327edd6"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("2ef48a47-7738-4166-961a-f0a6dcdb006e"+"b55c9303-72d9-424e-b682-e677e95223b2"), ShouldBeFalse)
	})
}


func TestMapAnyString_Get(t *testing.T) {
	Convey("TestMapAnyString.Get", t, func() {
		var k interface{} = "a119e71b-f17f-4eae-a359-024474d96e64"
		var v string = "532d2650-a4b9-490e-89ea-b83cf0b2ba5d"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("168f0011-2420-4041-9d7a-ac33db9ce725" + "16d5f5e8-0f5f-4457-8f13-03ac49412b1a")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyString_GetOpt(t *testing.T) {
	Convey("TestMapAnyString.GetOpt", t, func() {
		var k interface{} = "da1f2f58-5ad5-40ee-85ef-91bb916a14c6"
		var v string = "e6ac9dde-0058-4829-8b05-4fe4d19068a2"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("a88bff6d-a4df-4311-abce-f2c3cda87bfe" + "d466dbe6-7fcd-4fb6-8155-a96eed76a585")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyString_ForEach(t *testing.T) {
	Convey("TestMapAnyString.ForEach", t, func() {
		var k interface{} = "d795d315-37a9-4468-86b0-24b7812d7f07"
		var v string = "e6a081da-5b78-41cf-b9ca-24ebcd5aadef"
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
		var k interface{} = "3433e23e-8960-4c78-92b1-d5a9de7bf8a9"
		var v string = "509387cf-8bd2-4a2e-b0ee-b01bd6aa2a3b"

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
		var k interface{} = "8064aaa0-807f-4038-8125-0eaa0ba79c59"
		var v string = "bc1dd660-c5c4-42b9-9208-21c5fc96b7c6"

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
		var k interface{} = "5169fdea-7bdc-444c-8b02-6e9a71930283"
		var v string = "c7829f43-5233-4c5a-8e2d-b76309cb0633"

		test := omap.NewMapAnyString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("70273f72-a035-4cba-808a-a16a3cd80e10", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "2e93c478-8221-41a5-b63a-8a94cf5bb268"
		So(test.PutIfNotNil("8d1b7119-66c0-4408-a524-e41b7dc144c7", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyString.ReplaceIfExists", t, func() {
		var k interface{} = "774a7a5d-4d2e-4fb3-b1ae-a7c22f5805ec"
		var v string = "7a271a55-6427-4ece-b472-cc9d47c2ec83"
		var x string = "0a1b59f9-24dc-49e2-ae4d-81172c788042"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1d820a76-ba1b-4096-afa0-56f44579ba14", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyString.ReplaceOrPut", t, func() {
		var k interface{} = "7e3aa71c-d1e9-4414-856b-0a344982b4d7"
		var v string = "f1db279c-ee34-4eaa-ab25-5e4650c869de"
		var x string = "b5530af7-18b7-4ca1-9d3e-d48d6a1a2382"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("837c7d0a-0acf-4999-9e66-b022750033af", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyString.MarshalJSON", t, func() {
		var k interface{} = "617ab911-417e-4b03-b7d3-916f70eea424"
		var v string = "74d549db-8126-481e-ba5a-2bca09ca118b"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"617ab911-417e-4b03-b7d3-916f70eea424","value":"74d549db-8126-481e-ba5a-2bca09ca118b"}]`)
	})
}
