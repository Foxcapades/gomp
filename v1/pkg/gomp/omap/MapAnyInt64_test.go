package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt64_Put(t *testing.T) {
	Convey("TestMapAnyInt64.Put", t, func() {
		var k interface{} = "37d50105-f35a-49aa-a93f-d252dfc0ac23"
		var v int64 = 3889606077793664355

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt64_Delete(t *testing.T) {
	Convey("TestMapAnyInt64.Delete", t, func() {
		var k interface{} = "bb5e3035-1a58-44d9-855b-ac603f6a4952"
		var v int64 = 5501848630522582360

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt64_Has(t *testing.T) {
	Convey("TestMapAnyInt64.Has", t, func() {
		var k interface{} = "9f036f9f-0c47-43b6-a6ec-60494378e3be"
		var v int64 = 6291548689033019623

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("c0cc4e2f-e87e-4372-8b9d-bcf6a0242fdb"+"c2a8f93a-73b6-405b-a5eb-25e78f1aa99b"), ShouldBeFalse)
	})
}

func TestMapAnyInt64_Get(t *testing.T) {
	Convey("TestMapAnyInt64.Get", t, func() {
		var k interface{} = "5c548e44-06c1-48ed-b593-b576979ef375"
		var v int64 = 1672064798349613989

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("e82fb9fc-ac80-4108-b84d-548e78ee1321" + "144244f5-e88f-409e-b39a-48fbadbc05d5")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt64_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt64.GetOpt", t, func() {
		var k interface{} = "86c74e3f-d6df-4322-bea9-2860abd8cba1"
		var v int64 = 3421943623687498835

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("617e3fed-af1c-4a47-9477-b9e09b463341" + "aa05658e-8877-4475-953a-9b07504e124c")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt64_ForEach(t *testing.T) {
	Convey("TestMapAnyInt64.ForEach", t, func() {
		var k interface{} = "84a99397-9516-4984-9c73-92b48820ad7a"
		var v int64 = 8337851527556090167
		hits := 0

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt64.MarshalYAML", t, func() {
		var k interface{} = "77eedf9a-eeff-4f0f-bafb-1e6788c47aea"
		var v int64 = 2015882200778397599

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt64_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt64.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "05993b78-bc09-45bd-8063-38a7f9ca21be"
			var v int64 = 3858981334350643002

			test := omap.NewMapAnyInt64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k interface{} = "1e884a4b-f811-46fc-a630-baf712b8d0ab"
			var v int64 = 5690832199479679906

			test := omap.NewMapAnyInt64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapAnyInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt64.PutIfNotNil", t, func() {
		var k interface{} = "9197cb2c-d84b-4449-8be1-18b938a09996"
		var v int64 = 7845392995563650788

		test := omap.NewMapAnyInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("c5550c60-a9c3-4f14-a918-c71cf3762a1f", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 5705048072653940747
		So(test.PutIfNotNil("8a2a1bbc-187a-4adf-b9b5-e122f13baea5", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceIfExists", t, func() {
		var k interface{} = "5481402a-a63a-47a6-b8e9-0c4d8b0737b8"
		var v int64 = 6116431891140904738
		var x int64 = 292557180834894459

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("6739b39d-ec9f-406f-838b-ee40af39c583", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceOrPut", t, func() {
		var k interface{} = "4c45466b-7e3c-41d3-a27d-4fe534df8bdc"
		var v int64 = 6505212301422522525
		var x int64 = 7052400047550975925

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("cf0b918e-d2f3-4b95-bd48-9be2f518f713", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt64.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "7479c4bd-7e38-4687-9144-c2e46372b00e"
			var v int64 = 6064592121511653036

			test := omap.NewMapAnyInt64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"7479c4bd-7e38-4687-9144-c2e46372b00e","value":6064592121511653036}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "7479c4bd-7e38-4687-9144-c2e46372b00e"
			var v int64 = 6064592121511653036

			test := omap.NewMapAnyInt64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"7479c4bd-7e38-4687-9144-c2e46372b00e":6064592121511653036}`)
		})

	})
}
