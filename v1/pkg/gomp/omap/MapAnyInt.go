package omap

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapAnyInt defines an ordered map of interface{} to int.
type MapAnyInt interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v int) MapAnyInt

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *int) MapAnyInt

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v int) MapAnyInt

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v int) MapAnyInt

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value int, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of int if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Int

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyIntEntry

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

	Delete(k interface{}) MapAnyInt

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v int)) MapAnyInt

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) MapAnyInt
}

// MapAnyIntEntry is a single entry in an instance of
// MapAnyInt.
type MapAnyIntEntry struct {
	Key interface{} `json:"key"`
	Val int         `json:"value"`
}

// NewMapAnyInt creates a new instance of MapAnyInt presized to the
// given size.
func NewMapAnyInt(size int) MapAnyInt {
	return &implMapAnyInt{
		ordered:  make([]MapAnyIntEntry, 0, size),
		index:    make(map[interface{}]int, size),
		outOrder: true,
	}
}

// MapAnyInt is an ordered map interface{} to int.
type implMapAnyInt struct {
	ordered  []MapAnyIntEntry
	index    map[interface{}]int
	outOrder bool
}

func (i implMapAnyInt) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapAnyInt) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i implMapAnyInt) ToYAML() (*yaml.Node, error) {
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

func (i *implMapAnyInt) Put(k interface{}, v int) MapAnyInt {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyIntEntry{k, v})
	return i
}

func (i *implMapAnyInt) PutIfNotNil(k interface{}, v *int) MapAnyInt {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAnyInt) ReplaceOrPut(k interface{}, v int) MapAnyInt {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapAnyInt) ReplaceIfExists(k interface{}, v int) MapAnyInt {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAnyInt) Get(k interface{}) (value int, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAnyInt) GetOpt(k interface{}) option.Int {
	if v, ok := i.index[k]; ok {
		return option.NewInt(v)
	}

	return option.NewEmptyInt()
}

func (i *implMapAnyInt) At(j int) MapAnyIntEntry {
	return i.ordered[j]
}

func (i *implMapAnyInt) Len() int {
	return len(i.ordered)
}

func (i *implMapAnyInt) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAnyInt) IndexOf(k interface{}) int {
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

func (i *implMapAnyInt) Delete(k interface{}) MapAnyInt {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAnyInt) ForEach(f func(k interface{}, v int)) MapAnyInt {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implMapAnyInt) SerializeOrdered(b bool) MapAnyInt {
	i.outOrder = b
	return i
}
