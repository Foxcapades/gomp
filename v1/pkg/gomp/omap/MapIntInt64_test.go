package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt64_Put(t *testing.T) {
	Convey("TestMapIntInt64.Put", t, func() {
		var k int = 1019940614
		var v int64 = 7359819941723636277

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt64_Delete(t *testing.T) {
	Convey("TestMapIntInt64.Delete", t, func() {
		var k int = 1012872941
		var v int64 = 6263120719032436942

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt64_Has(t *testing.T) {
	Convey("TestMapIntInt64.Has", t, func() {
		var k int = 990980868
		var v int64 = 5308348682807948410

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2066207109+925855774), ShouldBeFalse)
	})
}


func TestMapIntInt64_Get(t *testing.T) {
	Convey("TestMapIntInt64.Get", t, func() {
		var k int = 448864905
		var v int64 = 8253946786223352973

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1253661421+451331916)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt64_GetOpt(t *testing.T) {
	Convey("TestMapIntInt64.GetOpt", t, func() {
		var k int = 1589488163
		var v int64 = 5973553408961331010

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(846622966+1204804964)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt64_ForEach(t *testing.T) {
	Convey("TestMapIntInt64.ForEach", t, func() {
		var k int = 1480756832
		var v int64 = 8572432729949104877
		hits := 0

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt64.MarshalYAML", t, func() {
		var k int = 324790493
		var v int64 = 5315149041077165732

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt64_ToYAML(t *testing.T) {
	Convey("TestMapIntInt64.ToYAML", t, func() {
		var k int = 1450366767
		var v int64 = 618383929038347771

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt64.PutIfNotNil", t, func() {
		var k int = 1586345893
		var v int64 = 7761965528827749883

		test := omap.NewMapIntInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(689517343, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 9144113976330165544
		So(test.PutIfNotNil(338656840, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceIfExists", t, func() {
		var k int = 341489324
		var v int64 = 1508718253340381471
		var x int64 = 2820844218222387120

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(550035386, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceOrPut", t, func() {
		var k int = 1679564689
		var v int64 = 2927292326986397066
		var x int64 = 1015128548958209735

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(160023041, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt64.MarshalJSON", t, func() {
		var k int = 1422013869
		var v int64 = 8054837581998665220

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1422013869,"value":8054837581998665220}]`)
	})
}

