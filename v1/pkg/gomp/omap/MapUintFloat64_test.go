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
		var k uint = 3263036276
		var v float64 = 0.337

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat64_Delete(t *testing.T) {
	Convey("TestMapUintFloat64.Delete", t, func() {
		var k uint = 2559239885
		var v float64 = 0.833

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat64_Has(t *testing.T) {
	Convey("TestMapUintFloat64.Has", t, func() {
		var k uint = 1204965195
		var v float64 = 0.775

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3898632397+3498196079), ShouldBeFalse)
	})
}


func TestMapUintFloat64_Get(t *testing.T) {
	Convey("TestMapUintFloat64.Get", t, func() {
		var k uint = 2524894148
		var v float64 = 0.785

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3733597319+2292999537)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintFloat64_GetOpt(t *testing.T) {
	Convey("TestMapUintFloat64.GetOpt", t, func() {
		var k uint = 3881568220
		var v float64 = 0.999

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3704179962+3929160943)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintFloat64_ForEach(t *testing.T) {
	Convey("TestMapUintFloat64.ForEach", t, func() {
		var k uint = 3483613217
		var v float64 = 0.731
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
		var k uint = 1045525849
		var v float64 = 0.669

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
		var k uint = 1188093490
		var v float64 = 0.267

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
		var k uint = 1969929076
		var v float64 = 0.578

		test := omap.NewMapUintFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2577815613, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.290
		So(test.PutIfNotNil(447997408, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceIfExists", t, func() {
		var k uint = 3582369540
		var v float64 = 0.605
		var x float64 = 0.829

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2040344712, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceOrPut", t, func() {
		var k uint = 3735217398
		var v float64 = 0.638
		var x float64 = 0.538

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1674797306, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintFloat64.MarshalJSON", t, func() {
		var k uint = 41158896
		var v float64 = 0.276

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":41158896,"value":0.276}]`)
	})
}

