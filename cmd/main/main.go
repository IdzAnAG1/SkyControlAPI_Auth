package main

import (
	"sc_auth/internal/entry/app"
)

func main() {
	a := app.NewApp(9060)
	err := a.Run()
	if err != nil {
		panic(err)
	}
}
