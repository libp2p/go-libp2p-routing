// package routing defines the interface for a routing system used by ipfs.
package routing

import (
	"context"
	"errors"

	cid "github.com/ipfs/go-cid"
	ci "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// ErrNotFound is returned when a search fails to find anything
var ErrNotFound = errors.New("routing: not found")

// ContentRouting is a value provider layer of indirection. It is used to find
// information about who has what content.
type ContentRouting interface {
	// Provide adds the given cid to the content routing system. If 'true' is
	// passed, it also announces it, otherwise it is just kept in the local
	// accounting of which objects are being provided.
	Provide(context.Context, *cid.Cid, bool) error

	// Search for peers who are able to provide a given key
	FindProvidersAsync(context.Context, *cid.Cid, int) <-chan pstore.PeerInfo
}

// PeerRouting is a way to find information about certain peers.
// This can be implemented by a simple lookup table, a tracking server,
// or even a DHT.
type PeerRouting interface {
	// Find specific Peer
	// FindPeer searches for a peer with given ID, returns a pstore.PeerInfo
	// with relevant addresses.
	FindPeer(context.Context, peer.ID) (pstore.PeerInfo, error)
}

type ValueStore interface {
	// Basic Put/Get

	// PutValue adds value corresponding to given Key.
	PutValue(context.Context, string, []byte) error

	// GetValue searches for the value corresponding to given Key.
	GetValue(context.Context, string) ([]byte, error)

	// GetValues searches for values corresponding to given Key.
	//
	// Passing a value of '0' for the count argument will cause the
	// routing interface to return values only from cached or local storage
	// and return an error if no cached value is found.
	//
	// Passing a value of '1' will return a local value if found, and query
	// the network for the first value it finds otherwise.
	// As a result, a value of '1' is mostly useful for cases where the record
	// in question has only one valid value (such as public keys)
	GetValues(c context.Context, k string, count int) ([]RecvdVal, error)

	// GetValues searches for values corresponding to given Key.
	//
	// Passing a negative value to the count argument, the search will not stop
	// until the context is cancelled or timeouts
	//
	// Passing a value of '0' for the count argument will cause the
	// routing interface to return values only from cached or local storage
	// and return an error if no cached value is found.
	//
	// Passing a value of '1' will return a local value if found, and query
	// the network for the first value it finds otherwise.
	// As a result, a value of '1' is mostly useful for cases where the record
	// in question has only one valid value (such as public keys)
	GetValuesAsync(c context.Context, k string, count int) <-chan *RecvdVal
}

// IpfsRouting is the combination of different routing types that ipfs
// uses. It can be satisfied by a single item (such as a DHT) or multiple
// different pieces that are more optimized to each task.
type IpfsRouting interface {
	ContentRouting
	PeerRouting
	ValueStore

	// Bootstrap allows callers to hint to the routing system to get into a
	// Boostrapped state
	Bootstrap(context.Context) error

	// TODO expose io.Closer or plain-old Close error
}

// RecvdVal represents a dht value record that has been received from a given peer
// it is used to track peers with expired records in order to correct them.
type RecvdVal struct {
	Error error
	From  peer.ID
	Val   []byte
}

type PubKeyFetcher interface {
	GetPublicKey(context.Context, peer.ID) (ci.PubKey, error)
}

// KeyForPublicKey returns the key used to retrieve public keys
// from the dht.
func KeyForPublicKey(id peer.ID) string {
	return "/pk/" + string(id)
}

func GetPublicKey(r ValueStore, ctx context.Context, pkhash []byte) (ci.PubKey, error) {
	if dht, ok := r.(PubKeyFetcher); ok {
		// If we have a DHT as our routing system, use optimized fetcher
		return dht.GetPublicKey(ctx, peer.ID(pkhash))
	} else {
		key := "/pk/" + string(pkhash)
		pkval, err := r.GetValue(ctx, key)
		if err != nil {
			return nil, err
		}

		// get PublicKey from node.Data
		return ci.UnmarshalPublicKey(pkval)
	}
}
