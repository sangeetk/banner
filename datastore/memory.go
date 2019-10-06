package datastore

import (
	"errors"
)

type Memory struct {
	storage map[string]interface{}
}

func NewMemoryStorage() *Memory {
	return &Memory{storage: make(map[string]interface{})}
}

func (m *Memory) Put(key string, data interface{}) error {
	m.storage[key] = data
	return nil
}

func (m *Memory) Get(key string) (interface{}, error) {
	var data interface{}
	var ok bool
	if data, ok = m.storage[key]; !ok {
		return nil, errors.New("Not found")
	}
	return data, nil
}

func (m *Memory) Del(key string) error {
	if _, ok := m.storage[key]; ok {
		delete(m.storage, key)
	}
	return nil
}
