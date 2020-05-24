package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt64_Put(t *testing.T) {
	Convey("TestMapIntInt64.Put", t, func() {
		var k int = 1671247339
		var v int64 = 5693472976477192771

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt64_Delete(t *testing.T) {
	Convey("TestMapIntInt64.Delete", t, func() {
		var k int = 1775506344
		var v int64 = 3899115990894810736

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt64_Has(t *testing.T) {
	Convey("TestMapIntInt64.Has", t, func() {
		var k int = 724506307
		var v int64 = 6485209121251283052

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1165093820+1269559807), ShouldBeFalse)
	})
}

func TestMapIntInt64_Get(t *testing.T) {
	Convey("TestMapIntInt64.Get", t, func() {
		var k int = 115131361
		var v int64 = 4396966651739443522

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(984933561 + 515245609)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt64_GetOpt(t *testing.T) {
	Convey("TestMapIntInt64.GetOpt", t, func() {
		var k int = 1257178502
		var v int64 = 6337052045107734754

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1882416956 + 18562208)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt64_ForEach(t *testing.T) {
	Convey("TestMapIntInt64.ForEach", t, func() {
		var k int = 404995794
		var v int64 = 8154173868427041222
		hits := 0

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt64.MarshalYAML", t, func() {
		var k int = 1029180308
		var v int64 = 1658127795731027567

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt64_ToYAML(t *testing.T) {
	Convey("TestMapIntInt64.ToYAML", t, func() {
		var k int = 2121364869
		var v int64 = 5162731650792020256

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt64.PutIfNotNil", t, func() {
		var k int = 1130715025
		var v int64 = 1829011504491268310

		test := omap.NewMapIntInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1562478709, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 2945453049288109883
		So(test.PutIfNotNil(1487884114, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceIfExists", t, func() {
		var k int = 1846512761
		var v int64 = 8265385922721510607
		var x int64 = 3218454964697304629

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(981012589, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceOrPut", t, func() {
		var k int = 1471173398
		var v int64 = 6692110918846474628
		var x int64 = 6116305707695133508

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1939863692, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt64.MarshalJSON", t, func() {
		var k int = 1414393598
		var v int64 = 8353910877168150309

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1414393598,"value":8353910877168150309}]`)
	})
}
