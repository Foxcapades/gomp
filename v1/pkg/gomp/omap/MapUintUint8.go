package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapUintUint8 defines an ordered map of uint to uint8.
type MapUintUint8 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k uint, v uint8) MapUintUint8

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k uint, v *uint8) MapUintUint8

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k uint, v uint8) MapUintUint8

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k uint, v uint8) MapUintUint8

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k uint) (value uint8, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of uint8 if an entry was found
	// with the key `k`.
	GetOpt(k uint) option.Uint8

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapUintUint8Entry

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

	Delete(k uint) MapUintUint8

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k uint, v uint8)) MapUintUint8
}

// MapUintUint8Entry is a single entry in an instance of
// MapUintUint8.
type MapUintUint8Entry struct {
	Key uint  `json:"key"`
	Val uint8 `json:"value"`
}

// NewMapUintUint8 creates a new instance of MapUintUint8 presized to the
// given size.
func NewMapUintUint8(size int) MapUintUint8 {
	return &implMapUintUint8{
		ordered: make([]MapUintUint8Entry, 0, size),
		index:   make(map[uint]uint8, size),
	}
}

// MapUintUint8 is an ordered map uint to uint8.
type implMapUintUint8 struct {
	ordered []MapUintUint8Entry
	index   map[uint]uint8
}

func (i implMapUintUint8) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapUintUint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapUintUint8) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapUintUint8) Put(k uint, v uint8) MapUintUint8 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapUintUint8Entry{k, v})
	return i
}

func (i *implMapUintUint8) PutIfNotNil(k uint, v *uint8) MapUintUint8 {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapUintUint8) ReplaceOrPut(k uint, v uint8) MapUintUint8 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapUintUint8) ReplaceIfExists(k uint, v uint8) MapUintUint8 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapUintUint8) Get(k uint) (value uint8, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapUintUint8) GetOpt(k uint) option.Uint8 {
	if v, ok := i.index[k]; ok {
		return option.NewUint8(v)
	}

	return option.NewEmptyUint8()
}

func (i *implMapUintUint8) At(j int) MapUintUint8Entry {
	return i.ordered[j]
}

func (i *implMapUintUint8) Len() int {
	return len(i.ordered)
}

func (i *implMapUintUint8) Has(k uint) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapUintUint8) IndexOf(k uint) int {
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

func (i *implMapUintUint8) Delete(k uint) MapUintUint8 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapUintUint8) ForEach(f func(k uint, v uint8)) MapUintUint8 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
