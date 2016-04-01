package main

import (
	"log"

	"github.com/itpkg/web"
	_ "github.com/itpkg/web/engines/blog"
	_ "github.com/itpkg/web/engines/books"
	_ "github.com/itpkg/web/engines/cms"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	if err := web.Main(); err != nil {
		log.Fatal(err)
	}
}
