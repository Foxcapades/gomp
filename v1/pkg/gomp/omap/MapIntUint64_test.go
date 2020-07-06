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
		var k int = 1621369613
		var v uint64 = 17730699129085194471

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint64_Delete(t *testing.T) {
	Convey("TestMapIntUint64.Delete", t, func() {
		var k int = 196976414
		var v uint64 = 8450519566476863007

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint64_Has(t *testing.T) {
	Convey("TestMapIntUint64.Has", t, func() {
		var k int = 804266981
		var v uint64 = 17354513181334675725

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1561182307+2071194540), ShouldBeFalse)
	})
}

func TestMapIntUint64_Get(t *testing.T) {
	Convey("TestMapIntUint64.Get", t, func() {
		var k int = 1091538663
		var v uint64 = 1544684489577286612

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1097287534 + 1481231596)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint64_GetOpt(t *testing.T) {
	Convey("TestMapIntUint64.GetOpt", t, func() {
		var k int = 1653793047
		var v uint64 = 11764275298180941428

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1653095337 + 1475884934)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint64_ForEach(t *testing.T) {
	Convey("TestMapIntUint64.ForEach", t, func() {
		var k int = 656154938
		var v uint64 = 2488096504678235315
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
		var k int = 1141923261
		var v uint64 = 12037638742972501615

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
		Convey("Ordered", func() {
			var k int = 1280999157
			var v uint64 = 6553775401025037962

			test := omap.NewMapIntUint64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k int = 114785933
			var v uint64 = 3253921881443470167

			test := omap.NewMapIntUint64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapIntUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint64.PutIfNotNil", t, func() {
		var k int = 1553105695
		var v uint64 = 3258171738466690641

		test := omap.NewMapIntUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1555717145, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 3602248266140131597
		So(test.PutIfNotNil(675586210, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceIfExists", t, func() {
		var k int = 2136739808
		var v uint64 = 5455881202225456031
		var x uint64 = 5717930730389443723

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(182889340, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceOrPut", t, func() {
		var k int = 196315624
		var v uint64 = 2716648391068126594
		var x uint64 = 11601865602662508995

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(994895484, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint64.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k int = 999008925
			var v uint64 = 373591153720254934

			test := omap.NewMapIntUint64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":999008925,"value":373591153720254934}]`)
		})

		Convey("Unordered", func() {
			var k int = 999008925
			var v uint64 = 373591153720254934

			test := omap.NewMapIntUint64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"999008925":373591153720254934}`)
		})

	})
}
