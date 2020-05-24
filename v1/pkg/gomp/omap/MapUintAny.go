package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapUintAny defines an ordered map of uint to interface{}.
type MapUintAny interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k uint, v interface{}) MapUintAny

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k uint, v *interface{}) MapUintAny

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k uint, v interface{}) MapUintAny

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k uint, v interface{}) MapUintAny

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k uint) (value interface{}, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of interface{} if an entry was found
	// with the key `k`.
	GetOpt(k uint) option.Untyped

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapUintAnyEntry

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

	Delete(k uint) MapUintAny

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k uint, v interface{})) MapUintAny
}

// MapUintAnyEntry is a single entry in an instance of
// MapUintAny.
type MapUintAnyEntry struct {
	Key uint        `json:"key"`
	Val interface{} `json:"value"`
}

// NewMapUintAny creates a new instance of MapUintAny presized to the
// given size.
func NewMapUintAny(size int) MapUintAny {
	return &implMapUintAny{
		ordered: make([]MapUintAnyEntry, 0, size),
		index:   make(map[uint]interface{}, size),
	}
}

// MapUintAny is an ordered map uint to interface{}.
type implMapUintAny struct {
	ordered []MapUintAnyEntry
	index   map[uint]interface{}
}

func (i implMapUintAny) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapUintAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapUintAny) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapUintAny) Put(k uint, v interface{}) MapUintAny {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapUintAnyEntry{k, v})
	return i
}

func (i *implMapUintAny) PutIfNotNil(k uint, v *interface{}) MapUintAny {
	if v != nil {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapUintAny) ReplaceOrPut(k uint, v interface{}) MapUintAny {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	} else {
		return i.Put(k, v)
	}
}

func (i *implMapUintAny) ReplaceIfExists(k uint, v interface{}) MapUintAny {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapUintAny) Get(k uint) (value interface{}, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapUintAny) GetOpt(k uint) option.Untyped {
	if v, ok := i.index[k]; ok {
		return option.NewUntyped(v)
	}

	return option.NewEmptyUntyped()
}

func (i *implMapUintAny) At(j int) MapUintAnyEntry {
	return i.ordered[j]
}

func (i *implMapUintAny) Len() int {
	return len(i.ordered)
}

func (i *implMapUintAny) Has(k uint) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapUintAny) IndexOf(k uint) int {
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

func (i *implMapUintAny) Delete(k uint) MapUintAny {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapUintAny) ForEach(f func(k uint, v interface{})) MapUintAny {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
