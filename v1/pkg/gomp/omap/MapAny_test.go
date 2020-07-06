package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAny_Put(t *testing.T) {
	Convey("TestMapAny.Put", t, func() {
		var k interface{} = "1031d2ba-36a2-4ea9-aa19-15d56aea39fd"
		var v interface{} = "8e3ef8fb-d120-447c-b871-f3b52a1b9bac"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAny_Delete(t *testing.T) {
	Convey("TestMapAny.Delete", t, func() {
		var k interface{} = "4802ad6d-3543-4df5-bc12-213e1cb1e1c8"
		var v interface{} = "082ee14d-01f0-4a46-a9b8-90df566ee35d"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAny_Has(t *testing.T) {
	Convey("TestMapAny.Has", t, func() {
		var k interface{} = "d2d4428e-030f-4a00-bb44-6715ec49306f"
		var v interface{} = "48cfd034-015c-4752-b233-e4576b334801"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("45b1ec14-0d02-4e87-a061-5010cbc0cf62"+"180dd586-183b-4488-b771-49f9535170f6"), ShouldBeFalse)
	})
}

func TestMapAny_Get(t *testing.T) {
	Convey("TestMapAny.Get", t, func() {
		var k interface{} = "ad50e0e5-163c-46ff-941b-341f57f1baee"
		var v interface{} = "526440d7-7228-41d1-91e1-559af7f84e73"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("6f45de9a-a3c1-42fa-a358-69c6ff29a99f" + "03bac1c4-8fbf-4d01-8baf-3a8edb3daeb5")
		So(b, ShouldBeFalse)
	})
}

func TestMapAny_GetOpt(t *testing.T) {
	Convey("TestMapAny.GetOpt", t, func() {
		var k interface{} = "085db6f3-7de3-41eb-bb88-914358f95138"
		var v interface{} = "43812aaa-933c-4262-b74b-69ced9c0a7f9"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("2afb4323-2e43-47fb-bc04-8d197c9aff18" + "677a28c4-f499-4f33-a85c-4e19b08e12c7")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAny_ForEach(t *testing.T) {
	Convey("TestMapAny.ForEach", t, func() {
		var k interface{} = "ba5bf060-d757-4958-81ee-3e19844dba0d"
		var v interface{} = "182cbbdb-058b-4101-9997-c25f57e42cad"
		hits := 0

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAny_MarshalYAML(t *testing.T) {
	Convey("TestMapAny.MarshalYAML", t, func() {
		var k interface{} = "a119f8cd-4d45-47bb-aeb8-779cf3087209"
		var v interface{} = "7e1059a0-58b3-4cde-858c-53789d6fb84b"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAny_ToYAML(t *testing.T) {
	Convey("TestMapAny.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "5bd9680b-b314-41ad-873c-ffd05f6e89a9"
			var v interface{} = "f8633a71-3c27-4818-ba3e-4000635da0f5"

			test := omap.NewMapAny(1)

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
			var k interface{} = "1b489cc3-7792-4fca-ac89-a3ed1fced955"
			var v interface{} = "1acbd0c9-f8f1-4002-9fa1-ee55a2188c0e"

			test := omap.NewMapAny(1)
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

func TestMapAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapAny.PutIfNotNil", t, func() {
		var k interface{} = "88659818-93ad-46c3-939e-8e0192e87091"
		var v interface{} = "f32dec28-d3c7-4e03-82aa-194f2ef88ecf"

		test := omap.NewMapAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("5d6aa478-e2c7-4a53-a5bc-d2038d1c37ed", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "57224c50-be67-44cd-8784-8230dda7fd63"
		So(test.PutIfNotNil("e13f0c21-66f6-41f2-8fb2-e8152e193691", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAny.ReplaceIfExists", t, func() {
		var k interface{} = "8207d810-6af0-4070-a3a0-5a1be56be3df"
		var v interface{} = "abc37b5f-46d2-4a08-93de-35b29d4c8f2e"
		var x interface{} = "f15ba035-356c-414f-9a65-32b816b2e684"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("eb8b36f4-1132-4480-8e1f-dc7ae9a0a8ce", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAny.ReplaceOrPut", t, func() {
		var k interface{} = "f2e62b65-914e-4c03-b51a-bc96b1f7bcc2"
		var v interface{} = "6482b722-9a89-43fb-a7a4-c52e8afca5e7"
		var x interface{} = "1af27903-c1b3-4520-b9c0-0368b256b7a4"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("288ab3b2-2ed3-4696-9610-f464ef635c0f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_MarshalJSON(t *testing.T) {
	Convey("TestMapAny.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "3118b896-d42a-4886-beae-cf0b57029317"
			var v interface{} = "2ca843af-9abd-465a-b01c-a2432491175a"

			test := omap.NewMapAny(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"3118b896-d42a-4886-beae-cf0b57029317","value":"2ca843af-9abd-465a-b01c-a2432491175a"}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "3118b896-d42a-4886-beae-cf0b57029317"
			var v interface{} = "2ca843af-9abd-465a-b01c-a2432491175a"

			test := omap.NewMapAny(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"3118b896-d42a-4886-beae-cf0b57029317":"2ca843af-9abd-465a-b01c-a2432491175a"}`)
		})

	})
}
