package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapAnyFloat32 defines an ordered map of interface{} to float32.
type MapAnyFloat32 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k interface{}, v float32) MapAnyFloat32

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k interface{}, v *float32) MapAnyFloat32

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k interface{}, v float32) MapAnyFloat32

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k interface{}, v float32) MapAnyFloat32

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k interface{}) (value float32, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of float32 if an entry was found
	// with the key `k`.
	GetOpt(k interface{}) option.Float32

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapAnyFloat32Entry

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

	Delete(k interface{}) MapAnyFloat32

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k interface{}, v float32)) MapAnyFloat32
}

// MapAnyFloat32Entry is a single entry in an instance of
// MapAnyFloat32.
type MapAnyFloat32Entry struct {
	Key interface{} `json:"key"`
	Val float32     `json:"value"`
}

// NewMapAnyFloat32 creates a new instance of MapAnyFloat32 presized to the
// given size.
func NewMapAnyFloat32(size int) MapAnyFloat32 {
	return &implMapAnyFloat32{
		ordered: make([]MapAnyFloat32Entry, 0, size),
		index:   make(map[interface{}]float32, size),
	}
}

// MapAnyFloat32 is an ordered map interface{} to float32.
type implMapAnyFloat32 struct {
	ordered []MapAnyFloat32Entry
	index   map[interface{}]float32
}

func (i implMapAnyFloat32) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapAnyFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapAnyFloat32) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapAnyFloat32) Put(k interface{}, v float32) MapAnyFloat32 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapAnyFloat32Entry{k, v})
	return i
}

func (i *implMapAnyFloat32) PutIfNotNil(k interface{}, v *float32) MapAnyFloat32 {
	if !gomp.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapAnyFloat32) ReplaceOrPut(k interface{}, v float32) MapAnyFloat32 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapAnyFloat32) ReplaceIfExists(k interface{}, v float32) MapAnyFloat32 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapAnyFloat32) Get(k interface{}) (value float32, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapAnyFloat32) GetOpt(k interface{}) option.Float32 {
	if v, ok := i.index[k]; ok {
		return option.NewFloat32(v)
	}

	return option.NewEmptyFloat32()
}

func (i *implMapAnyFloat32) At(j int) MapAnyFloat32Entry {
	return i.ordered[j]
}

func (i *implMapAnyFloat32) Len() int {
	return len(i.ordered)
}

func (i *implMapAnyFloat32) Has(k interface{}) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapAnyFloat32) IndexOf(k interface{}) int {
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

func (i *implMapAnyFloat32) Delete(k interface{}) MapAnyFloat32 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapAnyFloat32) ForEach(f func(k interface{}, v float32)) MapAnyFloat32 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
