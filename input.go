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

func (t *TextInput) RunInt() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	i, _ := strconv.Atoi(strings.TrimSpace(text))
	return i
}
func (t *TextInput) Run() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
