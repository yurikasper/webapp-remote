package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const windowWidth = 500
const settingsWidth = 320
const windowHeight = 520

func startUI() {
	go func() {
		w := app.NewWindow(
			app.Title("WebApp Remote"),
			app.MaxSize(unit.Dp(windowWidth), unit.Dp(windowHeight)),
			app.MinSize(unit.Dp(windowWidth), unit.Dp(windowHeight)),
		)
		err := home(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
}

func openSettings() {
	go func() {
		w := app.NewWindow(
			app.Title("Settings"),
			app.MaxSize(unit.Dp(settingsWidth), unit.Dp(windowHeight)),
			app.MinSize(unit.Dp(settingsWidth), unit.Dp(windowHeight)),
		)
		settings(w)
	}()
}

var darkGray = color.NRGBA{R: 37, G: 37, B: 37, A: 255}
var lightGray = color.NRGBA{R: 115, G: 115, B: 115, A: 255}
var lighterGray = color.NRGBA{R: 155, G: 155, B: 155, A: 255}
var settingsButtonClickable = new(widget.Clickable)

func home(w *app.Window) error {
	w.Perform(system.ActionCenter)
	theme := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			//home background
			defer clip.Rect{Max: image.Pt(windowWidth, windowHeight)}.Push(&ops).Pop()
			paint.ColorOp{Color: darkGray}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			//title placement
			defer op.Offset(image.Pt(0, 25)).Push(&ops).Pop()

			//home title
			title := material.H3(theme, "WebApp Remote")
			title.Color = lighterGray
			title.Alignment = text.Middle
			title.Layout(gtx)

			defer op.Offset(image.Pt(0, 75)).Push(&ops).Pop()

			//Draw IP address
			ipText := material.H4(theme, "http://"+localIp)

			ipText.Color = lightGray
			ipText.Alignment = text.Middle
			ipText.Layout(gtx)

			//center QR code
			defer op.Offset(image.Pt((windowWidth/2)-(QRimage.Bounds().Dx()/2), 75)).Push(&ops).Pop()

			//Draw IP QR code
			imageOp := paint.NewImageOp(QRimage)
			imageOp.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			//Settings button
			defer op.Offset(image.Pt(0, 275)).Push(&ops).Pop()

			settingsButton := material.Button(theme, settingsButtonClickable, "Settings")
			settingsButton.Background = lightGray
			settingsButton.Color = darkGray
			settingsButton.TextSize = unit.Sp(24)
			settingsButton.CornerRadius = unit.Dp(12)

			gtx.Constraints = layout.Exact(image.Point{X: QRimage.Bounds().Dx(), Y: 200})
			layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx, layout.Rigid(settingsButton.Layout))

			if settingsButtonClickable.Clicked() {
				openSettings()
			}

			//pointer cursor on hover
			if settingsButton.Button.Hovered() {
				pointer.CursorPointer.Add(&ops)
			} else {
				pointer.CursorDefault.Add(&ops)
			}

			e.Frame(gtx.Ops)
		}
	}
}

var volumeSettingsGroup = widget.Enum{Value: "default"}
var playSettingsGroup = widget.Enum{Value: "default"}
var seekSettingsGroup = widget.Enum{Value: "default"}
var list = &widget.List{
	List: layout.List{
		Axis: layout.Vertical,
	},
}

type (
	D = layout.Dimensions
	C = layout.Context
)

func settings(w *app.Window) error {
	w.Perform(system.ActionCenter)
	theme := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			fmt.Println("Closed settings")
			fmt.Println(config)
			saveConfig()
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			//home background
			defer clip.Rect{Max: image.Pt(windowWidth, windowHeight)}.Push(&ops).Pop()
			paint.ColorOp{Color: darkGray}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			//title placement
			defer op.Offset(image.Pt(0, 25)).Push(&ops).Pop()

			//home title
			title := material.H3(theme, "Settings")
			title.Color = lighterGray
			title.Alignment = text.Middle
			title.Layout(gtx)

			defer op.Offset(image.Pt(30, 70)).Push(&ops).Pop()

			//titles and radio buttons
			title1 := material.H5(theme, "Volume:")
			radio1a := material.RadioButton(theme, &volumeSettingsGroup, "alternate", "Up/Down")
			radio1b := material.RadioButton(theme, &volumeSettingsGroup, "default", "Media Keys")
			title1.Color = lighterGray
			radio1a.Color = lightGray
			radio1a.IconColor = lightGray
			radio1b.Color = lightGray
			radio1b.IconColor = lightGray

			title2 := material.H5(theme, "Play/Pause:")
			radio2a := material.RadioButton(theme, &playSettingsGroup, "default", "Spacebar")
			radio2b := material.RadioButton(theme, &playSettingsGroup, "alternate", "Media Key")
			title2.Color = lighterGray
			radio2a.Color = lightGray
			radio2a.IconColor = lightGray
			radio2b.Color = lightGray
			radio2b.IconColor = lightGray

			title3 := material.H5(theme, "Forwards/Backwards:")
			radio3a := material.RadioButton(theme, &seekSettingsGroup, "default", "Left/Right")
			radio3b := material.RadioButton(theme, &seekSettingsGroup, "alternate", "Media Keys")
			title3.Color = lighterGray
			radio3a.Color = lightGray
			radio3a.IconColor = lightGray
			radio3b.Color = lightGray
			radio3b.IconColor = lightGray

			widgets := []layout.Widget{
				func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(title1.Layout),
						layout.Rigid(radio1a.Layout),
						layout.Rigid(radio1b.Layout),
					)
				},
				func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(title2.Layout),
						layout.Rigid(radio2a.Layout),
						layout.Rigid(radio2b.Layout),
					)
				},
				func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(title3.Layout),
						layout.Rigid(radio3a.Layout),
						layout.Rigid(radio3b.Layout),
					)
				},
			}

			//pointer cursor on hover
			_, volumeHover := volumeSettingsGroup.Hovered()
			_, playHover := playSettingsGroup.Hovered()
			_, seekHover := seekSettingsGroup.Hovered()
			if volumeHover || playHover || seekHover {
				pointer.CursorPointer.Add(&ops)
			} else {
				pointer.CursorDefault.Add(&ops)
			}

			material.List(theme, list).Layout(gtx, len(widgets), func(gtx C, i int) D {
				return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
			})

			// check for settings changes
			if volumeSettingsGroup.Changed() || playSettingsGroup.Changed() || seekSettingsGroup.Changed() {
				//save config
				saveConfig()
				//update keymap
				refreshKeymap()
			}

			e.Frame(gtx.Ops)
		}
	}
}
