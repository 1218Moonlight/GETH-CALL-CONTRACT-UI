package GUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

type window struct {
	main      *ui.Window
	configure windowConfigure
}

type windowConfigure struct {
	title      string
	width      int
	height     int
	hasMenubar bool
}

func newWindow() window {
	win := window{
		main:      nil,
		configure: windowConfigure{title: title, height: height, width: width, hasMenubar: hasMenubar},
	}

	win.main = ui.NewWindow(win.configure.title, win.configure.width, win.configure.height, win.configure.hasMenubar)
	return win
}
