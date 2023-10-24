package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/ProductionPanic/go-pretty"
)

type TextInput struct {
	Prompt  string
	Default string
}

func GetText(prompt string) (string, error) {
	return NewTextInput(prompt).Run()
}

func GetInt(prompt string) (int, error) {
	return NewTextInput(prompt).RunInt()
}

func NewTextInput(prompt string) *TextInput {
	return &TextInput{
		Prompt: prompt,
	}
}

func (t *TextInput) __get_input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func (t *TextInput) RunInt() (int, error) {
	pretty.Println(t.Prompt)
	text, err := t.__get_input()
	if err != nil {
		return 0, err
	}
	text = strings.TrimSpace(text)
	if text == "" {
		text = t.Default
	}
	return strconv.Atoi(text)
}

func (t *TextInput) Run() (string, error) {
	pretty.Println(t.Prompt)
	text, err := t.__get_input()
	if err != nil {
		return "", err
	}
	text = strings.TrimSpace(text)
	if text == "" {
		text = t.Default
	}
	return text, nil
}
