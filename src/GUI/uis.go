package GUI

import "github.com/andlabs/ui"

type uiBox struct {
	mainVerticalBox   *ui.Box
}

func windowUI() {
	win := newWindow()
	win.onClosingAndQuit()

	boxManager := uiBox{mainVerticalBox:InitMainVerticalBox()}

	boxManager.Entry("CA")
	boxManager.Entry("EOA")
	boxManager.Erc20Info()

	win.setChild(boxManager.mainVerticalBox)

	win.show()
}

func InitMainVerticalBox()*ui.Box{
	box := ui.NewVerticalBox()
	box.SetPadded(true)
	return box
}

func (b uiBox) Entry(labelName string) {
	entry := ui.NewEntry()

	name := ui.NewLabel(" [ "+labelName+" ] ")

	hBox := ui.NewHorizontalBox()
	hBox.SetPadded(true)

	hBox.Append(name, false)
	hBox.Append(entry, true)

	b.mainVerticalBox.Append(hBox, false)
}

func(b uiBox) Erc20Info(){
	infoEntry := ui.NewEntry()
	infoEntry.SetReadOnly(true)
	infoEntry.SetText("hello")

	b.mainVerticalBox.Append(infoEntry, true)
}