package main

import (
	"github.com/go-gui-org/go-glyph"
	"github.com/go-gui-org/go-gui/gui"
	"github.com/go-gui-org/go-gui/gui/backend"
)

type App struct {
	Nombre     string
	Mensaje    string
	ColorTexto gui.Color
}

var app = App{}

func main() {
	gui.SetTheme(gui.ThemeDark.WithBorders(true))

	w := gui.NewWindow(gui.WindowCfg{
		State:  app,
		Title:  "Pruebas",
		Width:  400,
		Height: 300,
		OnInit: func(w *gui.Window) {
			w.UpdateView(mainView)
		},
	})
	app.Mensaje = "Hola majete!!"
	app.ColorTexto = gui.ColorFromString("#ff9f00")

	backend.Run(w)
}

func mainView(w *gui.Window) gui.View {
	ww, _ := w.WindowSize()
	//app := gui.State[App](w)

	return gui.Column(gui.ContainerCfg{
		Width:  float32(ww),
		Height: 300,
		Sizing: gui.FillFit,
		HAlign: gui.HAlignCenter,
		VAlign: gui.VAlignTop,
		Content: []gui.View{
			gui.Input(gui.InputCfg{
				ID:          "Mensaje",
				IDFocus:     1,
				Sizing:      gui.FillFit,
				Placeholder: "Mensaje",
				Text:        app.Nombre,
				OnTextChanged: func(l *gui.Layout, s string, w *gui.Window) {
					app.Nombre = s
				},
			}),
			gui.Text(gui.TextCfg{
				Text: app.Mensaje,
				TextStyle: gui.TextStyle{
					Size:     24,
					Color:    app.ColorTexto,
					Typeface: glyph.TypefaceBold,
					Align:    gui.TextAlignCenter,
				},
				Mode: gui.TextModeWrap,
			}),
			gui.Button(gui.ButtonCfg{
				Content: []gui.View{
					gui.Text(gui.TextCfg{Text: "Cambiar texto"}),
				},
				OnClick: func(l *gui.Layout, e *gui.Event, w *gui.Window) {
					app.Mensaje = "Eres un puto gilipollas😂, " + app.Nombre
					app.ColorTexto = gui.ColorFromString("#48ff00")
					app.Nombre = ""
					e.IsHandled = true
				},
			}),
		},
	})
}
