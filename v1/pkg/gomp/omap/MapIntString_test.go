package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntString_Put(t *testing.T) {
	Convey("TestMapIntString.Put", t, func() {
		var k int = 855634083
		var v string = "3e9538d0-e3c0-428b-ad18-41ff2ff261ea"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntString_Delete(t *testing.T) {
	Convey("TestMapIntString.Delete", t, func() {
		var k int = 1938969245
		var v string = "6344b73e-72f7-4488-86dc-2c5b56f287b9"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntString_Has(t *testing.T) {
	Convey("TestMapIntString.Has", t, func() {
		var k int = 788001965
		var v string = "99ae6ea7-6027-4d20-94c7-816389edddb0"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2115966092+1595450292), ShouldBeFalse)
	})
}

func TestMapIntString_Get(t *testing.T) {
	Convey("TestMapIntString.Get", t, func() {
		var k int = 2070048459
		var v string = "a1a1acd1-69f3-4ad4-9f23-7ce895a2d288"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1202935117 + 1274254368)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntString_GetOpt(t *testing.T) {
	Convey("TestMapIntString.GetOpt", t, func() {
		var k int = 1112372233
		var v string = "f05ff57d-5af6-4463-854f-7ce6102f2148"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(766206239 + 1902347842)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntString_ForEach(t *testing.T) {
	Convey("TestMapIntString.ForEach", t, func() {
		var k int = 1022779175
		var v string = "36050f84-a4dc-4bed-b2c7-a0e951203b0b"
		hits := 0

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntString_MarshalYAML(t *testing.T) {
	Convey("TestMapIntString.MarshalYAML", t, func() {
		var k int = 860072376
		var v string = "f7d8228a-d8a7-494c-8d2e-42425d07486c"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntString_ToYAML(t *testing.T) {
	Convey("TestMapIntString.ToYAML", t, func() {
		var k int = 1497606146
		var v string = "0f448b48-aebb-4881-87d4-78597bbf1019"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntString_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntString.PutIfNotNil", t, func() {
		var k int = 1966783090
		var v string = "66beef14-79e6-4f69-89ed-1244e5903168"

		test := omap.NewMapIntString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1075809432, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "46a5d62d-6e97-4ac0-b586-81bb98466cba"
		So(test.PutIfNotNil(1041034517, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntString.ReplaceIfExists", t, func() {
		var k int = 118210768
		var v string = "a618cff3-27db-4362-89e3-231b1fdd362f"
		var x string = "afd81e97-12a9-4462-975e-5e51a9a5f721"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1425953154, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntString.ReplaceOrPut", t, func() {
		var k int = 1387727189
		var v string = "9ee75af2-c42f-46a2-a3b8-d09cb04707ec"
		var x string = "e90e6a95-a413-4357-9bff-8e5fe6182837"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1761560237, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_MarshalJSON(t *testing.T) {
	Convey("TestMapIntString.MarshalJSON", t, func() {
		var k int = 1395699404
		var v string = "718473bc-4a2e-4f80-8d0f-acd1609cd260"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1395699404,"value":"718473bc-4a2e-4f80-8d0f-acd1609cd260"}]`)
	})
}
