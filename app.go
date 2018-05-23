package main

import (
	"github.com/fiscaluno/shaka/server"
	"github.com/fiscaluno/shaka/user"
)

func main() {
	user.Migrate()
	server.Listen()
}
