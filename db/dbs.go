package db

import (
	"github.com/haochend413/mantis/db/notedb"
)

var DBs *DataBases

type DataBases struct {
	NoteDB *notedb.NoteDB
}

// var _ db_models.DBController = (*NoteDB)(nil)

func InitAll() {
	DBs = &DataBases{}
	DBs.NoteDB = &notedb.NoteDB{}
	DBs.NoteDB.Db = DBInit("notes")
}

func CloseAll() {
	_ = DBs.NoteDB.Close()
}
