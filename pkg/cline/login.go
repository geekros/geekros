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
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/geekros/geekros/pkg/utils"
	"github.com/gookit/color"
)

type (
	sendCodeResponseMsg struct {
		success bool
		err     string
	}
	verifyCodeResponseMsg struct {
		success bool
		err     string
	}
)

type LoginModel struct {
	state      string
	phoneInput textinput.Model
	codeInput  textinput.Model
	Loading    spinner.Model
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

	loading := spinner.New()
	loading.Spinner = spinner.Dot
	loading.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return LoginModel{
		state:      "phone",
		phoneInput: phone,
		codeInput:  code,
		Loading:    loading,
	}
}

func (m LoginModel) Init() tea.Cmd {

	return tea.Batch(
		textinput.Blink,
		m.Loading.Tick,
	)
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
				m.state = "sending"
				return m, tea.Batch(
					m.Loading.Tick,
					sendCodeRequest(m.phoneInput.Value()),
				)
			case "code":
				m.state = "verifying"
				return m, tea.Batch(
					m.Loading.Tick,
					verifyCodeRequest(m.phoneInput.Value(), m.codeInput.Value()),
				)
			case "failed":
				m.codeInput.SetValue("")
				m.state = "code"
			}
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}
	case spinner.TickMsg:
		if m.state == "sending" || m.state == "verifying" {
			m.Loading, cmd = m.Loading.Update(msg)
			return m, cmd
		}
	case sendCodeResponseMsg:
		if msg.success {
			m.state = "code"
			m.codeInput.Focus()
		} else {
			m.err = msg.err
			m.state = "phone"
		}
	case verifyCodeResponseMsg:
		if msg.success {
			m.state = "success"
			return m, tea.Quit
		} else {
			m.err = msg.err
			m.state = "failed"
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
	case "sending":
		return fmt.Sprintf("Enter phone number:\n\n%s %s\n\n"+color.Gray.Text("Press Esc to exit."), m.Loading.View(), color.Gray.Text("Sending verification code..."))
	case "verifying":
		return fmt.Sprintf("Enter code (sent to %s):\n\n%s %s\n\n"+color.Gray.Text("Press Esc to exit."), utils.PhoneToFormat(m.phoneInput.Value()), m.Loading.View(), color.Gray.Text("Verifying code..."))
	case "success":
		return fmt.Sprintf(color.Gray.Text("Logged in successfully.") + "\n")
	case "failed":
		return fmt.Sprintf(color.Yellow.Text("Incorrect code. Press Enter to retry.") + "\n\n" + color.Gray.Text("Press Esc to exit."))
	}

	return ""
}

func sendCodeRequest(phone string) tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return sendCodeResponseMsg{true, ""}
	})
}

func verifyCodeRequest(phone, code string) tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return verifyCodeResponseMsg{true, ""}
	})
}
