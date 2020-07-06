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
		var k string = "0786e1d9-4ec2-4f75-b71d-e124f07de1a7"
		var v uint32 = 1871497568

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint32_Delete(t *testing.T) {
	Convey("TestMapStringUint32.Delete", t, func() {
		var k string = "57ec095d-ad1b-4a21-b07b-4c5383a0368d"
		var v uint32 = 453402708

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint32_Has(t *testing.T) {
	Convey("TestMapStringUint32.Has", t, func() {
		var k string = "53a30de3-2eed-4f3b-bd11-5218d27918b1"
		var v uint32 = 3303377348

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3b4299f1-7c9a-42c3-9981-779d83dba18c"+"414533c1-7d59-42a4-b895-3735a848877e"), ShouldBeFalse)
	})
}

func TestMapStringUint32_Get(t *testing.T) {
	Convey("TestMapStringUint32.Get", t, func() {
		var k string = "7a908a2c-58fa-427d-b7d3-b17352452d5c"
		var v uint32 = 754520866

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("fb095cde-7c64-4a70-8fe0-37e30e1a624a" + "59167c99-b973-44c4-b46d-34e5a20eb233")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint32_GetOpt(t *testing.T) {
	Convey("TestMapStringUint32.GetOpt", t, func() {
		var k string = "299b863e-4b44-4289-a3a4-c858893f2bad"
		var v uint32 = 2955010049

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("21e0cd12-9fb9-452b-b8e7-d5a10e096c3d" + "37485e85-a37e-42ad-bf82-fee13fd4b6cd")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint32_ForEach(t *testing.T) {
	Convey("TestMapStringUint32.ForEach", t, func() {
		var k string = "374909d9-ba11-4329-85a2-14798ec38600"
		var v uint32 = 3335820914
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
		var k string = "43d4740c-1c71-4984-afa3-f340894330b8"
		var v uint32 = 2503090174

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
		Convey("Ordered", func() {
			var k string = "e75f6581-4787-497d-bbae-55e8e3f8b6f0"
			var v uint32 = 2752285485

			test := omap.NewMapStringUint32(1)

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
			var k string = "a184ace9-85d4-47fb-b17e-172b24d7be5a"
			var v uint32 = 3249515147

			test := omap.NewMapStringUint32(1)
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

func TestMapStringUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint32.PutIfNotNil", t, func() {
		var k string = "994f5c20-c450-4ebf-988a-9fd2a035a813"
		var v uint32 = 3322998047

		test := omap.NewMapStringUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("67bee346-5316-454f-84fc-30e47934b4c2", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 2194166631
		So(test.PutIfNotNil("87c98845-5fa6-4473-88b6-400b53424779", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceIfExists", t, func() {
		var k string = "cbc3b50a-54e3-495f-8aad-a13c7bbc19b6"
		var v uint32 = 616679701
		var x uint32 = 3751486528

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("f7ff1915-9913-40dc-8670-16a50d7b1dd0", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceOrPut", t, func() {
		var k string = "7299bd38-f33e-4dd0-9e6a-a915aa6c1742"
		var v uint32 = 293766289
		var x uint32 = 750151446

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("28a56755-bbd4-4cd6-9ff4-3bfd42e41f57", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint32.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "3ed38dc0-ab8b-4c15-bede-3946d4e262f9"
			var v uint32 = 351142588

			test := omap.NewMapStringUint32(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"3ed38dc0-ab8b-4c15-bede-3946d4e262f9","value":351142588}]`)
		})

		Convey("Unordered", func() {
			var k string = "3ed38dc0-ab8b-4c15-bede-3946d4e262f9"
			var v uint32 = 351142588

			test := omap.NewMapStringUint32(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"3ed38dc0-ab8b-4c15-bede-3946d4e262f9":351142588}`)
		})

	})
}
