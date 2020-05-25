package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/internal/util"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapStringAny defines an ordered map of string to interface{}.
type MapStringAny interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v interface{}) MapStringAny

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v interface{}) MapStringAny

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v interface{}) MapStringAny

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v interface{}) MapStringAny

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value interface{}, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of interface{} if an entry was found
	// with the key `k`.
	GetOpt(k string) option.Untyped

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapStringAnyEntry

	// Len returns the current size of the map.
	Len() int

	// Has returns whether an entry exists with the key `k`.
	Has(k string) bool

	// Has returns the position in the map of the entry matching key `k`.
	//
	// If no entry exists in the map with key `k` this method returns -1.
	//
	// Note: this method may iterate, at most once, through all the entries in the
	// map.
	IndexOf(k string) int

	Delete(k string) MapStringAny

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v interface{})) MapStringAny
}

// MapStringAnyEntry is a single entry in an instance of
// MapStringAny.
type MapStringAnyEntry struct {
	Key string      `json:"key"`
	Val interface{} `json:"value"`
}

// NewMapStringAny creates a new instance of MapStringAny presized to the
// given size.
func NewMapStringAny(size int) MapStringAny {
	return &implMapStringAny{
		ordered: make([]MapStringAnyEntry, 0, size),
		index:   make(map[string]interface{}, size),
	}
}

// MapStringAny is an ordered map string to interface{}.
type implMapStringAny struct {
	ordered []MapStringAnyEntry
	index   map[string]interface{}
}

func (i implMapStringAny) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapStringAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapStringAny) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapStringAny) Put(k string, v interface{}) MapStringAny {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapStringAnyEntry{k, v})
	return i
}

func (i *implMapStringAny) PutIfNotNil(k string, v interface{}) MapStringAny {
	if !util.IsNil(v) {
		return i.Put(k, util.Deref(v))
	}

	return i
}

func (i *implMapStringAny) ReplaceOrPut(k string, v interface{}) MapStringAny {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapStringAny) ReplaceIfExists(k string, v interface{}) MapStringAny {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapStringAny) Get(k string) (value interface{}, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapStringAny) GetOpt(k string) option.Untyped {
	if v, ok := i.index[k]; ok {
		return option.NewUntyped(v)
	}

	return option.NewEmptyUntyped()
}

func (i *implMapStringAny) At(j int) MapStringAnyEntry {
	return i.ordered[j]
}

func (i *implMapStringAny) Len() int {
	return len(i.ordered)
}

func (i *implMapStringAny) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapStringAny) IndexOf(k string) int {
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

func (i *implMapStringAny) Delete(k string) MapStringAny {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapStringAny) ForEach(f func(k string, v interface{})) MapStringAny {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
