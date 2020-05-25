package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapAnyInt8 defines an ordered map of interface{} to int8.
type MapAnyInt8 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v int8) MapAnyInt8

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *int8) MapAnyInt8

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v int8) MapAnyInt8

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v int8) MapAnyInt8

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value int8, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of int8 if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Int8

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyInt8Entry

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

	Delete(k interface{}) MapAnyInt8

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v int8)) MapAnyInt8
}

// MapAnyInt8Entry is a single entry in an instance of
// MapAnyInt8.
type MapAnyInt8Entry struct {
	Key interface{} `json:"key"`
	Val int8        `json:"value"`
}

// NewMapAnyInt8 creates a new instance of MapAnyInt8 presized to the
// given size.
func NewMapAnyInt8(size int) MapAnyInt8 {
	return &implMapAnyInt8{
		ordered: make([]MapAnyInt8Entry, 0, size),
		index:   make(map[interface{}]int8, size),
	}
}

// MapAnyInt8 is an ordered map interface{} to int8.
type implMapAnyInt8 struct {
	ordered []MapAnyInt8Entry
	index   map[interface{}]int8
}

func (i implMapAnyInt8) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapAnyInt8) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapAnyInt8) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapAnyInt8) Put(k interface{}, v int8) MapAnyInt8 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyInt8Entry{k, v})
	return i
}

func (i *implMapAnyInt8) PutIfNotNil(k interface{}, v *int8) MapAnyInt8 {
	if !IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAnyInt8) ReplaceOrPut(k interface{}, v int8) MapAnyInt8 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapAnyInt8) ReplaceIfExists(k interface{}, v int8) MapAnyInt8 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAnyInt8) Get(k interface{}) (value int8, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAnyInt8) GetOpt(k interface{}) option.Int8 {
	if v, ok := i.index[k]; ok {
		return option.NewInt8(v)
	}

	return option.NewEmptyInt8()
}

func (i *implMapAnyInt8) At(j int) MapAnyInt8Entry {
	return i.ordered[j]
}

func (i *implMapAnyInt8) Len() int {
	return len(i.ordered)
}

func (i *implMapAnyInt8) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAnyInt8) IndexOf(k interface{}) int {
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

func (i *implMapAnyInt8) Delete(k interface{}) MapAnyInt8 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAnyInt8) ForEach(f func(k interface{}, v int8)) MapAnyInt8 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
