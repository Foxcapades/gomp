package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapUintUint16 defines an ordered map of uint to uint16.
type MapUintUint16 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k uint, v uint16) MapUintUint16

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k uint, v *uint16) MapUintUint16

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k uint, v uint16) MapUintUint16

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k uint, v uint16) MapUintUint16

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k uint) (value uint16, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of uint16 if an entry was found
	// with the key `k`.
	GetOpt(k uint) option.Uint16

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapUintUint16Entry

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

	Delete(k uint) MapUintUint16

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k uint, v uint16)) MapUintUint16
}

// MapUintUint16Entry is a single entry in an instance of
// MapUintUint16.
type MapUintUint16Entry struct {
	Key uint   `json:"key"`
	Val uint16 `json:"value"`
}

// NewMapUintUint16 creates a new instance of MapUintUint16 presized to the
// given size.
func NewMapUintUint16(size int) MapUintUint16 {
	return &implMapUintUint16{
		ordered: make([]MapUintUint16Entry, 0, size),
		index:   make(map[uint]uint16, size),
	}
}

// MapUintUint16 is an ordered map uint to uint16.
type implMapUintUint16 struct {
	ordered []MapUintUint16Entry
	index   map[uint]uint16
}

func (i implMapUintUint16) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapUintUint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapUintUint16) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapUintUint16) Put(k uint, v uint16) MapUintUint16 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapUintUint16Entry{k, v})
	return i
}

func (i *implMapUintUint16) PutIfNotNil(k uint, v *uint16) MapUintUint16 {
	if !IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapUintUint16) ReplaceOrPut(k uint, v uint16) MapUintUint16 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapUintUint16) ReplaceIfExists(k uint, v uint16) MapUintUint16 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapUintUint16) Get(k uint) (value uint16, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapUintUint16) GetOpt(k uint) option.Uint16 {
	if v, ok := i.index[k]; ok {
		return option.NewUint16(v)
	}

	return option.NewEmptyUint16()
}

func (i *implMapUintUint16) At(j int) MapUintUint16Entry {
	return i.ordered[j]
}

func (i *implMapUintUint16) Len() int {
	return len(i.ordered)
}

func (i *implMapUintUint16) Has(k uint) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapUintUint16) IndexOf(k uint) int {
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

func (i *implMapUintUint16) Delete(k uint) MapUintUint16 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapUintUint16) ForEach(f func(k uint, v uint16)) MapUintUint16 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
