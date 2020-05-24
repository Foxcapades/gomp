package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt64_Put(t *testing.T) {
	Convey("TestMapStringInt64.Put", t, func() {
		var k string = "7f97e3ef-1abe-43df-9042-d6e6548b6224"
		var v int64 = 2680322567411810914

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt64_Delete(t *testing.T) {
	Convey("TestMapStringInt64.Delete", t, func() {
		var k string = "bb596981-93f2-459c-8d5a-d112b57c8b6e"
		var v int64 = 1078641031860161745

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt64_Has(t *testing.T) {
	Convey("TestMapStringInt64.Has", t, func() {
		var k string = "b834cfb0-9b30-47a4-bdab-e71c22a22e45"
		var v int64 = 8946227838402883428

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b222e5cd-3cb3-497b-aed1-24ac38a31953"+"77534bf1-4938-4890-b723-79cf9455da75"), ShouldBeFalse)
	})
}

func TestMapStringInt64_Get(t *testing.T) {
	Convey("TestMapStringInt64.Get", t, func() {
		var k string = "1779e3de-fed7-4cdb-acca-0c621da6333b"
		var v int64 = 6851714880257551932

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("403611ea-9c7d-4a6d-b8a3-3c8f61426431" + "2b57aa2e-e89b-4bc7-982a-60b67cb14e58")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt64_GetOpt(t *testing.T) {
	Convey("TestMapStringInt64.GetOpt", t, func() {
		var k string = "36821319-0792-495d-a9a1-67e445965565"
		var v int64 = 7418674618600498689

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("93252414-e943-47d3-8097-12ae3f2c15e0" + "40626cb4-7f4d-4c90-a478-192d93246280")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt64_ForEach(t *testing.T) {
	Convey("TestMapStringInt64.ForEach", t, func() {
		var k string = "d3715292-e423-4654-95e4-1190358711da"
		var v int64 = 8448825487527157283
		hits := 0

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt64.MarshalYAML", t, func() {
		var k string = "9f2abe24-fab3-4c8f-aa21-850c74ea8217"
		var v int64 = 6643195232284515840

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt64_ToYAML(t *testing.T) {
	Convey("TestMapStringInt64.ToYAML", t, func() {
		var k string = "d12766ed-2057-49e4-8184-da66d22c9390"
		var v int64 = 5369951788412798583

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt64.PutIfNotNil", t, func() {
		var k string = "4ff2c4ec-ea86-4e5d-9e43-f14a1e6750d4"
		var v int64 = 6705530279660165628

		test := omap.NewMapStringInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("992c602a-ddc6-4e18-bb77-9adc298a18ea", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 50440987832445080
		So(test.PutIfNotNil("fe73f985-fe7a-4a63-8fd0-ee0685c96d30", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceIfExists", t, func() {
		var k string = "b10518dd-5e85-4b6b-8e92-e5bd0283e87e"
		var v int64 = 6934281221212117344
		var x int64 = 7405146930012810586

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("afa036e2-9e6d-4628-a1dc-a13d47244402", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceOrPut", t, func() {
		var k string = "8736e62d-e223-43f2-b711-cbbdac881eef"
		var v int64 = 7134451957941748034
		var x int64 = 8650665610669252330

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("38a251ab-9e68-46ac-8651-cd510ed07e8c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt64.MarshalJSON", t, func() {
		var k string = "d5fac4d2-df56-4f9d-94a0-5f229adeaec0"
		var v int64 = 5173280837258564623

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"d5fac4d2-df56-4f9d-94a0-5f229adeaec0","value":5173280837258564623}]`)
	})
}
