package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/internal/util"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapAnyBool defines an ordered map of interface{} to bool.
type MapAnyBool interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v bool) MapAnyBool

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *bool) MapAnyBool

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v bool) MapAnyBool

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v bool) MapAnyBool

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value bool, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of bool if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Bool

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyBoolEntry

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

	Delete(k interface{}) MapAnyBool

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v bool)) MapAnyBool
}

// MapAnyBoolEntry is a single entry in an instance of
// MapAnyBool.
type MapAnyBoolEntry struct {
	Key interface{} `json:"key"`
	Val bool        `json:"value"`
}

// NewMapAnyBool creates a new instance of MapAnyBool presized to the
// given size.
func NewMapAnyBool(size int) MapAnyBool {
	return &implMapAnyBool{
		ordered: make([]MapAnyBoolEntry, 0, size),
		index:   make(map[interface{}]bool, size),
	}
}

// MapAnyBool is an ordered map interface{} to bool.
type implMapAnyBool struct {
	ordered []MapAnyBoolEntry
	index   map[interface{}]bool
}

func (i implMapAnyBool) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapAnyBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapAnyBool) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapAnyBool) Put(k interface{}, v bool) MapAnyBool {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyBoolEntry{k, v})
	return i
}

func (i *implMapAnyBool) PutIfNotNil(k interface{}, v *bool) MapAnyBool {
	if !util.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAnyBool) ReplaceOrPut(k interface{}, v bool) MapAnyBool {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapAnyBool) ReplaceIfExists(k interface{}, v bool) MapAnyBool {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAnyBool) Get(k interface{}) (value bool, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAnyBool) GetOpt(k interface{}) option.Bool {
	if v, ok := i.index[k]; ok {
		return option.NewBool(v)
	}

	return option.NewEmptyBool()
}

func (i *implMapAnyBool) At(j int) MapAnyBoolEntry {
	return i.ordered[j]
}

func (i *implMapAnyBool) Len() int {
	return len(i.ordered)
}

func (i *implMapAnyBool) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAnyBool) IndexOf(k interface{}) int {
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

func (i *implMapAnyBool) Delete(k interface{}) MapAnyBool {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAnyBool) ForEach(f func(k interface{}, v bool)) MapAnyBool {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
