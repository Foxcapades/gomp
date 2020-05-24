package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintFloat64_Put(t *testing.T) {
	Convey("TestMapUintFloat64.Put", t, func() {
		var k uint = 234168343
		var v float64 = 0.800

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat64_Delete(t *testing.T) {
	Convey("TestMapUintFloat64.Delete", t, func() {
		var k uint = 457417944
		var v float64 = 0.143

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat64_Has(t *testing.T) {
	Convey("TestMapUintFloat64.Has", t, func() {
		var k uint = 1188145642
		var v float64 = 0.350

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3390518977+3146279013), ShouldBeFalse)
	})
}

func TestMapUintFloat64_Get(t *testing.T) {
	Convey("TestMapUintFloat64.Get", t, func() {
		var k uint = 3897008867
		var v float64 = 0.584

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3536732960 + 4244276899)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintFloat64_GetOpt(t *testing.T) {
	Convey("TestMapUintFloat64.GetOpt", t, func() {
		var k uint = 978775136
		var v float64 = 0.917

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2405597831 + 1701415282)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintFloat64_ForEach(t *testing.T) {
	Convey("TestMapUintFloat64.ForEach", t, func() {
		var k uint = 3112531910
		var v float64 = 0.925
		hits := 0

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapUintFloat64.MarshalYAML", t, func() {
		var k uint = 177547212
		var v float64 = 0.814

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintFloat64_ToYAML(t *testing.T) {
	Convey("TestMapUintFloat64.ToYAML", t, func() {
		var k uint = 357541071
		var v float64 = 0.723

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintFloat64.PutIfNotNil", t, func() {
		var k uint = 3312606693
		var v float64 = 0.262

		test := omap.NewMapUintFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1290243243, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.181
		So(test.PutIfNotNil(1159008527, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceIfExists", t, func() {
		var k uint = 3866915110
		var v float64 = 0.711
		var x float64 = 0.468

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3486056129, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceOrPut", t, func() {
		var k uint = 1718927026
		var v float64 = 0.008
		var x float64 = 0.919

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2757922698, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintFloat64.MarshalJSON", t, func() {
		var k uint = 4105291229
		var v float64 = 0.813

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":4105291229,"value":0.813}]`)
	})
}
