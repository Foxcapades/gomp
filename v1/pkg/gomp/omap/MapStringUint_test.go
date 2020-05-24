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
		var k string = "abe99a74-8be5-4c2c-9e20-7ca51755565e"
		var v uint = 3022005504

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint_Delete(t *testing.T) {
	Convey("TestMapStringUint.Delete", t, func() {
		var k string = "7b0a2866-e730-433c-9226-4fd8e8eb40fc"
		var v uint = 1999531610

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint_Has(t *testing.T) {
	Convey("TestMapStringUint.Has", t, func() {
		var k string = "57f6ee36-19f3-4df8-aa96-838c6af42863"
		var v uint = 2968910300

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d3da4b50-8cdf-4632-873d-9afbc39becd8"+"2ba15cbe-a0c5-4a03-96a2-e3037d5d9954"), ShouldBeFalse)
	})
}

func TestMapStringUint_Get(t *testing.T) {
	Convey("TestMapStringUint.Get", t, func() {
		var k string = "19cd8239-e917-4dec-9cff-5b0134d802ae"
		var v uint = 2499031540

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("7cfe5837-90e8-4181-b254-e04ea0e82613" + "121de855-5238-4377-a63d-6f1b82350e6d")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint_GetOpt(t *testing.T) {
	Convey("TestMapStringUint.GetOpt", t, func() {
		var k string = "1618e024-3396-49f3-a00b-fcad4a2aae89"
		var v uint = 14466746

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5d4e2479-0c25-4e7f-974b-dd4a5904e1a2" + "b28366e8-5cf1-4e16-8baf-76e17e93d0e5")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint_ForEach(t *testing.T) {
	Convey("TestMapStringUint.ForEach", t, func() {
		var k string = "f31faa82-b92d-4305-8c22-6a3d6a8d0633"
		var v uint = 3428210181
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
		var k string = "3a871f03-8f3c-4b46-a72c-99dcdb199dc8"
		var v uint = 125830215

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
		var k string = "ca6d3df1-532a-459e-a29f-0e2e5f49adb7"
		var v uint = 2886375792

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
		var k string = "5cac458b-7087-4d07-90e9-488bc0e1c558"
		var v uint = 1745776726

		test := omap.NewMapStringUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("f4757b44-6fc0-4351-a5c9-0f74d16e4b46", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 2147218647
		So(test.PutIfNotNil("602d5337-6c60-40f9-a050-551eee3cbdfc", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint.ReplaceIfExists", t, func() {
		var k string = "096b646f-0a4e-4935-b19f-8731ec91b8d4"
		var v uint = 3998568326
		var x uint = 1463824858

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e76c5669-19a6-4868-aff0-c49cb548fcd8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint.ReplaceOrPut", t, func() {
		var k string = "cbf8994a-288e-4e44-b5a8-eadf71b8ba70"
		var v uint = 736366825
		var x uint = 3850633460

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ee790845-f41c-495d-8c1d-f5a7873e4a22", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint.MarshalJSON", t, func() {
		var k string = "a8260ee2-7869-4421-8a9d-912a6586329a"
		var v uint = 3957231572

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a8260ee2-7869-4421-8a9d-912a6586329a","value":3957231572}]`)
	})
}
