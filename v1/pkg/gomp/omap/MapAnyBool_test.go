package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyBool_Put(t *testing.T) {
	Convey("TestMapAnyBool.Put", t, func() {
		var k interface{} = "18a29f31-288c-4442-9f55-f0ea60dde9ee"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyBool_Delete(t *testing.T) {
	Convey("TestMapAnyBool.Delete", t, func() {
		var k interface{} = "21866f0c-b4b9-429f-b4d0-c68b068fcf02"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyBool_Has(t *testing.T) {
	Convey("TestMapAnyBool.Has", t, func() {
		var k interface{} = "018db9ec-c683-4965-8c20-9c19f74181f6"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("fdd20f29-ce54-40ef-95d5-0d23fe73251c"+"739ce02b-dcb5-4903-88f1-f953478d528b"), ShouldBeFalse)
	})
}

func TestMapAnyBool_Get(t *testing.T) {
	Convey("TestMapAnyBool.Get", t, func() {
		var k interface{} = "a58fa819-0bee-48e9-811a-2790e0a953f1"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("55b00fbe-4369-4c11-996b-e191a2873e37" + "0fd885cb-31af-4187-9fb4-56be2692d085")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyBool_GetOpt(t *testing.T) {
	Convey("TestMapAnyBool.GetOpt", t, func() {
		var k interface{} = "9c372fab-164b-4bf5-9d34-bb3ef8f0b6e5"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("bb469398-463a-4478-bdfa-8d124c446a7e" + "380f1de8-a122-4c81-b993-ebbb8d418105")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyBool_ForEach(t *testing.T) {
	Convey("TestMapAnyBool.ForEach", t, func() {
		var k interface{} = "cf795b08-cb17-48a8-b62b-3c86c1af9b06"
		var v bool = false
		hits := 0

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyBool_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyBool.MarshalYAML", t, func() {
		var k interface{} = "89b12db5-f269-4a78-8a66-551784d938c0"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyBool_ToYAML(t *testing.T) {
	Convey("TestMapAnyBool.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "bdbba3dd-6060-47ff-bb25-c4368983b145"
			var v bool = false

			test := omap.NewMapAnyBool(1)

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
			var k interface{} = "d64511c3-eb93-4377-be6b-55e1b1171cb3"
			var v bool = false

			test := omap.NewMapAnyBool(1)
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

func TestMapAnyBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyBool.PutIfNotNil", t, func() {
		var k interface{} = "501ecc72-18dc-451b-9302-1c9a2bb9d9a5"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("5099e1c1-dc7e-477d-9a9f-60d38ba26fc1", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("75c67b10-607b-4b4f-8015-f5ee1b69e7b8", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceIfExists", t, func() {
		var k interface{} = "327227d4-9c4b-466c-83ce-5d6fc329cd89"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("2844daae-f3bb-4d8c-a630-0f6018282331", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceOrPut", t, func() {
		var k interface{} = "8b54fd89-69f7-4918-b1b3-ba0865cdb006"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("282df44c-75b8-450c-b0bb-d600debb8477", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyBool.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "fad6b79c-6bb7-45a8-8d98-ea773289c96d"
			var v bool = false

			test := omap.NewMapAnyBool(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"fad6b79c-6bb7-45a8-8d98-ea773289c96d","value":false}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "fad6b79c-6bb7-45a8-8d98-ea773289c96d"
			var v bool = false

			test := omap.NewMapAnyBool(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"fad6b79c-6bb7-45a8-8d98-ea773289c96d":false}`)
		})

	})
}
