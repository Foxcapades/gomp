package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntAny_Put(t *testing.T) {
	Convey("TestMapIntAny.Put", t, func() {
		var k int = 1583556300
		var v interface{} = "04069a67-7329-4892-927f-cc0e969c756d"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntAny_Delete(t *testing.T) {
	Convey("TestMapIntAny.Delete", t, func() {
		var k int = 1011555688
		var v interface{} = "44fb2b27-f858-4779-a112-9f731bf1c019"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntAny_Has(t *testing.T) {
	Convey("TestMapIntAny.Has", t, func() {
		var k int = 2082002977
		var v interface{} = "111d0454-bc34-4601-aa65-a4bcfb787696"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1982178578+104690118), ShouldBeFalse)
	})
}


func TestMapIntAny_Get(t *testing.T) {
	Convey("TestMapIntAny.Get", t, func() {
		var k int = 1403860244
		var v interface{} = "5884ba6a-4331-4bac-8c7b-2fabf7430b0c"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(374362679+611928803)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntAny_GetOpt(t *testing.T) {
	Convey("TestMapIntAny.GetOpt", t, func() {
		var k int = 1733150964
		var v interface{} = "06b20ccf-6a5e-463e-976d-9ee340303c0c"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(39151355+1256068204)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntAny_ForEach(t *testing.T) {
	Convey("TestMapIntAny.ForEach", t, func() {
		var k int = 1871360067
		var v interface{} = "4810147d-b992-40fc-8127-714274f56db7"
		hits := 0

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntAny_MarshalYAML(t *testing.T) {
	Convey("TestMapIntAny.MarshalYAML", t, func() {
		var k int = 497430063
		var v interface{} = "0a0f6e36-6e6c-4cc3-ba02-c6996c5a30c7"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntAny_ToYAML(t *testing.T) {
	Convey("TestMapIntAny.ToYAML", t, func() {
		var k int = 214241343
		var v interface{} = "9475f41e-fcdd-4907-a92b-3e146167e123"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntAny.PutIfNotNil", t, func() {
		var k int = 1193092063
		var v interface{} = "30b8ba5e-a46b-48a9-9e2a-3a212e83453b"

		test := omap.NewMapIntAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1567263993, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "0f2047f7-15b3-4171-bb05-2260f8be128e"
		So(test.PutIfNotNil(1640002911, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntAny.ReplaceIfExists", t, func() {
		var k int = 1601829606
		var v interface{} = "1bb3e344-ba02-4481-b0b1-e364d0bc389b"
		var x interface{} = "9d3ed3ce-8d2a-4b6b-880f-2060730bda9a"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(386822929, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntAny.ReplaceOrPut", t, func() {
		var k int = 1418143578
		var v interface{} = "04de29af-858b-4193-a431-159f159ff466"
		var x interface{} = "2d21a87d-e623-4086-a979-f8b4504e8961"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1282111499, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_MarshalJSON(t *testing.T) {
	Convey("TestMapIntAny.MarshalJSON", t, func() {
		var k int = 1241953641
		var v interface{} = "d73fb4d9-26e9-4096-b865-c638ee20e0e2"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1241953641,"value":"d73fb4d9-26e9-4096-b865-c638ee20e0e2"}]`)
	})
}

