package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/xvicarious/mango/crawler"
	"github.com/xvicarious/mango/schema"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Libraries []string `yaml:"libraries"`
}

func (c *conf) openConf() *conf {
	yamlFile, err := ioutil.ReadFile("mango.yml")
	if err != nil {
		log.Printf("readfile: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("unmarshal: %v", err)
	}
	return c
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {
	var c conf
	c.openConf()
	db, err := gorm.Open("sqlite3", "mango.sqlite3")
	if err != nil {
		panic("Failed to connect to db")
	}
	defer db.Close()
	db.AutoMigrate(&schema.Library{})
	db.AutoMigrate(&schema.Manga{})
	db.AutoMigrate(&schema.Chapter{})
	for _, libraryMod := range c.Libraries {
		db.FirstOrCreate(
			&schema.Library{Path: libraryMod},
		)
	}
	var library schema.Library
	db.First(&library, 1)
	crawler.ReadLibraryPath(&library, &db)
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":4545", nil))
}
