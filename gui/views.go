package gui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/haochend413/mantis/models"
)

// type windowInfoMapping struct {
// 	window *models.Window
// 	name   string
// }

// // Use pointer: manage the real state;
// func (gui *Gui) MapWindowNames() []*windowInfoMapping {
// 	return []*windowInfoMapping{
// 		{window: gui.windows.Note, name: "note"},
// 		{window: gui.windows.NoteHistory, name: "note-history"},
// 		{window: gui.windows.Cmd, name: "cmd"},
// 	}
// }

func (gui *Gui) CreateWindowTemplates() []*models.Window {
	maxX, maxY := gui.g.Size()
	//here, by default, we leave OnDisplay to be false and view to be uninitialized. This is just render info
	return []*models.Window{
		{
			//0
			Name:      "note",
			Title:     "[0]-Note",
			OnDisplay: true,
			X0:        1,
			Y0:        maxY / 5 * 4,
			X1:        maxX - 1,
			Y1:        maxY - 1,
			Editable:  true,
			Scroll:    true,
			Cursor:    true,
		},
		{
			//1
			Name:      "note-history",
			Title:     "[1]-Note-History",
			OnDisplay: true,
			X0:        1,
			Y0:        1,
			X1:        maxX / 3,
			Y1:        maxY/5*4 - 1,
			Scroll:    false,
			Cursor:    true,
			// Editable:  true,
		},
		{
			//2
			Name:     "cmd",
			Title:    "Cmd",
			X0:       20,
			Y0:       maxY/2 - 1,
			X1:       maxX - 20,
			Y1:       maxY/2 + 1,
			Editable: true,
			Cursor:   true,
		},
		{
			//3
			Name:      "note-detail",
			Title:     "[2]-Note-Detail",
			OnDisplay: true,
			X0:        maxX/3 + 1,
			Y0:        1,
			X1:        maxX - 1,
			Y1:        maxY/2 - 1,
			Editable:  true,
			Cursor:    true,
		},
	}
}

func (gui *Gui) prepareView(w *models.Window) (*gocui.View, error) {
	// arbitrarily giving the view enough size so that we don't get an error, but
	// it's expected that the view will be given the correct size before being shown
	return gui.g.SetView(w.Name, w.X0, w.Y0, w.X1, w.Y1, 0)
}

// // create windows and include configs
// func (gui *Gui) createAllWindows() error {
// 	gui.windows = gui.CreateWindowTemplates()
// 	return nil
// }
