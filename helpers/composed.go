package routinghelpers

import (
	"context"
	"sync"

	routing "github.com/libp2p/go-libp2p-routing"

	multierror "github.com/hashicorp/go-multierror"
	ci "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
)

// ComposedRouter composes the components into a single router.
//
// It also implements BootstrapRouting. All distinct components implementing
// BootstrapRouting will be bootstrapped in parallel.
type ComposedRouter struct {
	routing.ValueStore
	routing.PeerRouting
	routing.ContentRouting
}

// GetPublicKey returns the public key for the given peer.
func (cr *ComposedRouter) GetPublicKey(ctx context.Context, p peer.ID) (ci.PubKey, error) {
	return routing.GetPublicKey(cr.ValueStore, ctx, p)
}

// Bootstrap the router.
func (cr *ComposedRouter) Bootstrap(ctx context.Context) error {
	routers := make(map[routing.BootstrapRouting]struct{}, 3)
	for _, value := range []interface{}{
		cr.ValueStore,
		cr.ContentRouting,
		cr.PeerRouting,
	} {
		// No-oping out pieces is common so we might as well optimize
		// this.
		if _, ok := value.(NoopRouter); ok {
			continue
		}
		if b, ok := value.(routing.BootstrapRouting); ok {
			routers[b] = struct{}{}
		}
	}

	var wg sync.WaitGroup
	errs := make([]error, len(routers))
	wg.Add(len(routers))
	i := 0
	for b := range routers {
		i++
		go func(b routing.BootstrapRouting, i int) {
			errs[i] = b.Bootstrap(ctx)
			wg.Done()
		}(b, i)
	}
	wg.Wait()
	var me multierror.Error
	for _, err := range errs {
		if err != nil {
			me.Errors = append(me.Errors, err)
		}
	}
	return me.ErrorOrNil()
}

var _ routing.IpfsRouting = (*ComposedRouter)(nil)
var _ routing.PubKeyFetcher = (*ComposedRouter)(nil)
