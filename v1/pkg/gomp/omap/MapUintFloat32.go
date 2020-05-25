package omap

import (
	"encoding/json"

	"github.com/Foxcapades/gomp/v1/internal/util"
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// MapUintFloat32 defines an ordered map of uint to float32.
type MapUintFloat32 interface {
	yaml.Marshaler
	json.Marshaler
	xyml.Marshaler

	// Put appends the given key/value pair to the end of the map.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed.  The new value will still be appended to the end of the map.
	Put(k uint, v float32) MapUintFloat32

	// PutIfNotNil appends the given key/pair value to the end of the map if the
	// `v` pointer is not nil.
	//
	// If the map already contains an entry with the key `k`, then it will be
	// removed only if `v` is not nil.  Follows the same ordering/removal rules as
	// Put.
	PutIfNotNil(k uint, v *float32) MapUintFloat32

	// ReplaceOrPut either replaces the existing entry keyed at `k` without
	// changing the map ordering or appends the given key/value pair to the end of
	// the map if no entry with the key `k` exists.
	ReplaceOrPut(k uint, v float32) MapUintFloat32

	// ReplaceIfExists replaces the value at key `k` with the given value `v`
	// without changing the map ordering.
	//
	// If no entry in the map currently exists with the key `k` this method does
	// nothing.
	ReplaceIfExists(k uint, v float32) MapUintFloat32

	// Get looks up the value in the map with the given key `k`.
	//
	// Returns a value and a boolean value indicating whether the value was found.
	Get(k uint) (value float32, exists bool)

	// GetOpt looks up the value in the map with the given key `k` and returns
	// an option which will contain an option of float32 if an entry was found
	// with the key `k`.
	GetOpt(k uint) option.Float32

	// At returns the key/value pair at the given index.
	//
	// This method makes no attempt to verify that the index given actually exists
	// in the map.
	At(i int) MapUintFloat32Entry

	// Len returns the current size of the map.
	Len() int

	// Has returns whether an entry exists with the key `k`.
	Has(k uint) bool

	// Has returns the position in the map of the entry matching key `k`.
	//
	// If no entry exists in the map with key `k` this method returns -1.
	//
	// Note: this method may iterate, at most once, through all the entries in the
	// map.
	IndexOf(k uint) int

	Delete(k uint) MapUintFloat32

	// ForEach calls the given function for for every entry in the map.
	ForEach(func(k uint, v float32)) MapUintFloat32
}

// MapUintFloat32Entry is a single entry in an instance of
// MapUintFloat32.
type MapUintFloat32Entry struct {
	Key uint    `json:"key"`
	Val float32 `json:"value"`
}

// NewMapUintFloat32 creates a new instance of MapUintFloat32 presized to the
// given size.
func NewMapUintFloat32(size int) MapUintFloat32 {
	return &implMapUintFloat32{
		ordered: make([]MapUintFloat32Entry, 0, size),
		index:   make(map[uint]float32, size),
	}
}

// MapUintFloat32 is an ordered map uint to float32.
type implMapUintFloat32 struct {
	ordered []MapUintFloat32Entry
	index   map[uint]float32
}

func (i implMapUintFloat32) MarshalYAML() (interface{}, error) {
	return i.ToYAML()
}

func (i implMapUintFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ordered)
}

func (i *implMapUintFloat32) ToYAML() (*yaml.Node, error) {
	out := xyml.NewOrderedMapNode(i.Len())

	for j := range i.ordered {
		if err := xyml.MapAppend(out, i.ordered[j].Key, i.ordered[j].Val); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (i *implMapUintFloat32) Put(k uint, v float32) MapUintFloat32 {
	i.Delete(k)
	i.index[k] = v
	i.ordered = append(i.ordered, MapUintFloat32Entry{k, v})
	return i
}

func (i *implMapUintFloat32) PutIfNotNil(k uint, v *float32) MapUintFloat32 {
	if !util.IsNil(v) {
		return i.Put(k, *v)
	}

	return i
}

func (i *implMapUintFloat32) ReplaceOrPut(k uint, v float32) MapUintFloat32 {
	if i.Has(k) {
		return i.ReplaceIfExists(k, v)
	}

	return i.Put(k, v)
}

func (i *implMapUintFloat32) ReplaceIfExists(k uint, v float32) MapUintFloat32 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.index[k] = v
		i.ordered[pos].Val = v
	}

	return i
}

func (i *implMapUintFloat32) Get(k uint) (value float32, exists bool) {
	v, ok := i.index[k]
	return v, ok
}

func (i *implMapUintFloat32) GetOpt(k uint) option.Float32 {
	if v, ok := i.index[k]; ok {
		return option.NewFloat32(v)
	}

	return option.NewEmptyFloat32()
}

func (i *implMapUintFloat32) At(j int) MapUintFloat32Entry {
	return i.ordered[j]
}

func (i *implMapUintFloat32) Len() int {
	return len(i.ordered)
}

func (i *implMapUintFloat32) Has(k uint) bool {
	_, ok := i.index[k]
	return ok
}

func (i *implMapUintFloat32) IndexOf(k uint) int {
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

func (i *implMapUintFloat32) Delete(k uint) MapUintFloat32 {
	pos := i.IndexOf(k)

	if pos > -1 {
		i.ordered = append(i.ordered[:pos], i.ordered[pos+1:]...)
		delete(i.index, k)
	}

	return i
}

func (i *implMapUintFloat32) ForEach(f func(k uint, v float32)) MapUintFloat32 {
	for j := range i.ordered {
		f(i.ordered[j].Key, i.ordered[j].Val)
	}

	return i
}
