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
		var k uint = 135401319
		var v int64 = 328483009160204397

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt64_Delete(t *testing.T) {
	Convey("TestMapUintInt64.Delete", t, func() {
		var k uint = 2069359381
		var v int64 = 7615532348345487052

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt64_Has(t *testing.T) {
	Convey("TestMapUintInt64.Has", t, func() {
		var k uint = 858038823
		var v int64 = 7647242388558268358

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(39057696+3135349492), ShouldBeFalse)
	})
}


func TestMapUintInt64_Get(t *testing.T) {
	Convey("TestMapUintInt64.Get", t, func() {
		var k uint = 2220354183
		var v int64 = 3003204286748569704

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2051849340 + 2415143682)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt64_GetOpt(t *testing.T) {
	Convey("TestMapUintInt64.GetOpt", t, func() {
		var k uint = 3466924437
		var v int64 = 6272809646986057553

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3275296924 + 3972637319)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt64_ForEach(t *testing.T) {
	Convey("TestMapUintInt64.ForEach", t, func() {
		var k uint = 483309820
		var v int64 = 4405489958911279829
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
		var k uint = 1814863508
		var v int64 = 5787664814970307435

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
		var k uint = 3995501784
		var v int64 = 6336935614900356014

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
		var k uint = 2117319135
		var v int64 = 3824425110570840729

		test := omap.NewMapUintInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(337669519, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 8030643803593836268
		So(test.PutIfNotNil(238823248, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt64.ReplaceIfExists", t, func() {
		var k uint = 4125093488
		var v int64 = 1676763963911035688
		var x int64 = 6662062096925224633

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1805650689, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt64.ReplaceOrPut", t, func() {
		var k uint = 3023692740
		var v int64 = 6737009827702945446
		var x int64 = 5611690278100545178

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3829213176, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt64.MarshalJSON", t, func() {
		var k uint = 2646632720
		var v int64 = 4194552206459853255

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2646632720,"value":4194552206459853255}]`)
	})
}
