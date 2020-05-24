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
		var k uint = 1629666898
		var v rune = 1330408102

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintRune_Delete(t *testing.T) {
	Convey("TestMapUintRune.Delete", t, func() {
		var k uint = 1308780016
		var v rune = 1272693974

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintRune_Has(t *testing.T) {
	Convey("TestMapUintRune.Has", t, func() {
		var k uint = 230647372
		var v rune = 767862217

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2775924520+2731707222), ShouldBeFalse)
	})
}


func TestMapUintRune_Get(t *testing.T) {
	Convey("TestMapUintRune.Get", t, func() {
		var k uint = 67818396
		var v rune = 467700781

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(280054273 + 4062535855)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintRune_GetOpt(t *testing.T) {
	Convey("TestMapUintRune.GetOpt", t, func() {
		var k uint = 4264595071
		var v rune = 914114207

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(4249089470 + 1936964630)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintRune_ForEach(t *testing.T) {
	Convey("TestMapUintRune.ForEach", t, func() {
		var k uint = 4000570677
		var v rune = 1467476746
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
		var k uint = 2463848396
		var v rune = 1354308556

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
		var k uint = 2369495995
		var v rune = 34752242

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
		var k uint = 50594907
		var v rune = 1503016786

		test := omap.NewMapUintRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3417151693, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 845756697
		So(test.PutIfNotNil(3086260502, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintRune.ReplaceIfExists", t, func() {
		var k uint = 3130969275
		var v rune = 1419023925
		var x rune = 91149406

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3142282507, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintRune.ReplaceOrPut", t, func() {
		var k uint = 1529910171
		var v rune = 48271387
		var x rune = 2015439098

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1757049167, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_MarshalJSON(t *testing.T) {
	Convey("TestMapUintRune.MarshalJSON", t, func() {
		var k uint = 2566001360
		var v rune = 2140995361

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2566001360,"value":2140995361}]`)
	})
}
