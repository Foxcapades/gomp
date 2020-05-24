package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAny_Put(t *testing.T) {
	Convey("TestMapAny.Put", t, func() {
		var k interface{} = "a722dd02-ad32-4a76-b578-8caf9a62ab46"
		var v interface{} = "920b8666-a5cf-48dc-b1b9-b41e0f3bf9b9"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAny_Delete(t *testing.T) {
	Convey("TestMapAny.Delete", t, func() {
		var k interface{} = "69316d68-60a4-4332-bd6e-be949083581d"
		var v interface{} = "df0eba86-6477-4f25-8364-ae1225ce711b"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAny_Has(t *testing.T) {
	Convey("TestMapAny.Has", t, func() {
		var k interface{} = "857c90ac-fcc8-4fdb-b0a3-02a6eea05b64"
		var v interface{} = "8e02c29c-49e0-47e8-a6b9-da091b5901ef"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b410c169-0ec7-4a80-a582-b5974ea4dd92"+"578b9ab2-fd09-4459-a48c-e5081ce7e3b0"), ShouldBeFalse)
	})
}


func TestMapAny_Get(t *testing.T) {
	Convey("TestMapAny.Get", t, func() {
		var k interface{} = "a168cdfb-88d2-493c-8f1e-2c9a0b187d11"
		var v interface{} = "c1e682d7-8285-40d9-b988-e62e9a81db27"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("905824bb-456f-4657-899c-119fc974b7d8" + "2181dd64-8f2b-4364-9562-ebff2ac95d9f")
		So(b, ShouldBeFalse)
	})
}

func TestMapAny_GetOpt(t *testing.T) {
	Convey("TestMapAny.GetOpt", t, func() {
		var k interface{} = "0119bca9-9830-4a6f-9f9c-5ee147f85e15"
		var v interface{} = "18513789-61bc-4620-b279-3f09f43f7a13"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("7d15b313-f214-45e8-9af0-0761aadbc57e" + "e30f1ab4-df53-4fa6-a146-40c5b99feddf")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAny_ForEach(t *testing.T) {
	Convey("TestMapAny.ForEach", t, func() {
		var k interface{} = "b2a06042-8335-454a-884f-3f805e68e39f"
		var v interface{} = "31de345d-0a0a-44ba-923f-c1e236a0b64b"
		hits := 0

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAny_MarshalYAML(t *testing.T) {
	Convey("TestMapAny.MarshalYAML", t, func() {
		var k interface{} = "b20d20a9-b4fb-48aa-9d3b-e93da9d8427d"
		var v interface{} = "64411348-52f3-43ef-b5b0-f4eb813c6f05"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAny_ToYAML(t *testing.T) {
	Convey("TestMapAny.ToYAML", t, func() {
		var k interface{} = "44cc678e-ad32-4185-862c-bd0cae0391c3"
		var v interface{} = "81c90624-2342-460d-a32d-79408c8cbaf6"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapAny.PutIfNotNil", t, func() {
		var k interface{} = "86ad6075-be4a-404d-bf9c-63693e138011"
		var v interface{} = "e579e006-60c5-4926-bc95-7e2bef1957b0"

		test := omap.NewMapAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("33794cd6-a813-4236-9524-6ccaf6dfc5a0", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "4c96c401-5f6e-400a-9964-471b699b373d"
		So(test.PutIfNotNil("e180dcb6-7967-40f7-89b1-95ca5a71d18f", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAny.ReplaceIfExists", t, func() {
		var k interface{} = "e64bd335-9c6e-4a3b-bc91-982355ef783f"
		var v interface{} = "50716bf8-fbb0-4309-86db-e5431a2c76e1"
		var x interface{} = "db1dd677-2e10-4c6d-9338-781a3c38361c"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("837140ce-65e7-4dfe-be42-99e1153ea8e1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAny.ReplaceOrPut", t, func() {
		var k interface{} = "bf32849d-8527-41a1-b6a6-565fc662d623"
		var v interface{} = "2953010d-96e7-466f-9709-f41b67816f0f"
		var x interface{} = "5600adc4-ee07-4dfb-a92f-e7161dee9e35"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("8e1314d0-c9c9-457b-9e43-1d552a3b77fc", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_MarshalJSON(t *testing.T) {
	Convey("TestMapAny.MarshalJSON", t, func() {
		var k interface{} = "20aa9a2a-a376-47d1-ae23-506351363e46"
		var v interface{} = "6545d14f-3582-4700-8088-a9d9867de9fe"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"20aa9a2a-a376-47d1-ae23-506351363e46","value":"6545d14f-3582-4700-8088-a9d9867de9fe"}]`)
	})
}
