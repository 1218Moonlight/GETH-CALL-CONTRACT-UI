package GUI

import (
	"github.com/andlabs/ui"
	"fmt"
)

type uiBox struct {
	mainVerticalBox *ui.Box
	caEnter         *ui.Entry
	eoaEnter        *ui.Entry
}

func newUiBoxManager() uiBox {
	boxManager := uiBox{mainVerticalBox: func() *ui.Box {
		box := ui.NewVerticalBox()
		box.SetPadded(true)
		return box
	}(), caEnter: func() *ui.Entry {
		return ui.NewEntry()
	}(), eoaEnter: func() *ui.Entry {
		return ui.NewEntry()
	}()}
	return boxManager
}

func windowUI() {
	win := newWindow()
	win.onClosingAndQuit()

	///

	boxManager := newUiBoxManager()

	boxManager.caEntry()

	boxManager.eoaEntry()

	boxManager.erc20Info()

	///

	win.setChild(boxManager.mainVerticalBox)
	win.show()
}

func (b uiBox) caEntry() {
	name := ui.NewLabel(" [ CA ] ")

	hBox := ui.NewHorizontalBox()
	hBox.SetPadded(true)

	hBox.Append(name, false)
	hBox.Append(b.caEnter, true)

	b.mainVerticalBox.Append(hBox, false)
}

func (b uiBox) eoaEntry() {
	name := ui.NewLabel(" [ EOA ] ")

	hBox := ui.NewHorizontalBox()
	hBox.SetPadded(true)

	hBox.Append(name, false)
	hBox.Append(b.eoaEnter, true)

	b.mainVerticalBox.Append(hBox, false)
}

func (b uiBox) erc20Info() {
	requestBtn := ui.NewButton("request")
	b.mainVerticalBox.Append(requestBtn, false)

	infoEntry := ui.NewEntry()
	infoEntry.SetReadOnly(true)

	requestBtn.OnClicked(func(button *ui.Button) {
		if b.caEnter.Text() == "" {
			infoEntry.SetText("Empty CA")
			return
		} else if b.eoaEnter.Text() == "" {
			infoEntry.SetText("Empty EOA")
			return
		}

		symbol, err := symbol(b.caEnter.Text())
		if err != nil {
			fmt.Println(err)
		}

		balance, err := balanceOf(b.caEnter.Text(), b.eoaEnter.Text())
		if err != nil {
			fmt.Println(err)
		}

		infoEntry.SetText(fmt.Sprintf("[ symbol ] : %s, [ balanceOf ] : %v", *symbol, balance))
	})

	b.mainVerticalBox.Append(infoEntry, true)
}