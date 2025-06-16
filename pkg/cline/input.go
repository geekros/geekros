// Copyright 2025 GEEKROS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cline

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputModel struct {
	Input    textinput.Model
	Callback func()
	err      error
}

func InitInputModel(placeholder string, limit int, width int, callback func()) InputModel {

	textInput := textinput.New()
	textInput.Placeholder = placeholder
	textInput.CharLimit = limit
	textInput.Width = width
	textInput.Focus()

	return InputModel{
		Input:    textInput,
		Callback: callback,
		err:      nil,
	}
}

func (m InputModel) Init() tea.Cmd {

	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (InputModel, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			log.Println(msg.String())
			return m, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case error:
		m.err = msg
		return m, nil
	}

	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}

func (m InputModel) View() string {

	return fmt.Sprintf("What’s your favorite Pokémon?\n\n%s\n\n%s", m.Input.View(), "(esc to quit)") + "\n"
}
