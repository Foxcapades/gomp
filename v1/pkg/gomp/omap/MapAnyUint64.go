package omap

import (
	"encoding/json"
	"fmt"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapAnyUint64 defines an ordered map of interface{} to uint64.
type MapAnyUint64 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v uint64) MapAnyUint64

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *uint64) MapAnyUint64

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v uint64) MapAnyUint64

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v uint64) MapAnyUint64

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value uint64, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of uint64 if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Uint64

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyUint64Entry

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

	Delete(k interface{}) MapAnyUint64

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v uint64)) MapAnyUint64

	// SerializeOrdered sets whether or not the ordering should be enforced by
	// type when serializing the map.
	//
	// If set to true (the default value), the output will use an ordered type
	// when serializing (array for json, ordered map for yaml).  If set to false
	// the map will be serialized as a map/struct type and property ordering will
	// be determined by the serialization library.
	SerializeOrdered(bool) MapAnyUint64
}

// MapAnyUint64Entry is a single entry in an instance of
// MapAnyUint64.
type MapAnyUint64Entry struct {
	Key interface{} `json:"key"`
	Val uint64      `json:"value"`
}

// NewMapAnyUint64 creates a new instance of MapAnyUint64 presized to the
// given size.
func NewMapAnyUint64(size int) MapAnyUint64 {
	return &implMapAnyUint64{
		ordered:  make([]MapAnyUint64Entry, 0, size),
		index:    make(map[interface{}]uint64, size),
		outOrder: true,
	}
}

// MapAnyUint64 is an ordered map interface{} to uint64.
type implMapAnyUint64 struct {
	ordered  []MapAnyUint64Entry
	index    map[interface{}]uint64
	outOrder bool
}

func (i implMapAnyUint64) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapAnyUint64) MarshalJSON() ([]byte, error) {
	if i.outOrder {
		return json.Marshal(i.ordered)
	}

	out := make(map[string]interface{}, len(i.index))
	for k, v := range i.index {
		out[fmt.Sprint(k)] = v
	}

	return json.Marshal(out)
}

func (i implMapAnyUint64) ToYAML() (*yaml.Node, error) {
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

func (i *implMapAnyUint64) Put(k interface{}, v uint64) MapAnyUint64 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyUint64Entry{k, v})
	return i
}

func (i *implMapAnyUint64) PutIfNotNil(k interface{}, v *uint64) MapAnyUint64 {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAnyUint64) ReplaceOrPut(k interface{}, v uint64) MapAnyUint64 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapAnyUint64) ReplaceIfExists(k interface{}, v uint64) MapAnyUint64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAnyUint64) Get(k interface{}) (value uint64, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAnyUint64) GetOpt(k interface{}) option.Uint64 {
	if v, ok := i.index[k]; ok {
		return option.NewUint64(v)
	}

	return option.NewEmptyUint64()
}

func (i *implMapAnyUint64) At(j int) MapAnyUint64Entry {
	return i.ordered[j]
}

func (i *implMapAnyUint64) Len() int {
	return len(i.ordered)
}

func (i *implMapAnyUint64) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAnyUint64) IndexOf(k interface{}) int {
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

func (i *implMapAnyUint64) Delete(k interface{}) MapAnyUint64 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAnyUint64) ForEach(f func(k interface{}, v uint64)) MapAnyUint64 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}

func (i *implMapAnyUint64) SerializeOrdered(b bool) MapAnyUint64 {
	i.outOrder = b
	return i
}
