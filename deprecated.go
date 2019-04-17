// package routing defines the interface for a routing system used by ipfs.
package routing

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	moved "github.com/libp2p/go-libp2p-core/routing"
	ci "github.com/libp2p/go-libp2p-crypto"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.ErrNotFound instead.
var ErrNotFound = moved.ErrNotFound

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.ErrNotSupported instead.
var ErrNotSupported = moved.ErrNotSupported

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.ContentRouting instead.
type ContentRouting = moved.ContentRouting

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.PeerRouting instead.
type PeerRouting = moved.PeerRouting

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.ValueStore instead.
type ValueStore = moved.ValueStore

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.Routing instead.
type IpfsRouting = moved.Routing

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.PubKeyFetcher instead.
type PubKeyFetcher = moved.PubKeyFetcher

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.KeyForPublicKey instead.
func KeyForPublicKey(id peer.ID) string {
	return moved.KeyForPublicKey(id)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.GetPublicKey instead.
func GetPublicKey(r moved.ValueStore, ctx context.Context, p peer.ID) (ci.PubKey, error) {
	return moved.GetPublicKey(r, ctx, p)
}
