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

type LoginModel struct {
	state      string
	phoneInput textinput.Model
	codeInput  textinput.Model
	err        string
}

func InitModel() LoginModel {

	phone := textinput.New()
	phone.Placeholder = "Enter your phone number"
	phone.CharLimit = 11
	phone.Width = 100
	phone.Focus()

	code := textinput.New()
	code.Placeholder = "Enter verification code"
	code.CharLimit = 6
	code.Width = 50

	return LoginModel{
		state:      "phone",
		phoneInput: phone,
		codeInput:  code,
	}
}

func (m LoginModel) Init() tea.Cmd {

	return textinput.Blink
}

func (m LoginModel) Update(msg tea.Msg) (LoginModel, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			switch m.state {
			case "phone":
				if len(m.phoneInput.Value()) != 11 {
					m.err = color.Yellow.Text("Invalid phone number.")
					return m, nil
				}
				m.state = "code"
				m.codeInput.Focus()
			case "code":
				if m.codeInput.Value() == "123456" {
					m.state = "success"
					return m, tea.Quit
				} else {
					m.state = "failed"
				}
			case "failed":
				m.codeInput.SetValue("")
				m.state = "code"
			}
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	switch m.state {
	case "phone":
		m.phoneInput, cmd = m.phoneInput.Update(msg)
	case "code":
		m.codeInput, cmd = m.codeInput.Update(msg)
	}

	return m, cmd
}

func (m LoginModel) View() string {

	switch m.state {
	case "phone":
		return fmt.Sprintf("Enter phone number:\n\n%s\n\n"+color.Gray.Text("Press Esc to exit."), m.phoneInput.View())
	case "code":
		return fmt.Sprintf("Enter code (sent to %s):\n\n%s\n\n"+color.Gray.Text("Press Esc to exit."), utils.PhoneToFormat(m.phoneInput.Value()), m.codeInput.View())
	case "success":
		return fmt.Sprintf(color.Gray.Text("Logged in successfully.") + "\n")
	case "failed":
		return fmt.Sprintf(color.Yellow.Text("Incorrect code. Press Enter to retry.") + "\n\n" + color.Gray.Text("Press Esc to exit."))
	}

	return ""
}
