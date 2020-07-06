package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint_Put(t *testing.T) {
	Convey("TestMapAnyUint.Put", t, func() {
		var k interface{} = "813f61ed-a96a-433e-9cc9-a1c7306a513a"
		var v uint = 1873063792

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint_Delete(t *testing.T) {
	Convey("TestMapAnyUint.Delete", t, func() {
		var k interface{} = "3a738de6-f02e-4a3a-9c5c-fa56986d520c"
		var v uint = 1502738575

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint_Has(t *testing.T) {
	Convey("TestMapAnyUint.Has", t, func() {
		var k interface{} = "f283456b-b9c1-41b6-bc07-97ab30918d24"
		var v uint = 3465115378

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("68a5ead7-813b-4711-a1b8-47208dd751f8"+"fff3e64a-540c-4a32-ad19-cc06df7fadc9"), ShouldBeFalse)
	})
}

func TestMapAnyUint_Get(t *testing.T) {
	Convey("TestMapAnyUint.Get", t, func() {
		var k interface{} = "ca87bf98-78ce-4df1-baa4-f0e13a03c956"
		var v uint = 658778188

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("7ac3decb-85ef-48a6-80be-2998d52c0016" + "ecea1b5a-7ed2-4040-b4f2-1e9fb41e4960")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint.GetOpt", t, func() {
		var k interface{} = "a64d53c6-0bde-480d-aed5-7d5b6f8c1f2f"
		var v uint = 708406195

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("dffe130b-a33b-4e3b-bf2c-70dfdad745da" + "9c929126-90e9-477b-8705-4126554e3c19")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint_ForEach(t *testing.T) {
	Convey("TestMapAnyUint.ForEach", t, func() {
		var k interface{} = "ee8bf4bc-8550-4620-9977-01117de3749a"
		var v uint = 783699172
		hits := 0

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint.MarshalYAML", t, func() {
		var k interface{} = "8abaf9e6-a23b-47e3-92a9-80982f12862d"
		var v uint = 2509941998

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "558fe679-b213-494b-b783-366eb5844c11"
			var v uint = 2129677608

			test := omap.NewMapAnyUint(1)

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
			var k interface{} = "347176a7-ce2e-47a6-9b4d-b604854dd573"
			var v uint = 222902141

			test := omap.NewMapAnyUint(1)
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

func TestMapAnyUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint.PutIfNotNil", t, func() {
		var k interface{} = "efa68f2c-0bb2-4ec8-be41-2a84f07387d3"
		var v uint = 162067076

		test := omap.NewMapAnyUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b7a83c48-cf42-4a5e-aa62-4462fd556c6f", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 3362494769
		So(test.PutIfNotNil("a3f4cdb1-fd85-4648-84ba-6b2882b98202", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceIfExists", t, func() {
		var k interface{} = "587c3455-3533-41c8-a505-9b20f7ad5165"
		var v uint = 210307161
		var x uint = 3947298764

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("35febf5c-3593-45ff-8a44-4c7fea4e5323", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceOrPut", t, func() {
		var k interface{} = "4c374e65-a32d-421e-90de-1e70b8c8b77a"
		var v uint = 2799005704
		var x uint = 4132136312

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("a6c5f2d0-2dc7-4159-8cbd-c70399ffe0dc", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "069109a2-346a-42d2-ac78-560fc4c84f34"
			var v uint = 2125405271

			test := omap.NewMapAnyUint(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"069109a2-346a-42d2-ac78-560fc4c84f34","value":2125405271}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "069109a2-346a-42d2-ac78-560fc4c84f34"
			var v uint = 2125405271

			test := omap.NewMapAnyUint(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"069109a2-346a-42d2-ac78-560fc4c84f34":2125405271}`)
		})

	})
}
