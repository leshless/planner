package main

import "planner/backend/internal/app"

func main() {
	app, err := app.Init(app.NewDefaultPrimitives())
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
