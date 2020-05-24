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
		var k int = 255938856
		var v int32 = 1187575011

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt32_Delete(t *testing.T) {
	Convey("TestMapIntInt32.Delete", t, func() {
		var k int = 1949541974
		var v int32 = 739734223

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt32_Has(t *testing.T) {
	Convey("TestMapIntInt32.Has", t, func() {
		var k int = 120135319
		var v int32 = 1539482299

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1038024158+405538611), ShouldBeFalse)
	})
}

func TestMapIntInt32_Get(t *testing.T) {
	Convey("TestMapIntInt32.Get", t, func() {
		var k int = 385843300
		var v int32 = 1799241319

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1691472332 + 1147323170)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt32_GetOpt(t *testing.T) {
	Convey("TestMapIntInt32.GetOpt", t, func() {
		var k int = 1839433529
		var v int32 = 601655180

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(513969879 + 1522585766)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt32_ForEach(t *testing.T) {
	Convey("TestMapIntInt32.ForEach", t, func() {
		var k int = 1873056335
		var v int32 = 937762805
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
		var k int = 1628463779
		var v int32 = 2127872162

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
		var k int = 1496844744
		var v int32 = 2001231065

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
		var k int = 1370048759
		var v int32 = 985142638

		test := omap.NewMapIntInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2065960041, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 1673606771
		So(test.PutIfNotNil(49400005, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt32.ReplaceIfExists", t, func() {
		var k int = 1550974988
		var v int32 = 1835790453
		var x int32 = 884538465

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1832592762, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt32.ReplaceOrPut", t, func() {
		var k int = 640828284
		var v int32 = 1367087225
		var x int32 = 618087449

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1926189914, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt32.MarshalJSON", t, func() {
		var k int = 840334096
		var v int32 = 258703999

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":840334096,"value":258703999}]`)
	})
}
