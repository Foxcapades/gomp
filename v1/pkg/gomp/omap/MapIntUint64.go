package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapIntUint64 defines an ordered map of int to uint64.
type MapIntUint64 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k int, v uint64) MapIntUint64

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k int, v *uint64) MapIntUint64

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k int, v uint64) MapIntUint64

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k int, v uint64) MapIntUint64

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k int) (value uint64, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of uint64 if an entry was found
	// with the key `k`.
	GetOpt(k int) option.Uint64

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapIntUint64Entry

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

	Delete(k int) MapIntUint64

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k int, v uint64)) MapIntUint64
}

type MapIntUint64Entry struct {
	Key int    `json:"key"`
	Val uint64 `json:"value"`
}

func NewMapIntUint64(size int) MapIntUint64 {
	return &implMapIntUint64{
		ordered: make([]MapIntUint64Entry, 0, size),
		index:   make(map[int]uint64, size),
	}
}

// MapIntUint64 is an ordered map int to uint64.
type implMapIntUint64 struct {
	ordered []MapIntUint64Entry
	index   map[int]uint64
}

func (i implMapIntUint64) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapIntUint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapIntUint64) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapIntUint64) Put(k int, v uint64) MapIntUint64 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapIntUint64Entry{k, v})
	return i
}

func (i *implMapIntUint64) PutIfNotNil(k int, v *uint64) MapIntUint64 {
	if v != nil {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapIntUint64) ReplaceOrPut(k int, v uint64) MapIntUint64 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	} else {
		return i.Put(k, v)
	}
}

func (i *implMapIntUint64) ReplaceIfExists(k int, v uint64) MapIntUint64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapIntUint64) Get(k int) (value uint64, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapIntUint64) GetOpt(k int) option.Uint64 {
	if v, ok := i.index[k]; ok {
		return option.NewUint64(v)
	}

	return option.NewEmptyUint64()
}

func (i *implMapIntUint64) At(j int) MapIntUint64Entry {
	return i.ordered[j]
}

func (i *implMapIntUint64) Len() int {
	return len(i.ordered)
}

func (i *implMapIntUint64) Has(k int) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapIntUint64) IndexOf(k int) int {
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

func (i *implMapIntUint64) Delete(k int) MapIntUint64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapIntUint64) ForEach(f func(k int, v uint64)) MapIntUint64 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
