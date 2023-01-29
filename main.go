package main

import (
	. "data-collector/asset"
	_ "data-collector/database"
	orm "data-collector/database"
	"data-collector/router"
	. "data-collector/utils"
	"fmt"
)

func main() {

	defer orm.Conn.Close()
	router := router.InitRouter()
	fmt.Println(GetBanner())
	fmt.Println(LocalIP())
	router.Run(":8000")
}
