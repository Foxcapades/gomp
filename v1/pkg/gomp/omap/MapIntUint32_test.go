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
		var k int = 1369330783
		var v uint32 = 1017174586

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint32_Delete(t *testing.T) {
	Convey("TestMapIntUint32.Delete", t, func() {
		var k int = 1338590339
		var v uint32 = 1058897298

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint32_Has(t *testing.T) {
	Convey("TestMapIntUint32.Has", t, func() {
		var k int = 2073198936
		var v uint32 = 2644101608

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(446851931+673018088), ShouldBeFalse)
	})
}

func TestMapIntUint32_Get(t *testing.T) {
	Convey("TestMapIntUint32.Get", t, func() {
		var k int = 83593392
		var v uint32 = 1072549201

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2084693365 + 1601233438)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint32_GetOpt(t *testing.T) {
	Convey("TestMapIntUint32.GetOpt", t, func() {
		var k int = 1918372732
		var v uint32 = 3192597586

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(16255017 + 84206984)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint32_ForEach(t *testing.T) {
	Convey("TestMapIntUint32.ForEach", t, func() {
		var k int = 775046804
		var v uint32 = 1561772882
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
		var k int = 1412106024
		var v uint32 = 3014663917

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
		var k int = 1757830221
		var v uint32 = 3673419152

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
		var k int = 841549544
		var v uint32 = 3738294731

		test := omap.NewMapIntUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2086717914, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 270839960
		So(test.PutIfNotNil(131127086, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceIfExists", t, func() {
		var k int = 1332626161
		var v uint32 = 126832335
		var x uint32 = 1088344019

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1635185747, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceOrPut", t, func() {
		var k int = 195382815
		var v uint32 = 1721946021
		var x uint32 = 3744827158

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2025456857, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint32.MarshalJSON", t, func() {
		var k int = 1208203564
		var v uint32 = 2916695361

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1208203564,"value":2916695361}]`)
	})
}
