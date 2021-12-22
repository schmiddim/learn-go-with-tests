package main

type Store interface {
	Fetch() string
	Cancel()
}
type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {

	return s.response
}

func (s *StubStore) Cancel() {

}
