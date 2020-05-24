package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt16_Put(t *testing.T) {
	Convey("TestMapUintInt16.Put", t, func() {
		var k uint = 3145905352
		var v int16 = 8675

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt16_Delete(t *testing.T) {
	Convey("TestMapUintInt16.Delete", t, func() {
		var k uint = 3440206826
		var v int16 = 13305

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt16_Has(t *testing.T) {
	Convey("TestMapUintInt16.Has", t, func() {
		var k uint = 1195893129
		var v int16 = 14050

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3555394394+886988138), ShouldBeFalse)
	})
}


func TestMapUintInt16_Get(t *testing.T) {
	Convey("TestMapUintInt16.Get", t, func() {
		var k uint = 2402140926
		var v int16 = 15756

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(379485986 + 1746055028)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt16_GetOpt(t *testing.T) {
	Convey("TestMapUintInt16.GetOpt", t, func() {
		var k uint = 3180729121
		var v int16 = 5501

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2688113379 + 2905547360)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt16_ForEach(t *testing.T) {
	Convey("TestMapUintInt16.ForEach", t, func() {
		var k uint = 1761880344
		var v int16 = 24505
		hits := 0

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt16.MarshalYAML", t, func() {
		var k uint = 911528527
		var v int16 = 8167

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt16_ToYAML(t *testing.T) {
	Convey("TestMapUintInt16.ToYAML", t, func() {
		var k uint = 2690879797
		var v int16 = 19342

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt16.PutIfNotNil", t, func() {
		var k uint = 1488415602
		var v int16 = 5445

		test := omap.NewMapUintInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3061413585, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 14284
		So(test.PutIfNotNil(2316668717, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt16.ReplaceIfExists", t, func() {
		var k uint = 2594419365
		var v int16 = 23848
		var x int16 = 29553

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(799680918, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt16.ReplaceOrPut", t, func() {
		var k uint = 50741903
		var v int16 = 18364
		var x int16 = 18421

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2035076050, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt16.MarshalJSON", t, func() {
		var k uint = 208271515
		var v int16 = 18978

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":208271515,"value":18978}]`)
	})
}
