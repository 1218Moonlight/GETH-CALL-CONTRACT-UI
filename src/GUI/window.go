package GUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"log"
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

func (w window) onClosingAndQuit(){
	w.main.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		log.Println("OnClosing")
		return true
	})

	ui.OnShouldQuit(func() bool {
		w.main.Destroy()
		log.Println("OnShouldQuit")
		return true
	})
}

func (w window) show(){
	w.main.Show()
}

func (w window) setChild(child ui.Control){
	w.main.SetChild(child)
}