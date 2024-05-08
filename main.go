package main

import "github.com/forgewarden/falloutui/pipboy"

func main() {
	pipBoyOs := pipboy.NewPipBoyOS()

	err := pipBoyOs.Run()
	if err != nil {
		panic(err)
	}
}
