package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntBool_Put(t *testing.T) {
	Convey("TestMapIntBool.Put", t, func() {
		var k int = 230253769
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntBool_Delete(t *testing.T) {
	Convey("TestMapIntBool.Delete", t, func() {
		var k int = 1367193953
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntBool_Has(t *testing.T) {
	Convey("TestMapIntBool.Has", t, func() {
		var k int = 1364979433
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(971160591+102093755), ShouldBeFalse)
	})
}


func TestMapIntBool_Get(t *testing.T) {
	Convey("TestMapIntBool.Get", t, func() {
		var k int = 133915873
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2029848776 + 2003305208)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntBool_GetOpt(t *testing.T) {
	Convey("TestMapIntBool.GetOpt", t, func() {
		var k int = 256632614
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1242049558 + 1267700717)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntBool_ForEach(t *testing.T) {
	Convey("TestMapIntBool.ForEach", t, func() {
		var k int = 1918677922
		var v bool = false
		hits := 0

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntBool_MarshalYAML(t *testing.T) {
	Convey("TestMapIntBool.MarshalYAML", t, func() {
		var k int = 57765762
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntBool_ToYAML(t *testing.T) {
	Convey("TestMapIntBool.ToYAML", t, func() {
		var k int = 1680181838
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntBool.PutIfNotNil", t, func() {
		var k int = 1097937381
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1924737714, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil(1005470934, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntBool.ReplaceIfExists", t, func() {
		var k int = 1049546517
		var v bool = false
		var x bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(5952621, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntBool.ReplaceOrPut", t, func() {
		var k int = 1579287429
		var v bool = false
		var x bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1223762343, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntBool_MarshalJSON(t *testing.T) {
	Convey("TestMapIntBool.MarshalJSON", t, func() {
		var k int = 1895497357
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1895497357,"value":false}]`)
	})
}
