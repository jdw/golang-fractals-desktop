package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppSettings struct {
	Width               float64
	Height              float64
	MaxIter             float32
	Fullscreen          bool
	OrigoCap            float64
	ScreenOffset        PositionF64
	Scale               float64
	ModelDimensions     PositionI64
	SettingsWindowTitle string
	ViewerWindowTitle   string
	TeamName            string
	TeamPassword        string
}

func SettingsDisplay(w fyne.Window) {
	widthEntry := widget.NewEntry()
	heightEntry := widget.NewEntry()
	iterEntry := widget.NewEntry()
	fullEntry := widget.NewCheck("", func(b bool) {
		glob.Fullscreen = b

	})
	teamNameEntry := widget.NewEntry()
	teamNameEntry.SetPlaceHolder("beefbabe")
	teamPassword := widget.NewEntry()

	widthEntry.Text = strconv.Itoa(int(glob.Width))
	heightEntry.Text = strconv.Itoa(int(glob.Height))
	iterEntry.Text = strconv.Itoa(int(glob.MaxIter))

	currentErrorLabel := widget.NewLabel("")

	// Function to update the settings
	updateRet := func() {
		widthStr := widthEntry.Text

		if widthStr == "" {
			widthEntry.Text = strconv.Itoa(int(glob.Width))
			return
		}

		width, err := strconv.Atoi(widthStr)
		if err != nil {
			println("Error parsing width:", err.Error())
			currentErrorLabel.SetText("Error parsing width, see log for details!")
			return
		}

		if width > 100 {
			glob.Width = float64(width)
		} else {
			currentErrorLabel.SetText("Width must be atleast 100")
			return
		}

		heightStr := heightEntry.Text

		if heightStr == "" {
			heightEntry.Text = strconv.Itoa(int(glob.Height))
			return
		}

		height, err := strconv.Atoi(heightStr)
		if err != nil {
			println("Error parsing height:", err.Error())
			currentErrorLabel.SetText("Error parsing height, see log for details!")
			return
		}

		if height > 100 {
			glob.Height = float64(height)
		} else {
			currentErrorLabel.SetText("Height must be atleast 100")
			return
		}

		iterStr := iterEntry.Text

		if iterStr == "" {
			iterEntry.Text = strconv.Itoa(int(glob.MaxIter))
			return
		}

		iter, err := strconv.Atoi(iterStr)
		if err != nil {
			println("Error parsing iterations:", err.Error())
			currentErrorLabel.SetText("Error parsing iterations, see log for details!")
			return
		}

		if iter > 1 {
			glob.MaxIter = float32(iter)
		} else {
			currentErrorLabel.SetText("Iterations must be atleast 1")
			return
		}

		if teamNameEntry.Text != "" {
			glob.TeamName = teamNameEntry.Text
		} else {
			currentErrorLabel.SetText("Team name must not be empty")
			return
		}
		if teamPassword.Text != "" {
			glob.TeamPassword = teamPassword.Text
		} else {
			currentErrorLabel.SetText("Team password must not be empty")
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
		container.NewHBox(widget.NewLabel("Team name:"), teamNameEntry),
		container.NewHBox(widget.NewLabel("Team password:"), teamPassword),
		applyButton,
		currentErrorLabel,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 100)) // Initial window size
	w.CenterOnScreen()

	w.Show()
}
