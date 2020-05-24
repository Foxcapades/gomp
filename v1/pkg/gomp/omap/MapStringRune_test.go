package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringRune_Put(t *testing.T) {
	Convey("TestMapStringRune.Put", t, func() {
		var k string = "6a87f834-4341-4b83-8424-a5f59ea447ee"
		var v rune = 1894101640

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringRune_Delete(t *testing.T) {
	Convey("TestMapStringRune.Delete", t, func() {
		var k string = "5e796800-27d2-434e-ace2-545998072847"
		var v rune = 1177615822

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringRune_Has(t *testing.T) {
	Convey("TestMapStringRune.Has", t, func() {
		var k string = "d78f169b-3293-4b0d-9b74-8e2d8e85b2df"
		var v rune = 602131304

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("a204250a-b147-40a7-854d-c2028a67d86b"+"acd6a3aa-d2a2-4511-b1dd-b6306be14800"), ShouldBeFalse)
	})
}

func TestMapStringRune_Get(t *testing.T) {
	Convey("TestMapStringRune.Get", t, func() {
		var k string = "3014ddc5-8fbd-4f36-be13-3fa602d351bb"
		var v rune = 767895055

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("e2a11974-2134-438c-bd70-ba84efdd05aa" + "4474e42b-5d5b-4aad-8b88-f23112b8866e")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringRune_GetOpt(t *testing.T) {
	Convey("TestMapStringRune.GetOpt", t, func() {
		var k string = "18614fb3-85f7-42d8-a999-a7fb1f027bc6"
		var v rune = 825306448

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("8ad28034-b9cd-4d1a-a6be-c2eaba958966" + "6305ec4c-1bdc-4d27-8d02-5eb0ddf16e6f")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringRune_ForEach(t *testing.T) {
	Convey("TestMapStringRune.ForEach", t, func() {
		var k string = "0564b4e8-aad1-4ad3-99ad-7a67ab332b68"
		var v rune = 591612333
		hits := 0

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv rune) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringRune_MarshalYAML(t *testing.T) {
	Convey("TestMapStringRune.MarshalYAML", t, func() {
		var k string = "865c409b-7445-4427-a182-c6aacbd44d67"
		var v rune = 751687391

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringRune_ToYAML(t *testing.T) {
	Convey("TestMapStringRune.ToYAML", t, func() {
		var k string = "ad09ab85-73a9-4356-9dc0-93becbf1736c"
		var v rune = 1869025654

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringRune.PutIfNotNil", t, func() {
		var k string = "8f28aa4b-ab49-4baf-90fc-a527723e7bd8"
		var v rune = 1567242035

		test := omap.NewMapStringRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b2f50ab5-1142-4919-8b03-fd55c5139617", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1825788190
		So(test.PutIfNotNil("67555b7b-e4c6-4c10-9c5c-b221c2fc8635", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringRune.ReplaceIfExists", t, func() {
		var k string = "84fa6bad-ec7c-4e55-834d-960aa702ea75"
		var v rune = 659411236
		var x rune = 1740859166

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("34d0040f-fc30-4b3b-9329-dbfcdadee474", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringRune.ReplaceOrPut", t, func() {
		var k string = "9b764dd1-c7f3-451d-a286-adc6f8d5153f"
		var v rune = 765371729
		var x rune = 1695546381

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("f8d52fb1-4ba3-4689-aa44-65e931452280", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_MarshalJSON(t *testing.T) {
	Convey("TestMapStringRune.MarshalJSON", t, func() {
		var k string = "f69aac58-3279-4a48-9426-f298dfa9d000"
		var v rune = 304765991

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"f69aac58-3279-4a48-9426-f298dfa9d000","value":304765991}]`)
	})
}
