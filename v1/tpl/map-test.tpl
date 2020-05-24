{{- /*gotype: github.com/Foxcapades/gomp/v1/cmd/gomp-gen.MapValDefinition*/ -}}
{{define "tests" -}}
package {{.Package}}_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"{{.Repo}}"
)

func Test{{.Name}}_Put(t *testing.T) {
	Convey("Test{{.Name}}.Put", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func Test{{.Name}}_Delete(t *testing.T) {
	Convey("Test{{.Name}}.Delete", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func Test{{.Name}}_Has(t *testing.T) {
	Convey("Test{{.Name}}.Has", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has({{defVal .Key}}+{{defVal .Key}}), ShouldBeFalse)
	})
}


func Test{{.Name}}_Get(t *testing.T) {
	Convey("Test{{.Name}}.Get", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get({{defVal .Key}} + {{defVal .Key}})
		So(b, ShouldBeFalse)
	})
}

{{if isBase .Type -}}
func Test{{.Name}}_GetOpt(t *testing.T) {
	Convey("Test{{.Name}}.GetOpt", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt({{defVal .Key}} + {{defVal .Key}})
		So(a.IsNil(), ShouldBeTrue)
	})
}

{{end -}}

func Test{{.Name}}_ForEach(t *testing.T) {
	Convey("Test{{.Name}}.ForEach", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}
		hits := 0

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk {{.Key}}, vv {{.Type}}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func Test{{.Name}}_MarshalYAML(t *testing.T) {
	Convey("Test{{.Name}}.MarshalYAML", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func Test{{.Name}}_ToYAML(t *testing.T) {
	Convey("Test{{.Name}}.ToYAML", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func Test{{.Name}}_PutIfNotNil(t *testing.T) {
	Convey("Test{{.Name}}.PutIfNotNil", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*{{.Type}})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil({{defVal .Key}}, (*{{.Type}})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x {{.Type}} = {{defVal .Type}}
		So(test.PutIfNotNil({{defVal .Key}}, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func Test{{.Name}}_ReplaceIfExists(t *testing.T) {
	Convey("Test{{.Name}}.ReplaceIfExists", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}
		var x {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists({{defVal .Key}}, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func Test{{.Name}}_ReplaceOrPut(t *testing.T) {
	Convey("Test{{.Name}}.ReplaceOrPut", t, func() {
		var k {{.Key}} = {{defVal .Key}}
		var v {{.Type}} = {{defVal .Type}}
		var x {{.Type}} = {{defVal .Type}}

		test := {{.Package}}.New{{.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut({{defVal .Key}}, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

{{with $key := defVal .Key -}}
{{with $val := defVal $.Type -}}
func Test{{$.Name}}_MarshalJSON(t *testing.T) {
	Convey("Test{{$.Name}}.MarshalJSON", t, func() {
		var k {{$.Key}} = {{$key}}
		var v {{$.Type}} = {{$val}}

		test := {{$.Package}}.New{{$.Name}}(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		{{if eq $.Type "float32" "float64" -}}
		So(string(a), ShouldEqual, `[{"key":{{$key}},"value":{{trimR $val ".0"}}}]`)
		{{- else -}}
		So(string(a), ShouldEqual, `[{"key":{{$key}},"value":{{$val}}}]`)
		{{- end}}
	})
}{{end}}{{end}}
{{end -}}