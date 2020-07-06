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
		var k int = 433153065
		var v int64 = 366469222471490280

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt64_Delete(t *testing.T) {
	Convey("TestMapIntInt64.Delete", t, func() {
		var k int = 1406067068
		var v int64 = 5446481527874617251

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt64_Has(t *testing.T) {
	Convey("TestMapIntInt64.Has", t, func() {
		var k int = 2118857310
		var v int64 = 3823719460081585954

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1103713641+1178877299), ShouldBeFalse)
	})
}

func TestMapIntInt64_Get(t *testing.T) {
	Convey("TestMapIntInt64.Get", t, func() {
		var k int = 589579684
		var v int64 = 2136927701175997556

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(448333305 + 487342969)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt64_GetOpt(t *testing.T) {
	Convey("TestMapIntInt64.GetOpt", t, func() {
		var k int = 1463120049
		var v int64 = 8079709456011902152

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1341041668 + 873368936)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt64_ForEach(t *testing.T) {
	Convey("TestMapIntInt64.ForEach", t, func() {
		var k int = 1137705797
		var v int64 = 1902207744791194388
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
		var k int = 752844216
		var v int64 = 8498304471408521299

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
		Convey("Ordered", func() {
			var k int = 313905588
			var v int64 = 8760441771755723011

			test := omap.NewMapIntInt64(1)

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
			var k int = 1054128949
			var v int64 = 8790338813623994937

			test := omap.NewMapIntInt64(1)
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

func TestMapIntInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt64.PutIfNotNil", t, func() {
		var k int = 244205673
		var v int64 = 1434980408929680201

		test := omap.NewMapIntInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1069469417, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 4085752962653300826
		So(test.PutIfNotNil(1203722886, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceIfExists", t, func() {
		var k int = 472286591
		var v int64 = 3718371200677131898
		var x int64 = 2125018609138316871

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(742864008, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceOrPut", t, func() {
		var k int = 585975945
		var v int64 = 1692508824652831923
		var x int64 = 2010943223604047458

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1203102900, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt64.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k int = 339512660
			var v int64 = 3138298292617833240

			test := omap.NewMapIntInt64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":339512660,"value":3138298292617833240}]`)
		})

		Convey("Unordered", func() {
			var k int = 339512660
			var v int64 = 3138298292617833240

			test := omap.NewMapIntInt64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"339512660":3138298292617833240}`)
		})

	})
}
