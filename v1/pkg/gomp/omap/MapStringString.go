package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapStringString defines an ordered map of string to string.
type MapStringString interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v string) MapStringString

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *string) MapStringString

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v string) MapStringString

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v string) MapStringString

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value string, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of string if an entry was found
	// with the key `k`.
	GetOpt(k string) option.String

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapStringStringEntry

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

	Delete(k string) MapStringString

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v string)) MapStringString
}

type MapStringStringEntry struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

func NewMapStringString(size int) MapStringString {
	return &implMapStringString{
		ordered: make([]MapStringStringEntry, 0, size),
		index:   make(map[string]string, size),
	}
}

// MapStringString is an ordered map string to string.
type implMapStringString struct {
	ordered []MapStringStringEntry
	index   map[string]string
}

func (i implMapStringString) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapStringString) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapStringString) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapStringString) Put(k string, v string) MapStringString {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapStringStringEntry{k, v})
	return i
}

func (i *implMapStringString) PutIfNotNil(k string, v *string) MapStringString {
	if v != nil {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapStringString) ReplaceOrPut(k string, v string) MapStringString {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	} else {
		return i.Put(k, v)
	}
}

func (i *implMapStringString) ReplaceIfExists(k string, v string) MapStringString {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapStringString) Get(k string) (value string, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapStringString) GetOpt(k string) option.String {
	if v, ok := i.index[k]; ok {
		return option.NewString(v)
	}

	return option.NewEmptyString()
}

func (i *implMapStringString) At(j int) MapStringStringEntry {
	return i.ordered[j]
}

func (i *implMapStringString) Len() int {
	return len(i.ordered)
}

func (i *implMapStringString) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapStringString) IndexOf(k string) int {
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

func (i *implMapStringString) Delete(k string) MapStringString {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapStringString) ForEach(f func(k string, v string)) MapStringString {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
