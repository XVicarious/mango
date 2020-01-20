package schema

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Chapter struct {
	gorm.Model
	Path string
}

type Manga struct {
	gorm.Model
	Path     string
	Chapters []Chapter
}

type Library struct {
	gorm.Model
	Path  string
	Manga []Manga
}

func Makestuff() int {
	return 0
}
