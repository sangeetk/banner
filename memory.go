package banner

import (
	"errors"
)

// DataStoreMemory based implemntation of DataStore
type DataStoreMemory struct {
	storage map[string]interface{}
}

// NewDataStoreMemory allocates memory for in-memory datasotre
func NewDataStoreMemory() *DataStoreMemory {
	return &DataStoreMemory{storage: make(map[string]interface{})}
}

// Put stores the banner into datastore
func (m *DataStoreMemory) Put(key string, data interface{}) error {
	m.storage[key] = data
	return nil
}

// Get reads the banner from the datastore
func (m *DataStoreMemory) Get(key string) (interface{}, error) {
	var data interface{}
	var ok bool
	if data, ok = m.storage[key]; !ok {
		return nil, errors.New(ErrorNotFound)
	}
	return data, nil
}

// Del deletes the banner from the datastore
func (m *DataStoreMemory) Del(key string) error {
	if _, ok := m.storage[key]; ok {
		delete(m.storage, key)
	}
	return nil
}
