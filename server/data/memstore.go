package data

type memStore struct {
	Data map[string]WindowData 
}

type WindowData struct {
	Title  string
	Left   int
	Top    int   
	Right  int
	Bottom int
}       

func (m *memStore) StoreWindow(window WindowData) (err error) {
	if len(m.Data) == 0 {
		m.Data = make(map[string]WindowData)
	}
	m.Data[window.Title] = window
	return nil
}
