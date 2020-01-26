package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julienschmidt/httprouter"

	"github.com/xvicarious/mango/crawler"
	"github.com/xvicarious/mango/database"
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

func RouteManga(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "manga id %s", ps.ByName("manga_id"))
}

func main() {
	log.Printf("Starting ManGO")
	var c conf
	c.openConf()
	database.OpenDB()
	defer database.CloseDB()
	database.InitDB()
	for _, libraryMod := range c.Libraries {
		database.GetDB().FirstOrCreate(
			&database.Library{Path: libraryMod},
		)
	}
	log.Printf("Reading library.")
	var library database.Library
	database.GetDB().First(&library, 1)
	crawler.ReadLibraryPath(&library)
	// log.Printf("Setting up web server.")
	// router := httprouter.New()
	// router.GET("/manga/:manga_id", RouteManga)
	// router.GET("/manga/:manga_id/:chapter", RouteChapter)
	// router.GET("/manga/:manga_id/:chapter/:page", RoutePage)
	// log.Fatal(http.ListenAndServe(":4545", router))
}
