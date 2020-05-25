package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/internal/util"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapAnyFloat64 defines an ordered map of interface{} to float64.
type MapAnyFloat64 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v float64) MapAnyFloat64

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *float64) MapAnyFloat64

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v float64) MapAnyFloat64

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v float64) MapAnyFloat64

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value float64, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of float64 if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Float64

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyFloat64Entry

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

	Delete(k interface{}) MapAnyFloat64

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v float64)) MapAnyFloat64
}

// MapAnyFloat64Entry is a single entry in an instance of
// MapAnyFloat64.
type MapAnyFloat64Entry struct {
	Key interface{} `json:"key"`
	Val float64     `json:"value"`
}

// NewMapAnyFloat64 creates a new instance of MapAnyFloat64 presized to the
// given size.
func NewMapAnyFloat64(size int) MapAnyFloat64 {
	return &implMapAnyFloat64{
		ordered: make([]MapAnyFloat64Entry, 0, size),
		index:   make(map[interface{}]float64, size),
	}
}

// MapAnyFloat64 is an ordered map interface{} to float64.
type implMapAnyFloat64 struct {
	ordered []MapAnyFloat64Entry
	index   map[interface{}]float64
}

func (i implMapAnyFloat64) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapAnyFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapAnyFloat64) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapAnyFloat64) Put(k interface{}, v float64) MapAnyFloat64 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyFloat64Entry{k, v})
	return i
}

func (i *implMapAnyFloat64) PutIfNotNil(k interface{}, v *float64) MapAnyFloat64 {
	if !util.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAnyFloat64) ReplaceOrPut(k interface{}, v float64) MapAnyFloat64 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapAnyFloat64) ReplaceIfExists(k interface{}, v float64) MapAnyFloat64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAnyFloat64) Get(k interface{}) (value float64, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAnyFloat64) GetOpt(k interface{}) option.Float64 {
	if v, ok := i.index[k]; ok {
		return option.NewFloat64(v)
	}

	return option.NewEmptyFloat64()
}

func (i *implMapAnyFloat64) At(j int) MapAnyFloat64Entry {
	return i.ordered[j]
}

func (i *implMapAnyFloat64) Len() int {
	return len(i.ordered)
}

func (i *implMapAnyFloat64) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAnyFloat64) IndexOf(k interface{}) int {
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

func (i *implMapAnyFloat64) Delete(k interface{}) MapAnyFloat64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAnyFloat64) ForEach(f func(k interface{}, v float64)) MapAnyFloat64 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
