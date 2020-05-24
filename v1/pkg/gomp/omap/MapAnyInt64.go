package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

type MapAnyInt64 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v int64) MapAnyInt64

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *int64) MapAnyInt64

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v int64) MapAnyInt64

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v int64) MapAnyInt64

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value int64, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of int64 if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Int64

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyInt64Entry

	// Len returns the current size of the map.
	Len() int

	// Has returns whether an entry exists with the key `k`.
	Has(k interface{}) bool

	// Has returns the position in the map of the entry matching key `k`.
	//
	// If no entry exists in the map with key `k` this method returns -1.
	//
	// Note: this method may iterate, at most once, through all the entries in the
	// map.
	IndexOf(k interface{}) int

	Delete(k interface{}) MapAnyInt64

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v int64)) MapAnyInt64
}

type MapAnyInt64Entry struct {
	Key interface{}  `json:"key"`
	Val int64 `json:"value"`
}

func NewMapAnyInt64(size int) MapAnyInt64 {
	return &implMapAnyInt64{
		ordered: make([]MapAnyInt64Entry, 0, size),
		index:   make(map[interface{}]int64, size),
	}
}

// MapAnyInt64 is an ordered map interface{} to int64.
type implMapAnyInt64 struct {
	ordered []MapAnyInt64Entry
	index   map[interface{}]int64
}

func (i implMapAnyInt64) MarshalYAML() (interface{},  error) {
	return i.ToYAML()
}

func (i implMapAnyInt64) MarshalJSON() ([]byte,  error) {
	return json.Marshal(i.ordered)
}

func (i *implMapAnyInt64) ToYAML() (*yaml.Node,  error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapAnyInt64) Put(k interface{}, v int64) MapAnyInt64 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyInt64Entry{k, v})
	return i
}

func (i *implMapAnyInt64) PutIfNotNil(k interface{}, v *int64) MapAnyInt64 {
	if v != nil {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAnyInt64) ReplaceOrPut(k interface{}, v int64) MapAnyInt64 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	} else {
		return i.Put(k, v)
	}
}

func (i *implMapAnyInt64) ReplaceIfExists(k interface{}, v int64) MapAnyInt64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAnyInt64) Get(k interface{}) (value int64, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAnyInt64) GetOpt(k interface{}) option.Int64 {
	if v, ok := i.index[k]; ok {
		return option.NewInt64(v)
	}

	return option.NewEmptyInt64()
}

func (i *implMapAnyInt64) At(j int) MapAnyInt64Entry {
	return i.ordered[j]
}

func (i *implMapAnyInt64) Len() int {
	return len(i.ordered)
}

func (i *implMapAnyInt64) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAnyInt64) IndexOf(k interface{}) int {
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

func (i *implMapAnyInt64) Delete(k interface{}) MapAnyInt64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAnyInt64) ForEach(f func(k interface{}, v int64)) MapAnyInt64 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}