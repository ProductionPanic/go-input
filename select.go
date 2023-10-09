package input

import (
	"github.com/ProductionPanic/go-cursor"
	"github.com/ProductionPanic/go-pretty"
	"github.com/eiannone/keyboard"
)

type Select struct {
	Prompt   string
	List     []ListItem
	Selected int
}

type ListItem struct {
	Name  string
	Value any
}

func NewSelect(prompt string) *Select {
	return &Select{
		Prompt: prompt,
		List:   []ListItem{},
	}
}

func (s *Select) AddItem(name string, value any) {
	s.List = append(s.List, ListItem{name, value})
}

func (s *Select) getKeys() []string {
	var keys []string
	for key := range s.List {
		keys = append(keys, s.List[key].Name)
	}
	return keys
}

func (s *Select) getValueByKey(key string) any {
	for _, value := range s.List {
		if value.Name == key {
			return value
		}
	}
	return nil
}

func (s *Select) getValueByIndex(index int) any {
	var _i int
	for _, value := range s.List {
		if _i == index {
			return value.Value
		}
		_i++
	}
	return nil
}

func (s *Select) Run() any {
	cursor.Hide()
	defer cursor.Show()
	keys := s.getKeys()
	for {
		pretty.Println(s.Prompt)
		for i, key := range keys {
			if i == s.Selected {
				pretty.Printf("> %s\n", key)
			} else {
				pretty.Printf("  %s\n", key)
			}
		}
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyArrowUp {
			if s.Selected > 0 {
				s.Selected--
			}
		}
		if key == keyboard.KeyArrowDown {
			if s.Selected < len(keys)-1 {
				s.Selected++
			}
		}
		if key == keyboard.KeyEnter {
			return s.getValueByIndex(s.Selected)
		}
		if char == 'q' {
			return nil
		}

		cursor.Up(len(keys) + 1)
	}
}
