package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntByte_Put(t *testing.T) {
	Convey("TestMapIntByte.Put", t, func() {
		var k int = 420541753
		var v byte = 251

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntByte_Delete(t *testing.T) {
	Convey("TestMapIntByte.Delete", t, func() {
		var k int = 197411892
		var v byte = 37

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntByte_Has(t *testing.T) {
	Convey("TestMapIntByte.Has", t, func() {
		var k int = 1340360750
		var v byte = 178

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(501405086+49006327), ShouldBeFalse)
	})
}


func TestMapIntByte_Get(t *testing.T) {
	Convey("TestMapIntByte.Get", t, func() {
		var k int = 389917734
		var v byte = 63

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1163946576 + 606380185)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntByte_GetOpt(t *testing.T) {
	Convey("TestMapIntByte.GetOpt", t, func() {
		var k int = 1203753751
		var v byte = 61

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1114357751 + 1165603645)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntByte_ForEach(t *testing.T) {
	Convey("TestMapIntByte.ForEach", t, func() {
		var k int = 1900221699
		var v byte = 32
		hits := 0

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntByte_MarshalYAML(t *testing.T) {
	Convey("TestMapIntByte.MarshalYAML", t, func() {
		var k int = 1401201191
		var v byte = 217

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntByte_ToYAML(t *testing.T) {
	Convey("TestMapIntByte.ToYAML", t, func() {
		var k int = 146838768
		var v byte = 122

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntByte.PutIfNotNil", t, func() {
		var k int = 135434018
		var v byte = 174

		test := omap.NewMapIntByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1447404782, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 239
		So(test.PutIfNotNil(512932623, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntByte.ReplaceIfExists", t, func() {
		var k int = 1054336152
		var v byte = 239
		var x byte = 37

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1074671749, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntByte.ReplaceOrPut", t, func() {
		var k int = 656773514
		var v byte = 220
		var x byte = 50

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(534083993, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_MarshalJSON(t *testing.T) {
	Convey("TestMapIntByte.MarshalJSON", t, func() {
		var k int = 1646861794
		var v byte = 187

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1646861794,"value":187}]`)
	})
}

