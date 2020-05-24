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
		var k uint = 1711333347
		var v int16 = 19963

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt16_Delete(t *testing.T) {
	Convey("TestMapUintInt16.Delete", t, func() {
		var k uint = 1973907041
		var v int16 = 20601

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt16_Has(t *testing.T) {
	Convey("TestMapUintInt16.Has", t, func() {
		var k uint = 1347804474
		var v int16 = 19186

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3240861657+3451976742), ShouldBeFalse)
	})
}


func TestMapUintInt16_Get(t *testing.T) {
	Convey("TestMapUintInt16.Get", t, func() {
		var k uint = 3912073005
		var v int16 = 30000

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3849621866 + 53799930)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt16_GetOpt(t *testing.T) {
	Convey("TestMapUintInt16.GetOpt", t, func() {
		var k uint = 1230963194
		var v int16 = 27189

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1822554726 + 1580337869)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt16_ForEach(t *testing.T) {
	Convey("TestMapUintInt16.ForEach", t, func() {
		var k uint = 3352024999
		var v int16 = 14066
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
		var k uint = 824963225
		var v int16 = 31446

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
		var k uint = 783500497
		var v int16 = 14351

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
		var k uint = 3593633043
		var v int16 = 19332

		test := omap.NewMapUintInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1056177592, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 14492
		So(test.PutIfNotNil(2212246221, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt16.ReplaceIfExists", t, func() {
		var k uint = 862573146
		var v int16 = 14499
		var x int16 = 856

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2916966066, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt16.ReplaceOrPut", t, func() {
		var k uint = 2246938918
		var v int16 = 9861
		var x int16 = 424

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1757216200, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt16.MarshalJSON", t, func() {
		var k uint = 3901265205
		var v int16 = 31084

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3901265205,"value":31084}]`)
	})
}

