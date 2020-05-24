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
		var k string = "7b930d4c-6136-4101-98c0-942ee967d54a"
		var v uint32 = 133939227

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint32_Delete(t *testing.T) {
	Convey("TestMapStringUint32.Delete", t, func() {
		var k string = "32c998e9-8f1d-43a5-9bbb-726030147256"
		var v uint32 = 2672965537

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint32_Has(t *testing.T) {
	Convey("TestMapStringUint32.Has", t, func() {
		var k string = "a9f1889f-a8b8-44ed-9a8f-acc178bdfbf7"
		var v uint32 = 137328351

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("86e52f6b-fba3-456a-b025-7c480b182188"+"5a05eb97-0eed-4837-a0d7-44581cf51ddf"), ShouldBeFalse)
	})
}


func TestMapStringUint32_Get(t *testing.T) {
	Convey("TestMapStringUint32.Get", t, func() {
		var k string = "d7d2897b-f996-4059-bda5-c890f5dc0fdc"
		var v uint32 = 187473025

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("cb404119-4c4a-4603-af96-d7268560ea01" + "8de67f30-d670-4d21-827d-926ce18363f5")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint32_GetOpt(t *testing.T) {
	Convey("TestMapStringUint32.GetOpt", t, func() {
		var k string = "3ab3afd5-dc1b-48a5-9d4c-4614a8ae73fc"
		var v uint32 = 1826301408

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("bb1a0bd3-16f1-4350-bfe2-d4adb3f569b6" + "70d118a0-05f4-452f-a94f-fe1d6085c458")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint32_ForEach(t *testing.T) {
	Convey("TestMapStringUint32.ForEach", t, func() {
		var k string = "06261d03-258b-49eb-b31f-71ede6d6daa1"
		var v uint32 = 2712265326
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
		var k string = "2a1600ab-2421-4ad6-9b68-c7beaed44b0a"
		var v uint32 = 1844858749

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
		var k string = "bf3a8729-0bcd-4429-9d7d-8cb62126e8b3"
		var v uint32 = 375515611

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint32.PutIfNotNil", t, func() {
		var k string = "5e2792cb-83c5-4e05-a8fa-82491f24b3ae"
		var v uint32 = 2998757129

		test := omap.NewMapStringUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("7968aa10-acf5-462f-8fd9-adbf4cf298ee", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 3267631577
		So(test.PutIfNotNil("b4c145ca-5404-48d1-878d-4662bcf745ef", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceIfExists", t, func() {
		var k string = "614fcc6c-8393-4f98-897a-81e69dcad6d5"
		var v uint32 = 3706757716
		var x uint32 = 3974075629

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("d4a148e5-a335-432a-a4a9-aead16e33367", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceOrPut", t, func() {
		var k string = "a28781e1-3c8c-4464-912e-90ac1b9b4ce3"
		var v uint32 = 424652564
		var x uint32 = 831535429

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("8458b624-fb85-4e81-97a5-7eb55d253c82", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint32.MarshalJSON", t, func() {
		var k string = "7de8c088-8f77-458f-a173-64730271adc3"
		var v uint32 = 4097208018

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"7de8c088-8f77-458f-a173-64730271adc3","value":4097208018}]`)
	})
}

