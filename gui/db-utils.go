package gui

import (
	"github.com/haochend413/mantis/controllers"
	"github.com/haochend413/mantis/db"
)

// send note to db
func (gui *Gui) SendNote(content string) error {

	s := FetchContent(gui.windows[0], gui.G())

	if err := db.DBs.NoteDB.AddNote(s); err != nil {
		return err
	}

	gui.g.CurrentView().Clear()
	controllers.CursorOn(gui.g, gui.g.CurrentView())
	return nil
}
