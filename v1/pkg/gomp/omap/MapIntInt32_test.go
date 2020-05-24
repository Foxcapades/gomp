package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt32_Put(t *testing.T) {
	Convey("TestMapIntInt32.Put", t, func() {
		var k int = 300004641
		var v int32 = 1353268903

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt32_Delete(t *testing.T) {
	Convey("TestMapIntInt32.Delete", t, func() {
		var k int = 589933195
		var v int32 = 659399676

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt32_Has(t *testing.T) {
	Convey("TestMapIntInt32.Has", t, func() {
		var k int = 328446485
		var v int32 = 2125511370

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(714014492+525394231), ShouldBeFalse)
	})
}


func TestMapIntInt32_Get(t *testing.T) {
	Convey("TestMapIntInt32.Get", t, func() {
		var k int = 929666376
		var v int32 = 385810858

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(206642463+2029562338)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt32_GetOpt(t *testing.T) {
	Convey("TestMapIntInt32.GetOpt", t, func() {
		var k int = 466586676
		var v int32 = 712711405

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(850540351+1457179569)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt32_ForEach(t *testing.T) {
	Convey("TestMapIntInt32.ForEach", t, func() {
		var k int = 696502064
		var v int32 = 739021163
		hits := 0

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt32.MarshalYAML", t, func() {
		var k int = 950668738
		var v int32 = 1211537102

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt32_ToYAML(t *testing.T) {
	Convey("TestMapIntInt32.ToYAML", t, func() {
		var k int = 1651128055
		var v int32 = 300662033

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt32.PutIfNotNil", t, func() {
		var k int = 969327454
		var v int32 = 572972194

		test := omap.NewMapIntInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1455475158, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 105582142
		So(test.PutIfNotNil(1155735763, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt32.ReplaceIfExists", t, func() {
		var k int = 582309068
		var v int32 = 74755754
		var x int32 = 168486021

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2101200044, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt32.ReplaceOrPut", t, func() {
		var k int = 507860699
		var v int32 = 686534527
		var x int32 = 509417076

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(491760842, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt32.MarshalJSON", t, func() {
		var k int = 1149995596
		var v int32 = 1289091562

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1149995596,"value":1289091562}]`)
	})
}

