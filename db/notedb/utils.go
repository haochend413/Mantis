package notedb

import (
	"github.com/haochend413/mantis/db/dbstructs"
)

func (nd *NoteDB) AddNote(content string) error {
	//init note struct
	if content == "" {
		return nil
	}
	note := &dbstructs.Note{Content: content}
	//pass the string to database;
	result := nd.Db.Create(note)
	return result.Error
}
