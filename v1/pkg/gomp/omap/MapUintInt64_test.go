package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt64_Put(t *testing.T) {
	Convey("TestMapUintInt64.Put", t, func() {
		var k uint = 2552992446
		var v int64 = 4670018807430285986

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt64_Delete(t *testing.T) {
	Convey("TestMapUintInt64.Delete", t, func() {
		var k uint = 1683177803
		var v int64 = 4521672724270652074

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt64_Has(t *testing.T) {
	Convey("TestMapUintInt64.Has", t, func() {
		var k uint = 3990260763
		var v int64 = 899567037789221354

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1901170290+2571184182), ShouldBeFalse)
	})
}

func TestMapUintInt64_Get(t *testing.T) {
	Convey("TestMapUintInt64.Get", t, func() {
		var k uint = 1034080454
		var v int64 = 3587622827510794836

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1166232710 + 2258240018)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt64_GetOpt(t *testing.T) {
	Convey("TestMapUintInt64.GetOpt", t, func() {
		var k uint = 2176138909
		var v int64 = 8251211752311405533

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1845431437 + 3665604122)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt64_ForEach(t *testing.T) {
	Convey("TestMapUintInt64.ForEach", t, func() {
		var k uint = 3496366047
		var v int64 = 2142976131632003491
		hits := 0

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt64.MarshalYAML", t, func() {
		var k uint = 421149423
		var v int64 = 6067197025859569161

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt64_ToYAML(t *testing.T) {
	Convey("TestMapUintInt64.ToYAML", t, func() {
		var k uint = 4078497262
		var v int64 = 7521988515440189855

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt64.PutIfNotNil", t, func() {
		var k uint = 3769071951
		var v int64 = 4776183198100689321

		test := omap.NewMapUintInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(459568531, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 5533116617383968030
		So(test.PutIfNotNil(20551088, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt64.ReplaceIfExists", t, func() {
		var k uint = 2682419084
		var v int64 = 916534124652089221
		var x int64 = 2248008820376395770

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3468368656, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt64.ReplaceOrPut", t, func() {
		var k uint = 604304390
		var v int64 = 6253174709396420931
		var x int64 = 5423188013424231711

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(4005989840, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt64.MarshalJSON", t, func() {
		var k uint = 2911983873
		var v int64 = 3750255002226781859

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2911983873,"value":3750255002226781859}]`)
	})
}
