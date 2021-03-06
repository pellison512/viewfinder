package data

import "errors"

type MemStore struct {
	data map[string]WindowData
}

func NewMemStoreDataSvc() *MemStore {
	return &MemStore{
		data: make(map[string]WindowData),
	}
}

func (m *MemStore) StoreWindow(window WindowData) (err error) {
	m.data[window.Title] = window
	return nil
}

func (m *MemStore) GetWindow(windowName string) (WindowData, error) {
	data, found := m.data[windowName]
	if !found {
		return WindowData{}, errors.New("no window found")
	}
	return data, nil
}

func (m *MemStore) GetAllWindows() ([]WindowData, error) {
	windows := make([]WindowData, 0)
	for _, window := range m.data {
		windows = append(windows, window)
	}
	return windows, nil
}