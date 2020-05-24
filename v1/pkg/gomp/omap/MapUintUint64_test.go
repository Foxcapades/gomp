package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint64_Put(t *testing.T) {
	Convey("TestMapUintUint64.Put", t, func() {
		var k uint = 1306730440
		var v uint64 = 5012193545276043067

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint64_Delete(t *testing.T) {
	Convey("TestMapUintUint64.Delete", t, func() {
		var k uint = 4083957950
		var v uint64 = 3687845101242397209

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint64_Has(t *testing.T) {
	Convey("TestMapUintUint64.Has", t, func() {
		var k uint = 1695322787
		var v uint64 = 18235520595226599240

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(4086563272+747018758), ShouldBeFalse)
	})
}

func TestMapUintUint64_Get(t *testing.T) {
	Convey("TestMapUintUint64.Get", t, func() {
		var k uint = 1929908440
		var v uint64 = 16229134195468535456

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3499706524 + 2048669610)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint64_GetOpt(t *testing.T) {
	Convey("TestMapUintUint64.GetOpt", t, func() {
		var k uint = 1373656858
		var v uint64 = 16757116933223832447

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(916410736 + 3857156848)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint64_ForEach(t *testing.T) {
	Convey("TestMapUintUint64.ForEach", t, func() {
		var k uint = 589654770
		var v uint64 = 9739838658136962780
		hits := 0

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint64.MarshalYAML", t, func() {
		var k uint = 1324812498
		var v uint64 = 17381707371198119570

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint64_ToYAML(t *testing.T) {
	Convey("TestMapUintUint64.ToYAML", t, func() {
		var k uint = 2386274284
		var v uint64 = 18293291993701346351

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint64.PutIfNotNil", t, func() {
		var k uint = 1415881125
		var v uint64 = 16859647845599194150

		test := omap.NewMapUintUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3012032167, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 3288901875576871880
		So(test.PutIfNotNil(363389665, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint64.ReplaceIfExists", t, func() {
		var k uint = 2585182298
		var v uint64 = 11912204368289003740
		var x uint64 = 4295015966756549211

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1714867743, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint64.ReplaceOrPut", t, func() {
		var k uint = 382720825
		var v uint64 = 15333708793190373343
		var x uint64 = 5766504888187663009

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1692862820, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint64.MarshalJSON", t, func() {
		var k uint = 1352123643
		var v uint64 = 401773472351962684

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1352123643,"value":401773472351962684}]`)
	})
}
