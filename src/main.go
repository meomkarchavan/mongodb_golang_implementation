package main

import "mongo_go/src/controller"

func main() {

	r := controller.RegisterRoutes()
	r.Run(":8080")

}
