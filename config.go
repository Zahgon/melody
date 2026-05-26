package melody

import "time"

// Config melody configuration struct.
type Config struct {
	WriteWait                 time.Duration // Duration until write times out.
	PongWait                  time.Duration // Timeout for waiting on pong.
	PingPeriod                time.Duration // Duration between pings.
	MaxMessageSize            int64         // Maximum size in bytes of a message.
	MessageBufferSize         int           // The max amount of messages that can be in a sessions buffer before it starts dropping them.
	ConcurrentMessageHandling bool          // Handle messages from sessions concurrently.
}

func newConfig() *Config { _ = "STUB: not implemented"; return nil }
