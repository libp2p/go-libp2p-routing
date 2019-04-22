package notifications

import (
	"context"

	core "github.com/libp2p/go-libp2p-core/routing"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.QueryEventType instead.
type QueryEventType = core.QueryEventType

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.QueryEventBufferSize instead.
// Warning: it's impossible to alias a var in go, so reads and writes to this variable may be inaccurate
// or not have the intended effect.
var QueryEventBufferSize = core.QueryEventBufferSize

const (
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.SendingQuery instead.
	SendingQuery = core.SendingQuery
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.PeerResponse instead.
	PeerResponse = core.PeerResponse
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.FinalPeer instead.
	FinalPeer = core.FinalPeer
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.QueryError instead.
	QueryError = core.QueryError
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.Provider instead.
	Provider = core.Provider
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.Value instead.
	Value = core.Value
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.AddingPeer instead.
	AddingPeer = core.AddingPeer
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.DialingPeer instead.
	DialingPeer = core.DialingPeer
)

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.QueryEvent instead.
type QueryEvent = core.QueryEvent

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.RegisterForQueryEvents instead.
func RegisterForQueryEvents(ctx context.Context) (context.Context, <-chan *core.QueryEvent) {
	return core.RegisterForQueryEvents(ctx)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.PublishQueryEvent instead.
func PublishQueryEvent(ctx context.Context, ev *core.QueryEvent) {
	core.PublishQueryEvent(ctx, ev)
}
