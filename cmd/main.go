package main

import (
	"fmt"

	"github.com/patrickphan01/short-link/internal/database"
	"github.com/patrickphan01/short-link/internal/routers"
)

func main() {
	database.NewSQLite()
	r := routers.InitRouters()
	fmt.Println("Server listen on port:8080")
	r.Run()

}
