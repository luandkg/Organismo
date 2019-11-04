package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

var caminho string = "assets/organismos/"

func setup_window(title string) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(400, 600)
	win.SetPosition(gtk.WIN_POS_CENTER)
	return win
}

func setup_box(orient gtk.Orientation) *gtk.Box {
	box, err := gtk.BoxNew(orient, 0)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	return box
}



func setup_btn(label string, onClick func()) *gtk.Button {
	btn, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	btn.Connect("clicked", onClick)
	return btn
}




func main() {

	gtk.Init(nil)

	win := setup_window("Criador de Organismos")

	boxPrincipal,_:= gtk.BoxNew(gtk.ORIENTATION_VERTICAL,5)

	boxEspacador1,_:= gtk.BoxNew(gtk.ORIENTATION_VERTICAL,5)

	boxPrincipal.Add(boxEspacador1)

	produtores(boxPrincipal)
	consumidores(boxPrincipal)

	win.Add(boxPrincipal)

	win.ShowAll()

	gtk.Main()

}



