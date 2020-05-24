package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt_Put(t *testing.T) {
	Convey("TestMapIntInt.Put", t, func() {
		var k int = 1303083006
		var v int = 799558245

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt_Delete(t *testing.T) {
	Convey("TestMapIntInt.Delete", t, func() {
		var k int = 15392885
		var v int = 1516391847

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt_Has(t *testing.T) {
	Convey("TestMapIntInt.Has", t, func() {
		var k int = 1364215939
		var v int = 881807891

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(181365852+426391132), ShouldBeFalse)
	})
}

func TestMapIntInt_Get(t *testing.T) {
	Convey("TestMapIntInt.Get", t, func() {
		var k int = 1379369118
		var v int = 501727366

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1231148209 + 2038426189)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt_GetOpt(t *testing.T) {
	Convey("TestMapIntInt.GetOpt", t, func() {
		var k int = 1837636162
		var v int = 1259938110

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1685247440 + 1340273630)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt_ForEach(t *testing.T) {
	Convey("TestMapIntInt.ForEach", t, func() {
		var k int = 1610601153
		var v int = 69537724
		hits := 0

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt.MarshalYAML", t, func() {
		var k int = 235935890
		var v int = 1799971611

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt_ToYAML(t *testing.T) {
	Convey("TestMapIntInt.ToYAML", t, func() {
		var k int = 7797099
		var v int = 2001552468

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt.PutIfNotNil", t, func() {
		var k int = 944806653
		var v int = 1810036129

		test := omap.NewMapIntInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(130033308, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 894121312
		So(test.PutIfNotNil(1393028251, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt.ReplaceIfExists", t, func() {
		var k int = 1908465878
		var v int = 1379156267
		var x int = 1746346437

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(77381744, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt.ReplaceOrPut", t, func() {
		var k int = 572590760
		var v int = 721136590
		var x int = 1688869115

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(974219765, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt.MarshalJSON", t, func() {
		var k int = 438903309
		var v int = 1035268510

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":438903309,"value":1035268510}]`)
	})
}
