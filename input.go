package input

import (
	"fmt"
	"github.com/ProductionPanic/go-pretty"
)

type TextInput struct {
	Prompt  string
	Default string
}

func GetText(prompt string) string {
	return NewTextInput(prompt).Run()
}

func GetInt(prompt string) int {
	return NewTextInput(prompt).RunInt()
}

func NewTextInput(prompt string) *TextInput {
	return &TextInput{
		Prompt: prompt,
	}
}

func (t *TextInput) Run() string {
	var output string
	pretty.Println(t.Prompt)
	fmt.Scanln(&output)
	return output
}

func (t *TextInput) RunInt() int {
	output := t.Run()
	var outInt int
	fmt.Sscanf(output, "%d", &outInt)
	return outInt
}
