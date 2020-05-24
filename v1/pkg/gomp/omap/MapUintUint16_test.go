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
		var k uint = 3179096620
		var v uint16 = 62974

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint16_Delete(t *testing.T) {
	Convey("TestMapUintUint16.Delete", t, func() {
		var k uint = 2967054782
		var v uint16 = 16511

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint16_Has(t *testing.T) {
	Convey("TestMapUintUint16.Has", t, func() {
		var k uint = 1252777538
		var v uint16 = 50112

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3716076963+1463553829), ShouldBeFalse)
	})
}


func TestMapUintUint16_Get(t *testing.T) {
	Convey("TestMapUintUint16.Get", t, func() {
		var k uint = 4069311189
		var v uint16 = 10164

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2815402815 + 2973614366)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint16_GetOpt(t *testing.T) {
	Convey("TestMapUintUint16.GetOpt", t, func() {
		var k uint = 773748862
		var v uint16 = 48635

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2114295845 + 213586140)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint16_ForEach(t *testing.T) {
	Convey("TestMapUintUint16.ForEach", t, func() {
		var k uint = 2565492777
		var v uint16 = 5149
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
		var k uint = 78396451
		var v uint16 = 47833

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
		var k uint = 389013706
		var v uint16 = 13731

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
		var k uint = 3012929329
		var v uint16 = 25663

		test := omap.NewMapUintUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2513301228, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 10977
		So(test.PutIfNotNil(704885218, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceIfExists", t, func() {
		var k uint = 83993552
		var v uint16 = 4053
		var x uint16 = 64133

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3821762164, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint16.ReplaceOrPut", t, func() {
		var k uint = 22760059
		var v uint16 = 48434
		var x uint16 = 55532

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3706732744, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint16.MarshalJSON", t, func() {
		var k uint = 627174708
		var v uint16 = 25559

		test := omap.NewMapUintUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":627174708,"value":25559}]`)
	})
}

