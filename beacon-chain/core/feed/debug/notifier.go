package debug

import "github.com/OffchainLabs/prysm/v6/async/event"

// Notifier interface defines the methods of the service that provides debug events or updates to consumers.
type Notifier interface {
	DebugEventFeed() event.SubscriberSender
}
