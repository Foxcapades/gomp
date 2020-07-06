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
		var k int = 1943119621
		var v uint32 = 332695841

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint32_Delete(t *testing.T) {
	Convey("TestMapIntUint32.Delete", t, func() {
		var k int = 1572257376
		var v uint32 = 3296862190

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint32_Has(t *testing.T) {
	Convey("TestMapIntUint32.Has", t, func() {
		var k int = 623605465
		var v uint32 = 2509367776

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1285147468+1855903959), ShouldBeFalse)
	})
}

func TestMapIntUint32_Get(t *testing.T) {
	Convey("TestMapIntUint32.Get", t, func() {
		var k int = 101876368
		var v uint32 = 3396582635

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(406784919 + 137455727)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint32_GetOpt(t *testing.T) {
	Convey("TestMapIntUint32.GetOpt", t, func() {
		var k int = 552935224
		var v uint32 = 3942069441

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(165148128 + 1867231613)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint32_ForEach(t *testing.T) {
	Convey("TestMapIntUint32.ForEach", t, func() {
		var k int = 704957811
		var v uint32 = 2608894012
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
		var k int = 1778148667
		var v uint32 = 1324550729

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
		Convey("Ordered", func() {
			var k int = 1311350538
			var v uint32 = 3142762749

			test := omap.NewMapIntUint32(1)

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
			var k int = 2033775114
			var v uint32 = 636677281

			test := omap.NewMapIntUint32(1)
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

func TestMapIntUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint32.PutIfNotNil", t, func() {
		var k int = 949960700
		var v uint32 = 3393050751

		test := omap.NewMapIntUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2003223404, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 2866122979
		So(test.PutIfNotNil(1791997605, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceIfExists", t, func() {
		var k int = 563832060
		var v uint32 = 1914553709
		var x uint32 = 15890815

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1876525859, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceOrPut", t, func() {
		var k int = 1128073274
		var v uint32 = 2122621150
		var x uint32 = 2000776997

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(508438348, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint32.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k int = 682990891
			var v uint32 = 4265932093

			test := omap.NewMapIntUint32(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":682990891,"value":4265932093}]`)
		})

		Convey("Unordered", func() {
			var k int = 682990891
			var v uint32 = 4265932093

			test := omap.NewMapIntUint32(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"682990891":4265932093}`)
		})

	})
}
