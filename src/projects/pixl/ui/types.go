package ui

import (
	"fyne.io/fyne/v2"
	"zerotomastery.io/pixl/apptype"
	"zerotomastery.io/pixl/swatch"
)

type AppIntit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swtches    []*swatch.Swatch
}
