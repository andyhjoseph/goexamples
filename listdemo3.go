package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func listDemo3(pgui pg.ProntoGUI) {
	desc := pg.TextWith{Content: "This is what happens when not TemplateItem is specified."}.Make()
	li1 := pg.TextWith{Content: "A"}.Make()
	li2 := pg.TextWith{Content: "B"}.Make()
	li3 := pg.TextWith{Content: "C"}.Make()
	li4 := pg.TextWith{Content: "D"}.Make()
	li5 := pg.TextWith{Content: "E"}.Make()
	list := pg.ListWith{
		ListItems: []pg.Primitive{li1, li2, li3, li4, li5},
	}.Make()

	pgui.SetGUI(desc, list)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}
