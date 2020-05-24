package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt32_Put(t *testing.T) {
	Convey("TestMapIntInt32.Put", t, func() {
		var k int = 1160267556
		var v int32 = 1024108053

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt32_Delete(t *testing.T) {
	Convey("TestMapIntInt32.Delete", t, func() {
		var k int = 1690732907
		var v int32 = 154213233

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt32_Has(t *testing.T) {
	Convey("TestMapIntInt32.Has", t, func() {
		var k int = 2106215023
		var v int32 = 1381030323

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(919521377+8519057), ShouldBeFalse)
	})
}

func TestMapIntInt32_Get(t *testing.T) {
	Convey("TestMapIntInt32.Get", t, func() {
		var k int = 486180960
		var v int32 = 775566778

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1197556705 + 1454116234)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt32_GetOpt(t *testing.T) {
	Convey("TestMapIntInt32.GetOpt", t, func() {
		var k int = 1551427742
		var v int32 = 1669567988

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(482407643 + 71176756)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt32_ForEach(t *testing.T) {
	Convey("TestMapIntInt32.ForEach", t, func() {
		var k int = 1799569674
		var v int32 = 980497470
		hits := 0

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt32.MarshalYAML", t, func() {
		var k int = 1840101188
		var v int32 = 1942161216

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt32_ToYAML(t *testing.T) {
	Convey("TestMapIntInt32.ToYAML", t, func() {
		var k int = 2039918625
		var v int32 = 1217197904

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt32.PutIfNotNil", t, func() {
		var k int = 1413729708
		var v int32 = 721864198

		test := omap.NewMapIntInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1215659737, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 759406959
		So(test.PutIfNotNil(1524618018, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt32.ReplaceIfExists", t, func() {
		var k int = 1918640027
		var v int32 = 1803462539
		var x int32 = 1510691778

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(261552101, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt32.ReplaceOrPut", t, func() {
		var k int = 727241353
		var v int32 = 2032464848
		var x int32 = 1151939870

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(695259802, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt32.MarshalJSON", t, func() {
		var k int = 1538998240
		var v int32 = 1556263771

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1538998240,"value":1556263771}]`)
	})
}
