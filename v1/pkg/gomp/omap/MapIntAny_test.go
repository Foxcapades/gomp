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
		var k int = 1389948630
		var v interface{} = "180d352f-e13e-42dd-8ade-ebdc050b84b3"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntAny_Delete(t *testing.T) {
	Convey("TestMapIntAny.Delete", t, func() {
		var k int = 879456123
		var v interface{} = "7c08b05f-8cde-48ba-bfdf-b46ab217c3b9"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntAny_Has(t *testing.T) {
	Convey("TestMapIntAny.Has", t, func() {
		var k int = 857739316
		var v interface{} = "774afeae-bda5-49f0-b91d-31b0bc8cbe84"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2059265518+1198870136), ShouldBeFalse)
	})
}


func TestMapIntAny_Get(t *testing.T) {
	Convey("TestMapIntAny.Get", t, func() {
		var k int = 191345941
		var v interface{} = "0055dd05-5cf8-400c-8863-e0200602ad15"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1544589339 + 309833094)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntAny_GetOpt(t *testing.T) {
	Convey("TestMapIntAny.GetOpt", t, func() {
		var k int = 1001244371
		var v interface{} = "3c74d1fc-d78c-4ef0-8806-d6c59d6c404b"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1311664617 + 2085347018)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntAny_ForEach(t *testing.T) {
	Convey("TestMapIntAny.ForEach", t, func() {
		var k int = 1880509558
		var v interface{} = "ac58b3cc-fe4a-4c07-a876-acf2205d947e"
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
		var k int = 775688172
		var v interface{} = "c2bcd88d-55a2-4b93-a3a1-19b57dea7abf"

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
		var k int = 2140693689
		var v interface{} = "05abae5d-c66e-4d9c-87b2-d7eefe77062b"

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
		var k int = 1831026458
		var v interface{} = "3887d0c5-5631-40e6-9940-47e0665f104e"

		test := omap.NewMapIntAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(685976850, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "65d326a9-2b52-41af-b33e-fe0ce163fe63"
		So(test.PutIfNotNil(1457076915, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntAny.ReplaceIfExists", t, func() {
		var k int = 1238377461
		var v interface{} = "5cf9de64-ced6-4020-a638-8443b9ce7cb3"
		var x interface{} = "785c0c91-856f-4d24-b129-602b21ff6201"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1263792632, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntAny.ReplaceOrPut", t, func() {
		var k int = 1107292675
		var v interface{} = "fe37d89a-c8e5-4ccf-b079-6ab570eec473"
		var x interface{} = "38764bb1-8fc3-4e30-b8ba-2e9c88af29cc"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(944988974, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_MarshalJSON(t *testing.T) {
	Convey("TestMapIntAny.MarshalJSON", t, func() {
		var k int = 1495239384
		var v interface{} = "7c23852b-94b2-47fa-a40f-17c222243e26"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1495239384,"value":"7c23852b-94b2-47fa-a40f-17c222243e26"}]`)
	})
}

