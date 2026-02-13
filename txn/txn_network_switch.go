package txn

import (
	"fmt"
	"sync"
	"time"
)

type TxnState string

const (
	StateNew      TxnState = "NEW"
	StateAuthSent TxnState = "AUTH_SENT"
	StateAuthResp TxnState = "AUTH_RESPONDED"
	StateCleared  TxnState = "CLEARED"
	StateSettled  TxnState = "SETTLED"
	StateReversed TxnState = "REVERSED"
)

type Transaction struct {
	ID           string // txn_id / STAN composite
	State        TxnState
	AuthRequest  map[string]any
	AuthResponse map[string]any
	Timestamps   map[string]time.Time
	Lock         sync.Mutex // For atomic state transitions
}

// Partition in-memory state
type Partition struct {
	txns sync.Map // map[string]*Transaction
}

// Add or fetch txn
func (p *Partition) GetOrCreateTxn(id string) *Transaction {
	val, _ := p.txns.LoadOrStore(id, &Transaction{
		ID:         id,
		State:      StateNew,
		Timestamps: make(map[string]time.Time),
	})
	return val.(*Transaction)
}

func (t *Transaction) Transition(newState TxnState) error {
	t.Lock.Lock()
	defer t.Lock.Unlock()

	switch t.State {
	case StateNew:
		if newState != StateAuthSent {
			return fmt.Errorf("invalid transition: %s → %s", t.State, newState)
		}
	case StateAuthSent:
		if newState != StateAuthResp && newState != StateReversed {
			return fmt.Errorf("invalid transition: %s → %s", t.State, newState)
		}
	case StateAuthResp:
		if newState != StateCleared && newState != StateReversed {
			return fmt.Errorf("invalid transition: %s → %s", t.State, newState)
		}
	case StateCleared:
		if newState != StateSettled {
			return fmt.Errorf("invalid transition: %s → %s", t.State, newState)
		}
	default:
		return fmt.Errorf("cannot transition from terminal state: %s", t.State)
	}

	t.State = newState
	t.Timestamps[string(newState)] = time.Now()
	return nil
}
