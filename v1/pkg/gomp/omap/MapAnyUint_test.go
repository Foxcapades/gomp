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
		var k interface{} = "f43e9cef-3294-41d0-8a66-667a1f0ac4e7"
		var v uint = 1362444162

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint_Delete(t *testing.T) {
	Convey("TestMapAnyUint.Delete", t, func() {
		var k interface{} = "eda37463-04ee-4f3a-8de8-47fad62ce1b6"
		var v uint = 1012283765

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint_Has(t *testing.T) {
	Convey("TestMapAnyUint.Has", t, func() {
		var k interface{} = "66b2cfb9-fc00-4015-876c-0a080bf01793"
		var v uint = 1363027111

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("23c75225-57d7-4069-b929-275d0afa2641"+"5d5082e9-6ec9-47e2-88f7-04c752f25d03"), ShouldBeFalse)
	})
}

func TestMapAnyUint_Get(t *testing.T) {
	Convey("TestMapAnyUint.Get", t, func() {
		var k interface{} = "fd382584-874c-485f-a448-6ff4f96c0cea"
		var v uint = 866315631

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("68436f85-cb72-4d7a-9a42-71810885d9eb" + "7b53a2f4-b7a8-482f-97ec-567b26765a91")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint.GetOpt", t, func() {
		var k interface{} = "472f1f22-e3ed-4cee-bdc6-4c6345753253"
		var v uint = 633564189

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5799efaf-11ff-4ff1-8f11-3fbc48434230" + "a7146046-7c32-414e-91b5-829f3ae29c07")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint_ForEach(t *testing.T) {
	Convey("TestMapAnyUint.ForEach", t, func() {
		var k interface{} = "78283acb-2aa0-4d76-8397-f7dd08c4090b"
		var v uint = 3043365541
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
		var k interface{} = "4a553de4-8305-4041-8c06-4df6491d0434"
		var v uint = 3767432596

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
		var k interface{} = "93feb89c-67e9-4971-aaa5-d379acab255a"
		var v uint = 735208136

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapAnyUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint.PutIfNotNil", t, func() {
		var k interface{} = "2953c4e4-15f6-4339-9300-4b3855f9be84"
		var v uint = 3850416207

		test := omap.NewMapAnyUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("272f26ac-78bf-4fb0-bad4-9d9cf94ab57b", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 751971416
		So(test.PutIfNotNil("394cb349-f203-4bfa-b7d9-cb8c7f6f8f4e", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceIfExists", t, func() {
		var k interface{} = "42742195-17c2-418b-bbb2-024a8fbd4233"
		var v uint = 2712034569
		var x uint = 4062822941

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("5ea31b6c-44d9-4d5c-a91a-1a12cdd5b80f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceOrPut", t, func() {
		var k interface{} = "bea94c91-4dce-4fcc-881b-88f7a9f3286b"
		var v uint = 4261978238
		var x uint = 2055127513

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("b5145067-d45b-4b3a-9545-07b2a58ed528", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint.MarshalJSON", t, func() {
		var k interface{} = "dcf73668-5e92-40f5-bc10-df2dcdc43453"
		var v uint = 2643680703

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"dcf73668-5e92-40f5-bc10-df2dcdc43453","value":2643680703}]`)
	})
}
