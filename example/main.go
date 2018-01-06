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