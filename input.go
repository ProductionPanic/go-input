package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ProductionPanic/go-pretty"
	"github.com/eiannone/keyboard"
)

func SelectFromList(prompt string, list []string) int {
	// open a dev/tty file
	var selected int
	cursor := GetCursor()
	cursor.Hide()
	defer cursor.Show()
	for {
		pretty.Println(prompt)
		for i, item := range list {
			cursor.ClearLine()
			if i == selected {
				pretty.Println("[green,bold]>" + item)
			} else {
				pretty.Println(" " + item)
			}
			pretty.Print("[reset]")
		}
		_, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyArrowUp {
			if selected > 0 {
				selected--
			}
		} else if key == keyboard.KeyArrowDown {
			if selected < len(list)-1 {
				selected++
			}
		} else if key == keyboard.KeyEnter {
			break
		}
		// move up and clear line
		linesUp := len(list) + 1
		cursor.Up(linesUp)
	}
	return selected
}
