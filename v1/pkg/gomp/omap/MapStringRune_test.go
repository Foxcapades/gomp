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
		var k string = "2c0b8f3c-262b-40ee-8e35-a914deb35dc9"
		var v rune = 1151927716

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringRune_Delete(t *testing.T) {
	Convey("TestMapStringRune.Delete", t, func() {
		var k string = "69ee851a-2144-4608-8ab4-c12ca8e2de51"
		var v rune = 1311017689

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringRune_Has(t *testing.T) {
	Convey("TestMapStringRune.Has", t, func() {
		var k string = "b42b10e9-b556-4a5c-b6bd-7c394d2bf63e"
		var v rune = 1591714751

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("de8beaf3-5c4c-4497-92a5-387b3104e56c"+"7362c396-d698-4c69-8f26-c627f9b6b25d"), ShouldBeFalse)
	})
}


func TestMapStringRune_Get(t *testing.T) {
	Convey("TestMapStringRune.Get", t, func() {
		var k string = "13ab5efc-7ce6-431e-b0c0-aeb23e377f92"
		var v rune = 1675254706

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("27f2d3f3-f6b4-4044-824d-7ae4b4da43aa" + "08cfb2f6-fe2a-4ebc-87f6-8b61028fd0c8")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringRune_GetOpt(t *testing.T) {
	Convey("TestMapStringRune.GetOpt", t, func() {
		var k string = "ebbb7461-8a2c-413e-a51f-ad764c2d32d9"
		var v rune = 15199526

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("e185c5fc-557c-4b16-88a6-4bc745fa0829" + "1b64f756-cc80-4489-93f1-0f88906cfdcb")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringRune_ForEach(t *testing.T) {
	Convey("TestMapStringRune.ForEach", t, func() {
		var k string = "bc1626ad-2788-49c4-97bb-5566c2214890"
		var v rune = 891479487
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
		var k string = "c6867310-ae80-47ca-adff-5cb82a0b1dd4"
		var v rune = 1169448512

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
		var k string = "6655bf21-d92f-4e19-9089-77473dab75bd"
		var v rune = 756167563

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
		var k string = "b4a690f4-bcbc-4b36-8ce6-7d6fc6f12bde"
		var v rune = 2137354936

		test := omap.NewMapStringRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("a48bfddc-3f62-48af-a780-a5ba492e6ebc", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1734148692
		So(test.PutIfNotNil("4b1fc80b-c510-4508-9f15-c345acf761ff", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringRune.ReplaceIfExists", t, func() {
		var k string = "6420f9b7-1870-4c7e-b639-592d683753a1"
		var v rune = 1604181834
		var x rune = 1945753449

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("0b979742-6bc7-4cf1-8c64-43200ac2d387", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringRune.ReplaceOrPut", t, func() {
		var k string = "d75df9f3-25f9-4980-838d-fb58c6bd1e00"
		var v rune = 997388694
		var x rune = 352317810

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("94aac99f-e9a1-4139-aa35-d498f4d2cd3e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_MarshalJSON(t *testing.T) {
	Convey("TestMapStringRune.MarshalJSON", t, func() {
		var k string = "3b879928-bb8b-4f8a-9f18-3215eb2989a1"
		var v rune = 1434240159

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"3b879928-bb8b-4f8a-9f18-3215eb2989a1","value":1434240159}]`)
	})
}

