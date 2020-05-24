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
		var k string = "e559d651-5ab4-4eba-aa3a-5c6394deefb1"
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
		var k string = "b11593c8-877c-4666-9a33-a78866c6b953"
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
		var k string = "bedd4343-29ff-4d14-821b-d6c8ee5d7d6d"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("80ef0734-192e-40c4-b08b-925b87eeace8"+"c8aeee53-ae62-4528-9914-f98854d90b6e"), ShouldBeFalse)
	})
}


func TestMapStringBool_Get(t *testing.T) {
	Convey("TestMapStringBool.Get", t, func() {
		var k string = "630eb4e9-ba93-478e-9a6a-244ced88b46e"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("317184e4-a442-4660-95cf-b85296db7fd4" + "5ba542c8-223a-4268-8b81-4344a5a943f8")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringBool_GetOpt(t *testing.T) {
	Convey("TestMapStringBool.GetOpt", t, func() {
		var k string = "f1f4378e-e9d1-4374-abad-316cf54ad3fe"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("1d61dc86-266a-4dec-9ca1-b9e116c6ca61" + "6a193b21-2101-44dc-ad72-2a636ce0acff")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringBool_ForEach(t *testing.T) {
	Convey("TestMapStringBool.ForEach", t, func() {
		var k string = "a8af972b-a0cb-4886-bfff-7fabd527be1b"
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
		var k string = "3d1fa416-9218-4761-aa8d-ddc4720f9bbe"
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
		var k string = "36c114b1-7997-4879-a398-ab80188306e4"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringBool.PutIfNotNil", t, func() {
		var k string = "84de57f1-b0f7-4f2f-8f53-464c1cb31449"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("d30ab1fc-dbb4-45b8-b270-773ed3ac8c68", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("d8a4d3d4-2a2d-4442-81eb-ffddf479f7a6", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringBool.ReplaceIfExists", t, func() {
		var k string = "3ba31e82-9123-4ba9-b1b4-8dfc2eb68bc2"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("3287aec8-b342-4316-9079-685b35b3dd5c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringBool.ReplaceOrPut", t, func() {
		var k string = "01683962-5803-40d6-9c53-9f79a932353c"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("5ed7959f-7c58-402d-bfbb-5d31d16e7c54", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_MarshalJSON(t *testing.T) {
	Convey("TestMapStringBool.MarshalJSON", t, func() {
		var k string = "51aef942-e138-4eb0-a814-bae8d6652040"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"51aef942-e138-4eb0-a814-bae8d6652040","value":false}]`)
	})
}
