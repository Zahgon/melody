package melody

import (
	"sync"
	"sync/atomic"
)

type hub struct {
	mu       sync.RWMutex
	sessions map[*Session]struct{}
	open     atomic.Bool
}

func newHub() *hub { _ = "STUB: not implemented"; return nil }

func (h *hub) closed() bool { _ = "STUB: not implemented"; return false }

func (h *hub) len() int { _ = "STUB: not implemented"; return 0 }

func (h *hub) all() []*Session { _ = "STUB: not implemented"; return nil }

func (h *hub) register(s *Session) { _ = "STUB: not implemented"; return }

func (h *hub) unregister(s *Session) { _ = "STUB: not implemented"; return }

func (h *hub) exit(msg envelope) { _ = "STUB: not implemented"; return }

func (h *hub) broadcast(msg envelope) { _ = "STUB: not implemented"; return }
