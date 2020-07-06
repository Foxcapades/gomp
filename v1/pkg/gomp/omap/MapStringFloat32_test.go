package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringFloat32_Put(t *testing.T) {
	Convey("TestMapStringFloat32.Put", t, func() {
		var k string = "ac12dd32-b795-4c3a-9675-347b887119bf"
		var v float32 = 0.914

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat32_Delete(t *testing.T) {
	Convey("TestMapStringFloat32.Delete", t, func() {
		var k string = "69f01a36-5ac8-412c-80e3-51c61d172b71"
		var v float32 = 0.604

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat32_Has(t *testing.T) {
	Convey("TestMapStringFloat32.Has", t, func() {
		var k string = "9818c50c-926b-41f4-9721-036c46aab5b7"
		var v float32 = 0.360

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("ed957537-5a83-4c6e-94ae-29e2c0c21e5a"+"4bc7beef-4aaa-4f9f-8c9e-e1e91552c7d4"), ShouldBeFalse)
	})
}

func TestMapStringFloat32_Get(t *testing.T) {
	Convey("TestMapStringFloat32.Get", t, func() {
		var k string = "1374e89b-d915-413f-8773-cdb56336acc9"
		var v float32 = 0.046

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("bf04ea04-06cf-4b6d-a377-4afe29321816" + "798c7d60-50c5-4e43-92fa-5b645fe7203d")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat32_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat32.GetOpt", t, func() {
		var k string = "832a3bea-53ce-4765-9b7a-3b82dff4ac47"
		var v float32 = 0.308

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("52b35189-611d-4e51-b1ff-a873bf10922d" + "8d4ef76d-8dee-4923-aac2-c272f887b9fd")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat32_ForEach(t *testing.T) {
	Convey("TestMapStringFloat32.ForEach", t, func() {
		var k string = "a37f159e-5eac-40b5-94bf-780a585db8a2"
		var v float32 = 0.165
		hits := 0

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringFloat32.MarshalYAML", t, func() {
		var k string = "a50bec07-f586-4e82-906f-6531824c7f71"
		var v float32 = 0.784

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringFloat32_ToYAML(t *testing.T) {
	Convey("TestMapStringFloat32.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k string = "432df241-0353-4d51-8160-4584d803636d"
			var v float32 = 0.909

			test := omap.NewMapStringFloat32(1)

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
			var k string = "128f3f53-06c4-41d1-8347-1541e7c8514b"
			var v float32 = 0.686

			test := omap.NewMapStringFloat32(1)
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

func TestMapStringFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringFloat32.PutIfNotNil", t, func() {
		var k string = "4c1dd20a-088c-4785-9bd5-2191ae2dee68"
		var v float32 = 0.970

		test := omap.NewMapStringFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("0bbb19a4-351a-4c23-92e9-29f1a6d04eab", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.161
		So(test.PutIfNotNil("15ee2554-00de-4c7d-8643-85c31ef463fa", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceIfExists", t, func() {
		var k string = "420d8ab0-d9a3-4633-b670-16fb7a706699"
		var v float32 = 0.251
		var x float32 = 0.530

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("f0affe5c-4d74-4d67-af54-93ab778a8203", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceOrPut", t, func() {
		var k string = "4371fbd0-b39a-4dde-ba97-7a50ac67a27c"
		var v float32 = 0.876
		var x float32 = 0.307

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("8d3e853b-a546-402d-a151-2887ddf8fe1a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat32.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "c2e26933-3340-479e-9b20-b4e1a4ad0dab"
			var v float32 = 0.843

			test := omap.NewMapStringFloat32(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"c2e26933-3340-479e-9b20-b4e1a4ad0dab","value":0.843}]`)
		})

		Convey("Unordered", func() {
			var k string = "c2e26933-3340-479e-9b20-b4e1a4ad0dab"
			var v float32 = 0.843

			test := omap.NewMapStringFloat32(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"c2e26933-3340-479e-9b20-b4e1a4ad0dab":0.843}`)
		})

	})
}
