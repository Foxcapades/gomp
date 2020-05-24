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
		var k interface{} = "eeadc58f-3f57-438c-b7af-f34065946ace"
		var v uint64 = 3652946243721370785

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint64_Delete(t *testing.T) {
	Convey("TestMapAnyUint64.Delete", t, func() {
		var k interface{} = "584817ae-760f-4394-b5ad-a17822e67acf"
		var v uint64 = 5532630613333149746

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint64_Has(t *testing.T) {
	Convey("TestMapAnyUint64.Has", t, func() {
		var k interface{} = "f7bfa6b9-b771-453e-a09a-a9461328583b"
		var v uint64 = 12905877509749377022

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("ed271287-76da-43c3-8364-f3fafe6d2161"+"b5143e9f-e98e-44ed-b5fa-39041b96068a"), ShouldBeFalse)
	})
}


func TestMapAnyUint64_Get(t *testing.T) {
	Convey("TestMapAnyUint64.Get", t, func() {
		var k interface{} = "a6dbd2de-4db6-4d78-ad69-fc7e8288eb20"
		var v uint64 = 5137237447044453363

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("088dec78-692a-4e1a-999c-07cbac1e6712" + "e793a97c-b02e-4679-ba95-9dcbdd7ae8e6")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint64_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint64.GetOpt", t, func() {
		var k interface{} = "17149bcd-0cf8-4cfc-aa78-847215e82de5"
		var v uint64 = 10947745498287601617

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("96e79457-8bfb-436e-9d9d-f3350481e8ff" + "281647a9-b606-4a34-9f79-a50ad810e83d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint64_ForEach(t *testing.T) {
	Convey("TestMapAnyUint64.ForEach", t, func() {
		var k interface{} = "d1db239d-a864-40b6-a7f4-7098ca17776e"
		var v uint64 = 12650839492427545026
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
		var k interface{} = "83fd84e9-5b7d-4c94-9d02-85f8d0db83ea"
		var v uint64 = 7619329899311508249

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
		var k interface{} = "e3324fe3-cad2-459b-8dab-7b85e0233a8d"
		var v uint64 = 6609212514127681035

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
		var k interface{} = "d92932c5-48c7-4567-873c-41799d415366"
		var v uint64 = 10235480788556987121

		test := omap.NewMapAnyUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b0b8dfb0-1882-4129-931e-8db6b0044c64", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 7312371458810952410
		So(test.PutIfNotNil("1c21f0ef-8042-47f8-9c5b-46a419944e29", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceIfExists", t, func() {
		var k interface{} = "a048c71a-b360-4c11-add1-e7385b8cdca4"
		var v uint64 = 6189457323250170387
		var x uint64 = 12749018766818845301

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1fa69593-fe7f-437a-bbbf-a3970cba2630", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceOrPut", t, func() {
		var k interface{} = "e9d33cc5-ab61-4b30-9230-7f7a64fb06d1"
		var v uint64 = 12911047404492501700
		var x uint64 = 11374479186443514411

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("864c2669-caad-4c6b-9961-72001a24102f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint64.MarshalJSON", t, func() {
		var k interface{} = "a804f41c-f1d9-4d5b-a4bb-1bb2caca085e"
		var v uint64 = 9732078110730841165

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a804f41c-f1d9-4d5b-a4bb-1bb2caca085e","value":9732078110730841165}]`)
	})
}

