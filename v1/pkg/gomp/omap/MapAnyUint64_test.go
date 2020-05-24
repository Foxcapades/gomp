package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint64_Put(t *testing.T) {
	Convey("TestMapAnyUint64.Put", t, func() {
		var k interface{} = "2b7a5a09-9f09-4182-9106-ebf557849eaf"
		var v uint64 = 13734329548015794725

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint64_Delete(t *testing.T) {
	Convey("TestMapAnyUint64.Delete", t, func() {
		var k interface{} = "6da4974b-5b3c-4573-bf92-baefe75efb87"
		var v uint64 = 1696042973491484818

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint64_Has(t *testing.T) {
	Convey("TestMapAnyUint64.Has", t, func() {
		var k interface{} = "49d04844-124d-48e4-974d-58e4165f77d1"
		var v uint64 = 6269490645566490063

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("8e49caab-e4d7-4685-ba35-05631d7e5601"+"e3cee60c-03f7-40ba-9e03-336c66b3f8d4"), ShouldBeFalse)
	})
}

func TestMapAnyUint64_Get(t *testing.T) {
	Convey("TestMapAnyUint64.Get", t, func() {
		var k interface{} = "8d217596-ce40-4783-b02d-720b04e945f4"
		var v uint64 = 4855783388361913440

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("96b3ef4a-21b0-4f74-b69f-83d9d09f3d91" + "345b45c5-298c-4947-b835-1be9d6953227")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint64_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint64.GetOpt", t, func() {
		var k interface{} = "91e72a75-e0df-48ad-9211-2fc44e298e76"
		var v uint64 = 16212273862719570376

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b604479b-c184-454f-8e8b-bf0d5a7d1c39" + "329669b0-9711-47f1-bc41-88697e3894a9")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint64_ForEach(t *testing.T) {
	Convey("TestMapAnyUint64.ForEach", t, func() {
		var k interface{} = "201fb6c7-f5bf-4dc9-972d-7f175c6d7f7b"
		var v uint64 = 14645549531512902908
		hits := 0

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint64.MarshalYAML", t, func() {
		var k interface{} = "d56257c9-901f-451f-8077-157b55b30379"
		var v uint64 = 3172240397331464829

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint64_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint64.ToYAML", t, func() {
		var k interface{} = "3ab1534c-9b94-4768-9826-768c85993b6b"
		var v uint64 = 15092181199285216496

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint64.PutIfNotNil", t, func() {
		var k interface{} = "a31363c5-5013-4060-adfb-fce748945b95"
		var v uint64 = 8087184039116674683

		test := omap.NewMapAnyUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("a3440da3-cc40-49f1-bf2d-0726535eea2e", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 11009898156963372993
		So(test.PutIfNotNil("bc04a2d9-ff0c-4537-a349-a324a82e0872", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceIfExists", t, func() {
		var k interface{} = "944cfe2a-3653-4583-b26f-a68e3d9bcdcc"
		var v uint64 = 7565985999390146299
		var x uint64 = 6207068707466406954

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("b64ed5fc-f45b-49bd-ae30-5e7771b3d760", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceOrPut", t, func() {
		var k interface{} = "4e08d688-4954-470d-822f-3ab075a13c7d"
		var v uint64 = 10216995852740003072
		var x uint64 = 3522909439706154917

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("7e83e2cc-38cc-4ef9-a073-1cab130c0026", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint64.MarshalJSON", t, func() {
		var k interface{} = "a96ca186-77bf-4ea1-bf29-f0f6544020bd"
		var v uint64 = 3990237244928198610

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a96ca186-77bf-4ea1-bf29-f0f6544020bd","value":3990237244928198610}]`)
	})
}
