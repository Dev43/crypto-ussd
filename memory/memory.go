package memory

import (
	"errors"
	"time"

	"github.com/orcaman/concurrent-map"
)

type InMemorySessionStore struct {
	lastWriteTime time.Time
	data          cmap.ConcurrentMap
}

// NewInMemorySessionStore
func NewInMemorySessionStore() *InMemorySessionStore {
	return &InMemorySessionStore{
		lastWriteTime: time.Now(),
		data:          cmap.New(),
	}
}

// Write
func (m *InMemorySessionStore) Write(sessionID string, value interface{}) error {
	m.data.Set(sessionID, value)
	m.lastWriteTime = time.Now()
	return nil
}

// Delete
func (m *InMemorySessionStore) Delete(sessionID string) {
	m.data.Remove(sessionID)
}

// Get
func (m *InMemorySessionStore) Get(sessionID string) (interface{}, error) {
	e, ok := m.data.Get(sessionID)
	if !ok {
		return nil, errors.New("Session does not exist in SessionStore")
	}

	return e, nil
}
