// package routing defines the interface for a routing system used by ipfs.
package routing

import moved "github.com/libp2p/go-libp2p-core/routing"

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
var KeyForPublicKey = moved.KeyForPublicKey

// Deprecated: use github.com/libp2p/go-libp2p-core/routing.GetPublicKey instead.
var GetPublicKey = moved.GetPublicKey
