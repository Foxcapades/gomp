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
		var k string = "754a9807-547f-4827-9b3b-38221e111e8d"
		var v rune = 1231150869

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringRune_Delete(t *testing.T) {
	Convey("TestMapStringRune.Delete", t, func() {
		var k string = "80dadf21-9ded-4a47-b3a9-7b194a0098d0"
		var v rune = 330942844

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringRune_Has(t *testing.T) {
	Convey("TestMapStringRune.Has", t, func() {
		var k string = "1a0e9c32-8319-41dd-ba30-789368c7cba0"
		var v rune = 1888683888

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("82d47027-5271-420c-a966-a507072b9742"+"49f3abd9-5bbe-4443-b4e5-8444413f38ec"), ShouldBeFalse)
	})
}

func TestMapStringRune_Get(t *testing.T) {
	Convey("TestMapStringRune.Get", t, func() {
		var k string = "538e53f1-633b-46a8-926f-386f0a577743"
		var v rune = 691355755

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("0b96ca22-2a85-42e2-8192-53fed1b3303d" + "3153adfc-3e85-4b63-8860-7c77f76f7e5a")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringRune_GetOpt(t *testing.T) {
	Convey("TestMapStringRune.GetOpt", t, func() {
		var k string = "75c5f22f-bad0-4a2f-bdc7-84538ad346dc"
		var v rune = 1761901064

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("10d55c40-1571-42f7-b678-a2fef309452c" + "607be883-250b-416d-acd8-eb3384abcad3")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringRune_ForEach(t *testing.T) {
	Convey("TestMapStringRune.ForEach", t, func() {
		var k string = "c78cb20f-dc0a-49b0-8f12-4d024d269209"
		var v rune = 220388893
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
		var k string = "1657d3a0-058c-4531-a0fb-f745b79551de"
		var v rune = 604882710

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
		var k string = "58cc1fd2-cb8f-42f6-b1d8-c8ff173e7396"
		var v rune = 227898411

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
		var k string = "44f9b57a-72e3-4110-a1d3-865593b40657"
		var v rune = 1389317323

		test := omap.NewMapStringRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("cf00d00f-95fb-499f-905f-7a0c12dd9ab6", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1219893338
		So(test.PutIfNotNil("dbecb85d-e196-46c4-aa96-6cdaae60240f", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringRune.ReplaceIfExists", t, func() {
		var k string = "ef56ce3d-7b12-445d-920c-0a8e4f843bdf"
		var v rune = 956024248
		var x rune = 960927477

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("bab4890b-a0a2-4baf-804c-0bd429370089", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringRune.ReplaceOrPut", t, func() {
		var k string = "163b7963-bf01-4d9e-a544-9b4df0564885"
		var v rune = 1817635889
		var x rune = 1093226037

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("f35dd645-83aa-4509-a341-eedaac5cc3db", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_MarshalJSON(t *testing.T) {
	Convey("TestMapStringRune.MarshalJSON", t, func() {
		var k string = "f845313b-4511-4686-9d87-cf354a68f722"
		var v rune = 710687178

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"f845313b-4511-4686-9d87-cf354a68f722","value":710687178}]`)
	})
}
