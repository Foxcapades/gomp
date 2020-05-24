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
		var k int = 562692428
		var v uint64 = 15644931789604474125

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint64_Delete(t *testing.T) {
	Convey("TestMapIntUint64.Delete", t, func() {
		var k int = 1786008206
		var v uint64 = 2460007574878616666

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint64_Has(t *testing.T) {
	Convey("TestMapIntUint64.Has", t, func() {
		var k int = 278213765
		var v uint64 = 12906300501163022095

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(94788319+1031320799), ShouldBeFalse)
	})
}


func TestMapIntUint64_Get(t *testing.T) {
	Convey("TestMapIntUint64.Get", t, func() {
		var k int = 60508039
		var v uint64 = 3493290361130998527

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1125929559 + 2053753815)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint64_GetOpt(t *testing.T) {
	Convey("TestMapIntUint64.GetOpt", t, func() {
		var k int = 115453858
		var v uint64 = 18410931598399830223

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(760072859 + 1683736911)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint64_ForEach(t *testing.T) {
	Convey("TestMapIntUint64.ForEach", t, func() {
		var k int = 623362239
		var v uint64 = 2674846641474448153
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
		var k int = 1536213372
		var v uint64 = 4491536912160116937

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
		var k int = 244613959
		var v uint64 = 4695395392627831400

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
		var k int = 166079112
		var v uint64 = 2809632241093857033

		test := omap.NewMapIntUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(308936345, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 5721420325634072076
		So(test.PutIfNotNil(654221909, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceIfExists", t, func() {
		var k int = 2016995386
		var v uint64 = 17066267289317486383
		var x uint64 = 5002848016394407376

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(585752631, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceOrPut", t, func() {
		var k int = 1210744622
		var v uint64 = 15690279029985064700
		var x uint64 = 3800161135833544866

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1458876063, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint64.MarshalJSON", t, func() {
		var k int = 771114079
		var v uint64 = 208738535520370429

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":771114079,"value":208738535520370429}]`)
	})
}
