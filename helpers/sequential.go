package routinghelpers

import (
	"context"

	routing "github.com/libp2p/go-libp2p-routing"
	ropts "github.com/libp2p/go-libp2p-routing/options"

	multierror "github.com/hashicorp/go-multierror"
	cid "github.com/ipfs/go-cid"
	ci "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// SequentialRouter is like the ParallelRouter except that GetValue and FindPeer
// are called in series.
type SequentialRouter []routing.IpfsRouting

func (r SequentialRouter) PutValue(ctx context.Context, key string, value []byte, opts ...ropts.Option) error {
	return ParallelRouter(r).PutValue(ctx, key, value, opts...)
}

func (r SequentialRouter) get(ctx context.Context, do func(routing.IpfsRouting) (interface{}, error)) (interface{}, error) {
	var errs []error
	for _, ri := range r {
		val, err := do(ri)
		switch err {
		case nil:
			return val, nil
		case routing.ErrNotFound, routing.ErrNotSupported:
			continue
		}
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		errs = append(errs, err)
	}
	switch len(errs) {
	case 0:
		return nil, routing.ErrNotFound
	case 1:
		return nil, errs[0]
	default:
		return nil, &multierror.Error{Errors: errs}
	}
}

func (r SequentialRouter) GetValue(ctx context.Context, key string, opts ...ropts.Option) ([]byte, error) {
	valInt, err := r.get(ctx, func(ri routing.IpfsRouting) (interface{}, error) {
		return ri.GetValue(ctx, key, opts...)
	})
	val, _ := valInt.([]byte)
	return val, err
}

func (r SequentialRouter) GetPublicKey(ctx context.Context, p peer.ID) (ci.PubKey, error) {
	vInt, err := r.get(ctx, func(ri routing.IpfsRouting) (interface{}, error) {
		return routing.GetPublicKey(ri, ctx, p)
	})
	val, _ := vInt.(ci.PubKey)
	return val, err
}

func (r SequentialRouter) Provide(ctx context.Context, c *cid.Cid, local bool) error {
	return ParallelRouter(r).Provide(ctx, c, local)
}

func (r SequentialRouter) FindProvidersAsync(ctx context.Context, c *cid.Cid, count int) <-chan pstore.PeerInfo {
	return ParallelRouter(r).FindProvidersAsync(ctx, c, count)
}

func (r SequentialRouter) FindPeer(ctx context.Context, p peer.ID) (pstore.PeerInfo, error) {
	valInt, err := r.get(ctx, func(ri routing.IpfsRouting) (interface{}, error) {
		return ri.FindPeer(ctx, p)
	})
	val, _ := valInt.(pstore.PeerInfo)
	return val, err
}

func (r SequentialRouter) Bootstrap(ctx context.Context) error {
	return ParallelRouter(r).Bootstrap(ctx)
}

var _ routing.IpfsRouting = (SequentialRouter)(nil)
