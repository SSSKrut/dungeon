package main

type Storage interface {
	Add(Node) *Node
	Get(id string) (Node, bool)
	Delete(id string)
}

type InMemoryStorage struct {
	store map[string]Node
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{store: make(map[string]Node)}
}

func (s *InMemoryStorage) Add(n Node) *Node {
	s.store[n.ID] = n
	return &n
}

func (s *InMemoryStorage) Get(id string) (Node, bool) {
	e, exists := s.store[id]
	return e, exists
}

func (s *InMemoryStorage) Delete(id string) {
	delete(s.store, id)
}
