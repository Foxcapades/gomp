package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint32_Put(t *testing.T) {
	Convey("TestMapStringUint32.Put", t, func() {
		var k string = "a68e536e-b725-48cc-8760-acee167f6f30"
		var v uint32 = 560984230

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint32_Delete(t *testing.T) {
	Convey("TestMapStringUint32.Delete", t, func() {
		var k string = "ad4d1e37-a4a1-4d8f-a289-c1b8b5d2ad7d"
		var v uint32 = 2250125341

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint32_Has(t *testing.T) {
	Convey("TestMapStringUint32.Has", t, func() {
		var k string = "7df44e3e-fe24-4d66-8967-e91259184288"
		var v uint32 = 1663271769

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("777485bb-9db4-4322-bea1-d49a72f44367"+"5577dd9d-def3-403b-88f3-4f2c8e6f8f24"), ShouldBeFalse)
	})
}

func TestMapStringUint32_Get(t *testing.T) {
	Convey("TestMapStringUint32.Get", t, func() {
		var k string = "61d53884-b574-4b5c-9efa-082121753b4a"
		var v uint32 = 1800937238

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("a626720d-f875-42b9-8b75-9df6bb8a1f48" + "c03cabc5-968a-49b1-bdf5-c2bcd03a96bf")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint32_GetOpt(t *testing.T) {
	Convey("TestMapStringUint32.GetOpt", t, func() {
		var k string = "395dda0e-e02c-4d80-8050-14bdc9cd676c"
		var v uint32 = 1597431445

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5a5e6048-afc6-4dc3-8aa0-bc5f3ebcec70" + "ca535a2e-b2b7-484c-9300-3bdabdb997a2")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint32_ForEach(t *testing.T) {
	Convey("TestMapStringUint32.ForEach", t, func() {
		var k string = "6c5e8152-abbe-4020-b600-8e6f96f5f5f7"
		var v uint32 = 1509917945
		hits := 0

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint32.MarshalYAML", t, func() {
		var k string = "66d51acf-f2e1-4a77-84b3-293ee61c7441"
		var v uint32 = 3540818714

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint32_ToYAML(t *testing.T) {
	Convey("TestMapStringUint32.ToYAML", t, func() {
		var k string = "38ef381e-2b0a-4905-a9bd-72352a8a94f4"
		var v uint32 = 131154239

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapStringUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint32.PutIfNotNil", t, func() {
		var k string = "931712c3-c372-4cf5-a5ba-483ac18afc9b"
		var v uint32 = 2053256572

		test := omap.NewMapStringUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("c22c65ba-bf8b-4da2-b43a-04a9c41206ee", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 3382747576
		So(test.PutIfNotNil("e94dfb5f-58e8-41f6-80ba-ec7abc521d91", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceIfExists", t, func() {
		var k string = "40f1f3b8-7eb8-40a6-bdf7-af2a182e37ae"
		var v uint32 = 2817194005
		var x uint32 = 4292012428

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1c0c708e-fb49-4bbd-a5ed-49ee1a6fbdfd", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceOrPut", t, func() {
		var k string = "83fd89e0-81cf-4e30-b60d-86e8cd972e00"
		var v uint32 = 2051482348
		var x uint32 = 3975976088

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("6b59f10f-e9f3-44ee-a78f-3345ff6dd49c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint32.MarshalJSON", t, func() {
		var k string = "bbcf8395-6d94-413a-b4e5-5e262e834854"
		var v uint32 = 2277409137

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"bbcf8395-6d94-413a-b4e5-5e262e834854","value":2277409137}]`)
	})
}
