package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapIntRune defines an ordered map of int to rune.
type MapIntRune interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k int, v rune) MapIntRune

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k int, v *rune) MapIntRune

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k int, v rune) MapIntRune

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k int, v rune) MapIntRune

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k int) (value rune, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of rune if an entry was found
	// with the key `k`.
	GetOpt(k int) option.Rune

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapIntRuneEntry

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

	Delete(k int) MapIntRune

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k int, v rune)) MapIntRune
}

// MapIntRuneEntry is a single entry in an instance of
// MapIntRune.
type MapIntRuneEntry struct {
	Key int  `json:"key"`
	Val rune `json:"value"`
}

// NewMapIntRune creates a new instance of MapIntRune presized to the
// given size.
func NewMapIntRune(size int) MapIntRune {
	return &implMapIntRune{
		ordered: make([]MapIntRuneEntry, 0, size),
		index:   make(map[int]rune, size),
	}
}

// MapIntRune is an ordered map int to rune.
type implMapIntRune struct {
	ordered []MapIntRuneEntry
	index   map[int]rune
}

func (i implMapIntRune) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapIntRune) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapIntRune) ToYAML() (*yaml.Node, error) {
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

func (i *implMapIntRune) Put(k int, v rune) MapIntRune {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapIntRuneEntry{k, v})
	return i
}

func (i *implMapIntRune) PutIfNotNil(k int, v *rune) MapIntRune {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapIntRune) ReplaceOrPut(k int, v rune) MapIntRune {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapIntRune) ReplaceIfExists(k int, v rune) MapIntRune {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapIntRune) Get(k int) (value rune, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapIntRune) GetOpt(k int) option.Rune {
	if v, ok := i.index[k]; ok {
		return option.NewRune(v)
	}

	return option.NewEmptyRune()
}

func (i *implMapIntRune) At(j int) MapIntRuneEntry {
	return i.ordered[j]
}

func (i *implMapIntRune) Len() int {
	return len(i.ordered)
}

func (i *implMapIntRune) Has(k int) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapIntRune) IndexOf(k int) int {
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

func (i *implMapIntRune) Delete(k int) MapIntRune {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapIntRune) ForEach(f func(k int, v rune)) MapIntRune {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
