package omap

import (
	"encoding/json"

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapIntString defines an ordered map of int to string.
type MapIntString interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k int, v string) MapIntString

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k int, v *string) MapIntString

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k int, v string) MapIntString

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k int, v string) MapIntString

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k int) (value string, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of string if an entry was found
	// with the key `k`.
	GetOpt(k int) option.String

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapIntStringEntry

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

	Delete(k int) MapIntString

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k int, v string)) MapIntString
}

// MapIntStringEntry is a single entry in an instance of
// MapIntString.
type MapIntStringEntry struct {
	Key int    `json:"key"`
	Val string `json:"value"`
}

// NewMapIntString creates a new instance of MapIntString presized to the
// given size.
func NewMapIntString(size int) MapIntString {
	return &implMapIntString{
		ordered: make([]MapIntStringEntry, 0, size),
		index:   make(map[int]string, size),
	}
}

// MapIntString is an ordered map int to string.
type implMapIntString struct {
	ordered []MapIntStringEntry
	index   map[int]string
}

func (i implMapIntString) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapIntString) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapIntString) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapIntString) Put(k int, v string) MapIntString {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapIntStringEntry{k, v})
	return i
}

func (i *implMapIntString) PutIfNotNil(k int, v *string) MapIntString {
	if v != nil {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapIntString) ReplaceOrPut(k int, v string) MapIntString {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapIntString) ReplaceIfExists(k int, v string) MapIntString {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapIntString) Get(k int) (value string, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapIntString) GetOpt(k int) option.String {
	if v, ok := i.index[k]; ok {
		return option.NewString(v)
	}

	return option.NewEmptyString()
}

func (i *implMapIntString) At(j int) MapIntStringEntry {
	return i.ordered[j]
}

func (i *implMapIntString) Len() int {
	return len(i.ordered)
}

func (i *implMapIntString) Has(k int) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapIntString) IndexOf(k int) int {
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

func (i *implMapIntString) Delete(k int) MapIntString {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapIntString) ForEach(f func(k int, v string)) MapIntString {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
