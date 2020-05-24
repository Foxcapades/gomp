package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint_Put(t *testing.T) {
	Convey("TestMapIntUint.Put", t, func() {
		var k int = 266639867
		var v uint = 2103227403

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint_Delete(t *testing.T) {
	Convey("TestMapIntUint.Delete", t, func() {
		var k int = 1274198078
		var v uint = 2638409453

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint_Has(t *testing.T) {
	Convey("TestMapIntUint.Has", t, func() {
		var k int = 401455393
		var v uint = 2164100149

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1328407469+1655359846), ShouldBeFalse)
	})
}


func TestMapIntUint_Get(t *testing.T) {
	Convey("TestMapIntUint.Get", t, func() {
		var k int = 2139146555
		var v uint = 2893309480

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1236173989 + 872118402)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint_GetOpt(t *testing.T) {
	Convey("TestMapIntUint.GetOpt", t, func() {
		var k int = 2108100626
		var v uint = 3582155426

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(172388060 + 1416492511)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint_ForEach(t *testing.T) {
	Convey("TestMapIntUint.ForEach", t, func() {
		var k int = 1470347120
		var v uint = 3564207246
		hits := 0

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint.MarshalYAML", t, func() {
		var k int = 1175895024
		var v uint = 2335296118

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint_ToYAML(t *testing.T) {
	Convey("TestMapIntUint.ToYAML", t, func() {
		var k int = 352939872
		var v uint = 296814707

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint.PutIfNotNil", t, func() {
		var k int = 1277744162
		var v uint = 1696172023

		test := omap.NewMapIntUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1139095192, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 951091454
		So(test.PutIfNotNil(191621251, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint.ReplaceIfExists", t, func() {
		var k int = 636066250
		var v uint = 1158907516
		var x uint = 1063771003

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1544772184, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint.ReplaceOrPut", t, func() {
		var k int = 972645004
		var v uint = 1814318312
		var x uint = 909974371

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1423105789, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint.MarshalJSON", t, func() {
		var k int = 704941966
		var v uint = 641770967

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":704941966,"value":641770967}]`)
	})
}

