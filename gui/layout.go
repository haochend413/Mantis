package gui

import (
	"fmt"

	"github.com/haochend413/mantis/controllers"
	"github.com/haochend413/mantis/db"
	"github.com/haochend413/mantis/db/dbstructs"
	"github.com/jroimartin/gocui"
)

var FIRST_INIT_CHECK bool = true

// Define layout for all views;
func (gui *Gui) layout(g *gocui.Gui) error {
	//init template
	if FIRST_INIT_CHECK {
		gui.windows = gui.CreateWindowTemplates()
	}

	//here, only check logic
	// init views
	for _, w := range gui.windows {
		// fmt.Fprint(os.Stdout, w.Name)
		if !w.OnDisplay {
			// Don't show views that are off
			continue
		}
		//here it set up view prepare;
		// Only initialize if the view was just created

		v, err := gui.prepareView(w)
		if !w.OnDisplay {
			g.DeleteView(w.Name)
		}
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}

		//check: init if only first created
		if err == gocui.ErrUnknownView {
			//view config
			v.Title = w.Title
			w.View = v
			if w.Editable {
				v.Editable = true
			}
			if w.Scroll {
				v.Autoscroll = true
			}
			if w.Cursor {
				controllers.CursorOn(g, v)
			}
		}

		if w.Name == "note-history" {
			nh, e := g.View("note-history")
			nh.Clear()
			if e != nil {
				return e
			}
			var history []dbstructs.Note
			result := db.DBs.NoteDB.Db.Find(&history)

			//display history
			for _, note := range history {
				fmt.Fprintln(v, note.Content)
			}

			if result.Error != nil {
				return result.Error
			}
		}

	}

	//setstartview
	if FIRST_INIT_CHECK {
		g.SetCurrentView("note")
		FIRST_INIT_CHECK = false
	}

	return nil
}
