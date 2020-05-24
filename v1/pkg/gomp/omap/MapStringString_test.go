package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringString_Put(t *testing.T) {
	Convey("TestMapStringString.Put", t, func() {
		var k string = "5129233e-8933-4e49-bafe-bb96b446bd56"
		var v string = "09e4ea7d-95e0-4bb9-8b6e-69e16ec469ee"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringString_Delete(t *testing.T) {
	Convey("TestMapStringString.Delete", t, func() {
		var k string = "2f8b1d6b-91be-4c3e-ab70-a70bea63844a"
		var v string = "a4a20a39-6b91-4037-9627-c27695b161b5"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringString_Has(t *testing.T) {
	Convey("TestMapStringString.Has", t, func() {
		var k string = "67af9b2d-fa80-45c1-b934-a814b87144b8"
		var v string = "45120b7c-4f08-45cd-94a8-cabea71c0e17"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3ac5a1f7-3f6f-4e75-a556-614d3febb3a7"+"f45534b9-b631-4a11-a52c-d88529d629a6"), ShouldBeFalse)
	})
}

func TestMapStringString_Get(t *testing.T) {
	Convey("TestMapStringString.Get", t, func() {
		var k string = "f4ba6365-996f-473e-9069-a61d25ddc579"
		var v string = "89770807-974e-4fe3-9474-bd2cc9655436"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("1ade343f-1a03-45d5-8221-25273efcf3e7" + "fc232e94-8830-4357-84a7-21a28e2faca2")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringString_GetOpt(t *testing.T) {
	Convey("TestMapStringString.GetOpt", t, func() {
		var k string = "e5aee69c-afee-4ab8-a280-87210a776f7e"
		var v string = "1a634486-814b-4716-9764-dafc5b622719"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("97094c94-efeb-4cce-a631-b4449462f2b8" + "5d859b8e-4ae5-4525-be8a-9f15ea069d24")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringString_ForEach(t *testing.T) {
	Convey("TestMapStringString.ForEach", t, func() {
		var k string = "6517800b-17a6-4ab9-8333-fe73ace399cc"
		var v string = "e7230e34-6875-4abf-9e39-f8fd72c6e4fc"
		hits := 0

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringString_MarshalYAML(t *testing.T) {
	Convey("TestMapStringString.MarshalYAML", t, func() {
		var k string = "85665903-c15c-43b0-b01a-51a9c97995e7"
		var v string = "3a961ac4-8db6-4b12-a63b-94ffac27b977"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringString_ToYAML(t *testing.T) {
	Convey("TestMapStringString.ToYAML", t, func() {
		var k string = "75be2637-a218-4f5e-a2ed-37faaf5bd8a9"
		var v string = "0183dd73-fdca-4220-9586-fd6929b303f4"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringString_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringString.PutIfNotNil", t, func() {
		var k string = "a8018f5b-5762-4623-ad4b-996a1f573595"
		var v string = "d39f13d4-20af-4b2f-8a16-c1e6d75eea57"

		test := omap.NewMapStringString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("d3e07a41-d2aa-4a4c-b8e5-99a4ce15cc28", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "ae5e9590-43b5-40b9-8bc4-718ffabe07f6"
		So(test.PutIfNotNil("fc341e7a-f062-43a5-90f9-83eae4cd787f", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringString.ReplaceIfExists", t, func() {
		var k string = "9ef79ea9-a975-4f80-8d14-d94ae5d4e56d"
		var v string = "055089d9-2ed6-4e68-bd92-5620afcf2266"
		var x string = "e76b6505-11f1-4966-bb44-9ca9e0c27e6c"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("f44855d3-b046-4f1a-94c3-416cf51ff2ff", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringString.ReplaceOrPut", t, func() {
		var k string = "1f7ff5b6-ba87-4cb8-92ec-5abe87cfc9db"
		var v string = "84763348-ffe4-401f-a19b-e3f181129848"
		var x string = "0c9ab61c-e99d-4ce5-8a33-6ff45a84b182"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("5cf786eb-156f-4bae-8a98-29b874648793", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_MarshalJSON(t *testing.T) {
	Convey("TestMapStringString.MarshalJSON", t, func() {
		var k string = "8f5a49b4-876d-4593-9762-687fcd74b783"
		var v string = "d2e5b3fb-bdec-41fb-98b4-4ecea80898ee"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"8f5a49b4-876d-4593-9762-687fcd74b783","value":"d2e5b3fb-bdec-41fb-98b4-4ecea80898ee"}]`)
	})
}
