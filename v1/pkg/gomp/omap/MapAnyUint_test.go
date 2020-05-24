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
		var k interface{} = "abd184f0-2a1c-42bf-a5c5-11fcfa8b0a69"
		var v uint = 1845346338

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint_Delete(t *testing.T) {
	Convey("TestMapAnyUint.Delete", t, func() {
		var k interface{} = "796757f1-6f81-4532-b4cb-3c4d86fb2fa4"
		var v uint = 2329843981

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint_Has(t *testing.T) {
	Convey("TestMapAnyUint.Has", t, func() {
		var k interface{} = "dfcb41cd-6c3c-4892-a307-65b153344afe"
		var v uint = 2430929289

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("c93f6051-f1cb-41b5-a473-6c42839a1fad"+"e97c6bbb-54e1-4da5-877a-2ed3b15ece05"), ShouldBeFalse)
	})
}

func TestMapAnyUint_Get(t *testing.T) {
	Convey("TestMapAnyUint.Get", t, func() {
		var k interface{} = "71331733-8bca-4fec-b24b-7e7effa9166e"
		var v uint = 2519194814

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("5fb31c06-c73d-4e21-bda2-62d1526b2c45" + "efd0c455-2da2-4d06-bbd0-d74eb1ff80bf")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint.GetOpt", t, func() {
		var k interface{} = "0426e101-7def-4132-8c29-dcc0b770ca3f"
		var v uint = 3959513815

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("2131442e-651e-4a92-a357-c0eb515942da" + "650e104c-4463-4564-9338-55b0a6f3fda5")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint_ForEach(t *testing.T) {
	Convey("TestMapAnyUint.ForEach", t, func() {
		var k interface{} = "7ff9b8e1-a7c5-4986-b16c-12283acd1452"
		var v uint = 515997580
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
		var k interface{} = "435a1dba-2cc9-457a-9696-56294ed40c86"
		var v uint = 3541145664

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
		var k interface{} = "cd8d6f98-7ddd-4013-b9b8-22aeedbdd73c"
		var v uint = 1387567404

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint.PutIfNotNil", t, func() {
		var k interface{} = "1dc53674-963d-4b3b-be6e-ca781475dadf"
		var v uint = 732310034

		test := omap.NewMapAnyUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("0bf1458d-1f3b-47db-9459-0d6acfa9713e", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 1226460612
		So(test.PutIfNotNil("17a0c912-39f4-4ad2-909a-0f7f304e0fea", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceIfExists", t, func() {
		var k interface{} = "cb86fa32-7377-4049-800e-bbb5c6e3c2d7"
		var v uint = 3115391708
		var x uint = 3888155357

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("4a0b2159-3ed4-40b1-b71f-ada4066602b9", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceOrPut", t, func() {
		var k interface{} = "25614599-447a-476b-b1bd-b67ac2dd3a62"
		var v uint = 124089063
		var x uint = 4092266567

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("b427f77d-9c7a-442e-81b5-b82e1629f92a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint.MarshalJSON", t, func() {
		var k interface{} = "fcfa9fa4-7373-40cf-8188-d5fe772ab42d"
		var v uint = 2914496026

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"fcfa9fa4-7373-40cf-8188-d5fe772ab42d","value":2914496026}]`)
	})
}
