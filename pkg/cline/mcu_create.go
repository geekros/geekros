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

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type McuCreateModel struct {
	state   string
	keyword textinput.Model
	loading spinner.Model
	err     string
}

func InitMcuCreateModel() McuCreateModel {

	keyword := textinput.New()
	keyword.Placeholder = "Enter search keywords"
	keyword.CharLimit = 32
	keyword.Width = 100
	keyword.Focus()

	loading := spinner.New()
	loading.Spinner = spinner.Dot
	loading.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return McuCreateModel{
		state:   "home",
		keyword: keyword,
		loading: loading,
	}
}

func (m McuCreateModel) Init() tea.Cmd {

	return tea.Batch(
		textinput.Blink,
		m.loading.Tick,
	)
}

func (m McuCreateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	switch m.state {
	case "home":
		m.keyword, cmd = m.keyword.Update(msg)
	}

	return m, cmd
}

func (m McuCreateModel) View() string {

	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Margin(1, 0)

	switch m.state {
	case "home":
		text := fmt.Sprintf("Please select a basic microcontroller\n\n")
		text += fmt.Sprintf("Microcontroller model:\n\n%s\n"+helpStyle.Render("Press Esc to exit."), m.keyword.View())
		return text
	}

	return ""
}
