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
	sel := &Select{
		Prompt:   prompt,
		Selected: 0,
		List:     make([]ListItem, 0),
	}
	return sel
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
	keys := s.getKeys()
	for {
		pretty.Println("[blue,bold]" + s.Prompt + "[reset,dim,white] (q to quit)[reset]")
		for i, key := range keys {
			if i == s.Selected {
				pretty.Printf("[cyan]>[reset]  %s\n", key)
			} else {
				pretty.Printf("   %s\n", key)
			}
		}
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			cursor.Show()
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
			cursor.Show()
			return s.getValueByIndex(s.Selected)
		}
		if char == 'q' {
			cursor.Show()
			return nil
		}

		cursor.Up(len(keys) + 1)
	}

}
