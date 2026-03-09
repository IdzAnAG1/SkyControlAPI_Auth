package main

import (
	"sc_auth/internal/entry/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	err = a.Run()
	if err != nil {
		panic(err)
	}
}
