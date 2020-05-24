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
		var k uint = 2849370580
		var v byte = 165

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintByte_Delete(t *testing.T) {
	Convey("TestMapUintByte.Delete", t, func() {
		var k uint = 3114173321
		var v byte = 135

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintByte_Has(t *testing.T) {
	Convey("TestMapUintByte.Has", t, func() {
		var k uint = 2712770495
		var v byte = 50

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(626053040+1501498755), ShouldBeFalse)
	})
}


func TestMapUintByte_Get(t *testing.T) {
	Convey("TestMapUintByte.Get", t, func() {
		var k uint = 3320949694
		var v byte = 174

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1862811773 + 2650735461)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintByte_GetOpt(t *testing.T) {
	Convey("TestMapUintByte.GetOpt", t, func() {
		var k uint = 811204662
		var v byte = 182

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(993206684 + 2536074831)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintByte_ForEach(t *testing.T) {
	Convey("TestMapUintByte.ForEach", t, func() {
		var k uint = 969968889
		var v byte = 68
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
		var k uint = 213063559
		var v byte = 219

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
		var k uint = 88086615
		var v byte = 109

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
		var k uint = 3406047188
		var v byte = 103

		test := omap.NewMapUintByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3774870328, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 138
		So(test.PutIfNotNil(678267720, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintByte.ReplaceIfExists", t, func() {
		var k uint = 1860562815
		var v byte = 58
		var x byte = 1

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1276002745, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintByte.ReplaceOrPut", t, func() {
		var k uint = 1471882429
		var v byte = 53
		var x byte = 211

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2413951579, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_MarshalJSON(t *testing.T) {
	Convey("TestMapUintByte.MarshalJSON", t, func() {
		var k uint = 2002384542
		var v byte = 212

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2002384542,"value":212}]`)
	})
}
