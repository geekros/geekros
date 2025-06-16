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
	state   string
	phone   textinput.Model
	code    textinput.Model
	loading spinner.Model
	err     string
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
		state:   "phone",
		phone:   phone,
		code:    code,
		loading: loading,
	}
}

func (m LoginModel) Init() tea.Cmd {

	return tea.Batch(
		textinput.Blink,
		m.loading.Tick,
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
				if len(m.phone.Value()) != 11 {
					m.err = color.Yellow.Text("Invalid phone number.")
					return m, nil
				}
				m.state = "sending"
				return m, tea.Batch(
					m.loading.Tick,
					sendCodeRequest(m.phone.Value()),
				)
			case "code":
				if m.code.Value() != "123456" {
					m.state = "failed"
					return m, nil
				}
				m.state = "verifying"
				return m, tea.Batch(
					m.loading.Tick,
					verifyCodeRequest(m.phone.Value(), m.code.Value()),
				)
			case "failed":
				m.code.SetValue("")
				m.state = "code"
				m.code.Focus()
			}
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}
	case spinner.TickMsg:
		if m.state == "sending" || m.state == "verifying" {
			m.loading, cmd = m.loading.Update(msg)
			return m, cmd
		}
	case sendCodeResponseMsg:
		if msg.success {
			m.state = "code"
			m.code.Focus()
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
		m.phone, cmd = m.phone.Update(msg)
	case "code":
		m.code, cmd = m.code.Update(msg)
	}

	return m, cmd
}

func (m LoginModel) View() string {
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Margin(1, 0)
	switch m.state {
	case "phone":
		return fmt.Sprintf("Enter phone number:\n\n%s\n"+helpStyle.Render("Press Esc to exit."), m.phone.View())
	case "code":
		return fmt.Sprintf("Enter code (sent to %s):\n\n%s\n"+helpStyle.Render("Press Esc to exit."), utils.PhoneToFormat(m.phone.Value()), m.code.View())
	case "sending":
		return fmt.Sprintf("Enter phone number:\n\n%s%s\n"+helpStyle.Render("Press Esc to exit."), m.loading.View(), color.Gray.Text("Sending verification code..."))
	case "verifying":
		return fmt.Sprintf("Enter code (sent to %s):\n\n%s%s\n"+helpStyle.Render("Press Esc to exit."), utils.PhoneToFormat(m.phone.Value()), m.loading.View(), color.Gray.Text("Verifying code..."))
	case "success":
		return fmt.Sprintf(color.Gray.Text("Logged in successfully.") + "\n")
	case "failed":
		return fmt.Sprintf(color.Yellow.Text("Incorrect code. Press Enter to retry.") + "\n" + helpStyle.Render("Press Esc to exit."))
	}

	return ""
}

func sendCodeRequest(phone string) tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return sendCodeResponseMsg{true, ""}
	})
}

func verifyCodeRequest(phone string, code string) tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return verifyCodeResponseMsg{true, ""}
	})
}
