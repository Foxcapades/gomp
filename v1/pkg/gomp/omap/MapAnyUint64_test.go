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
		var k interface{} = "f501795f-1603-4295-a7ab-ffb04ae7f6e7"
		var v uint64 = 15119194464845422083

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint64_Delete(t *testing.T) {
	Convey("TestMapAnyUint64.Delete", t, func() {
		var k interface{} = "37181000-40c9-4067-89d4-a336bea70edd"
		var v uint64 = 1700926619572515183

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint64_Has(t *testing.T) {
	Convey("TestMapAnyUint64.Has", t, func() {
		var k interface{} = "2a27a26f-6507-437d-89d3-6f7e68dc133b"
		var v uint64 = 3646548946484008717

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("88b12b2c-86cc-40cf-b21c-eefbb8399ac5"+"c770e0c3-fe28-4789-a623-e93dcd73a459"), ShouldBeFalse)
	})
}


func TestMapAnyUint64_Get(t *testing.T) {
	Convey("TestMapAnyUint64.Get", t, func() {
		var k interface{} = "39308708-a9fd-4c34-af81-a6abd75f12ee"
		var v uint64 = 9362112000656537374

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("7a304d00-1d78-4970-9938-2c5acc1fc634" + "6b11dd99-26c4-425c-a91c-90db88427c1e")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint64_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint64.GetOpt", t, func() {
		var k interface{} = "fddbaf9f-10ca-4136-a9cb-a903c21c6af7"
		var v uint64 = 2699321915174859821

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("70f3dfc4-1d96-4018-8533-cba64051f8db" + "b0bf62fb-8d26-414e-b3d0-fed620d3cef8")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint64_ForEach(t *testing.T) {
	Convey("TestMapAnyUint64.ForEach", t, func() {
		var k interface{} = "14425485-b2d0-4cbd-9458-2588b93c2f21"
		var v uint64 = 16428132416057887858
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
		var k interface{} = "c8525252-bddd-47ae-913f-3fe6673920ab"
		var v uint64 = 2717804400238079477

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
		var k interface{} = "0e109ee5-13d9-4d71-babe-0ed508174cf7"
		var v uint64 = 16237274911114966559

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
		var k interface{} = "20202196-e973-406e-851d-b257064192f9"
		var v uint64 = 13406962860324283119

		test := omap.NewMapAnyUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("6aa00a92-489c-4839-b383-0ddf46280c22", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 11208199338989538219
		So(test.PutIfNotNil("693faa60-808f-4d90-a762-3a76c05e5518", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceIfExists", t, func() {
		var k interface{} = "9c32c174-01b5-49d7-88a6-45f19355c185"
		var v uint64 = 3721640223038801333
		var x uint64 = 1248901234712503694

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1faa8ffb-511a-40e2-90c7-45f1b0f2217c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceOrPut", t, func() {
		var k interface{} = "e7d41110-c16b-4817-8a0c-2d36b9e95b6e"
		var v uint64 = 18398343534181659728
		var x uint64 = 5936013382165576932

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("91fa48ca-6e3b-4386-9e2f-b8ef9f37653b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint64.MarshalJSON", t, func() {
		var k interface{} = "970daa71-a80f-4d5d-baa7-014fff1a9a7c"
		var v uint64 = 4633300009221191750

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"970daa71-a80f-4d5d-baa7-014fff1a9a7c","value":4633300009221191750}]`)
	})
}
