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
		var k int = 1790385408
		var v uint = 2628747773

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint_Delete(t *testing.T) {
	Convey("TestMapIntUint.Delete", t, func() {
		var k int = 45499999
		var v uint = 1939766156

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint_Has(t *testing.T) {
	Convey("TestMapIntUint.Has", t, func() {
		var k int = 2027766993
		var v uint = 544826431

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(403368039+137857195), ShouldBeFalse)
	})
}

func TestMapIntUint_Get(t *testing.T) {
	Convey("TestMapIntUint.Get", t, func() {
		var k int = 1117769476
		var v uint = 2688850062

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1971465292 + 1408865170)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint_GetOpt(t *testing.T) {
	Convey("TestMapIntUint.GetOpt", t, func() {
		var k int = 981075078
		var v uint = 1675715375

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1147133501 + 961950432)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint_ForEach(t *testing.T) {
	Convey("TestMapIntUint.ForEach", t, func() {
		var k int = 943502390
		var v uint = 3324013479
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
		var k int = 1459349558
		var v uint = 1281852413

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
		var k int = 287894772
		var v uint = 2340591235

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
		var k int = 835034819
		var v uint = 4051042541

		test := omap.NewMapIntUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1244829314, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 3348570161
		So(test.PutIfNotNil(943678558, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint.ReplaceIfExists", t, func() {
		var k int = 1622706712
		var v uint = 3749176992
		var x uint = 1254025000

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1977203224, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint.ReplaceOrPut", t, func() {
		var k int = 66963754
		var v uint = 2735447243
		var x uint = 314457050

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2037198806, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint.MarshalJSON", t, func() {
		var k int = 104578992
		var v uint = 1449703606

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":104578992,"value":1449703606}]`)
	})
}
