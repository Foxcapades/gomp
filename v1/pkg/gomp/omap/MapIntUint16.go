package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

type MapIntUint16 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k int, v uint16) MapIntUint16

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k int, v *uint16) MapIntUint16

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k int, v uint16) MapIntUint16

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k int, v uint16) MapIntUint16

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k int) (value uint16, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of uint16 if an entry was found
	// with the key `k`.
	GetOpt(k int) option.Uint16

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapIntUint16Entry

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

	Delete(k int) MapIntUint16

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k int, v uint16)) MapIntUint16
}

type MapIntUint16Entry struct {
	Key int    `json:"key"`
	Val uint16 `json:"value"`
}

func NewMapIntUint16(size int) MapIntUint16 {
	return &implMapIntUint16{
		ordered: make([]MapIntUint16Entry, 0, size),
		index:   make(map[int]uint16, size),
	}
}

// MapIntUint16 is an ordered map int to uint16.
type implMapIntUint16 struct {
	ordered []MapIntUint16Entry
	index   map[int]uint16
}

func (i implMapIntUint16) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapIntUint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapIntUint16) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapIntUint16) Put(k int, v uint16) MapIntUint16 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapIntUint16Entry{k, v})
	return i
}

func (i *implMapIntUint16) PutIfNotNil(k int, v *uint16) MapIntUint16 {
	if v != nil {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapIntUint16) ReplaceOrPut(k int, v uint16) MapIntUint16 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	} else {
		return i.Put(k, v)
	}
}

func (i *implMapIntUint16) ReplaceIfExists(k int, v uint16) MapIntUint16 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapIntUint16) Get(k int) (value uint16, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapIntUint16) GetOpt(k int) option.Uint16 {
	if v, ok := i.index[k]; ok {
		return option.NewUint16(v)
	}

	return option.NewEmptyUint16()
}

func (i *implMapIntUint16) At(j int) MapIntUint16Entry {
	return i.ordered[j]
}

func (i *implMapIntUint16) Len() int {
	return len(i.ordered)
}

func (i *implMapIntUint16) Has(k int) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapIntUint16) IndexOf(k int) int {
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

func (i *implMapIntUint16) Delete(k int) MapIntUint16 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapIntUint16) ForEach(f func(k int, v uint16)) MapIntUint16 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
