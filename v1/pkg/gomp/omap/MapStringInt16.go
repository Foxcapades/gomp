package omap

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapStringInt16 defines an ordered map of string to int16.
type MapStringInt16 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v int16) MapStringInt16

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *int16) MapStringInt16

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v int16) MapStringInt16

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v int16) MapStringInt16

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value int16, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of int16 if an entry was found
	// with the key `k`.
	GetOpt(k string) option.Int16

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapStringInt16Entry

	// Len returns the current size of the map.
	Len() int

	// Has returns whether an entry exists with the key `k`.
	Has(k string) bool

	// Has returns the position in the map of the entry matching key `k`.
	//
	// If no entry exists in the map with key `k` this method returns -1.
	//
	// Note: this method may iterate, at most once, through all the entries in the
	// map.
	IndexOf(k string) int

	Delete(k string) MapStringInt16

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v int16)) MapStringInt16

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) MapStringInt16
}

// MapStringInt16Entry is a single entry in an instance of
// MapStringInt16.
type MapStringInt16Entry struct {
	Key string `json:"key"`
	Val int16  `json:"value"`
}

// NewMapStringInt16 creates a new instance of MapStringInt16 presized to the
// given size.
func NewMapStringInt16(size int) MapStringInt16 {
	return &implMapStringInt16{
		ordered:  make([]MapStringInt16Entry, 0, size),
		index:    make(map[string]int16, size),
		outOrder: true,
	}
}

// MapStringInt16 is an ordered map string to int16.
type implMapStringInt16 struct {
	ordered  []MapStringInt16Entry
	index    map[string]int16
	outOrder bool
}

func (i implMapStringInt16) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapStringInt16) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i *implMapStringInt16) ToYAML() (*yaml.Node, error) {
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

func (i *implMapStringInt16) Put(k string, v int16) MapStringInt16 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapStringInt16Entry{k, v})
	return i
}

func (i *implMapStringInt16) PutIfNotNil(k string, v *int16) MapStringInt16 {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapStringInt16) ReplaceOrPut(k string, v int16) MapStringInt16 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapStringInt16) ReplaceIfExists(k string, v int16) MapStringInt16 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapStringInt16) Get(k string) (value int16, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapStringInt16) GetOpt(k string) option.Int16 {
	if v, ok := i.index[k]; ok {
		return option.NewInt16(v)
	}

	return option.NewEmptyInt16()
}

func (i *implMapStringInt16) At(j int) MapStringInt16Entry {
	return i.ordered[j]
}

func (i *implMapStringInt16) Len() int {
	return len(i.ordered)
}

func (i *implMapStringInt16) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapStringInt16) IndexOf(k string) int {
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

func (i *implMapStringInt16) Delete(k string) MapStringInt16 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapStringInt16) ForEach(f func(k string, v int16)) MapStringInt16 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implMapStringInt16) SerializeOrdered(b bool) MapStringInt16 {
	i.outOrder = b
	return i
}
