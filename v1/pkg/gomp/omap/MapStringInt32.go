package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

type MapStringInt32 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v int32) MapStringInt32

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *int32) MapStringInt32

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v int32) MapStringInt32

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v int32) MapStringInt32

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value int32, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of int32 if an entry was found
	// with the key `k`.
	GetOpt(k string) option.Int32

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapStringInt32Entry

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

	Delete(k string) MapStringInt32

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v int32)) MapStringInt32
}

type MapStringInt32Entry struct {
	Key string `json:"key"`
	Val int32  `json:"value"`
}

func NewMapStringInt32(size int) MapStringInt32 {
	return &implMapStringInt32{
		ordered: make([]MapStringInt32Entry, 0, size),
		index:   make(map[string]int32, size),
	}
}

// MapStringInt32 is an ordered map string to int32.
type implMapStringInt32 struct {
	ordered []MapStringInt32Entry
	index   map[string]int32
}

func (i implMapStringInt32) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapStringInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapStringInt32) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapStringInt32) Put(k string, v int32) MapStringInt32 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapStringInt32Entry{k, v})
	return i
}

func (i *implMapStringInt32) PutIfNotNil(k string, v *int32) MapStringInt32 {
	if v != nil {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapStringInt32) ReplaceOrPut(k string, v int32) MapStringInt32 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	} else {
		return i.Put(k, v)
	}
}

func (i *implMapStringInt32) ReplaceIfExists(k string, v int32) MapStringInt32 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapStringInt32) Get(k string) (value int32, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapStringInt32) GetOpt(k string) option.Int32 {
	if v, ok := i.index[k]; ok {
		return option.NewInt32(v)
	}

	return option.NewEmptyInt32()
}

func (i *implMapStringInt32) At(j int) MapStringInt32Entry {
	return i.ordered[j]
}

func (i *implMapStringInt32) Len() int {
	return len(i.ordered)
}

func (i *implMapStringInt32) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapStringInt32) IndexOf(k string) int {
	if _, ok := i.index[k]; ok {
		return -1
	}
	for j := range i.ordered {
		if i.ordered[j].Key == k {
			return j
		}
	}
	panic("invalid map state, out of sync")
}

func (i *implMapStringInt32) Delete(k string) MapStringInt32 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapStringInt32) ForEach(f func(k string, v int32)) MapStringInt32 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
