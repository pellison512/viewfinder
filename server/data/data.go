package data

type DataSvc interface {
	StoreWindow(window WindowData) (err error)
	GetWindow(windowName string) (WindowData, error)
	//TODO remove this after testing
	GetAllWindows() ([]WindowData, error)
}

type WindowData struct {
	Title  string
	Left   int
	Top    int
	Right  int
	Bottom int
}

type Polygon struct {
	Coordinates []Coodinate
}

type Coodinate struct {
	X int;
	Y int;
}