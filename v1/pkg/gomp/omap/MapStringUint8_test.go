package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint8_Put(t *testing.T) {
	Convey("TestMapStringUint8.Put", t, func() {
		var k string = "5dda0bfa-c24c-4cff-8e36-d31002a0e0fb"
		var v uint8 = 175

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint8_Delete(t *testing.T) {
	Convey("TestMapStringUint8.Delete", t, func() {
		var k string = "b1f0a517-e141-4298-a2cc-fb5a2c057f91"
		var v uint8 = 82

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint8_Has(t *testing.T) {
	Convey("TestMapStringUint8.Has", t, func() {
		var k string = "b9bb4473-b2b7-462e-b901-f2af4d3a1011"
		var v uint8 = 179

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("8d925050-35d6-4ae1-89af-a66b01388f08"+"cd71e904-68f4-4212-861d-6df4018600a4"), ShouldBeFalse)
	})
}


func TestMapStringUint8_Get(t *testing.T) {
	Convey("TestMapStringUint8.Get", t, func() {
		var k string = "55388174-59ec-42d4-9a54-74ab7117379a"
		var v uint8 = 232

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("d09f4f4a-8108-43d3-84f5-7370ac0a2bc1"+"06d67177-2637-4f95-856e-b89af4434a48")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint8_GetOpt(t *testing.T) {
	Convey("TestMapStringUint8.GetOpt", t, func() {
		var k string = "efea3dd2-4e38-4baf-be72-71e0c474ac00"
		var v uint8 = 43

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("4c1ad172-ebef-41a9-9417-64f1fff7b7e7"+"0d0eda66-602d-43bc-becd-87660a1bdb80")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint8_ForEach(t *testing.T) {
	Convey("TestMapStringUint8.ForEach", t, func() {
		var k string = "106a4cfd-eb76-44a9-811e-f4046a89a098"
		var v uint8 = 138
		hits := 0

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint8.MarshalYAML", t, func() {
		var k string = "dc28e699-5e76-4a50-8eba-a89c65acc25b"
		var v uint8 = 187

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint8_ToYAML(t *testing.T) {
	Convey("TestMapStringUint8.ToYAML", t, func() {
		var k string = "1d07574a-0c7f-4724-a455-8b61844a996c"
		var v uint8 = 209

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint8.PutIfNotNil", t, func() {
		var k string = "c0a38034-af23-4405-bd06-44303058ef3a"
		var v uint8 = 174

		test := omap.NewMapStringUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("d720b276-840c-4978-8a6f-497a4f43e336", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 215
		So(test.PutIfNotNil("efdefe41-6540-4e3c-b4f7-8d4e03ded290", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceIfExists", t, func() {
		var k string = "8bee313f-74a3-4ada-9315-807660e4d041"
		var v uint8 = 229
		var x uint8 = 82

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1c00f18c-af71-4202-b4fe-d86baf0271df", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceOrPut", t, func() {
		var k string = "91910038-9657-48c8-b8c9-d3c0eed2df44"
		var v uint8 = 119
		var x uint8 = 221

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ee2b8a97-6590-4686-98f4-60dafbf0057b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint8.MarshalJSON", t, func() {
		var k string = "d6383128-75bd-4d88-83d9-cf47e5bccbaa"
		var v uint8 = 235

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"d6383128-75bd-4d88-83d9-cf47e5bccbaa","value":235}]`)
	})
}

