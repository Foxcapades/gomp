package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint16_Put(t *testing.T) {
	Convey("TestMapAnyUint16.Put", t, func() {
		var k interface{} = "7649806b-a905-4395-8901-74eb52221b48"
		var v uint16 = 37511

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint16_Delete(t *testing.T) {
	Convey("TestMapAnyUint16.Delete", t, func() {
		var k interface{} = "3d69f94f-4ed0-4d12-be6a-52eb19a64226"
		var v uint16 = 11807

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint16_Has(t *testing.T) {
	Convey("TestMapAnyUint16.Has", t, func() {
		var k interface{} = "adb43d23-0c1f-48c3-8567-4c7635c383b1"
		var v uint16 = 19610

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("9154315e-57bb-45a1-b195-c157cef5ec01"+"2fc0fdda-5fdc-4fb6-a099-485a8a262084"), ShouldBeFalse)
	})
}

func TestMapAnyUint16_Get(t *testing.T) {
	Convey("TestMapAnyUint16.Get", t, func() {
		var k interface{} = "d5ca47bb-0137-487e-9295-6c5350c9ceea"
		var v uint16 = 12122

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("2f5652c0-1daa-40e3-8cff-043c55ef3763" + "00ddb89d-8b19-4651-95fe-a2099f92d75a")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint16_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint16.GetOpt", t, func() {
		var k interface{} = "da9d66e4-db34-47fa-9d0c-9327290d63aa"
		var v uint16 = 36892

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("963c5cab-0e21-46d7-a823-e1b2e10d5164" + "3f7cc8f3-da4d-48aa-81ff-829b4c537259")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint16_ForEach(t *testing.T) {
	Convey("TestMapAnyUint16.ForEach", t, func() {
		var k interface{} = "3983ed8b-ca13-4bb6-b811-3c0ebf013cba"
		var v uint16 = 23805
		hits := 0

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalYAML", t, func() {
		var k interface{} = "fe2e1ec0-161a-4e31-98aa-a4ddbc4ce3aa"
		var v uint16 = 15677

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint16_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint16.ToYAML", t, func() {
		var k interface{} = "2fdbf9d0-dbf4-4fde-9d24-921d375106e8"
		var v uint16 = 42392

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint16.PutIfNotNil", t, func() {
		var k interface{} = "81b96d44-0b4b-443c-bc54-7723cef99977"
		var v uint16 = 5507

		test := omap.NewMapAnyUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("3c5c1de8-07ef-49d3-b58a-91662be6dda2", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 63159
		So(test.PutIfNotNil("c335066b-b7bb-45a7-9b36-83d30505a258", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceIfExists", t, func() {
		var k interface{} = "994e38d3-c6cb-42ca-b37b-77e38e15ce12"
		var v uint16 = 51738
		var x uint16 = 43396

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("2bdf10eb-6656-459f-969c-83e4d7722106", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceOrPut", t, func() {
		var k interface{} = "44f9ff17-bb69-4a5b-829d-bbd95c4e83ea"
		var v uint16 = 33272
		var x uint16 = 60913

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("881a9855-13b6-4ae2-961c-b716ea082bde", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalJSON", t, func() {
		var k interface{} = "d8d87422-dcb9-4381-a71e-6b0360e0181e"
		var v uint16 = 1470

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"d8d87422-dcb9-4381-a71e-6b0360e0181e","value":1470}]`)
	})
}
