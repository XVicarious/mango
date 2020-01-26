package crawler

import (
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/xvicarious/mango/schema"
)

func ReadLibraryPath(library *schema.Library, db **gorm.DB) {
	directory, err := os.Open(library.Path)
	if err != nil {
		log.Fatalf("error open directory: %v", err)
	}
	defer directory.Close()
	list, _ := directory.Readdirnames(0)
	log.Printf("Found %s manga", strconv.Itoa(len(list)))
	for _, mangaDir := range list {
		var manga schema.Manga
		(*db).FirstOrCreate(&manga, schema.Manga{
			Path: mangaDir,
		})
		(*db).Model(&library).Association("Mangas").Append(&manga)
		(*db).Model(&manga).Association("Library").Append(&(*library)) // Lol. Idk what I'm doing
		// I'm sure there is a better way to do this, but this is only day 2 of Go.
		ReadMangaPath(&manga, db)
	}
}

func ReadMangaPath(manga *schema.Manga, db **gorm.DB) {
	directory, err := os.Open(manga.FullPath())
	if err != nil {
		log.Fatalf("error opening directory: %v", err)
	}
	defer directory.Close()
	list, _ := directory.Readdirnames(0)
	log.Printf("%s: found %s chapters", manga.Path, strconv.Itoa(len(list)))
	for _, chapterDir := range list {
		var chapter schema.Chapter
		(*db).FirstOrCreate(&chapter, schema.Chapter{
			Path: chapterDir,
		})
		(*db).Model(&manga).Association("Chapters").Append(&chapter)
		(*db).Model(&chapter).Association("Manga").Append(&(*manga))
	}
}
