{{- /*gotype: github.com/Foxcapades/gomp/v1/cmd/gomp-gen.MapValDefinition*/ -}}
{{define "interface" -}}
package {{.Package}}

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	{{if isBase .Type -}}"github.com/Foxcapades/goop/v1/pkg/option"
	{{end -}}
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// {{.Name}} defines an ordered map of {{.Key}} to {{.Type}}.
type {{.Name}} interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k {{.Key}}, v {{.Type}}) {{.Name}}

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k {{.Key}}, v {{if ne .Type "interface{}"}}*{{end}}{{.Type}}) {{.Name}}

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k {{.Key}}, v {{.Type}}) {{.Name}}

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k {{.Key}}, v {{.Type}}) {{.Name}}

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k {{.Key}}) (value {{.Type}}, exists bool)

	{{if isBase .Type -}}
	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of {{.Type}} if an entry was found
	// with the key `k`.
	GetOpt(k {{.Key}}) option.{{if eq .Type "interface{}"}}Untyped{{else}}{{titleCap .Type}}{{end}}
	{{- end}}

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) {{.Name}}Entry

	// Len returns the current size of the map.
	Len() int

	// Has returns whether an entry exists with the key `k`.
	Has(k {{.Key}}) bool

	// Has returns the position in the map of the entry matching key `k`.
	//
	// If no entry exists in the map with key `k` this method returns -1.
	//
	// Note: this method may iterate, at most once, through all the entries in the
	// map.
	IndexOf(k {{.Key}}) int

	Delete(k {{.Key}}) {{.Name}}

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k {{.Key}}, v {{.Type}})) {{.Name}}

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) {{.Name}}
}

// {{.Name}}Entry is a single entry in an instance of
// {{.Name}}.
type {{.Name}}Entry struct {
	Key {{.Key}}{{pad .Key .Type}} `json:"key"`
	Val {{.Type}}{{pad .Type .Key}} `json:"value"`
}

// New{{.Name}} creates a new instance of {{.Name}} presized to the
// given size.
func New{{.Name}}(size int) {{.Name}} {
	return &impl{{.Name}}{
		ordered:  make([]{{.Name}}Entry, 0, size),
		index:    make(map[{{.Key}}]{{.Type}}, size),
		outOrder: true,
	}
}

// {{.Name}} is an ordered map {{.Key}} to {{.Type}}.
type impl{{.Name}} struct {
	ordered  []{{.Name}}Entry
	index    map[{{.Key}}]{{.Type}}
	outOrder bool
}

func (i impl{{.Name}}) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i impl{{.Name}}) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i *impl{{.Name}}) ToYAML() (*yaml.Node, error) {
	if i.outOrder {
		out := xyml.NewOrderedMapNode(i.Len())

		for j := range i.ordered {
			tmp := xyml.NewMapNode(1)
			if e := xyml.MapAppend(tmp, i.ordered[j].Key, i.ordered[j].Val); e != nil {
				return nil, e
			}
			if err := xyml.SequenceAppend(out, tmp); err != nil {
				return nil, err
			}
		}

		return out, nil
	}

	out := xyml.NewMapNode(i.Len())

	for j := range i.ordered {
		if e := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); e != nil {
			return nil, e
		}
	}

	return out, nil
}

func (i *impl{{.Name}}) Put(k {{.Key}}, v {{.Type}}) {{.Name}} {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, {{.Name}}Entry{k, v})
	return i
}

func (i *impl{{.Name}}) PutIfNotNil(k {{.Key}}, v {{if ne .Type "interface{}"}}*{{end}}{{.Type}}) {{.Name}} {
	if !gomp.IsNil(v) {
		{{if ne .Type "interface{}" -}}
		return i.Put(k, *v)
		{{- else -}}
		return i.Put(k, gomp.Deref(v))
		{{- end}}
	}

	return i
}

func (i *impl{{.Name}}) ReplaceOrPut(k {{.Key}}, v {{.Type}}) {{.Name}} {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *impl{{.Name}}) ReplaceIfExists(k {{.Key}}, v {{.Type}}) {{.Name}} {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *impl{{.Name}}) Get(k {{.Key}}) (value {{.Type}}, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

{{if isBase .Type -}}
func (i *impl{{.Name}}) GetOpt(k {{.Key}}) option.{{if eq .Type "interface{}"}}Untyped{{else}}{{titleCap .Type}}{{end}} {
	if v, ok := i.index[k]; ok {
		return option.New{{if eq .Type "interface{}"}}Untyped{{else}}{{titleCap .Type}}{{end}}(v)
	}

	return option.NewEmpty{{if eq .Type "interface{}"}}Untyped{{else}}{{titleCap .Type}}{{end}}()
}
{{- end}}

func (i *impl{{.Name}}) At(j int) {{.Name}}Entry {
	return i.ordered[j]
}

func (i *impl{{.Name}}) Len() int {
	return len(i.ordered)
}

func (i *impl{{.Name}}) Has(k {{.Key}}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *impl{{.Name}}) IndexOf(k {{.Key}}) int {
	if _, ok := i.index[k]; !ok {
		return -1
	}
	for j := range i.ordered {
		if i.ordered[j].Key == k {
			return j
		}
	}
	panic("invalid map state, out of sync")
}

func (i *impl{{.Name}}) Delete(k {{.Key}}) {{.Name}} {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *impl{{.Name}}) ForEach(f func(k {{.Key}}, v {{.Type}})) {{.Name}} {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *impl{{.Name}}) SerializeOrdered(b bool) {{.Name}} {
	i.outOrder = b
	return i
}
{{end -}}