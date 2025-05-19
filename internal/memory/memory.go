package memory

import "sync"

type Store struct {
    mu       sync.Mutex
    sessions map[string][]string
}

func NewStore() *Store {
    return &Store{
        sessions: make(map[string][]string),
    }
}

func (s *Store) Store(session, message string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.sessions[session] = append(s.sessions[session], message)
}

func (s *Store) Retrieve(session string) []string {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.sessions[session]
}