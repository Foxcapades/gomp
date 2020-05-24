package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt16_Put(t *testing.T) {
	Convey("TestMapIntInt16.Put", t, func() {
		var k int = 1483764985
		var v int16 = 28730

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt16_Delete(t *testing.T) {
	Convey("TestMapIntInt16.Delete", t, func() {
		var k int = 1987490753
		var v int16 = 15372

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt16_Has(t *testing.T) {
	Convey("TestMapIntInt16.Has", t, func() {
		var k int = 1789268374
		var v int16 = 21873

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(409372262+338679595), ShouldBeFalse)
	})
}

func TestMapIntInt16_Get(t *testing.T) {
	Convey("TestMapIntInt16.Get", t, func() {
		var k int = 1470951842
		var v int16 = 22619

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1163704765 + 607191253)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt16_GetOpt(t *testing.T) {
	Convey("TestMapIntInt16.GetOpt", t, func() {
		var k int = 1407367712
		var v int16 = 32254

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1697877100 + 1673869804)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt16_ForEach(t *testing.T) {
	Convey("TestMapIntInt16.ForEach", t, func() {
		var k int = 130431739
		var v int16 = 25511
		hits := 0

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt16.MarshalYAML", t, func() {
		var k int = 1638585455
		var v int16 = 6099

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt16_ToYAML(t *testing.T) {
	Convey("TestMapIntInt16.ToYAML", t, func() {
		var k int = 543496848
		var v int16 = 22425

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt16.PutIfNotNil", t, func() {
		var k int = 614530460
		var v int16 = 26024

		test := omap.NewMapIntInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(701599696, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 25336
		So(test.PutIfNotNil(2088725275, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceIfExists", t, func() {
		var k int = 968587621
		var v int16 = 6797
		var x int16 = 2618

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(903039628, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceOrPut", t, func() {
		var k int = 768131943
		var v int16 = 4224
		var x int16 = 21054

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1569864937, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt16.MarshalJSON", t, func() {
		var k int = 529114798
		var v int16 = 11739

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":529114798,"value":11739}]`)
	})
}
