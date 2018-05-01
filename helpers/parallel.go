package routinghelpers

import (
	"context"
	"reflect"
	"sync"

	routing "github.com/libp2p/go-libp2p-routing"
	ropts "github.com/libp2p/go-libp2p-routing/options"

	multierror "github.com/hashicorp/go-multierror"
	cid "github.com/ipfs/go-cid"
	ci "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// ParallelRouter operates on the slice of routers in parallel.
type ParallelRouter []routing.IpfsRouting

// Helper function that sees through router composition to avoid unnecessary
// go routines.
func supportsKey(vs routing.ValueStore, key string) bool {
	switch vs := vs.(type) {
	case NoopRouter:
		return false
	case *ComposedRouter:
		return supportsKey(vs.ValueStore, key)
	case ParallelRouter:
		for _, ri := range vs {
			if supportsKey(ri, key) {
				return true
			}
		}
		return false
	case SequentialRouter:
		for _, ri := range vs {
			if supportsKey(ri, key) {
				return true
			}
		}
		return false
	case *LimitedValueStore:
		return vs.KeySupported(key) && supportsKey(vs.ValueStore, key)
	default:
		return true
	}
}

func supportsPeer(vs routing.PeerRouting) bool {
	switch vs := vs.(type) {
	case NoopRouter:
		return false
	case *ComposedRouter:
		return supportsPeer(vs.PeerRouting)
	case ParallelRouter:
		for _, ri := range vs {
			if supportsPeer(ri) {
				return true
			}
		}
		return false
	case SequentialRouter:
		for _, ri := range vs {
			if supportsPeer(ri) {
				return true
			}
		}
		return false
	default:
		return true
	}
}

func supportsContent(vs routing.ContentRouting) bool {
	switch vs := vs.(type) {
	case NoopRouter:
		return false
	case *ComposedRouter:
		return supportsContent(vs.ContentRouting)
	case ParallelRouter:
		for _, ri := range vs {
			if supportsContent(ri) {
				return true
			}
		}
		return false
	case SequentialRouter:
		for _, ri := range vs {
			if supportsContent(ri) {
				return true
			}
		}
		return false
	default:
		return true
	}
}

func (r ParallelRouter) filter(filter func(routing.IpfsRouting) bool) ParallelRouter {
	cpy := make(ParallelRouter, 0, len(r))
	for _, ri := range r {
		if filter(ri) {
			cpy = append(cpy, ri)
		}
	}
	return cpy
}

func (r ParallelRouter) put(do func(routing.IpfsRouting) error) error {
	switch len(r) {
	case 0:
		return routing.ErrNotSupported
	case 1:
		return do(r[0])
	}

	var wg sync.WaitGroup
	results := make([]error, len(r))
	wg.Add(len(r))
	for i, ri := range r {
		go func(ri routing.IpfsRouting, i int) {
			results[i] = do(ri)
			wg.Done()
		}(ri, i)
	}
	wg.Wait()

	var errs []error
	for _, err := range results {
		switch err {
		case nil:
			// Success!
			return nil
		case routing.ErrNotSupported:
		default:
			errs = append(errs, err)
		}
	}

	switch len(errs) {
	case 0:
		return routing.ErrNotSupported
	case 1:
		return errs[0]
	default:
		return &multierror.Error{Errors: errs}
	}
}

func (r ParallelRouter) get(ctx context.Context, do func(routing.IpfsRouting) (interface{}, error)) (interface{}, error) {
	switch len(r) {
	case 0:
		return nil, routing.ErrNotFound
	case 1:
		return do(r[0])
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	results := make(chan struct {
		val interface{}
		err error
	})
	for _, ri := range r {
		go func(ri routing.IpfsRouting) {
			value, err := do(ri)
			select {
			case results <- struct {
				val interface{}
				err error
			}{
				val: value,
				err: err,
			}:
			case <-ctx.Done():
			}
		}(ri)
	}

	var errs []error
	for _ = range r {
		select {
		case res := <-results:
			switch res.err {
			case nil:
				return res.val, nil
			case routing.ErrNotFound, routing.ErrNotSupported:
				continue
			}
			// If the context has expired, just return that error
			// and ignore the other errors.
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			errs = append(errs, res.err)
		case <-ctx.Done():
			return nil, ctx.Err()
		}
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

func (r ParallelRouter) forKey(key string) ParallelRouter {
	return r.filter(func(ri routing.IpfsRouting) bool {
		return supportsKey(ri, key)
	})
}

func (r ParallelRouter) PutValue(ctx context.Context, key string, value []byte, opts ...ropts.Option) error {
	return r.forKey(key).put(func(ri routing.IpfsRouting) error {
		return ri.PutValue(ctx, key, value, opts...)
	})
}

func (r ParallelRouter) GetValue(ctx context.Context, key string, opts ...ropts.Option) ([]byte, error) {
	vInt, err := r.forKey(key).get(ctx, func(ri routing.IpfsRouting) (interface{}, error) {
		return ri.GetValue(ctx, key, opts...)
	})
	val, _ := vInt.([]byte)
	return val, err
}

func (r ParallelRouter) GetPublicKey(ctx context.Context, p peer.ID) (ci.PubKey, error) {
	vInt, err := r.
		forKey(routing.KeyForPublicKey(p)).
		get(ctx, func(ri routing.IpfsRouting) (interface{}, error) {
			return routing.GetPublicKey(ri, ctx, p)
		})
	val, _ := vInt.(ci.PubKey)
	return val, err
}

func (r ParallelRouter) FindPeer(ctx context.Context, p peer.ID) (pstore.PeerInfo, error) {
	vInt, err := r.filter(func(ri routing.IpfsRouting) bool {
		return supportsPeer(ri)
	}).get(ctx, func(ri routing.IpfsRouting) (interface{}, error) {
		return ri.FindPeer(ctx, p)
	})
	pi, _ := vInt.(pstore.PeerInfo)
	return pi, err
}

func (r ParallelRouter) Provide(ctx context.Context, c *cid.Cid, local bool) error {
	return r.filter(func(ri routing.IpfsRouting) bool {
		return supportsContent(ri)
	}).put(func(ri routing.IpfsRouting) error {
		return ri.Provide(ctx, c, local)
	})
}

func (r ParallelRouter) FindProvidersAsync(ctx context.Context, c *cid.Cid, count int) <-chan pstore.PeerInfo {
	routers := r.filter(func(ri routing.IpfsRouting) bool {
		return supportsContent(ri)
	})

	switch len(routers) {
	case 0:
		ch := make(chan pstore.PeerInfo)
		close(ch)
		return ch
	case 1:
		return r.FindProvidersAsync(ctx, c, count)
	}

	out := make(chan pstore.PeerInfo, count)

	ctx, cancel := context.WithCancel(ctx)

	providers := make([]reflect.SelectCase, 0, len(routers))
	for _, ri := range routers {
		ch := ri.FindProvidersAsync(ctx, c, count)

		// Always send all values immediately available.
		if drain(out, ch) {
			providers = append(providers, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			})
		}

		// Short-circuit finish
		if len(out) == cap(out) {
			close(out)
			cancel()
			return out
		}
	}
	count -= len(out)
	go func() {
		defer cancel()
		defer close(out)
		for count > 0 && len(providers) > 0 {
			chosen, val, ok := reflect.Select(providers)
			if !ok {
				providers[chosen] = providers[len(providers)-1]
				providers = providers[:len(providers)-1]
			} else {
				out <- val.Interface().(pstore.PeerInfo)
				count--
			}
		}
	}()
	return out
}

func drain(out chan<- pstore.PeerInfo, in <-chan pstore.PeerInfo) bool {
	for len(out) < cap(out) {
		select {
		case p, ok := <-in:
			if !ok {
				return false
			}
			out <- p
		default:
			return true
		}
	}
	return true
}

func (r ParallelRouter) Bootstrap(ctx context.Context) error {
	var wg sync.WaitGroup
	errs := make([]error, len(r))
	wg.Add(len(r))
	for i, b := range r {
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

var _ routing.IpfsRouting = (ParallelRouter)(nil)
