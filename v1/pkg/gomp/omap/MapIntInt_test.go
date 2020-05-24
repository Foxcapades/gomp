package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt_Put(t *testing.T) {
	Convey("TestMapIntInt.Put", t, func() {
		var k int = 664848701
		var v int = 1071814662

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt_Delete(t *testing.T) {
	Convey("TestMapIntInt.Delete", t, func() {
		var k int = 1577022956
		var v int = 1643918694

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt_Has(t *testing.T) {
	Convey("TestMapIntInt.Has", t, func() {
		var k int = 1640509625
		var v int = 825641800

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1430911378+1076777096), ShouldBeFalse)
	})
}


func TestMapIntInt_Get(t *testing.T) {
	Convey("TestMapIntInt.Get", t, func() {
		var k int = 493791115
		var v int = 2133759492

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1373594322+1037533452)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt_GetOpt(t *testing.T) {
	Convey("TestMapIntInt.GetOpt", t, func() {
		var k int = 2066025381
		var v int = 1313275210

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1189579127+391276791)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt_ForEach(t *testing.T) {
	Convey("TestMapIntInt.ForEach", t, func() {
		var k int = 1520422710
		var v int = 845504510
		hits := 0

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt.MarshalYAML", t, func() {
		var k int = 1347166839
		var v int = 1669901499

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt_ToYAML(t *testing.T) {
	Convey("TestMapIntInt.ToYAML", t, func() {
		var k int = 1605797917
		var v int = 1493900588

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt.PutIfNotNil", t, func() {
		var k int = 91026754
		var v int = 1308910557

		test := omap.NewMapIntInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(550364043, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 935145028
		So(test.PutIfNotNil(1490768993, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt.ReplaceIfExists", t, func() {
		var k int = 378912996
		var v int = 2113560097
		var x int = 15932478

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2077143718, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt.ReplaceOrPut", t, func() {
		var k int = 615612256
		var v int = 1142565239
		var x int = 344194074

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(634797, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt.MarshalJSON", t, func() {
		var k int = 824634512
		var v int = 597327680

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":824634512,"value":597327680}]`)
	})
}

