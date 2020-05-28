package main

import (
	_ "github.com/Rhizomyidae/rat-server/boot"
	_ "github.com/Rhizomyidae/rat-server/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
