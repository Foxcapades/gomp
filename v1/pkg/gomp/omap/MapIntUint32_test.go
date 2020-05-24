package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint32_Put(t *testing.T) {
	Convey("TestMapIntUint32.Put", t, func() {
		var k int = 515365047
		var v uint32 = 1769398787

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint32_Delete(t *testing.T) {
	Convey("TestMapIntUint32.Delete", t, func() {
		var k int = 1936876794
		var v uint32 = 2305060755

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint32_Has(t *testing.T) {
	Convey("TestMapIntUint32.Has", t, func() {
		var k int = 635629804
		var v uint32 = 498897588

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1963703424+813035196), ShouldBeFalse)
	})
}

func TestMapIntUint32_Get(t *testing.T) {
	Convey("TestMapIntUint32.Get", t, func() {
		var k int = 44662426
		var v uint32 = 2050204257

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(949641876 + 610904914)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint32_GetOpt(t *testing.T) {
	Convey("TestMapIntUint32.GetOpt", t, func() {
		var k int = 447023948
		var v uint32 = 3338906437

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1601596210 + 1143503434)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint32_ForEach(t *testing.T) {
	Convey("TestMapIntUint32.ForEach", t, func() {
		var k int = 554096030
		var v uint32 = 3218890487
		hits := 0

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint32.MarshalYAML", t, func() {
		var k int = 407093202
		var v uint32 = 665526698

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint32_ToYAML(t *testing.T) {
	Convey("TestMapIntUint32.ToYAML", t, func() {
		var k int = 1587080871
		var v uint32 = 1265743556

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint32.PutIfNotNil", t, func() {
		var k int = 1116755698
		var v uint32 = 2896932825

		test := omap.NewMapIntUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1554619879, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 3549907122
		So(test.PutIfNotNil(808816989, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceIfExists", t, func() {
		var k int = 98727845
		var v uint32 = 1347393259
		var x uint32 = 591656455

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(857321466, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceOrPut", t, func() {
		var k int = 1824491377
		var v uint32 = 3309611132
		var x uint32 = 1796066534

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2145227769, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint32.MarshalJSON", t, func() {
		var k int = 992284717
		var v uint32 = 191323863

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":992284717,"value":191323863}]`)
	})
}
