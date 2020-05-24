package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt32_Put(t *testing.T) {
	Convey("TestMapAnyInt32.Put", t, func() {
		var k interface{} = "8e73ba61-e64a-4c50-ac61-7e772e5ab319"
		var v int32 = 1310098614

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt32_Delete(t *testing.T) {
	Convey("TestMapAnyInt32.Delete", t, func() {
		var k interface{} = "66d82e5f-3e70-4312-baef-af33615e1c7b"
		var v int32 = 2136873074

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt32_Has(t *testing.T) {
	Convey("TestMapAnyInt32.Has", t, func() {
		var k interface{} = "2c1937e6-0c6e-4c30-a335-ade9f0ce2125"
		var v int32 = 1474427690

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("f20fc19f-6f75-40f2-953b-fe94134c3a06"+"0f04a839-c4ec-4d1d-85c6-875e6ff72586"), ShouldBeFalse)
	})
}

func TestMapAnyInt32_Get(t *testing.T) {
	Convey("TestMapAnyInt32.Get", t, func() {
		var k interface{} = "c86c8b28-7968-47fe-a606-6c603d775561"
		var v int32 = 2063107473

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("5c4fd897-cb47-4f53-91bb-cdb572259060" + "ad5714c8-a250-4898-bfe6-6d12fa406956")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt32_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt32.GetOpt", t, func() {
		var k interface{} = "f484d05a-f5ff-4f27-861c-bfd8e70614da"
		var v int32 = 2101967381

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("e4648c03-9e03-4075-aff0-540ece8c880d" + "acc36652-f8ae-40b5-ae1e-4ff9f2edebea")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt32_ForEach(t *testing.T) {
	Convey("TestMapAnyInt32.ForEach", t, func() {
		var k interface{} = "364a9da3-cfdb-4341-9f9d-80c67cff74d9"
		var v int32 = 2134799828
		hits := 0

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalYAML", t, func() {
		var k interface{} = "5e1b2ed8-ebaf-4ca7-be0a-d84ac5930203"
		var v int32 = 1432571935

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt32_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt32.ToYAML", t, func() {
		var k interface{} = "c3f4417b-30af-472a-8d5e-bf6367e9612f"
		var v int32 = 1936259592

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt32.PutIfNotNil", t, func() {
		var k interface{} = "60ee56dd-0135-49df-8930-f9ce010e6ca4"
		var v int32 = 694514449

		test := omap.NewMapAnyInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("47e42b76-c086-4140-9c52-dc6f3ee64cfd", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 631087398
		So(test.PutIfNotNil("9a91aea7-7e2c-4193-8bbb-0dcd6d9f58cc", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceIfExists", t, func() {
		var k interface{} = "98f3071d-4ecb-4b89-bfeb-9ebad3e4ed68"
		var v int32 = 40748936
		var x int32 = 1954998383

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("4ff625a6-bd79-41d7-bf2a-807ea1d42450", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceOrPut", t, func() {
		var k interface{} = "9484028d-eaf9-481c-8873-2461ac431077"
		var v int32 = 3573177
		var x int32 = 171190164

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("9c5a3d7d-242e-4d9c-9222-b0fbaa87345b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalJSON", t, func() {
		var k interface{} = "aed9a771-000b-49db-9fd7-599f877b394b"
		var v int32 = 1688841067

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"aed9a771-000b-49db-9fd7-599f877b394b","value":1688841067}]`)
	})
}
