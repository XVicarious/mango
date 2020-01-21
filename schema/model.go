package schema

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Chapter struct {
	gorm.Model
	Path    string
	MangaID uint
	Manga   Manga
}

type Manga struct {
	gorm.Model
	Path      string
	Chapters  []Chapter
	LibraryID uint
	Library   Library
}

type Library struct {
	gorm.Model
	Path   string
	Mangas []Manga
}

func (m *Manga) FullPath() string {
	return (m.Library.Path + "/" + m.Path)
}
