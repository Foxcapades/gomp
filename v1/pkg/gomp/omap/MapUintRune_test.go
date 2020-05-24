package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintRune_Put(t *testing.T) {
	Convey("TestMapUintRune.Put", t, func() {
		var k uint = 3862378213
		var v rune = 1298566174

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintRune_Delete(t *testing.T) {
	Convey("TestMapUintRune.Delete", t, func() {
		var k uint = 1264093799
		var v rune = 669679913

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintRune_Has(t *testing.T) {
	Convey("TestMapUintRune.Has", t, func() {
		var k uint = 1011771309
		var v rune = 740338652

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(4194563016+12532987), ShouldBeFalse)
	})
}


func TestMapUintRune_Get(t *testing.T) {
	Convey("TestMapUintRune.Get", t, func() {
		var k uint = 2914259337
		var v rune = 387895839

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(4030333591 + 2196776213)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintRune_GetOpt(t *testing.T) {
	Convey("TestMapUintRune.GetOpt", t, func() {
		var k uint = 2631254259
		var v rune = 23695061

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3138391608 + 3755023146)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintRune_ForEach(t *testing.T) {
	Convey("TestMapUintRune.ForEach", t, func() {
		var k uint = 873536538
		var v rune = 924140006
		hits := 0

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv rune) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintRune_MarshalYAML(t *testing.T) {
	Convey("TestMapUintRune.MarshalYAML", t, func() {
		var k uint = 2595935926
		var v rune = 119279734

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintRune_ToYAML(t *testing.T) {
	Convey("TestMapUintRune.ToYAML", t, func() {
		var k uint = 476821039
		var v rune = 896019341

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintRune.PutIfNotNil", t, func() {
		var k uint = 221552599
		var v rune = 1230061743

		test := omap.NewMapUintRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(539789691, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1406175602
		So(test.PutIfNotNil(696763143, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintRune.ReplaceIfExists", t, func() {
		var k uint = 3898767596
		var v rune = 1563120512
		var x rune = 488055680

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2883882456, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintRune.ReplaceOrPut", t, func() {
		var k uint = 3470699150
		var v rune = 245610691
		var x rune = 253103435

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1828229412, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_MarshalJSON(t *testing.T) {
	Convey("TestMapUintRune.MarshalJSON", t, func() {
		var k uint = 684164423
		var v rune = 650483291

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":684164423,"value":650483291}]`)
	})
}

