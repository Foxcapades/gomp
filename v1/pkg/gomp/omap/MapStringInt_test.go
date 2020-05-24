package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt_Put(t *testing.T) {
	Convey("TestMapStringInt.Put", t, func() {
		var k string = "2caca0a5-7b2c-4a9d-a567-fd4a8c537f3b"
		var v int = 1560649528

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt_Delete(t *testing.T) {
	Convey("TestMapStringInt.Delete", t, func() {
		var k string = "fa5ba860-b557-4a1a-953e-15527bc876d3"
		var v int = 2087184648

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt_Has(t *testing.T) {
	Convey("TestMapStringInt.Has", t, func() {
		var k string = "29eaeb8f-072d-47cd-8543-55c1d3ace68a"
		var v int = 1042343197

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("f8b5c5c1-4d9b-4cd8-97d2-a5921f3b6db8"+"17ac638f-479b-4ae5-baba-25eae8aa9aa6"), ShouldBeFalse)
	})
}


func TestMapStringInt_Get(t *testing.T) {
	Convey("TestMapStringInt.Get", t, func() {
		var k string = "9d6e9f18-fe98-4b0a-98b4-c613f6a2316d"
		var v int = 1411424126

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("a7405cc7-2cdf-4e66-bb72-e5bc6d3ec4f9" + "9d3cffbd-cd4a-402b-8d9c-f915bf7854b9")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt_GetOpt(t *testing.T) {
	Convey("TestMapStringInt.GetOpt", t, func() {
		var k string = "4c00da56-d373-4bde-9554-e111cc37da9f"
		var v int = 1865623677

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("90543cf4-deed-408b-ac63-c93253039d4d" + "b7c9c20e-4c46-4082-aa0e-097af65e84f9")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt_ForEach(t *testing.T) {
	Convey("TestMapStringInt.ForEach", t, func() {
		var k string = "33cc0553-8745-4f07-b631-ec638bc5140c"
		var v int = 1993427197
		hits := 0

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt.MarshalYAML", t, func() {
		var k string = "1e657f73-8190-4744-ae89-edcef916159e"
		var v int = 651158321

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt_ToYAML(t *testing.T) {
	Convey("TestMapStringInt.ToYAML", t, func() {
		var k string = "0481b8d0-6fa8-49f2-b1d5-c3c44f47b5b3"
		var v int = 2111181750

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt.PutIfNotNil", t, func() {
		var k string = "b9fde339-18a9-4f25-b730-802c334c1a01"
		var v int = 1826707586

		test := omap.NewMapStringInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("911d8053-7924-479d-9e75-0426761d45fc", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 617198390
		So(test.PutIfNotNil("91cc3fda-3861-4214-9295-a90a430bfbb8", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt.ReplaceIfExists", t, func() {
		var k string = "2572472f-0ba3-42d4-8d36-409f49a2c2ef"
		var v int = 1450752501
		var x int = 2057067701

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("586d5cc5-29b7-4f6a-8d0b-0ca5d531d528", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt.ReplaceOrPut", t, func() {
		var k string = "f1bc53e9-7d09-40d3-ae69-2056ba6fe3fa"
		var v int = 802423073
		var x int = 1672268189

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("c06228b0-122c-4b04-bb27-d348ac2f8fed", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt.MarshalJSON", t, func() {
		var k string = "f9309ee5-063b-45af-b411-9a040d1dd973"
		var v int = 2005024524

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"f9309ee5-063b-45af-b411-9a040d1dd973","value":2005024524}]`)
	})
}
