package routinghelpers

import (
	"context"

	routing "github.com/libp2p/go-libp2p-routing"
	ropts "github.com/libp2p/go-libp2p-routing/options"

	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// NoopRouter always returns errors.
type NoopRouter struct{}

// PutValue returns ErrNotSupported
func (nr NoopRouter) PutValue(context.Context, string, []byte, ...ropts.Option) error {
	return routing.ErrNotSupported
}

// GetValue returns ErrNotFound
func (nr NoopRouter) GetValue(context.Context, string, ...ropts.Option) ([]byte, error) {
	return nil, routing.ErrNotFound
}

// Provide returns ErrNotSupported
func (nr NoopRouter) Provide(context.Context, *cid.Cid, bool) error {
	return routing.ErrNotSupported
}

// FindProvidersAsync returns a closed channel
func (nr NoopRouter) FindProvidersAsync(context.Context, *cid.Cid, int) <-chan pstore.PeerInfo {
	ch := make(chan pstore.PeerInfo)
	close(ch)
	return ch
}

// FindPeer returns ErrNotFound
func (nr NoopRouter) FindPeer(context.Context, peer.ID) (pstore.PeerInfo, error) {
	return pstore.PeerInfo{}, routing.ErrNotFound
}

// Bootstrap returns nil (no point in returning an error, nothing really failed).
func (nr NoopRouter) Bootstrap(context.Context) error {
	return nil
}

var _ routing.IpfsRouting = NoopRouter{}
