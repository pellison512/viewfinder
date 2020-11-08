package data

type DataSvc interface {
	StoreWindow(window WindowData) (err error)
	GetWindow(windowName string) (WindowData, error)
}

type WindowData struct {
	Title  string
	Left   int
	Top    int
	Right  int
	Bottom int
}
