package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapAny defines an ordered map of interface{} to interface{}.
type MapAny interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v interface{}) MapAny

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *interface{}) MapAny

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v interface{}) MapAny

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v interface{}) MapAny

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value interface{}, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of interface{} if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Untyped

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyEntry

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

	Delete(k interface{}) MapAny

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v interface{})) MapAny
}

// MapAnyEntry is a single entry in an instance of
// MapAny.
type MapAnyEntry struct {
	Key interface{} `json:"key"`
	Val interface{} `json:"value"`
}

// NewMapAny creates a new instance of MapAny presized to the
// given size.
func NewMapAny(size int) MapAny {
	return &implMapAny{
		ordered: make([]MapAnyEntry, 0, size),
		index:   make(map[interface{}]interface{}, size),
	}
}

// MapAny is an ordered map interface{} to interface{}.
type implMapAny struct {
	ordered []MapAnyEntry
	index   map[interface{}]interface{}
}

func (i implMapAny) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapAny) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapAny) Put(k interface{}, v interface{}) MapAny {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyEntry{k, v})
	return i
}

func (i *implMapAny) PutIfNotNil(k interface{}, v *interface{}) MapAny {
	if v != nil {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAny) ReplaceOrPut(k interface{}, v interface{}) MapAny {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapAny) ReplaceIfExists(k interface{}, v interface{}) MapAny {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAny) Get(k interface{}) (value interface{}, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAny) GetOpt(k interface{}) option.Untyped {
	if v, ok := i.index[k]; ok {
		return option.NewUntyped(v)
	}

	return option.NewEmptyUntyped()
}

func (i *implMapAny) At(j int) MapAnyEntry {
	return i.ordered[j]
}

func (i *implMapAny) Len() int {
	return len(i.ordered)
}

func (i *implMapAny) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAny) IndexOf(k interface{}) int {
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

func (i *implMapAny) Delete(k interface{}) MapAny {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAny) ForEach(f func(k interface{}, v interface{})) MapAny {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
