package omap

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapAnyRune defines an ordered map of interface{} to rune.
type MapAnyRune interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v rune) MapAnyRune

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *rune) MapAnyRune

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v rune) MapAnyRune

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v rune) MapAnyRune

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value rune, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of rune if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Rune

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyRuneEntry

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

	Delete(k interface{}) MapAnyRune

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v rune)) MapAnyRune

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) MapAnyRune
}

// MapAnyRuneEntry is a single entry in an instance of
// MapAnyRune.
type MapAnyRuneEntry struct {
	Key interface{} `json:"key"`
	Val rune        `json:"value"`
}

// NewMapAnyRune creates a new instance of MapAnyRune presized to the
// given size.
func NewMapAnyRune(size int) MapAnyRune {
	return &implMapAnyRune{
		ordered:  make([]MapAnyRuneEntry, 0, size),
		index:    make(map[interface{}]rune, size),
		outOrder: true,
	}
}

// MapAnyRune is an ordered map interface{} to rune.
type implMapAnyRune struct {
	ordered  []MapAnyRuneEntry
	index    map[interface{}]rune
	outOrder bool
}

func (i implMapAnyRune) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapAnyRune) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i implMapAnyRune) ToYAML() (*yaml.Node, error) {
	if i.outOrder {
		out := xyml.NewOrderedMapNode(i.Len())

		for j := range i.ordered {
			tmp := xyml.NewMapNode(1)
			if e := xyml.MapAppend(tmp, i.ordered[j].Key, i.ordered[j].Val); e != nil {
				return nil, e
			}
			if err := xyml.SequenceAppend(out, tmp); err != nil {
				return nil, err
			}
		}

		return out, nil
	}

	out := xyml.NewMapNode(i.Len())

	for j := range i.ordered {
		if e := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); e != nil {
			return nil, e
		}
	}

	return out, nil
}

func (i *implMapAnyRune) Put(k interface{}, v rune) MapAnyRune {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyRuneEntry{k, v})
	return i
}

func (i *implMapAnyRune) PutIfNotNil(k interface{}, v *rune) MapAnyRune {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAnyRune) ReplaceOrPut(k interface{}, v rune) MapAnyRune {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapAnyRune) ReplaceIfExists(k interface{}, v rune) MapAnyRune {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAnyRune) Get(k interface{}) (value rune, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAnyRune) GetOpt(k interface{}) option.Rune {
	if v, ok := i.index[k]; ok {
		return option.NewRune(v)
	}

	return option.NewEmptyRune()
}

func (i *implMapAnyRune) At(j int) MapAnyRuneEntry {
	return i.ordered[j]
}

func (i *implMapAnyRune) Len() int {
	return len(i.ordered)
}

func (i *implMapAnyRune) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAnyRune) IndexOf(k interface{}) int {
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

func (i *implMapAnyRune) Delete(k interface{}) MapAnyRune {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAnyRune) ForEach(f func(k interface{}, v rune)) MapAnyRune {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implMapAnyRune) SerializeOrdered(b bool) MapAnyRune {
	i.outOrder = b
	return i
}
