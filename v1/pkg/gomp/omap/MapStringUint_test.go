package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint_Put(t *testing.T) {
	Convey("TestMapStringUint.Put", t, func() {
		var k string = "23210ba3-4700-4868-9038-591ec02c405b"
		var v uint = 816864974

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint_Delete(t *testing.T) {
	Convey("TestMapStringUint.Delete", t, func() {
		var k string = "09bd63c6-d9b0-4179-901b-74c7e1497370"
		var v uint = 2813652978

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint_Has(t *testing.T) {
	Convey("TestMapStringUint.Has", t, func() {
		var k string = "b4f0a840-645e-4c9c-b015-6c5ee5fcbadd"
		var v uint = 1473182912

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("ad2c6967-8805-480d-98a2-b29f67b5830a"+"b585852f-4c42-4b38-bcac-0e7e6520ecbb"), ShouldBeFalse)
	})
}


func TestMapStringUint_Get(t *testing.T) {
	Convey("TestMapStringUint.Get", t, func() {
		var k string = "42c2dc75-4404-4991-b3fb-ef4e3789c318"
		var v uint = 490938106

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("1e5b883b-96e1-4d4b-8ecd-24049b87a5b7" + "0765657f-8199-4deb-932c-67c830e07df5")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint_GetOpt(t *testing.T) {
	Convey("TestMapStringUint.GetOpt", t, func() {
		var k string = "f39d31ef-99f5-4bb1-88c2-0b63f65b2e2d"
		var v uint = 538483604

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("52b9bb27-a8ff-4df8-8ab5-2a218b51fbcf" + "0f5b7cfe-1713-483c-96ba-d329ccf85b4c")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint_ForEach(t *testing.T) {
	Convey("TestMapStringUint.ForEach", t, func() {
		var k string = "4783fe70-c4c3-4d90-9e6a-1898e630f02c"
		var v uint = 1632036196
		hits := 0

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint.MarshalYAML", t, func() {
		var k string = "0240f784-bcd6-47c9-a809-c404f8c4fda8"
		var v uint = 2500915857

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint_ToYAML(t *testing.T) {
	Convey("TestMapStringUint.ToYAML", t, func() {
		var k string = "c96b8630-a2ed-4993-b717-047db7dbb7c9"
		var v uint = 3927052582

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint.PutIfNotNil", t, func() {
		var k string = "79264690-278e-4a78-9a8f-cc5ff50a732f"
		var v uint = 1810588640

		test := omap.NewMapStringUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("2bd9e2cc-bbcf-4922-94bb-cf914d4e0bb9", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 1676058700
		So(test.PutIfNotNil("5c88ecbd-7aa3-4850-bc6f-17c66b5b3c80", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint.ReplaceIfExists", t, func() {
		var k string = "d429c89e-00f1-4598-975c-cdd38eba46af"
		var v uint = 2120986135
		var x uint = 2385791621

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("d047f945-f78f-46b2-b9fe-12bf4788a25a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint.ReplaceOrPut", t, func() {
		var k string = "29b5be15-d46e-49cd-a3b1-4af44e6cae23"
		var v uint = 2744985755
		var x uint = 3978068760

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("4b8a7aad-75a7-4dea-b456-12e425331096", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint.MarshalJSON", t, func() {
		var k string = "f03a61e2-a154-4711-9237-ca965723b0ec"
		var v uint = 4004934139

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"f03a61e2-a154-4711-9237-ca965723b0ec","value":4004934139}]`)
	})
}

