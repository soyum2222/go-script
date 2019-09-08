package main

import (
	"fmt"
	"go-script/scripte"
	"time"
)

func main() {

	for {
		script := scripte.NewScript("example/script/script")

		script.Arge("go-script")
		decoder, err := script.Run()
		if err != nil {
			panic(err)
		}

		var r int
		decoder.Decode(&r)

		fmt.Println("go-script return :", r)

		time.Sleep(5 * time.Second)

	}

}
