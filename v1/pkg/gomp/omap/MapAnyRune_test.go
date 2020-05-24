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
		var k interface{} = "03309e74-366e-4b74-9096-c64a600854b6"
		var v rune = 1447606903

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyRune_Delete(t *testing.T) {
	Convey("TestMapAnyRune.Delete", t, func() {
		var k interface{} = "e441a608-ce73-4fdc-aba8-97dab0991a12"
		var v rune = 752491631

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyRune_Has(t *testing.T) {
	Convey("TestMapAnyRune.Has", t, func() {
		var k interface{} = "cc52ef25-905e-4562-a1a0-f6779c5e7f52"
		var v rune = 209339970

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("0b608b68-2cf9-4154-8779-40b9759c83ee"+"2eba4a5c-0891-4504-8172-9faba319dda3"), ShouldBeFalse)
	})
}


func TestMapAnyRune_Get(t *testing.T) {
	Convey("TestMapAnyRune.Get", t, func() {
		var k interface{} = "c315aac5-5f45-429e-a377-74ec507c38bf"
		var v rune = 195669610

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("392e6d2f-36d2-4f88-a2a1-7843201590c6" + "c836e78f-e7a5-4ae2-82e6-fb433ae31e44")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyRune_GetOpt(t *testing.T) {
	Convey("TestMapAnyRune.GetOpt", t, func() {
		var k interface{} = "73dad752-91fd-4919-b3ac-f502e80060c4"
		var v rune = 149424881

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("e9a1d3c2-8ca4-4cc4-bb27-496e2523bfd6" + "94651ed8-6722-4157-a167-1e45e91454ad")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyRune_ForEach(t *testing.T) {
	Convey("TestMapAnyRune.ForEach", t, func() {
		var k interface{} = "d4682605-8fda-4fee-8baa-d392a51e95d0"
		var v rune = 1681108911
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
		var k interface{} = "d04b9b79-f792-410f-b512-c6dcbe21a006"
		var v rune = 1291526192

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
		var k interface{} = "e41a8484-e5c5-44c2-a9c8-077d172196e2"
		var v rune = 469867980

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
		var k interface{} = "9454c079-8af2-460c-9b54-6af5d5838e6c"
		var v rune = 1631673899

		test := omap.NewMapAnyRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("4f52bef3-c99e-4f97-8a03-acf8ca003b58", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 35791402
		So(test.PutIfNotNil("851e11ce-ac54-4009-a5f2-6a659b6732a3", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceIfExists", t, func() {
		var k interface{} = "40e2b1fe-646d-4048-b4a4-0a18eacaec71"
		var v rune = 193003129
		var x rune = 315951715

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("4be5e7fa-7225-4c27-a26f-6377e94f8c3e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceOrPut", t, func() {
		var k interface{} = "6cdd13a7-45ef-4414-a8cf-23e10031ce86"
		var v rune = 121805308
		var x rune = 1545716625

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("3dbffc8b-9511-4aad-bce4-a1d055bcb592", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyRune.MarshalJSON", t, func() {
		var k interface{} = "42ad2ff8-4143-4c56-bdd7-23674bd6df66"
		var v rune = 2080297290

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"42ad2ff8-4143-4c56-bdd7-23674bd6df66","value":2080297290}]`)
	})
}

