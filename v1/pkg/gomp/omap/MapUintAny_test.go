package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintAny_Put(t *testing.T) {
	Convey("TestMapUintAny.Put", t, func() {
		var k uint = 3028382565
		var v interface{} = "8d5e5a59-3c59-47d9-ad9d-60b5260a6de1"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintAny_Delete(t *testing.T) {
	Convey("TestMapUintAny.Delete", t, func() {
		var k uint = 3962255268
		var v interface{} = "4a27345d-ce96-4f10-a717-0aa107b1b034"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintAny_Has(t *testing.T) {
	Convey("TestMapUintAny.Has", t, func() {
		var k uint = 3422946123
		var v interface{} = "fef6564c-df63-454c-b37c-d4fbce523f6b"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(4259488858+3632758635), ShouldBeFalse)
	})
}

func TestMapUintAny_Get(t *testing.T) {
	Convey("TestMapUintAny.Get", t, func() {
		var k uint = 3357258095
		var v interface{} = "bd1258a1-e38c-4fda-9344-80ad0cd97d60"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(71863880 + 249719080)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintAny_GetOpt(t *testing.T) {
	Convey("TestMapUintAny.GetOpt", t, func() {
		var k uint = 631672581
		var v interface{} = "3ace76a9-0af8-4d44-b1c1-11c7c7562be7"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2620337805 + 3447753577)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintAny_ForEach(t *testing.T) {
	Convey("TestMapUintAny.ForEach", t, func() {
		var k uint = 2117575855
		var v interface{} = "2ef18178-0976-426a-aef3-2a38406055e7"
		hits := 0

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintAny_MarshalYAML(t *testing.T) {
	Convey("TestMapUintAny.MarshalYAML", t, func() {
		var k uint = 2204478220
		var v interface{} = "a9432050-0ddb-48a0-8bb7-e8033eb5e738"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintAny_ToYAML(t *testing.T) {
	Convey("TestMapUintAny.ToYAML", t, func() {
		var k uint = 3596288187
		var v interface{} = "77ef1fa2-6515-4a98-8a18-a5c95b8a3b6b"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintAny.PutIfNotNil", t, func() {
		var k uint = 743517721
		var v interface{} = "033c42b7-c59f-4837-8164-b8c5500a04b3"

		test := omap.NewMapUintAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(927795805, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "30e49fbf-e7f8-402a-a606-e11e9b6d7bb5"
		So(test.PutIfNotNil(3956602918, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintAny.ReplaceIfExists", t, func() {
		var k uint = 2530411918
		var v interface{} = "cb608e9a-c129-4ac9-a050-c87b3ee0b69e"
		var x interface{} = "e2d2b097-a28f-46b7-a60f-7f36d02344f3"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2581469680, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintAny.ReplaceOrPut", t, func() {
		var k uint = 3628811830
		var v interface{} = "32e86c19-8f7c-4f28-8591-972301e20931"
		var x interface{} = "1f7c2ee3-3002-4b7b-a6b8-e9fa9a5e8ce2"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(450843694, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_MarshalJSON(t *testing.T) {
	Convey("TestMapUintAny.MarshalJSON", t, func() {
		var k uint = 4033784224
		var v interface{} = "0507e8da-0695-4e6a-8b0e-45b9510bf050"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":4033784224,"value":"0507e8da-0695-4e6a-8b0e-45b9510bf050"}]`)
	})
}
