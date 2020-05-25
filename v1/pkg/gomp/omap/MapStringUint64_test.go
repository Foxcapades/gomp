package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint64_Put(t *testing.T) {
	Convey("TestMapStringUint64.Put", t, func() {
		var k string = "04f91cb5-9277-4caa-81e6-77f7e6559010"
		var v uint64 = 3911034999264027952

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint64_Delete(t *testing.T) {
	Convey("TestMapStringUint64.Delete", t, func() {
		var k string = "299fd6e6-dd1d-41c9-8799-daabd870bef5"
		var v uint64 = 9288195605908003992

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint64_Has(t *testing.T) {
	Convey("TestMapStringUint64.Has", t, func() {
		var k string = "a699d8c0-07c7-4157-a26a-587e05456818"
		var v uint64 = 7009572161341695372

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("94698cb5-546c-4d73-b232-612f802cba7c"+"6cd58794-906f-49a1-b24c-322d9d51a99d"), ShouldBeFalse)
	})
}

func TestMapStringUint64_Get(t *testing.T) {
	Convey("TestMapStringUint64.Get", t, func() {
		var k string = "ad9d2dc3-2ada-4f74-9f72-5f1ccb990306"
		var v uint64 = 11649460395924522732

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("da915aa6-a1eb-47f9-a13e-0601c68391fc" + "e64c9175-1f56-4306-9ffc-2eee43392018")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint64_GetOpt(t *testing.T) {
	Convey("TestMapStringUint64.GetOpt", t, func() {
		var k string = "7b92d23b-f916-4a9e-9265-e23365fb8381"
		var v uint64 = 11890022778509834146

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("f3daaf87-d1ef-48a0-bcd2-42775a597857" + "4101b3c2-d6a3-437b-b0c1-083f9bcaf2b5")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint64_ForEach(t *testing.T) {
	Convey("TestMapStringUint64.ForEach", t, func() {
		var k string = "4d6ec09a-9230-4c52-a88b-af3207736e9b"
		var v uint64 = 5708910800049734314
		hits := 0

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint64.MarshalYAML", t, func() {
		var k string = "463bc565-a384-407f-9b42-617928f3d172"
		var v uint64 = 5472115075367678277

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint64_ToYAML(t *testing.T) {
	Convey("TestMapStringUint64.ToYAML", t, func() {
		var k string = "fb84ae8d-1c4b-4982-9f34-73bfb923c4dc"
		var v uint64 = 16920361473180088492

		test := omap.NewMapStringUint64(1)

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

func TestMapStringUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint64.PutIfNotNil", t, func() {
		var k string = "fc472b87-75e9-48b9-81fc-161883aeaf13"
		var v uint64 = 4199558752411461125

		test := omap.NewMapStringUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("5625e185-ad9d-4083-aba1-6e77cf359080", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 3606824379071177063
		So(test.PutIfNotNil("1e3a357c-68b3-4643-9867-81fe618d0816", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceIfExists", t, func() {
		var k string = "42424479-fa6b-482a-9e34-827d2368a43f"
		var v uint64 = 8285770198546684964
		var x uint64 = 14285587401751261016

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("0b6a2672-184d-4e02-bde7-fe13f0855d25", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceOrPut", t, func() {
		var k string = "3baf64d1-8b24-4e1f-af4d-1709cd3f3a47"
		var v uint64 = 596948744303767923
		var x uint64 = 5513909365758624632

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("bcd0728b-10db-4be7-bf4a-9bf8f8ef5905", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint64.MarshalJSON", t, func() {
		var k string = "a21c7307-c36b-4e72-a2f8-56d784b368ef"
		var v uint64 = 11166707762328533612

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a21c7307-c36b-4e72-a2f8-56d784b368ef","value":11166707762328533612}]`)
	})
}
