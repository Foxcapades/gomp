package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapStringRune defines an ordered map of string to rune.
type MapStringRune interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k string, v rune) MapStringRune

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k string, v *rune) MapStringRune

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k string, v rune) MapStringRune

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k string, v rune) MapStringRune

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k string) (value rune, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of rune if an entry was found
	// with the key `k`.
	GetOpt(k string) option.Rune

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapStringRuneEntry

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

	Delete(k string) MapStringRune

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k string, v rune)) MapStringRune
}

// MapStringRuneEntry is a single entry in an instance of
// MapStringRune.
type MapStringRuneEntry struct {
	Key string `json:"key"`
	Val rune   `json:"value"`
}

// NewMapStringRune creates a new instance of MapStringRune presized to the
// given size.
func NewMapStringRune(size int) MapStringRune {
	return &implMapStringRune{
		ordered: make([]MapStringRuneEntry, 0, size),
		index:   make(map[string]rune, size),
	}
}

// MapStringRune is an ordered map string to rune.
type implMapStringRune struct {
	ordered []MapStringRuneEntry
	index   map[string]rune
}

func (i implMapStringRune) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapStringRune) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapStringRune) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapStringRune) Put(k string, v rune) MapStringRune {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapStringRuneEntry{k, v})
	return i
}

func (i *implMapStringRune) PutIfNotNil(k string, v *rune) MapStringRune {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapStringRune) ReplaceOrPut(k string, v rune) MapStringRune {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapStringRune) ReplaceIfExists(k string, v rune) MapStringRune {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapStringRune) Get(k string) (value rune, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapStringRune) GetOpt(k string) option.Rune {
	if v, ok := i.index[k]; ok {
		return option.NewRune(v)
	}

	return option.NewEmptyRune()
}

func (i *implMapStringRune) At(j int) MapStringRuneEntry {
	return i.ordered[j]
}

func (i *implMapStringRune) Len() int {
	return len(i.ordered)
}

func (i *implMapStringRune) Has(k string) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapStringRune) IndexOf(k string) int {
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

func (i *implMapStringRune) Delete(k string) MapStringRune {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapStringRune) ForEach(f func(k string, v rune)) MapStringRune {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
