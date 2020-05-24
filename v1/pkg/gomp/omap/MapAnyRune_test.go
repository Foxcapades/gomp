package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyRune_Put(t *testing.T) {
	Convey("TestMapAnyRune.Put", t, func() {
		var k interface{} = "231f2163-4d13-4df9-b202-0754f808c861"
		var v rune = 617534466

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyRune_Delete(t *testing.T) {
	Convey("TestMapAnyRune.Delete", t, func() {
		var k interface{} = "b75ab4fd-f737-4139-b9a8-9496691f4a5d"
		var v rune = 1992857840

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyRune_Has(t *testing.T) {
	Convey("TestMapAnyRune.Has", t, func() {
		var k interface{} = "24002334-7f3f-4cc2-b21f-f34f7181e27f"
		var v rune = 2007811917

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("30c0ff29-a78e-496f-b790-4b1368c75f5f"+"7298dd61-1dc9-4fdc-b312-52539feafa0b"), ShouldBeFalse)
	})
}

func TestMapAnyRune_Get(t *testing.T) {
	Convey("TestMapAnyRune.Get", t, func() {
		var k interface{} = "2be263c6-c2d0-4342-8ed3-112aeb97203e"
		var v rune = 1867617829

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("f758ca6e-5f5c-415e-b235-7858639eada6" + "3bb70895-62e4-49a0-acea-ff9798f1d633")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyRune_GetOpt(t *testing.T) {
	Convey("TestMapAnyRune.GetOpt", t, func() {
		var k interface{} = "971336f7-d05c-42bf-8234-02c61f2d1534"
		var v rune = 612490345

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("fccbdd3b-15ed-4a9a-ad86-9c44cdb8a3e8" + "9f279127-9674-491e-a63e-0656dd3a24ae")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyRune_ForEach(t *testing.T) {
	Convey("TestMapAnyRune.ForEach", t, func() {
		var k interface{} = "427ffe6c-602b-4b01-aa5f-cc41c10ab3f6"
		var v rune = 218849497
		hits := 0

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv rune) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyRune_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyRune.MarshalYAML", t, func() {
		var k interface{} = "d92233e4-053a-4a71-b320-3de1b43fcdaf"
		var v rune = 493203177

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyRune_ToYAML(t *testing.T) {
	Convey("TestMapAnyRune.ToYAML", t, func() {
		var k interface{} = "e7089b9d-3348-4eda-9e62-d78693f0f105"
		var v rune = 767087571

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyRune.PutIfNotNil", t, func() {
		var k interface{} = "c3b8e8a1-6686-4501-ad1b-379df3a02a1d"
		var v rune = 1641279179

		test := omap.NewMapAnyRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("83afb151-749d-4924-ac68-b7c8d0648d0b", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 387242072
		So(test.PutIfNotNil("0690fa4a-5214-4571-90bf-5a629d1c411f", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceIfExists", t, func() {
		var k interface{} = "46b8af1f-2033-4be7-ace3-47bd0786d470"
		var v rune = 79174656
		var x rune = 632850166

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("4262209d-2a02-4cab-9246-760865415c66", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceOrPut", t, func() {
		var k interface{} = "2ef4d9b4-f359-4ddd-b1a6-09771c0fa1fb"
		var v rune = 412956273
		var x rune = 941185245

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("424ae489-022f-4c21-af9c-d3074f6f7113", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyRune.MarshalJSON", t, func() {
		var k interface{} = "0148d480-2eae-43ef-b278-0a0594386ff8"
		var v rune = 453539675

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"0148d480-2eae-43ef-b278-0a0594386ff8","value":453539675}]`)
	})
}
