package main

import (
	"fmt"
	"hostinfo/app"
	"log"
	"os"
)

func main() {
	fmt.Println("")

	aplication := app.Generate()

	if err := aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
