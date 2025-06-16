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

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/geekros/geekros/pkg/utils"
	"github.com/gookit/color"
)

type state int

const (
	inputPhone state = iota
	inputCode
	success
	failed
)

const mockCode = "123456"

type model struct {
	state      state
	phoneInput textinput.Model
	codeInput  textinput.Model
	err        string
	msg        string
}

func InitModel() model {
	phone := textinput.New()
	phone.Placeholder = "Enter your phone number"
	phone.CharLimit = 11
	phone.Width = 100
	phone.Focus()

	code := textinput.New()
	code.Placeholder = "Enter verification code"
	code.CharLimit = 6
	code.Width = 50

	return model{
		state:      inputPhone,
		phoneInput: phone,
		codeInput:  code,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			switch m.state {
			case inputPhone:
				if len(m.phoneInput.Value()) != 11 {
					m.err = color.Yellow.Text("Invalid phone number.")
					return m, nil
				}
				m.state = inputCode
				m.codeInput.Focus()
			case inputCode:
				if m.codeInput.Value() == mockCode {
					m.state = success
					return m, tea.Quit
				} else {
					m.state = failed
				}
			case failed:
				m.codeInput.SetValue("")
				m.state = inputCode
			}
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	switch m.state {
	case inputPhone:
		m.phoneInput, cmd = m.phoneInput.Update(msg)
	case inputCode:
		m.codeInput, cmd = m.codeInput.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	switch m.state {
	case inputPhone:
		return fmt.Sprintf("Enter phone number:\n\n%s"+color.Gray.Text("Press Esc to exit."), m.phoneInput.View())
	case inputCode:
		return fmt.Sprintf("Enter code (sent to %s):\n\n%s\n\n"+color.Gray.Text("Press Esc to exit."), utils.PhoneToFormat(m.phoneInput.Value()), m.codeInput.View())
	case success:
		return color.Green.Text("Logged in successfully.")
	case failed:
		return color.Yellow.Text("Incorrect code. Press Enter to retry.") + "\n\n" + color.Gray.Text("Press Esc to exit.")
	}
	return ""
}
