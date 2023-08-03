package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("hii !!")

	text := widget.NewLabel("\n\n\n\nhii !! \nanother app using go \nalso first time using fyne package!!  ")

	img := canvas.NewImageFromFile("image.png")
	img.FillMode = canvas.ImageFillOriginal
	img.SetMinSize(fyne.NewSize(200, 200))

	content := container.NewHBox(img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
