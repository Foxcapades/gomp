package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringBool_Put(t *testing.T) {
	Convey("TestMapStringBool.Put", t, func() {
		var k string = "498d1efc-b841-4737-a62b-fb0f4f1d1323"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringBool_Delete(t *testing.T) {
	Convey("TestMapStringBool.Delete", t, func() {
		var k string = "b380f060-8941-4df6-8eee-8d27694e6979"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringBool_Has(t *testing.T) {
	Convey("TestMapStringBool.Has", t, func() {
		var k string = "5c58c119-9d38-424f-b0eb-6863921601d7"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("40870dae-e12d-40e3-935f-683641ca138a"+"4b496a11-0c22-4248-8fdc-cf47f84ad661"), ShouldBeFalse)
	})
}

func TestMapStringBool_Get(t *testing.T) {
	Convey("TestMapStringBool.Get", t, func() {
		var k string = "a6b527e1-ac34-4b83-9f11-5325e70ecca8"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("8e942b7c-3488-481b-8e83-28c35b7c6c3c" + "9a59535f-aed2-4710-8ac0-e55241397f76")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringBool_GetOpt(t *testing.T) {
	Convey("TestMapStringBool.GetOpt", t, func() {
		var k string = "a4e0a92c-2f0b-4bbd-9fba-720891a8dcb7"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("2a236fb0-8744-4d95-89f4-5e1c3757338f" + "fb7643bd-61fa-4a0f-b832-cc9c3bf16087")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringBool_ForEach(t *testing.T) {
	Convey("TestMapStringBool.ForEach", t, func() {
		var k string = "15dffd83-55fb-4004-b7f1-3bdd6d491468"
		var v bool = false
		hits := 0

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringBool_MarshalYAML(t *testing.T) {
	Convey("TestMapStringBool.MarshalYAML", t, func() {
		var k string = "b0b00d4f-6bf5-4e4b-8fd0-b257052052f0"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringBool_ToYAML(t *testing.T) {
	Convey("TestMapStringBool.ToYAML", t, func() {
		var k string = "740393f2-00bc-4e99-945a-9fdd5494bc7d"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapStringBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringBool.PutIfNotNil", t, func() {
		var k string = "feeb6ac2-7767-493f-9f2e-24b7323e4861"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("45fdaab1-b4ca-4b12-82c9-1a16c5961de8", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("b2cf991b-58aa-45d1-ad33-7999360ee7bc", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringBool.ReplaceIfExists", t, func() {
		var k string = "1765bc7f-6b60-4d15-aaa8-c1747b04645b"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("79b3c676-4358-41e7-a3a0-198fd0367655", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringBool.ReplaceOrPut", t, func() {
		var k string = "bc50ebf2-12b2-4365-942f-16d306c91b92"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ad3560a9-4a77-492b-b5a1-f9229aeed04e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_MarshalJSON(t *testing.T) {
	Convey("TestMapStringBool.MarshalJSON", t, func() {
		var k string = "8af58ee9-980d-4dc8-98a9-5afdeccf4c6c"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"8af58ee9-980d-4dc8-98a9-5afdeccf4c6c","value":false}]`)
	})
}
