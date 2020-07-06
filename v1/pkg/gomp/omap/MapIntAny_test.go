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
		var k int = 914926001
		var v interface{} = "bcf67178-ef6d-403c-8f8b-d45dcf541e16"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntAny_Delete(t *testing.T) {
	Convey("TestMapIntAny.Delete", t, func() {
		var k int = 1397557191
		var v interface{} = "ba9aba02-8af9-4b9c-a24a-a9dfb1ad47e6"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntAny_Has(t *testing.T) {
	Convey("TestMapIntAny.Has", t, func() {
		var k int = 1354348693
		var v interface{} = "ce53ef7d-2a7f-4de9-9b6f-d1e792d50580"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(153484554+233139240), ShouldBeFalse)
	})
}

func TestMapIntAny_Get(t *testing.T) {
	Convey("TestMapIntAny.Get", t, func() {
		var k int = 1143112935
		var v interface{} = "dc4cbba4-1bed-45ce-86b4-9d3e2a29f283"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2018418498 + 16498760)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntAny_GetOpt(t *testing.T) {
	Convey("TestMapIntAny.GetOpt", t, func() {
		var k int = 894137148
		var v interface{} = "f2dfc949-7762-4d29-b11c-7b4ace09b51a"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(495418517 + 17249227)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntAny_ForEach(t *testing.T) {
	Convey("TestMapIntAny.ForEach", t, func() {
		var k int = 552042551
		var v interface{} = "ca523db8-d01c-4a01-b581-9e2594803d44"
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
		var k int = 594320773
		var v interface{} = "5124194a-47f3-4f14-89f2-3f83856eb80c"

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
		Convey("Ordered", func() {
			var k int = 719019044
			var v interface{} = "7fb898ac-2ae1-4a49-a422-0bd9f168c677"

			test := omap.NewMapIntAny(1)

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
			var k int = 631525818
			var v interface{} = "94c7d94e-cf83-44f5-a7bb-0780e6350621"

			test := omap.NewMapIntAny(1)
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

func TestMapIntAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntAny.PutIfNotNil", t, func() {
		var k int = 624754617
		var v interface{} = "d00714de-1e0e-48d6-a68d-50607f1ace26"

		test := omap.NewMapIntAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(88977473, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "ec51ba62-8cca-44a9-b78d-1e34cb6ee5a8"
		So(test.PutIfNotNil(1067665248, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntAny.ReplaceIfExists", t, func() {
		var k int = 966187941
		var v interface{} = "ed6c0f61-3f9f-4038-8736-6e9106781de8"
		var x interface{} = "43a2a79f-10d1-4088-b3ca-9d88e0795770"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1806689158, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntAny.ReplaceOrPut", t, func() {
		var k int = 93335093
		var v interface{} = "1d5d9202-40b5-4370-a412-b0914af52c02"
		var x interface{} = "c26d3f5c-056a-4a5a-8f9a-ee5a57a06db3"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(830486101, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_MarshalJSON(t *testing.T) {
	Convey("TestMapIntAny.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k int = 1929997107
			var v interface{} = "aa1c5163-6ead-433a-9a5d-2756906d7ada"

			test := omap.NewMapIntAny(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":1929997107,"value":"aa1c5163-6ead-433a-9a5d-2756906d7ada"}]`)
		})

		Convey("Unordered", func() {
			var k int = 1929997107
			var v interface{} = "aa1c5163-6ead-433a-9a5d-2756906d7ada"

			test := omap.NewMapIntAny(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"1929997107":"aa1c5163-6ead-433a-9a5d-2756906d7ada"}`)
		})

	})
}
