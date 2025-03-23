package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppSettings struct {
	Width        int
	Height       int
	MaxIter      int
	Fullscreen   bool
	OrigoCap     float64
	ScreenOffset PositionI64
	Scale        float64
}

func SettingsDisplay(w fyne.Window) {
	widthEntry := widget.NewEntry()
	//widthEntry.Resize(fyne.NewSize(200, widthEntry.MinSize().Height))
	heightEntry := widget.NewEntry()
	iterEntry := widget.NewEntry()
	fullEntry := widget.NewCheck("", func(b bool) {
		glob.Fullscreen = b

	})
	widthEntry.Text = strconv.Itoa(glob.Width)
	heightEntry.Text = strconv.Itoa(glob.Height)
	iterEntry.Text = strconv.Itoa(glob.MaxIter)

	currentErrorLabel := widget.NewLabel("")

	// Function to update the settings
	updateRet := func() {
		widthStr := widthEntry.Text

		if widthStr == "" {
			widthEntry.Text = strconv.Itoa(glob.Width)
			return
		}

		width, err := strconv.Atoi(widthStr)
		if err != nil {
			println("Error parsing width:", err.Error())
			currentErrorLabel.SetText("Error parsing width, see log for details!")
			return
		}

		if width > 100 {
			glob.Width = width
		} else {
			currentErrorLabel.SetText("Width must be atleast 100")
			return
		}

		heightStr := heightEntry.Text

		if heightStr == "" {
			heightEntry.Text = strconv.Itoa(glob.Height)
			return
		}

		height, err := strconv.Atoi(heightStr)
		if err != nil {
			println("Error parsing height:", err.Error())
			currentErrorLabel.SetText("Error parsing height, see log for details!")
			return
		}

		if height > 100 {
			glob.Height = height
		} else {
			currentErrorLabel.SetText("Height must be atleast 100")
			return
		}

		iterStr := iterEntry.Text

		if iterStr == "" {
			iterEntry.Text = strconv.Itoa(glob.MaxIter)
			return
		}

		iter, err := strconv.Atoi(iterStr)
		if err != nil {
			println("Error parsing iterations:", err.Error())
			currentErrorLabel.SetText("Error parsing iterations, see log for details!")
			return
		}

		if iter > 1 {
			glob.MaxIter = iter
		} else {
			currentErrorLabel.SetText("Iterations must be atleast 1")
			return
		}

		w.Close()
	}

	// Create a button to trigger the width update
	applyButton := widget.NewButton("Render!", updateRet)

	// Arrange the widgets in a container
	content := container.NewVBox(
		container.NewHBox(widget.NewLabel("Viewer window width (int):"), widthEntry),
		container.NewHBox(widget.NewLabel("Viewer window heigth (int):"), heightEntry),
		container.NewHBox(widget.NewLabel("Max iterations (int):"), iterEntry),
		container.NewHBox(widget.NewLabel("Fullscreen:"), fullEntry),
		applyButton,
		currentErrorLabel,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 100)) // Initial window size
	w.CenterOnScreen()

	w.Show()
}
