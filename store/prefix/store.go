package prefix

import (
	"github.com/cosmos/cosmos-sdk/store/types"
)

var _ types.KVStore = Store{}

type Store struct {
	parent types.KVStore
	prefix []byte
}

func NewStore(parent types.KVStore, prefix []byte) Store {
	return Store{
		parent: parent,
		prefix: prefix,
	}
}

// Implements KVStore
func (s Store) Get(key []byte) []byte {
	return s.parent.Get(append(s.prefix, key...))
}

// Implements KVStore
func (s Store) Has(key []byte) bool {
	return s.parent.Has(append(s.prefix, key...))
}

// Implements KVStore
func (s Store) Set(key, value []byte) {
	s.parent.Set(append(s.prefix, key...), value)
}

// Implements KVStore
func (s Store) Delete(key []byte) {
	s.parent.Delete(append(s.prefix, key...))
}

// Implements KVStore
func (s Store) Iterator(start, end []byte) types.Iterator {
	if end == nil {
		end = types.PrefixEndBytes(s.prefix)
	} else {
		end = append(s.prefix, end...)
	}
	return prefixIterator{
		prefix: s.prefix,
		iter:   s.parent.Iterator(append(s.prefix, start...), end),
	}
}

// Implements KVStore
func (s Store) ReverseIterator(start, end []byte) types.Iterator {
	if end == nil {
		end = types.PrefixEndBytes(s.prefix)
	} else {
		end = append(s.prefix, end...)
	}
	return prefixIterator{
		prefix: s.prefix,
		iter:   s.parent.ReverseIterator(start, end),
	}
}

type prefixIterator struct {
	prefix []byte

	iter types.Iterator
}

// Implements Iterator
func (iter prefixIterator) Domain() (start []byte, end []byte) {
	start, end = iter.iter.Domain()
	start = start[len(iter.prefix):]
	end = end[len(iter.prefix):]
	return
}

// Implements Iterator
func (iter prefixIterator) Valid() bool {
	return iter.iter.Valid()
}

// Implements Iterator
func (iter prefixIterator) Next() {
	iter.iter.Next()
}

// Implements Iterator
func (iter prefixIterator) Key() (key []byte) {
	key = iter.iter.Key()
	key = key[len(iter.prefix):]
	return
}

// Implements Iterator
func (iter prefixIterator) Value() []byte {
	return iter.iter.Value()
}

// Implements Iterator
func (iter prefixIterator) Close() {
	iter.iter.Close()
}
