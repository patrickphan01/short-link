package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)




func main(){

	r := gin.Default()

	fmt.Println("Server listen on port:8080")
	r.Run()

}