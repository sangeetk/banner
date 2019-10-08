package banner

import (
	"errors"
)

// DataStoreMemory based implemntation of DataStore
type DataStoreMemory struct {
	storage map[string]Banner
}

// NewDataStoreMemory allocates memory for in-memory datasotre
func NewDataStoreMemory() *DataStoreMemory {
	return &DataStoreMemory{storage: make(map[string]Banner)}
}

// Put stores the banner into datastore
func (m *DataStoreMemory) Put(b Banner) error {
	m.storage[b.ID] = b
	return nil
}

// Get reads the banner from the datastore
func (m *DataStoreMemory) Get(id string) (Banner, error) {
	if _, ok := m.storage[id]; !ok {
		return Banner{}, errors.New(ErrorNotFound)
	}
	return m.storage[id], nil
}

// List returns all banners in the datastore.
func (m *DataStoreMemory) List() []Banner {
	banners := []Banner{}
	for _, b := range m.storage {
		banners = append(banners, b)
	}
	return banners
}

// Del deletes the banner from the datastore
func (m *DataStoreMemory) Del(key string) error {
	if _, ok := m.storage[key]; ok {
		delete(m.storage, key)
		return nil
	}
	return errors.New(ErrorNotFound)
}
