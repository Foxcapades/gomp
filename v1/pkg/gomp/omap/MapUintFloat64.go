package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapUintFloat64 defines an ordered map of uint to float64.
type MapUintFloat64 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k uint, v float64) MapUintFloat64

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k uint, v *float64) MapUintFloat64

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k uint, v float64) MapUintFloat64

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k uint, v float64) MapUintFloat64

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k uint) (value float64, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of float64 if an entry was found
	// with the key `k`.
	GetOpt(k uint) option.Float64

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapUintFloat64Entry

	// Len returns the current size of the map.
	Len() int

	// Has returns whether an entry exists with the key `k`.
	Has(k uint) bool

	// Has returns the position in the map of the entry matching key `k`.
	//
	// If no entry exists in the map with key `k` this method returns -1.
	//
	// Note: this method may iterate, at most once, through all the entries in the
	// map.
	IndexOf(k uint) int

	Delete(k uint) MapUintFloat64

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k uint, v float64)) MapUintFloat64
}

// MapUintFloat64Entry is a single entry in an instance of
// MapUintFloat64.
type MapUintFloat64Entry struct {
	Key uint    `json:"key"`
	Val float64 `json:"value"`
}

// NewMapUintFloat64 creates a new instance of MapUintFloat64 presized to the
// given size.
func NewMapUintFloat64(size int) MapUintFloat64 {
	return &implMapUintFloat64{
		ordered: make([]MapUintFloat64Entry, 0, size),
		index:   make(map[uint]float64, size),
	}
}

// MapUintFloat64 is an ordered map uint to float64.
type implMapUintFloat64 struct {
	ordered []MapUintFloat64Entry
	index   map[uint]float64
}

func (i implMapUintFloat64) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapUintFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapUintFloat64) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		tmp := xyml.NewMapNode(1)
		_ = xyml.MapAppend(tmp, i.ordered[j].Key, i.ordered[j].Val)
		if err := xyml.SequenceAppend(out, tmp); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapUintFloat64) Put(k uint, v float64) MapUintFloat64 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapUintFloat64Entry{k, v})
	return i
}

func (i *implMapUintFloat64) PutIfNotNil(k uint, v *float64) MapUintFloat64 {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapUintFloat64) ReplaceOrPut(k uint, v float64) MapUintFloat64 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapUintFloat64) ReplaceIfExists(k uint, v float64) MapUintFloat64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapUintFloat64) Get(k uint) (value float64, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapUintFloat64) GetOpt(k uint) option.Float64 {
	if v, ok := i.index[k]; ok {
		return option.NewFloat64(v)
	}

	return option.NewEmptyFloat64()
}

func (i *implMapUintFloat64) At(j int) MapUintFloat64Entry {
	return i.ordered[j]
}

func (i *implMapUintFloat64) Len() int {
	return len(i.ordered)
}

func (i *implMapUintFloat64) Has(k uint) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapUintFloat64) IndexOf(k uint) int {
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

func (i *implMapUintFloat64) Delete(k uint) MapUintFloat64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapUintFloat64) ForEach(f func(k uint, v float64)) MapUintFloat64 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
