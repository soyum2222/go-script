package main

import (
	"fmt"
	"github.com/soyum2222/go-script/script"
	"time"
)

func main() {

	for {
		script := script.NewScript("example/script/script")

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
