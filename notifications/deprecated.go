package notifications

import (
	"context"

	moved "github.com/libp2p/go-libp2p-core/routing/notifications"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.QueryEventType instead.
type QueryEventType = moved.QueryEventType

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.QueryEventBufferSize instead.
// Warning: it's impossible to alias a var in go, so reads and writes to this variable may be inaccurate
// or not have the intended effect.
var QueryEventBufferSize = moved.QueryEventBufferSize

const (
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.SendingQuery instead.
	SendingQuery = moved.SendingQuery
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.PeerResponse instead.
	PeerResponse = moved.PeerResponse
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.FinalPeer instead.
	FinalPeer = moved.FinalPeer
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.QueryError instead.
	QueryError = moved.QueryError
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.Provider instead.
	Provider = moved.Provider
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.Value instead.
	Value = moved.Value
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.AddingPeer instead.
	AddingPeer = moved.AddingPeer
	// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.DialingPeer instead.
	DialingPeer = moved.DialingPeer
)

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.QueryEvent instead.
type QueryEvent = moved.QueryEvent

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.RegisterForQueryEvents instead.
func RegisterForQueryEvents(ctx context.Context) (context.Context, <-chan *moved.QueryEvent) {
	return moved.RegisterForQueryEvents(ctx)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.PublishQueryEvent instead.
func PublishQueryEvent(ctx context.Context, ev *moved.QueryEvent) {
	moved.PublishQueryEvent(ctx, ev)
}
