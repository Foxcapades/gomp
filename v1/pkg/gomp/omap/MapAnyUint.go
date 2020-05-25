package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapAnyUint defines an ordered map of interface{} to uint.
type MapAnyUint interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v uint) MapAnyUint

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *uint) MapAnyUint

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v uint) MapAnyUint

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v uint) MapAnyUint

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value uint, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of uint if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Uint

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyUintEntry

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

	Delete(k interface{}) MapAnyUint

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v uint)) MapAnyUint
}

// MapAnyUintEntry is a single entry in an instance of
// MapAnyUint.
type MapAnyUintEntry struct {
	Key interface{} `json:"key"`
	Val uint        `json:"value"`
}

// NewMapAnyUint creates a new instance of MapAnyUint presized to the
// given size.
func NewMapAnyUint(size int) MapAnyUint {
	return &implMapAnyUint{
		ordered: make([]MapAnyUintEntry, 0, size),
		index:   make(map[interface{}]uint, size),
	}
}

// MapAnyUint is an ordered map interface{} to uint.
type implMapAnyUint struct {
	ordered []MapAnyUintEntry
	index   map[interface{}]uint
}

func (i implMapAnyUint) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapAnyUint) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapAnyUint) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapAnyUint) Put(k interface{}, v uint) MapAnyUint {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyUintEntry{k, v})
	return i
}

func (i *implMapAnyUint) PutIfNotNil(k interface{}, v *uint) MapAnyUint {
	if !IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAnyUint) ReplaceOrPut(k interface{}, v uint) MapAnyUint {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapAnyUint) ReplaceIfExists(k interface{}, v uint) MapAnyUint {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAnyUint) Get(k interface{}) (value uint, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAnyUint) GetOpt(k interface{}) option.Uint {
	if v, ok := i.index[k]; ok {
		return option.NewUint(v)
	}

	return option.NewEmptyUint()
}

func (i *implMapAnyUint) At(j int) MapAnyUintEntry {
	return i.ordered[j]
}

func (i *implMapAnyUint) Len() int {
	return len(i.ordered)
}

func (i *implMapAnyUint) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAnyUint) IndexOf(k interface{}) int {
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

func (i *implMapAnyUint) Delete(k interface{}) MapAnyUint {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAnyUint) ForEach(f func(k interface{}, v uint)) MapAnyUint {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
