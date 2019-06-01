Fontawesome5 definitions for use in Go.

To update, download the latest fontawesome5 in this directory and name it "fontawesome-free".
Now run "make" to generate the definitions and compile.

License: MIT

Example:


	package main

	import (
		"log"

		fa "github.com/mjl-/fontawesome5"
		"github.com/mjl-/duit"
	)

	func check(err error, msg string) {
		if err != nil {
			log.Fatalf("%s: %s\n", msg, err)
		}
	}

	func main() {
		dui, err := duit.NewDUI("fontawesome5", "800x600")
		check(err, "new dui")

		awesome, err := dui.Env.Display.OpenFont("/mnt/font/FontAwesome5FreeRegular/15a/font")
		check(err, "open fontawesome")

		dui.Top = &duit.Field{Font: awesome, Text: string(fa.ThumbsUp)}
		dui.Render()

		for {
			select {
			case e := <-dui.Events:
				dui.Event(e)
			}
		}
	}
