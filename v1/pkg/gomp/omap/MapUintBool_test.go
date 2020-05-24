package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintBool_Put(t *testing.T) {
	Convey("TestMapUintBool.Put", t, func() {
		var k uint = 2152559191
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintBool_Delete(t *testing.T) {
	Convey("TestMapUintBool.Delete", t, func() {
		var k uint = 2098702851
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintBool_Has(t *testing.T) {
	Convey("TestMapUintBool.Has", t, func() {
		var k uint = 4060086538
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1208828854+1305137858), ShouldBeFalse)
	})
}


func TestMapUintBool_Get(t *testing.T) {
	Convey("TestMapUintBool.Get", t, func() {
		var k uint = 3163532185
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1819680956 + 484883583)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintBool_GetOpt(t *testing.T) {
	Convey("TestMapUintBool.GetOpt", t, func() {
		var k uint = 4211340127
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(829311775 + 3447209036)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintBool_ForEach(t *testing.T) {
	Convey("TestMapUintBool.ForEach", t, func() {
		var k uint = 2341261293
		var v bool = false
		hits := 0

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintBool_MarshalYAML(t *testing.T) {
	Convey("TestMapUintBool.MarshalYAML", t, func() {
		var k uint = 3384786957
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintBool_ToYAML(t *testing.T) {
	Convey("TestMapUintBool.ToYAML", t, func() {
		var k uint = 364456774
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintBool.PutIfNotNil", t, func() {
		var k uint = 2749978987
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3328576339, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil(4125878684, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintBool.ReplaceIfExists", t, func() {
		var k uint = 2535997690
		var v bool = false
		var x bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1744926987, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintBool.ReplaceOrPut", t, func() {
		var k uint = 3576154116
		var v bool = false
		var x bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(450417676, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintBool_MarshalJSON(t *testing.T) {
	Convey("TestMapUintBool.MarshalJSON", t, func() {
		var k uint = 2109644049
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2109644049,"value":false}]`)
	})
}

