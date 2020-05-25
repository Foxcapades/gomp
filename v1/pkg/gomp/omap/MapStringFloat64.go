package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapStringFloat64 defines an ordered map of string to float64.
type MapStringFloat64 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v float64) MapStringFloat64

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *float64) MapStringFloat64

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v float64) MapStringFloat64

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v float64) MapStringFloat64

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value float64, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of float64 if an entry was found
	// with the key `k`.
	GetOpt(k string) option.Float64

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapStringFloat64Entry

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

	Delete(k string) MapStringFloat64

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v float64)) MapStringFloat64
}

// MapStringFloat64Entry is a single entry in an instance of
// MapStringFloat64.
type MapStringFloat64Entry struct {
	Key string  `json:"key"`
	Val float64 `json:"value"`
}

// NewMapStringFloat64 creates a new instance of MapStringFloat64 presized to the
// given size.
func NewMapStringFloat64(size int) MapStringFloat64 {
	return &implMapStringFloat64{
		ordered: make([]MapStringFloat64Entry, 0, size),
		index:   make(map[string]float64, size),
	}
}

// MapStringFloat64 is an ordered map string to float64.
type implMapStringFloat64 struct {
	ordered []MapStringFloat64Entry
	index   map[string]float64
}

func (i implMapStringFloat64) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapStringFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapStringFloat64) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapStringFloat64) Put(k string, v float64) MapStringFloat64 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapStringFloat64Entry{k, v})
	return i
}

func (i *implMapStringFloat64) PutIfNotNil(k string, v *float64) MapStringFloat64 {
	if !IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapStringFloat64) ReplaceOrPut(k string, v float64) MapStringFloat64 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapStringFloat64) ReplaceIfExists(k string, v float64) MapStringFloat64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapStringFloat64) Get(k string) (value float64, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapStringFloat64) GetOpt(k string) option.Float64 {
	if v, ok := i.index[k]; ok {
		return option.NewFloat64(v)
	}

	return option.NewEmptyFloat64()
}

func (i *implMapStringFloat64) At(j int) MapStringFloat64Entry {
	return i.ordered[j]
}

func (i *implMapStringFloat64) Len() int {
	return len(i.ordered)
}

func (i *implMapStringFloat64) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapStringFloat64) IndexOf(k string) int {
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

func (i *implMapStringFloat64) Delete(k string) MapStringFloat64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapStringFloat64) ForEach(f func(k string, v float64)) MapStringFloat64 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
