package notifications

import moved "github.com/libp2p/go-libp2p-core/routing/notifications"

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
var RegisterForQueryEvents = moved.RegisterForQueryEvents

// Deprecated: use github.com/libp2p/go-libp2p-core/routing/notifications.PublishQueryEvent instead.
var PublishQueryEvent = moved.PublishQueryEvent
