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
		var k int = 705151794
		var v byte = 246

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntByte_Delete(t *testing.T) {
	Convey("TestMapIntByte.Delete", t, func() {
		var k int = 694112319
		var v byte = 161

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntByte_Has(t *testing.T) {
	Convey("TestMapIntByte.Has", t, func() {
		var k int = 545642118
		var v byte = 66

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(932684219+657040551), ShouldBeFalse)
	})
}


func TestMapIntByte_Get(t *testing.T) {
	Convey("TestMapIntByte.Get", t, func() {
		var k int = 960963712
		var v byte = 100

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(224884318 + 2060693868)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntByte_GetOpt(t *testing.T) {
	Convey("TestMapIntByte.GetOpt", t, func() {
		var k int = 2062408163
		var v byte = 65

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1382671073 + 1100334884)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntByte_ForEach(t *testing.T) {
	Convey("TestMapIntByte.ForEach", t, func() {
		var k int = 288214577
		var v byte = 27
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
		var k int = 615707769
		var v byte = 243

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
		var k int = 676225477
		var v byte = 93

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
		var k int = 362598446
		var v byte = 138

		test := omap.NewMapIntByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(406559674, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 68
		So(test.PutIfNotNil(1556238673, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntByte.ReplaceIfExists", t, func() {
		var k int = 1270692325
		var v byte = 185
		var x byte = 27

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(402430363, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntByte.ReplaceOrPut", t, func() {
		var k int = 1355082969
		var v byte = 98
		var x byte = 70

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1348044332, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_MarshalJSON(t *testing.T) {
	Convey("TestMapIntByte.MarshalJSON", t, func() {
		var k int = 1770973063
		var v byte = 125

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1770973063,"value":125}]`)
	})
}
