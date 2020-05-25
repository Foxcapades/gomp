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
		var k int = 439781094
		var v byte = 143

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntByte_Delete(t *testing.T) {
	Convey("TestMapIntByte.Delete", t, func() {
		var k int = 1604753471
		var v byte = 129

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntByte_Has(t *testing.T) {
	Convey("TestMapIntByte.Has", t, func() {
		var k int = 2071285875
		var v byte = 92

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1284687696+1783883947), ShouldBeFalse)
	})
}

func TestMapIntByte_Get(t *testing.T) {
	Convey("TestMapIntByte.Get", t, func() {
		var k int = 1360351751
		var v byte = 88

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1038367382 + 1175716234)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntByte_GetOpt(t *testing.T) {
	Convey("TestMapIntByte.GetOpt", t, func() {
		var k int = 1793343839
		var v byte = 107

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(713866536 + 1343040561)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntByte_ForEach(t *testing.T) {
	Convey("TestMapIntByte.ForEach", t, func() {
		var k int = 129710114
		var v byte = 11
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
		var k int = 1951213866
		var v byte = 69

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
		var k int = 1948087446
		var v byte = 95

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapIntByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntByte.PutIfNotNil", t, func() {
		var k int = 1690477897
		var v byte = 128

		test := omap.NewMapIntByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(898595302, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 134
		So(test.PutIfNotNil(1124630118, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntByte.ReplaceIfExists", t, func() {
		var k int = 1835616714
		var v byte = 22
		var x byte = 100

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(620602173, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntByte.ReplaceOrPut", t, func() {
		var k int = 1917464443
		var v byte = 184
		var x byte = 183

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1559906219, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_MarshalJSON(t *testing.T) {
	Convey("TestMapIntByte.MarshalJSON", t, func() {
		var k int = 469450036
		var v byte = 58

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":469450036,"value":58}]`)
	})
}
