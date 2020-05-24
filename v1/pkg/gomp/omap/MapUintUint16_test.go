package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint16_Put(t *testing.T) {
	Convey("TestMapUintUint16.Put", t, func() {
		var k uint = 3776216502
		var v uint16 = 42458

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint16_Delete(t *testing.T) {
	Convey("TestMapUintUint16.Delete", t, func() {
		var k uint = 3655428210
		var v uint16 = 55692

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint16_Has(t *testing.T) {
	Convey("TestMapUintUint16.Has", t, func() {
		var k uint = 3935614559
		var v uint16 = 63369

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(531801443+1435333635), ShouldBeFalse)
	})
}

func TestMapUintUint16_Get(t *testing.T) {
	Convey("TestMapUintUint16.Get", t, func() {
		var k uint = 3818857359
		var v uint16 = 2563

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1301811764 + 3555233204)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint16_GetOpt(t *testing.T) {
	Convey("TestMapUintUint16.GetOpt", t, func() {
		var k uint = 557334338
		var v uint16 = 7512

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3479219747 + 1730111753)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint16_ForEach(t *testing.T) {
	Convey("TestMapUintUint16.ForEach", t, func() {
		var k uint = 2404287903
		var v uint16 = 53158
		hits := 0

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint16.MarshalYAML", t, func() {
		var k uint = 2950057433
		var v uint16 = 32887

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint16_ToYAML(t *testing.T) {
	Convey("TestMapUintUint16.ToYAML", t, func() {
		var k uint = 4058331345
		var v uint16 = 35046

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint16.PutIfNotNil", t, func() {
		var k uint = 1884666413
		var v uint16 = 60428

		test := omap.NewMapUintUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1595402735, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 21887
		So(test.PutIfNotNil(3398805391, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceIfExists", t, func() {
		var k uint = 1891230918
		var v uint16 = 50147
		var x uint16 = 38386

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3963302653, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceOrPut", t, func() {
		var k uint = 25373011
		var v uint16 = 422
		var x uint16 = 22357

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3146969610, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint16.MarshalJSON", t, func() {
		var k uint = 846881044
		var v uint16 = 53988

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":846881044,"value":53988}]`)
	})
}
