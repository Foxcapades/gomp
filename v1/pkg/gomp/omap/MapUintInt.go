package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapUintInt defines an ordered map of uint to int.
type MapUintInt interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k uint, v int) MapUintInt

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k uint, v *int) MapUintInt

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k uint, v int) MapUintInt

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k uint, v int) MapUintInt

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k uint) (value int, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of int if an entry was found
	// with the key `k`.
	GetOpt(k uint) option.Int

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapUintIntEntry

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

	Delete(k uint) MapUintInt

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k uint, v int)) MapUintInt
}

// MapUintIntEntry is a single entry in an instance of
// MapUintInt.
type MapUintIntEntry struct {
	Key uint `json:"key"`
	Val int  `json:"value"`
}

// NewMapUintInt creates a new instance of MapUintInt presized to the
// given size.
func NewMapUintInt(size int) MapUintInt {
	return &implMapUintInt{
		ordered: make([]MapUintIntEntry, 0, size),
		index:   make(map[uint]int, size),
	}
}

// MapUintInt is an ordered map uint to int.
type implMapUintInt struct {
	ordered []MapUintIntEntry
	index   map[uint]int
}

func (i implMapUintInt) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapUintInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapUintInt) ToYAML() (*yaml.Node, error) {
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

func (i *implMapUintInt) Put(k uint, v int) MapUintInt {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapUintIntEntry{k, v})
	return i
}

func (i *implMapUintInt) PutIfNotNil(k uint, v *int) MapUintInt {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapUintInt) ReplaceOrPut(k uint, v int) MapUintInt {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapUintInt) ReplaceIfExists(k uint, v int) MapUintInt {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapUintInt) Get(k uint) (value int, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapUintInt) GetOpt(k uint) option.Int {
	if v, ok := i.index[k]; ok {
		return option.NewInt(v)
	}

	return option.NewEmptyInt()
}

func (i *implMapUintInt) At(j int) MapUintIntEntry {
	return i.ordered[j]
}

func (i *implMapUintInt) Len() int {
	return len(i.ordered)
}

func (i *implMapUintInt) Has(k uint) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapUintInt) IndexOf(k uint) int {
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

func (i *implMapUintInt) Delete(k uint) MapUintInt {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapUintInt) ForEach(f func(k uint, v int)) MapUintInt {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
