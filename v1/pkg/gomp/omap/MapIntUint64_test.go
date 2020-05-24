package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint64_Put(t *testing.T) {
	Convey("TestMapIntUint64.Put", t, func() {
		var k int = 1960637462
		var v uint64 = 10251921531232531227

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint64_Delete(t *testing.T) {
	Convey("TestMapIntUint64.Delete", t, func() {
		var k int = 64715545
		var v uint64 = 1582327560773928322

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint64_Has(t *testing.T) {
	Convey("TestMapIntUint64.Has", t, func() {
		var k int = 1049815438
		var v uint64 = 4837764552569930053

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2033423697+1033573964), ShouldBeFalse)
	})
}

func TestMapIntUint64_Get(t *testing.T) {
	Convey("TestMapIntUint64.Get", t, func() {
		var k int = 1647856968
		var v uint64 = 6769900928052053669

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(273254799 + 756265196)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint64_GetOpt(t *testing.T) {
	Convey("TestMapIntUint64.GetOpt", t, func() {
		var k int = 1636354712
		var v uint64 = 1435168975135690243

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1600807082 + 1025996722)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint64_ForEach(t *testing.T) {
	Convey("TestMapIntUint64.ForEach", t, func() {
		var k int = 877694366
		var v uint64 = 11475658371762791454
		hits := 0

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint64.MarshalYAML", t, func() {
		var k int = 960238716
		var v uint64 = 13857850745051305456

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint64_ToYAML(t *testing.T) {
	Convey("TestMapIntUint64.ToYAML", t, func() {
		var k int = 680583034
		var v uint64 = 3684109403812587433

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint64.PutIfNotNil", t, func() {
		var k int = 1258979550
		var v uint64 = 4494771008306132298

		test := omap.NewMapIntUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1410802560, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 11663546591096917057
		So(test.PutIfNotNil(1932613110, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceIfExists", t, func() {
		var k int = 1945802998
		var v uint64 = 17647628506144631874
		var x uint64 = 4914530124824691578

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(956662845, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceOrPut", t, func() {
		var k int = 112699478
		var v uint64 = 8208293185156894277
		var x uint64 = 7762900489765343963

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1489555093, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint64.MarshalJSON", t, func() {
		var k int = 2037084747
		var v uint64 = 4901618223786470568

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2037084747,"value":4901618223786470568}]`)
	})
}
