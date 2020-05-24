{{- /*gotype: github.com/Foxcapades/gomp/v1/cmd/gomp-gen.MapValDefinition*/ -}}
{{define "tests" -}}
package {{.Package}}_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

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
		So(test.Has(k + k), ShouldBeFalse)
	})
}

{{- end}}