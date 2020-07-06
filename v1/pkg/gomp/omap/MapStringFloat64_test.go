package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringFloat64_Put(t *testing.T) {
	Convey("TestMapStringFloat64.Put", t, func() {
		var k string = "b07edb71-1343-4416-82c8-589421b6aeee"
		var v float64 = 0.072

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat64_Delete(t *testing.T) {
	Convey("TestMapStringFloat64.Delete", t, func() {
		var k string = "8f7e83a3-2c09-4410-8314-aeebeb5b876d"
		var v float64 = 0.552

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat64_Has(t *testing.T) {
	Convey("TestMapStringFloat64.Has", t, func() {
		var k string = "13b28187-e1d9-4a09-bbf8-d73b77c49495"
		var v float64 = 0.589

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("08edb19f-05c3-48d2-a8f7-e9a6329ddaf9"+"055b5370-a303-479d-b848-a13c3588cb55"), ShouldBeFalse)
	})
}

func TestMapStringFloat64_Get(t *testing.T) {
	Convey("TestMapStringFloat64.Get", t, func() {
		var k string = "7c433a5e-644f-4fc4-81ab-ce92d5602f69"
		var v float64 = 0.119

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("2bae1a5c-a6ff-4237-81ac-ca9b6b304906" + "d8d25e99-0e46-48dc-82e9-9dc28220a578")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat64_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat64.GetOpt", t, func() {
		var k string = "d2d8d807-b521-469b-bd9a-743a75a6352f"
		var v float64 = 0.245

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("3f26fb8e-6e08-46f1-84cf-4b9fc9d29da4" + "fb29325a-28ac-4ed4-86ca-0d43c777f13f")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat64_ForEach(t *testing.T) {
	Convey("TestMapStringFloat64.ForEach", t, func() {
		var k string = "4056de9f-2820-4d0d-a91f-b8681bdd780c"
		var v float64 = 0.548
		hits := 0

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringFloat64.MarshalYAML", t, func() {
		var k string = "23e5865d-4842-4ad1-9249-6e3af3f9d82a"
		var v float64 = 0.400

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringFloat64_ToYAML(t *testing.T) {
	Convey("TestMapStringFloat64.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k string = "a72d7994-12c3-450e-bbba-b0d54afd98f9"
			var v float64 = 0.070

			test := omap.NewMapStringFloat64(1)

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
			var k string = "24bc6be8-6e94-4b06-bbec-25342d41792e"
			var v float64 = 0.216

			test := omap.NewMapStringFloat64(1)
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

func TestMapStringFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringFloat64.PutIfNotNil", t, func() {
		var k string = "908793c9-ce7d-4c76-bba2-0ac43c5d7201"
		var v float64 = 0.845

		test := omap.NewMapStringFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("d6e045d4-93fe-4fc3-b1d5-c47e24de7fee", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.670
		So(test.PutIfNotNil("b28c588a-c65a-49ee-8a05-fd6f57d2bc44", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceIfExists", t, func() {
		var k string = "20649884-be72-4630-99a5-6a192c174f38"
		var v float64 = 0.182
		var x float64 = 0.944

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("4817669b-baaf-4abd-a4da-2b40a385765c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceOrPut", t, func() {
		var k string = "1fa43288-36a6-4923-9c1a-b2f02df5824f"
		var v float64 = 0.263
		var x float64 = 0.703

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("3bec9123-9559-4095-be8e-c120595b7258", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat64.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "3c7bfcff-5839-45d6-bcb7-433fb67c02a6"
			var v float64 = 0.138

			test := omap.NewMapStringFloat64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"3c7bfcff-5839-45d6-bcb7-433fb67c02a6","value":0.138}]`)
		})

		Convey("Unordered", func() {
			var k string = "3c7bfcff-5839-45d6-bcb7-433fb67c02a6"
			var v float64 = 0.138

			test := omap.NewMapStringFloat64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"3c7bfcff-5839-45d6-bcb7-433fb67c02a6":0.138}`)
		})

	})
}
