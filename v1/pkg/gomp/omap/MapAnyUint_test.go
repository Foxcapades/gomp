package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint_Put(t *testing.T) {
	Convey("TestMapAnyUint.Put", t, func() {
		var k interface{} = "14a7a2d6-d97b-415b-8c94-b42a66627c69"
		var v uint = 709845755

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint_Delete(t *testing.T) {
	Convey("TestMapAnyUint.Delete", t, func() {
		var k interface{} = "a0aa7759-1372-42f8-92e9-26abf5bcdf92"
		var v uint = 566079820

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint_Has(t *testing.T) {
	Convey("TestMapAnyUint.Has", t, func() {
		var k interface{} = "f3b9873e-e279-4995-b165-0eee3d3961e3"
		var v uint = 932220678

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("adca44a6-539b-4e1c-8029-ad53f7d4ad75"+"0dd9e6dc-3c96-4f17-80d6-7a9557748e9c"), ShouldBeFalse)
	})
}


func TestMapAnyUint_Get(t *testing.T) {
	Convey("TestMapAnyUint.Get", t, func() {
		var k interface{} = "a9250a33-21fa-4c1a-b869-61aa620c7531"
		var v uint = 227172725

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("ac207023-aa5b-477e-a699-912487dac0e0"+"aa4fce82-15fc-455c-9bb8-50c02c91b66c")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint.GetOpt", t, func() {
		var k interface{} = "ebd9b5f9-32af-407e-8da2-9b8a1838b305"
		var v uint = 3424690375

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("a1e4a76d-0ace-4878-861d-3721b041b773"+"a9684884-2310-4a24-9c2f-edd6c3de17b6")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint_ForEach(t *testing.T) {
	Convey("TestMapAnyUint.ForEach", t, func() {
		var k interface{} = "86ae6db8-9235-4edf-a7b3-9983fc81a966"
		var v uint = 1932103871
		hits := 0

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint.MarshalYAML", t, func() {
		var k interface{} = "6e20ec10-08d4-40e2-9653-4b395e3212ab"
		var v uint = 506245773

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint.ToYAML", t, func() {
		var k interface{} = "7d95d2b1-90de-4544-9d3e-59700a7ec391"
		var v uint = 254685618

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint.PutIfNotNil", t, func() {
		var k interface{} = "66d2a4a1-985a-429b-b04f-f89368990552"
		var v uint = 2430799222

		test := omap.NewMapAnyUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("da93dd05-96b0-4735-83fe-ecba27d465b2", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 1511057163
		So(test.PutIfNotNil("12bc43ef-71e8-4209-b6cc-12983bc0fa65", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceIfExists", t, func() {
		var k interface{} = "15645974-a583-4d21-9ae0-7037be5f147c"
		var v uint = 1483174034
		var x uint = 2099364077

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("dcb7da4c-8c31-42ab-8d4e-a7b064a07e82", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceOrPut", t, func() {
		var k interface{} = "0bc30b0f-8659-47ef-b741-67f43c5bf5c5"
		var v uint = 4182188788
		var x uint = 507842378

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("4c8fb3d4-ed2f-48b5-861d-8cf780745514", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint.MarshalJSON", t, func() {
		var k interface{} = "704fd9ce-1e29-4171-a9a9-83c3a4d41da5"
		var v uint = 674671353

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"704fd9ce-1e29-4171-a9a9-83c3a4d41da5","value":674671353}]`)
	})
}

