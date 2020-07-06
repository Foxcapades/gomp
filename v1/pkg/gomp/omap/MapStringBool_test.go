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
		var k string = "b502c98d-3403-451c-adb8-4c197252d402"
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
		var k string = "c203f784-2b28-44db-9385-83e080451a2d"
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
		var k string = "b42b3f67-f09d-4fca-9d2d-84fa2a13d887"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5806fcc3-d30b-4634-a398-25a992159d42"+"deb55f2a-2994-4c4a-9b26-7e20e054a9f2"), ShouldBeFalse)
	})
}

func TestMapStringBool_Get(t *testing.T) {
	Convey("TestMapStringBool.Get", t, func() {
		var k string = "bb700f26-d940-4d4c-911d-62964125df81"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("503ca6ef-dc8e-4fbb-bcc7-9f4f083c1e94" + "71c4bc98-5484-4b77-b0c4-b83bc0ac94b5")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringBool_GetOpt(t *testing.T) {
	Convey("TestMapStringBool.GetOpt", t, func() {
		var k string = "1039024b-a947-4988-9a8b-b35b57712ad4"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("e3aa7dbf-81ec-48b3-b37a-6b232e6da879" + "6f94e1fd-c62e-4896-8d4e-60858c2d2d26")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringBool_ForEach(t *testing.T) {
	Convey("TestMapStringBool.ForEach", t, func() {
		var k string = "4f1084f3-b346-488d-b08c-e8e6e03db591"
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
		var k string = "76101233-ebdb-4a30-a757-cdd3573dd966"
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
		Convey("Ordered", func() {
			var k string = "96452275-6555-4540-a8bb-4df484b17a0f"
			var v bool = false

			test := omap.NewMapStringBool(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k string = "5d0a11d5-17d8-4ff1-99d5-a033ddf35abd"
			var v bool = false

			test := omap.NewMapStringBool(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapStringBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringBool.PutIfNotNil", t, func() {
		var k string = "73dc97c4-e6a3-44bc-b0cd-8ca8b0863820"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("1be717f1-c42e-4bbe-802e-e5df966d8dd3", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("7e4ab92d-8d48-432b-809d-57f4a6f93742", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringBool.ReplaceIfExists", t, func() {
		var k string = "ca7e3fa3-4a61-43ff-a3a6-c612a7e13f7d"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("b43d94c5-55ad-4ea9-8c5e-4e5f51c1e2d0", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringBool.ReplaceOrPut", t, func() {
		var k string = "f0b90c8a-3eb7-4b5a-9b5f-f3603ce7c500"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d2fa2692-ca32-4bb9-b4a5-69ce9d2f48e7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_MarshalJSON(t *testing.T) {
	Convey("TestMapStringBool.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "693a0f61-e14e-491c-894d-2cf22eb12056"
			var v bool = false

			test := omap.NewMapStringBool(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"693a0f61-e14e-491c-894d-2cf22eb12056","value":false}]`)
		})

		Convey("Unordered", func() {
			var k string = "693a0f61-e14e-491c-894d-2cf22eb12056"
			var v bool = false

			test := omap.NewMapStringBool(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"693a0f61-e14e-491c-894d-2cf22eb12056":false}`)
		})

	})
}
