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
		var k uint = 553315997
		var v float64 = 0.896

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat64_Delete(t *testing.T) {
	Convey("TestMapUintFloat64.Delete", t, func() {
		var k uint = 151537740
		var v float64 = 0.542

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat64_Has(t *testing.T) {
	Convey("TestMapUintFloat64.Has", t, func() {
		var k uint = 1967187841
		var v float64 = 0.849

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(4273700353+2783957110), ShouldBeFalse)
	})
}


func TestMapUintFloat64_Get(t *testing.T) {
	Convey("TestMapUintFloat64.Get", t, func() {
		var k uint = 2085218125
		var v float64 = 0.303

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2248827506 + 2416041532)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintFloat64_GetOpt(t *testing.T) {
	Convey("TestMapUintFloat64.GetOpt", t, func() {
		var k uint = 2007565983
		var v float64 = 0.383

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(4178917544 + 3563094634)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintFloat64_ForEach(t *testing.T) {
	Convey("TestMapUintFloat64.ForEach", t, func() {
		var k uint = 979820963
		var v float64 = 0.463
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
		var k uint = 3324713207
		var v float64 = 0.246

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
		var k uint = 310314027
		var v float64 = 0.387

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
		var k uint = 1789248610
		var v float64 = 0.581

		test := omap.NewMapUintFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1203946432, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.262
		So(test.PutIfNotNil(669513801, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceIfExists", t, func() {
		var k uint = 2407463472
		var v float64 = 0.351
		var x float64 = 0.582

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(4014544296, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceOrPut", t, func() {
		var k uint = 1552112382
		var v float64 = 0.675
		var x float64 = 0.032

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1460603781, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintFloat64.MarshalJSON", t, func() {
		var k uint = 2349490854
		var v float64 = 0.225

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2349490854,"value":0.225}]`)
	})
}

