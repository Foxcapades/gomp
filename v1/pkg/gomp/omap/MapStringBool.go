package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapStringBool defines an ordered map of string to bool.
type MapStringBool interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v bool) MapStringBool

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *bool) MapStringBool

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v bool) MapStringBool

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v bool) MapStringBool

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value bool, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of bool if an entry was found
	// with the key `k`.
	GetOpt(k string) option.Bool

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapStringBoolEntry

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

	Delete(k string) MapStringBool

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v bool)) MapStringBool
}

// MapStringBoolEntry is a single entry in an instance of
// MapStringBool.
type MapStringBoolEntry struct {
	Key string `json:"key"`
	Val bool   `json:"value"`
}

// NewMapStringBool creates a new instance of MapStringBool presized to the
// given size.
func NewMapStringBool(size int) MapStringBool {
	return &implMapStringBool{
		ordered: make([]MapStringBoolEntry, 0, size),
		index:   make(map[string]bool, size),
	}
}

// MapStringBool is an ordered map string to bool.
type implMapStringBool struct {
	ordered []MapStringBoolEntry
	index   map[string]bool
}

func (i implMapStringBool) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapStringBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapStringBool) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapStringBool) Put(k string, v bool) MapStringBool {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapStringBoolEntry{k, v})
	return i
}

func (i *implMapStringBool) PutIfNotNil(k string, v *bool) MapStringBool {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapStringBool) ReplaceOrPut(k string, v bool) MapStringBool {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapStringBool) ReplaceIfExists(k string, v bool) MapStringBool {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapStringBool) Get(k string) (value bool, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapStringBool) GetOpt(k string) option.Bool {
	if v, ok := i.index[k]; ok {
		return option.NewBool(v)
	}

	return option.NewEmptyBool()
}

func (i *implMapStringBool) At(j int) MapStringBoolEntry {
	return i.ordered[j]
}

func (i *implMapStringBool) Len() int {
	return len(i.ordered)
}

func (i *implMapStringBool) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapStringBool) IndexOf(k string) int {
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

func (i *implMapStringBool) Delete(k string) MapStringBool {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapStringBool) ForEach(f func(k string, v bool)) MapStringBool {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
