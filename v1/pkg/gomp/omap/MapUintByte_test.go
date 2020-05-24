package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintByte_Put(t *testing.T) {
	Convey("TestMapUintByte.Put", t, func() {
		var k uint = 176493560
		var v byte = 92

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintByte_Delete(t *testing.T) {
	Convey("TestMapUintByte.Delete", t, func() {
		var k uint = 1881038527
		var v byte = 233

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintByte_Has(t *testing.T) {
	Convey("TestMapUintByte.Has", t, func() {
		var k uint = 711940867
		var v byte = 166

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3185001244+2981469343), ShouldBeFalse)
	})
}


func TestMapUintByte_Get(t *testing.T) {
	Convey("TestMapUintByte.Get", t, func() {
		var k uint = 2950080463
		var v byte = 220

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2785130945 + 1233790498)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintByte_GetOpt(t *testing.T) {
	Convey("TestMapUintByte.GetOpt", t, func() {
		var k uint = 2240891254
		var v byte = 136

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(4141277341 + 3691745437)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintByte_ForEach(t *testing.T) {
	Convey("TestMapUintByte.ForEach", t, func() {
		var k uint = 1143854159
		var v byte = 147
		hits := 0

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintByte_MarshalYAML(t *testing.T) {
	Convey("TestMapUintByte.MarshalYAML", t, func() {
		var k uint = 2686130782
		var v byte = 18

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintByte_ToYAML(t *testing.T) {
	Convey("TestMapUintByte.ToYAML", t, func() {
		var k uint = 3254670590
		var v byte = 99

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintByte.PutIfNotNil", t, func() {
		var k uint = 969769737
		var v byte = 184

		test := omap.NewMapUintByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1453313294, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 147
		So(test.PutIfNotNil(37197957, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintByte.ReplaceIfExists", t, func() {
		var k uint = 4133954352
		var v byte = 32
		var x byte = 149

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3686572284, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintByte.ReplaceOrPut", t, func() {
		var k uint = 3695900932
		var v byte = 12
		var x byte = 21

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(4060857778, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_MarshalJSON(t *testing.T) {
	Convey("TestMapUintByte.MarshalJSON", t, func() {
		var k uint = 4233214992
		var v byte = 29

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":4233214992,"value":29}]`)
	})
}

