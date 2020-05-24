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
		var k int = 1476491898
		var v int16 = 23643

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt16_Delete(t *testing.T) {
	Convey("TestMapIntInt16.Delete", t, func() {
		var k int = 892281606
		var v int16 = 22917

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt16_Has(t *testing.T) {
	Convey("TestMapIntInt16.Has", t, func() {
		var k int = 1600804661
		var v int16 = 8405

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(503127931+870794514), ShouldBeFalse)
	})
}


func TestMapIntInt16_Get(t *testing.T) {
	Convey("TestMapIntInt16.Get", t, func() {
		var k int = 1423076984
		var v int16 = 6378

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1635747835 + 1692821916)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt16_GetOpt(t *testing.T) {
	Convey("TestMapIntInt16.GetOpt", t, func() {
		var k int = 1050715989
		var v int16 = 31896

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(621839919 + 1743168524)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt16_ForEach(t *testing.T) {
	Convey("TestMapIntInt16.ForEach", t, func() {
		var k int = 846517188
		var v int16 = 25319
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
		var k int = 1037556135
		var v int16 = 5805

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
		var k int = 1281785693
		var v int16 = 10580

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
		var k int = 1427144178
		var v int16 = 12653

		test := omap.NewMapIntInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1297865422, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 4136
		So(test.PutIfNotNil(1232032020, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceIfExists", t, func() {
		var k int = 1911166761
		var v int16 = 3278
		var x int16 = 9225

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(135470712, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt16.ReplaceOrPut", t, func() {
		var k int = 457878249
		var v int16 = 31131
		var x int16 = 1254

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(792973845, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt16.MarshalJSON", t, func() {
		var k int = 1050573036
		var v int16 = 12440

		test := omap.NewMapIntInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1050573036,"value":12440}]`)
	})
}

