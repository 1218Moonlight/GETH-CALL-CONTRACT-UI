package GUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"log"
)

func windowUI() {
	win := newWindow()

	win.main.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		log.Println("OnClosing")
		return true
	})

	ui.OnShouldQuit(func() bool {
		win.main.Destroy()
		log.Println("OnShouldQuit")
		return true
	})

	win.main.Show()
}
