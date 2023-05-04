package sqlite

import (
	"tgbot/internal/clients"
	"tgbot/internal/storage"
)

type slStorage struct {
	path string
}

func New(path string) *slStorage {
	return &slStorage{path: path}
}

func (s *slStorage) Save(m *storage.Msg) (err error) {
	defer func() { err = clients.Wrap("error save", err) }()

	return nil
}

func (s *slStorage) Show(m *storage.Msg) (err error) {
	defer func() { err = clients.Wrap("error show", err) }()
	return nil
}

func (s *slStorage) Remove(tag string) (err error) {
	defer func() { err = clients.Wrap("error remove ", err) }()
	return nil
}
