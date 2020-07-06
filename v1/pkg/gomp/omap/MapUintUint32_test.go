package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint32_Put(t *testing.T) {
	Convey("TestMapUintUint32.Put", t, func() {
		var k uint = 2137989576
		var v uint32 = 706136360

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint32_Delete(t *testing.T) {
	Convey("TestMapUintUint32.Delete", t, func() {
		var k uint = 3937505745
		var v uint32 = 1573144540

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint32_Has(t *testing.T) {
	Convey("TestMapUintUint32.Has", t, func() {
		var k uint = 2057497264
		var v uint32 = 3662958275

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(927688153+2914696213), ShouldBeFalse)
	})
}

func TestMapUintUint32_Get(t *testing.T) {
	Convey("TestMapUintUint32.Get", t, func() {
		var k uint = 2047797753
		var v uint32 = 2445715435

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(3787716360 + 1503497259)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint32_GetOpt(t *testing.T) {
	Convey("TestMapUintUint32.GetOpt", t, func() {
		var k uint = 1572213389
		var v uint32 = 4059827579

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1541346297 + 243595859)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint32_ForEach(t *testing.T) {
	Convey("TestMapUintUint32.ForEach", t, func() {
		var k uint = 3213499863
		var v uint32 = 2939669991
		hits := 0

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint32.MarshalYAML", t, func() {
		var k uint = 3949956779
		var v uint32 = 3300841800

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint32_ToYAML(t *testing.T) {
	Convey("TestMapUintUint32.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k uint = 371040210
			var v uint32 = 1310752095

			test := omap.NewMapUintUint32(1)

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
			var k uint = 1625345479
			var v uint32 = 2880919341

			test := omap.NewMapUintUint32(1)
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

func TestMapUintUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint32.PutIfNotNil", t, func() {
		var k uint = 3922268660
		var v uint32 = 1823148853

		test := omap.NewMapUintUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2767001486, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 310385755
		So(test.PutIfNotNil(3011423973, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceIfExists", t, func() {
		var k uint = 1965584818
		var v uint32 = 3642911920
		var x uint32 = 1028150908

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1014090087, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceOrPut", t, func() {
		var k uint = 1414746595
		var v uint32 = 495019393
		var x uint32 = 1259520646

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(343773088, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint32.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k uint = 4284914301
			var v uint32 = 3277211403

			test := omap.NewMapUintUint32(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":4284914301,"value":3277211403}]`)
		})

		Convey("Unordered", func() {
			var k uint = 4284914301
			var v uint32 = 3277211403

			test := omap.NewMapUintUint32(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"4284914301":3277211403}`)
		})

	})
}
