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
		var k uint = 1965455059
		var v rune = 243378315

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintRune_Delete(t *testing.T) {
	Convey("TestMapUintRune.Delete", t, func() {
		var k uint = 3045187743
		var v rune = 15062297

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintRune_Has(t *testing.T) {
	Convey("TestMapUintRune.Has", t, func() {
		var k uint = 2763193076
		var v rune = 1431007489

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3092711207+4144024451), ShouldBeFalse)
	})
}

func TestMapUintRune_Get(t *testing.T) {
	Convey("TestMapUintRune.Get", t, func() {
		var k uint = 3516427375
		var v rune = 1070405205

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2407080106 + 4135174439)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintRune_GetOpt(t *testing.T) {
	Convey("TestMapUintRune.GetOpt", t, func() {
		var k uint = 3715852350
		var v rune = 123308577

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3541107543 + 2484226171)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintRune_ForEach(t *testing.T) {
	Convey("TestMapUintRune.ForEach", t, func() {
		var k uint = 2607168341
		var v rune = 1229930966
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
		var k uint = 1670347778
		var v rune = 1625559436

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
		var k uint = 2079037960
		var v rune = 1225382877

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
		var k uint = 1148826760
		var v rune = 1707889275

		test := omap.NewMapUintRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1213335233, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1864369346
		So(test.PutIfNotNil(4158205620, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintRune.ReplaceIfExists", t, func() {
		var k uint = 1160985738
		var v rune = 869354660
		var x rune = 1734867627

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2791325215, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintRune.ReplaceOrPut", t, func() {
		var k uint = 2913732349
		var v rune = 1554170920
		var x rune = 1691886958

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2889547357, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_MarshalJSON(t *testing.T) {
	Convey("TestMapUintRune.MarshalJSON", t, func() {
		var k uint = 3762197377
		var v rune = 986857875

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3762197377,"value":986857875}]`)
	})
}
