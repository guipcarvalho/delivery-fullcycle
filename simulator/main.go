package main

import (
	"fmt"

	"github.com/guipcarvalho/delivery-fullcycle-simulator/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()

	stringJson, _ := route.ExportJsonPositions()

	fmt.Println(stringJson[0])
}
