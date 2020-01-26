package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDB() {
	db.AutoMigrate(&Library{})
	db.AutoMigrate(&Manga{})
	db.AutoMigrate(&Chapter{})
}

func OpenDB() {
	db, _ = gorm.Open("sqlite3", "mango.sqlite3")
}

func CloseDB() {
	db.Close()
}

func CreateManga(mangaPath string, library *Library) (Manga, error) {
	var manga Manga
	db.FirstOrInit(&manga, Manga{
		Path: mangaPath,
	})
	if manga.ID != 0 {
		return manga, fmt.Errorf("%s is already in the database.", mangaPath)
	}
	db.Create(&manga).Association("Library").Append(library)
	db.Model(library).Association("Mangas").Append(&manga)
	return manga, nil
}

func CreateChapter(chapterPath string, manga *Manga) (Chapter, error) {
	var chapter Chapter
	db.FirstOrInit(&chapter, Chapter{
		Path: chapterPath,
	})
	if chapter.ID != 0 {
		return chapter, errors.New("The Chapter is alredy in the database.")
	}
	log.Printf("Creating %s %s", manga.Path, chapter.Path)
	db.Create(&chapter).Association("Manga").Append(manga)
	db.Model(manga).Association("Chapters").Append(&chapter)
	return chapter, nil
}

func GetLibrary(library Library, libraryId uint) {
	db.First(&library, libraryId)
}

func GetDB() *gorm.DB {
	return db
}
