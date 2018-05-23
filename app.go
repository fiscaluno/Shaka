package main

import (
	"github.com/fiscaluno/shaka/course"
	"github.com/fiscaluno/shaka/server"
)

func main() {
	course.Migrate()
	server.Listen()
}
