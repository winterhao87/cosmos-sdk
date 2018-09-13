package prefix

import (
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/tracekv"
)

var _ sdk.KVStore = Store{}

type Store struct {
	parent sdk.KVStore
	prefix []byte
}

// Implements Store
func (s Store) GetStoreType() sdk.StoreType {
	return s.parent.GetStoreType()
}

// Implements CacheWrap
func (s Store) CacheWrap() sdk.CacheWrap {
	return cachekv.NewStore(s)
}

// CacheWrapWithTrace implements the KVStore interface.
func (s Store) CacheWrapWithTrace(w io.Writer, tc sdk.TraceContext) sdk.CacheWrap {
	return cachekv.NewStore(tracekv.NewStore(s, w, tc))
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
func (s Store) Prefix(prefix []byte) sdk.KVStore {
	return Store{s, prefix}
}

// Implements KVStore
func (s Store) Gas(meter sdk.GasMeter, config sdk.GasConfig) sdk.KVStore {
	return NewGasKVStore(meter, config, s)
}

// Implements KVStore
func (s Store) Iterator(start, end []byte) sdk.Iterator {
	if end == nil {
		end = sdk.PrefixEndBytes(s.prefix)
	} else {
		end = append(s.prefix, end...)
	}
	return prefixIterator{
		prefix: s.prefix,
		iter:   s.parent.Iterator(append(s.prefix, start...), end),
	}
}

// Implements KVStore
func (s Store) ReverseIterator(start, end []byte) sdk.Iterator {
	if end == nil {
		end = sdk.PrefixEndBytes(s.prefix)
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

	iter sdk.Iterator
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
