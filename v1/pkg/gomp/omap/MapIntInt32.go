package omap

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapIntInt32 defines an ordered map of int to int32.
type MapIntInt32 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k int, v int32) MapIntInt32

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k int, v *int32) MapIntInt32

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k int, v int32) MapIntInt32

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k int, v int32) MapIntInt32

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k int) (value int32, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of int32 if an entry was found
	// with the key `k`.
	GetOpt(k int) option.Int32

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapIntInt32Entry

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

	Delete(k int) MapIntInt32

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k int, v int32)) MapIntInt32

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) MapIntInt32
}

// MapIntInt32Entry is a single entry in an instance of
// MapIntInt32.
type MapIntInt32Entry struct {
	Key int   `json:"key"`
	Val int32 `json:"value"`
}

// NewMapIntInt32 creates a new instance of MapIntInt32 presized to the
// given size.
func NewMapIntInt32(size int) MapIntInt32 {
	return &implMapIntInt32{
		ordered:  make([]MapIntInt32Entry, 0, size),
		index:    make(map[int]int32, size),
		outOrder: true,
	}
}

// MapIntInt32 is an ordered map int to int32.
type implMapIntInt32 struct {
	ordered  []MapIntInt32Entry
	index    map[int]int32
	outOrder bool
}

func (i implMapIntInt32) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapIntInt32) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i implMapIntInt32) ToYAML() (*yaml.Node, error) {
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

func (i *implMapIntInt32) Put(k int, v int32) MapIntInt32 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapIntInt32Entry{k, v})
	return i
}

func (i *implMapIntInt32) PutIfNotNil(k int, v *int32) MapIntInt32 {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapIntInt32) ReplaceOrPut(k int, v int32) MapIntInt32 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapIntInt32) ReplaceIfExists(k int, v int32) MapIntInt32 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapIntInt32) Get(k int) (value int32, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapIntInt32) GetOpt(k int) option.Int32 {
	if v, ok := i.index[k]; ok {
		return option.NewInt32(v)
	}

	return option.NewEmptyInt32()
}

func (i *implMapIntInt32) At(j int) MapIntInt32Entry {
	return i.ordered[j]
}

func (i *implMapIntInt32) Len() int {
	return len(i.ordered)
}

func (i *implMapIntInt32) Has(k int) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapIntInt32) IndexOf(k int) int {
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

func (i *implMapIntInt32) Delete(k int) MapIntInt32 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapIntInt32) ForEach(f func(k int, v int32)) MapIntInt32 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implMapIntInt32) SerializeOrdered(b bool) MapIntInt32 {
	i.outOrder = b
	return i
}
