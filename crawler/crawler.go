package crawler

import (
	"log"
	"os"
	"strconv"

	"github.com/xvicarious/mango/database"
)

func ReadLibraryPath(library *database.Library) {
	directory, err := os.Open(library.Path)
	if err != nil {
		log.Fatalf("error open directory: %v", err)
	}
	defer directory.Close()
	list, _ := directory.Readdirnames(0)
	log.Printf("Found %s manga", strconv.Itoa(len(list)))
	for _, mangaDir := range list {
		manga, err := database.CreateManga(mangaDir, library)
		if err != nil {
			log.Print(err)
		}
		if manga.Library.ID == 0 {
			log.Printf("This is odd, the library ID is 0")
			continue
		}
		ReadMangaPath(&manga)
	}
}

func ReadMangaPath(manga *database.Manga) {
	directory, err := os.Open(manga.FullPath() + "/")
	if err != nil {
		log.Fatalf("error opening directory: %v", err)
	}
	defer directory.Close()
	list, _ := directory.Readdirnames(0)
	log.Printf("%s: found %s chapters", manga.Path, strconv.Itoa(len(list)))
	for _, chapterDir := range list {
		_, err := database.CreateChapter(chapterDir, manga)
		if err != nil {
			log.Print(err)
		}
	}
}
