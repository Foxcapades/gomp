package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint64_Put(t *testing.T) {
	Convey("TestMapAnyUint64.Put", t, func() {
		var k interface{} = "338eb23c-c73b-4104-8aa7-c9098f1dc7b3"
		var v uint64 = 10500886801579450830

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint64_Delete(t *testing.T) {
	Convey("TestMapAnyUint64.Delete", t, func() {
		var k interface{} = "be825c79-7878-4de8-8064-bdf42f5d9074"
		var v uint64 = 8970283770307293298

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint64_Has(t *testing.T) {
	Convey("TestMapAnyUint64.Has", t, func() {
		var k interface{} = "3b29fd99-7451-478a-a570-f6c362fafea1"
		var v uint64 = 1315273897383495858

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("81f27302-9106-4d3c-a4e1-314b0a232e7c"+"6a5f236e-2c56-4c31-a47c-0c88ebb05061"), ShouldBeFalse)
	})
}


func TestMapAnyUint64_Get(t *testing.T) {
	Convey("TestMapAnyUint64.Get", t, func() {
		var k interface{} = "c199cd6f-1c73-405f-8e27-bc2d83b6eadd"
		var v uint64 = 13658743227360610018

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("331a3231-acfc-40e4-a3ff-030eba82bac5"+"ed829f1e-8e51-4b00-bccb-d6920092638a")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint64_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint64.GetOpt", t, func() {
		var k interface{} = "95325aa2-728f-4cbb-bf9b-14fe2aa8194a"
		var v uint64 = 11820628251275304084

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("27c3a3c0-236d-40de-b334-e0327630f914"+"16a02d23-dd97-4c66-9a43-29c0944fa630")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint64_ForEach(t *testing.T) {
	Convey("TestMapAnyUint64.ForEach", t, func() {
		var k interface{} = "45f95cc3-f496-4421-931d-469b865d6c60"
		var v uint64 = 13930910558366050251
		hits := 0

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint64.MarshalYAML", t, func() {
		var k interface{} = "c7c1abd6-155a-462e-b4c3-355301725a26"
		var v uint64 = 17952099498886657325

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint64_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint64.ToYAML", t, func() {
		var k interface{} = "2a4d666e-d4f1-406f-8a71-3aecf43a9f83"
		var v uint64 = 6480889680256733463

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint64.PutIfNotNil", t, func() {
		var k interface{} = "2bb41f3c-cf67-4d7e-9890-9ec685a43d71"
		var v uint64 = 12839087567821022593

		test := omap.NewMapAnyUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b9d52167-5acc-4f82-969b-0b6ba112f09d", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 16952403611396069064
		So(test.PutIfNotNil("fb258300-dd4c-47c8-acc1-3344c569ce54", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceIfExists", t, func() {
		var k interface{} = "efeade69-2bf0-417b-85a5-b58e58e201eb"
		var v uint64 = 10326336439728712730
		var x uint64 = 16034565454778641629

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("94fc9934-2bd6-4dea-a403-f4110df67d77", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceOrPut", t, func() {
		var k interface{} = "30e204e3-f81e-4227-860f-7ac079a5ca73"
		var v uint64 = 15910446798445057246
		var x uint64 = 15820351207795962262

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("f9f8b99a-8126-40d8-b2d2-9e29809f965d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint64.MarshalJSON", t, func() {
		var k interface{} = "758ee6c7-0a68-42bf-bf44-7cb79c7d27af"
		var v uint64 = 15710855728756961094

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"758ee6c7-0a68-42bf-bf44-7cb79c7d27af","value":15710855728756961094}]`)
	})
}

