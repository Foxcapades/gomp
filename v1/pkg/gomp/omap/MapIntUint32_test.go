package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint32_Put(t *testing.T) {
	Convey("TestMapIntUint32.Put", t, func() {
		var k int = 1730975170
		var v uint32 = 1817957059

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint32_Delete(t *testing.T) {
	Convey("TestMapIntUint32.Delete", t, func() {
		var k int = 1543879668
		var v uint32 = 2385046537

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint32_Has(t *testing.T) {
	Convey("TestMapIntUint32.Has", t, func() {
		var k int = 1412612364
		var v uint32 = 727498218

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(673474412+1080055552), ShouldBeFalse)
	})
}


func TestMapIntUint32_Get(t *testing.T) {
	Convey("TestMapIntUint32.Get", t, func() {
		var k int = 1469507612
		var v uint32 = 3468774731

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2110282199+1518289480)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint32_GetOpt(t *testing.T) {
	Convey("TestMapIntUint32.GetOpt", t, func() {
		var k int = 1680052947
		var v uint32 = 3218463936

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1845859253+1110039631)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint32_ForEach(t *testing.T) {
	Convey("TestMapIntUint32.ForEach", t, func() {
		var k int = 1780690128
		var v uint32 = 695889578
		hits := 0

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint32.MarshalYAML", t, func() {
		var k int = 1035841837
		var v uint32 = 3980453656

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint32_ToYAML(t *testing.T) {
	Convey("TestMapIntUint32.ToYAML", t, func() {
		var k int = 1397360134
		var v uint32 = 3335194967

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint32.PutIfNotNil", t, func() {
		var k int = 1895840062
		var v uint32 = 1568789249

		test := omap.NewMapIntUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1181677850, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 3594160083
		So(test.PutIfNotNil(2094548180, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceIfExists", t, func() {
		var k int = 699607064
		var v uint32 = 3242026653
		var x uint32 = 3734158685

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1132909312, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceOrPut", t, func() {
		var k int = 234338898
		var v uint32 = 3251386075
		var x uint32 = 4035398249

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2088247948, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint32.MarshalJSON", t, func() {
		var k int = 1244918406
		var v uint32 = 1654595515

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1244918406,"value":1654595515}]`)
	})
}

