package omap

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapIntInt16 defines an ordered map of int to int16.
type MapIntInt16 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k int, v int16) MapIntInt16

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k int, v *int16) MapIntInt16

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k int, v int16) MapIntInt16

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k int, v int16) MapIntInt16

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k int) (value int16, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of int16 if an entry was found
	// with the key `k`.
	GetOpt(k int) option.Int16

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapIntInt16Entry

	// Len returns the current size of the map.
	Len() int

	// Has returns whether an entry exists with the key `k`.
	Has(k int) bool

	// Has returns the position in the map of the entry matching key `k`.
	//
	// If no entry exists in the map with key `k` this method returns -1.
	//
	// Note: this method may iterate, at most once, through all the entries in the
	// map.
	IndexOf(k int) int

	Delete(k int) MapIntInt16

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k int, v int16)) MapIntInt16

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) MapIntInt16
}

// MapIntInt16Entry is a single entry in an instance of
// MapIntInt16.
type MapIntInt16Entry struct {
	Key int   `json:"key"`
	Val int16 `json:"value"`
}

// NewMapIntInt16 creates a new instance of MapIntInt16 presized to the
// given size.
func NewMapIntInt16(size int) MapIntInt16 {
	return &implMapIntInt16{
		ordered:  make([]MapIntInt16Entry, 0, size),
		index:    make(map[int]int16, size),
		outOrder: true,
	}
}

// MapIntInt16 is an ordered map int to int16.
type implMapIntInt16 struct {
	ordered  []MapIntInt16Entry
	index    map[int]int16
	outOrder bool
}

func (i implMapIntInt16) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapIntInt16) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i *implMapIntInt16) ToYAML() (*yaml.Node, error) {
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

func (i *implMapIntInt16) Put(k int, v int16) MapIntInt16 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapIntInt16Entry{k, v})
	return i
}

func (i *implMapIntInt16) PutIfNotNil(k int, v *int16) MapIntInt16 {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapIntInt16) ReplaceOrPut(k int, v int16) MapIntInt16 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapIntInt16) ReplaceIfExists(k int, v int16) MapIntInt16 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapIntInt16) Get(k int) (value int16, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapIntInt16) GetOpt(k int) option.Int16 {
	if v, ok := i.index[k]; ok {
		return option.NewInt16(v)
	}

	return option.NewEmptyInt16()
}

func (i *implMapIntInt16) At(j int) MapIntInt16Entry {
	return i.ordered[j]
}

func (i *implMapIntInt16) Len() int {
	return len(i.ordered)
}

func (i *implMapIntInt16) Has(k int) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapIntInt16) IndexOf(k int) int {
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

func (i *implMapIntInt16) Delete(k int) MapIntInt16 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapIntInt16) ForEach(f func(k int, v int16)) MapIntInt16 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implMapIntInt16) SerializeOrdered(b bool) MapIntInt16 {
	i.outOrder = b
	return i
}
