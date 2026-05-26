package melody

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Close codes defined in RFC 6455, section 11.7.
// Duplicate of codes from gorilla/websocket for convenience.
const (
	CloseNormalClosure           = 1000
	CloseGoingAway               = 1001
	CloseProtocolError           = 1002
	CloseUnsupportedData         = 1003
	CloseNoStatusReceived        = 1005
	CloseAbnormalClosure         = 1006
	CloseInvalidFramePayloadData = 1007
	ClosePolicyViolation         = 1008
	CloseMessageTooBig           = 1009
	CloseMandatoryExtension      = 1010
	CloseInternalServerErr       = 1011
	CloseServiceRestart          = 1012
	CloseTryAgainLater           = 1013
	CloseTLSHandshake            = 1015
)

type handleMessageFunc func(*Session, []byte)
type handleErrorFunc func(*Session, error)
type handleCloseFunc func(*Session, int, string) error
type handleSessionFunc func(*Session)
type filterFunc func(*Session) bool

// Melody implements a websocket manager.
type Melody struct {
	Config                   *Config
	Upgrader                 *websocket.Upgrader
	messageHandler           handleMessageFunc
	messageHandlerBinary     handleMessageFunc
	messageSentHandler       handleMessageFunc
	messageSentHandlerBinary handleMessageFunc
	errorHandler             handleErrorFunc
	closeHandler             handleCloseFunc
	connectHandler           handleSessionFunc
	disconnectHandler        handleSessionFunc
	pongHandler              handleSessionFunc
	hub                      *hub
}

// New creates a new melody instance with default Upgrader and Config.
func New() *Melody { _ = "STUB: not implemented"; return nil }

// HandleConnect fires fn when a session connects.
func (m *Melody) HandleConnect(fn func(*Session)) { _ = "STUB: not implemented"; return }

// HandleDisconnect fires fn when a session disconnects.
func (m *Melody) HandleDisconnect(fn func(*Session)) { _ = "STUB: not implemented"; return }

// HandlePong fires fn when a pong is received from a session.
func (m *Melody) HandlePong(fn func(*Session)) {
	_ = "STUB: not implemented"

	// HandleMessage fires fn when a text message comes in.
	// NOTE: by default Melody handles messages sequentially for each
	// session. This has the effect that a message handler exceeding the
	// read deadline (Config.PongWait, by default 1 minute) will time out
	// the session. Concurrent message handling can be turned on by setting
	// Config.ConcurrentMessageHandling to true.
	return
}

func (m *Melody) HandleMessage(fn func(*Session, []byte)) { _ = "STUB: not implemented"; return }

// HandleMessageBinary fires fn when a binary message comes in.
func (m *Melody) HandleMessageBinary(fn func(*Session, []byte)) { _ = "STUB: not implemented"; return }

// HandleSentMessage fires fn when a text message is successfully sent.
func (m *Melody) HandleSentMessage(fn func(*Session, []byte)) { _ = "STUB: not implemented"; return }

// HandleSentMessageBinary fires fn when a binary message is successfully sent.
func (m *Melody) HandleSentMessageBinary(fn func(*Session, []byte)) {
	_ = "STUB: not implemented"
	return
}

// HandleError fires fn when a session has an error.
func (m *Melody) HandleError(fn func(*Session, error)) { _ = "STUB: not implemented"; return }

// HandleClose sets the handler for close messages received from the session.
// The code argument to h is the received close code or CloseNoStatusReceived
// if the close message is empty. The default close handler sends a close frame
// back to the session.
//
// The application must read the connection to process close messages as
// described in the section on Control Frames above.
//
// The connection read methods return a CloseError when a close frame is
// received. Most applications should handle close messages as part of their
// normal error handling. Applications should only set a close handler when the
// application must perform some action before sending a close frame back to
// the session.
func (m *Melody) HandleClose(fn func(*Session, int, string) error) {
	_ = "STUB: not implemented"
	return
}

// HandleRequest upgrades http requests to websocket connections and dispatches them to be handled by the melody instance.
func (m *Melody) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleRequestWithKeys does the same as HandleRequest but populates session.Keys with keys.
func (m *Melody) HandleRequestWithKeys(w http.ResponseWriter, r *http.Request, keys map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// Broadcast broadcasts a text message to all sessions.
func (m *Melody) Broadcast(msg []byte) error { _ = "STUB: not implemented"; return nil }

// BroadcastFilter broadcasts a text message to all sessions that fn returns true for.
func (m *Melody) BroadcastFilter(msg []byte, fn func(*Session) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// BroadcastOthers broadcasts a text message to all sessions except session s.
func (m *Melody) BroadcastOthers(msg []byte, s *Session) error {
	_ = "STUB: not implemented"
	return nil
}

// BroadcastMultiple broadcasts a text message to multiple sessions given in the sessions slice.
func (m *Melody) BroadcastMultiple(msg []byte, sessions []*Session) error {
	_ = "STUB: not implemented"
	return nil
}

// BroadcastBinary broadcasts a binary message to all sessions.
func (m *Melody) BroadcastBinary(msg []byte) error { _ = "STUB: not implemented"; return nil }

// BroadcastBinaryFilter broadcasts a binary message to all sessions that fn returns true for.
func (m *Melody) BroadcastBinaryFilter(msg []byte, fn func(*Session) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// BroadcastBinaryOthers broadcasts a binary message to all sessions except session s.
func (m *Melody) BroadcastBinaryOthers(msg []byte, s *Session) error {
	_ = "STUB: not implemented"
	return nil
}

// Sessions returns all sessions. An error is returned if the melody session is closed.
func (m *Melody) Sessions() ([]*Session, error) { _ = "STUB: not implemented"; return nil, nil }

// Close closes the melody instance and all connected sessions.
func (m *Melody) Close() error { _ = "STUB: not implemented"; return nil }

// CloseWithMsg closes the melody instance with the given close payload and all connected sessions.
// Use the FormatCloseMessage function to format a proper close message payload.
func (m *Melody) CloseWithMsg(msg []byte) error { _ = "STUB: not implemented"; return nil }

// Len return the number of connected sessions.
func (m *Melody) Len() int {
	_ = "STUB: not implemented"

	// IsClosed returns the status of the melody instance.
	return 0
}

func (m *Melody) IsClosed() bool { _ = "STUB: not implemented"; return false }

// FormatCloseMessage formats closeCode and text as a WebSocket close message.
func FormatCloseMessage(closeCode int, text string) []byte { _ = "STUB: not implemented"; return nil }

// BroadcastWithDeadline broadcasts a text message with a custom write deadline.
// If deadline is 0, uses Config.WriteWait.
func (m *Melody) BroadcastWithDeadline(msg []byte, deadline time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// BroadcastFilterWithDeadline broadcasts a text message to filtered sessions with a custom write deadline.
// If deadline is 0, uses Config.WriteWait.
func (m *Melody) BroadcastFilterWithDeadline(msg []byte, deadline time.Duration, fn func(*Session) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// BroadcastBinaryWithDeadline broadcasts a binary message with a custom write deadline.
// If deadline is 0, uses Config.WriteWait.
func (m *Melody) BroadcastBinaryWithDeadline(msg []byte, deadline time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// BroadcastBinaryFilterWithDeadline broadcasts a binary message to filtered sessions with a custom write deadline.
// If deadline is 0, uses Config.WriteWait.
func (m *Melody) BroadcastBinaryFilterWithDeadline(msg []byte, deadline time.Duration, fn func(*Session) bool) error {
	_ = "STUB: not implemented"
	return nil
}
