package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapIntUint8 defines an ordered map of int to uint8.
type MapIntUint8 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k int, v uint8) MapIntUint8

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k int, v *uint8) MapIntUint8

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k int, v uint8) MapIntUint8

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k int, v uint8) MapIntUint8

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k int) (value uint8, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of uint8 if an entry was found
	// with the key `k`.
	GetOpt(k int) option.Uint8

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapIntUint8Entry

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

	Delete(k int) MapIntUint8

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k int, v uint8)) MapIntUint8
}

// MapIntUint8Entry is a single entry in an instance of
// MapIntUint8.
type MapIntUint8Entry struct {
	Key int   `json:"key"`
	Val uint8 `json:"value"`
}

// NewMapIntUint8 creates a new instance of MapIntUint8 presized to the
// given size.
func NewMapIntUint8(size int) MapIntUint8 {
	return &implMapIntUint8{
		ordered: make([]MapIntUint8Entry, 0, size),
		index:   make(map[int]uint8, size),
	}
}

// MapIntUint8 is an ordered map int to uint8.
type implMapIntUint8 struct {
	ordered []MapIntUint8Entry
	index   map[int]uint8
}

func (i implMapIntUint8) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapIntUint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapIntUint8) ToYAML() (*yaml.Node, error) {
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

func (i *implMapIntUint8) Put(k int, v uint8) MapIntUint8 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapIntUint8Entry{k, v})
	return i
}

func (i *implMapIntUint8) PutIfNotNil(k int, v *uint8) MapIntUint8 {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapIntUint8) ReplaceOrPut(k int, v uint8) MapIntUint8 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapIntUint8) ReplaceIfExists(k int, v uint8) MapIntUint8 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapIntUint8) Get(k int) (value uint8, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapIntUint8) GetOpt(k int) option.Uint8 {
	if v, ok := i.index[k]; ok {
		return option.NewUint8(v)
	}

	return option.NewEmptyUint8()
}

func (i *implMapIntUint8) At(j int) MapIntUint8Entry {
	return i.ordered[j]
}

func (i *implMapIntUint8) Len() int {
	return len(i.ordered)
}

func (i *implMapIntUint8) Has(k int) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapIntUint8) IndexOf(k int) int {
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

func (i *implMapIntUint8) Delete(k int) MapIntUint8 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapIntUint8) ForEach(f func(k int, v uint8)) MapIntUint8 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
