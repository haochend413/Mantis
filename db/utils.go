package db

import (
	"log"

	"github.com/haochend413/mantis/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// // transfer input data to database;
// func AddNote(content string) error {
// 	//init note struct
// 	note := &Note{Content: content}
// 	//pass the string to database;
// 	result := NoteDB.Create(note)
// 	return result.Error
// }

// // // Display all notes added;
// // func Display() error {

// // }

// func (nd *DataBase) DBAdd(content string) error {
// 	//init note struct
// 	note := &dbstructs.Note{Content: content}
// 	//pass the string to database;
// 	result := nd.Db.Create(note)
// 	return result.Error
// }

func DBInit(name string) *gorm.DB {
	// open notes database
	n, err := gorm.Open(sqlite.Open(name+".db"), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	//assign
	n.AutoMigrate(&models.Note{})
	return n
}
