package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ProductionPanic/go-pretty"
	"golang.org/x/crypto/ssh/terminal"
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

type PasswordInput struct {
	Prompt  string
	Default string
}

func NewPasswordInput(prompt string) *PasswordInput {
	return &PasswordInput{
		Prompt: prompt,
	}
}

func (p *PasswordInput) Run() (string, error) {
	fmt.Print(p.Prompt)
	bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	password := string(bytePassword)
	password = strings.TrimSpace(password)
	if password == "" {
		password = p.Default
	}
	return password, nil
}

func GetPassword(prompt string) (string, error) {
	return NewPasswordInput(prompt).Run()
}
