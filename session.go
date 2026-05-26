package melody

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Session wrapper around websocket connections.
type Session struct {
	Request    *http.Request
	Keys       map[string]any
	conn       *websocket.Conn
	output     chan envelope
	outputDone chan struct{}
	melody     *Melody
	open       bool
	rwmutex    sync.RWMutex
}

func (s *Session) writeMessage(message envelope) { _ = "STUB: not implemented"; return }

func (s *Session) writeRaw(message envelope) error { _ = "STUB: not implemented"; return nil }

// Use per-message deadline if specified, otherwise use global config

func (s *Session) closed() bool { _ = "STUB: not implemented"; return false }

func (s *Session) close() { _ = "STUB: not implemented"; return }

func (s *Session) ping() { _ = "STUB: not implemented"; return }

func (s *Session) writePump() { _ = "STUB: not implemented"; return }

func (s *Session) readPump() { _ = "STUB: not implemented"; return }

func (s *Session) handleMessage(t int, message []byte) { _ = "STUB: not implemented"; return }

// Write writes message to session.
func (s *Session) Write(msg []byte) error { _ = "STUB: not implemented"; return nil }

// WriteBinary writes a binary message to session.
func (s *Session) WriteBinary(msg []byte) error { _ = "STUB: not implemented"; return nil }

// Close closes session.
func (s *Session) Close() error { _ = "STUB: not implemented"; return nil }

// CloseWithMsg closes the session with the provided payload.
// Use the FormatCloseMessage function to format a proper close message payload.
func (s *Session) CloseWithMsg(msg []byte) error { _ = "STUB: not implemented"; return nil }

// Set is used to store a new key/value pair exclusively for this session.
// It also lazy initializes s.Keys if it was not used previously.
func (s *Session) Set(key string, value any) { _ = "STUB: not implemented"; return }

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (s *Session) Get(key string) (value any, exists bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// MustGet returns the value for the given key if it exists, otherwise it panics.
func (s *Session) MustGet(key string) any { _ = "STUB: not implemented"; return *new(any) }

// UnSet will delete the key and has no return value
func (s *Session) UnSet(key string) { _ = "STUB: not implemented"; return }

// IsClosed returns the status of the connection.
func (s *Session) IsClosed() bool {
	_ = "STUB: not implemented"

	// LocalAddr returns the local addr of the connection.
	return false
}

func (s *Session) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// RemoteAddr returns the remote addr of the connection.
	new(net.Addr)
}

func (s *Session) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// WebsocketConnection returns the underlying websocket connection.
// This can be used to e.g. set/read additional websocket options or to write sychronous messages.
func (s *Session) WebsocketConnection() *websocket.Conn {
	_ = "STUB: not implemented"

	// WriteWithDeadline writes a text message to the session with a custom write deadline.
	// If deadline is 0, uses Config.WriteWait.
	return nil
}

func (s *Session) WriteWithDeadline(msg []byte, deadline time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteBinaryWithDeadline writes a binary message to the session with a custom write deadline.
// If deadline is 0, uses Config.WriteWait.
func (s *Session) WriteBinaryWithDeadline(msg []byte, deadline time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
