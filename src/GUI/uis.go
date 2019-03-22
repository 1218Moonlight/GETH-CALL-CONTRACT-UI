package GUI

import "github.com/andlabs/ui"

func windowUI() {
	win := newWindow()
	win.onClosingAndQuit()

	topVbox := ui.NewVerticalBox()
	topVbox.SetPadded(true)

	ipBtn := ui.NewButton("testBtn")
	topVbox.Append(ipBtn, true)

	win.setChild(topVbox)

	win.show()
}