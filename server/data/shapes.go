package data



type Rect struct {
	
}

func CreateOverlappingShape(baseWindow WindowData, openWindows []WindowData) {
	//Each window can be deconstructed to 4 coordinates
	baseTopLeft := Coodinate{
		X: baseWindow.Left,
		Y: baseWindow.Top,
	}
	baseTopRight := Coodinate{
		X: baseWindow.Right,
		Y: baseWindow.Top,
	}
	baseBotLeft := Coodinate{
		X: baseWindow.Left,
		Y: baseWindow.Bottom,
	}
	baseBotRight := Coodinate{
		X: baseWindow.Right,
		Y: baseWindow.Bottom,
	}


	for _, window := range openWindows {
		recCords := make([]Coodinate, 0)

		//Each window can be deconstructed to 4 coordinates
		topLeft := Coodinate{
			X: window.Left,
			Y: window.Top,
		}
		topRight := Coodinate{
			X: window.Right,
			Y: window.Top,
		}
		botLeft := Coodinate{
			X: window.Left,
			Y: window.Bottom,
		}
		botRight := Coodinate{
			X: window.Right,
			Y: window.Bottom,
		}


	}
}