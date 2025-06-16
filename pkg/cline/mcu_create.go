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

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gookit/color"
)

type (
	onItemsRequestMsg struct {
		success bool
		err     string
		items   []list.Item
	}
)

type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type McuCreateModel struct {
	state   string
	keyword textinput.Model
	items   list.Model
	loading spinner.Model
	width   int
	height  int
	err     string
}

func InitMcuCreateModel() McuCreateModel {

	keyword := textinput.New()
	keyword.Placeholder = "Enter search keywords"
	keyword.CharLimit = 32
	keyword.Width = 100
	keyword.Focus()

	items := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	items.Title = "Search Results"
	items.SetShowHelp(false)

	loading := spinner.New()
	loading.Spinner = spinner.Dot
	loading.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return McuCreateModel{
		state:   "home",
		keyword: keyword,
		items:   items,
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
		case tea.KeyEnter:
			switch m.state {
			case "home":
				if len(m.keyword.Value()) > 0 {
					m.state = "loading"
					return m, tea.Batch(
						m.loading.Tick,
						m.onItemsRequest(m.keyword.Value()),
					)
				}
			case "items":
				i, ok := m.items.SelectedItem().(item)
				if ok {

				}
			}
		case tea.KeyEsc, tea.KeyCtrlC:
			if m.state == "items" {
				m.state = "home"
				m.keyword.Focus()
				return m, cmd
			}
			return m, tea.Quit
		}
	case spinner.TickMsg:
		if m.state == "loading" {
			m.loading, cmd = m.loading.Update(msg)
			return m, cmd
		}
	case onItemsRequestMsg:
		if msg.success {
			if len(msg.items) == 0 {
				m.state = "home"
				m.keyword.Focus()
				return m, cmd
			}
			m.items = list.New(msg.items, list.NewDefaultDelegate(), 0, 0)
			m.items.Title = "Search Results"
			m.items.SetShowHelp(false)
			m.items.SetSize(m.width, m.height-(m.height/2))
			m.state = "items"
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.items.SetSize(msg.Width, m.height-(m.height/2))
	}

	switch m.state {
	case "home":
		m.keyword, cmd = m.keyword.Update(msg)
	case "items":
		m.items, cmd = m.items.Update(msg)
	}

	return m, cmd
}

func (m McuCreateModel) View() string {

	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Margin(1, 0)

	switch m.state {
	case "home":
		return fmt.Sprintf("Please select a basic microcontroller model:\n\n%s\n"+helpStyle.Render("Press Esc to exit."), m.keyword.View())
	case "loading":
		return fmt.Sprintf("Please select a basic microcontroller model:\n\n%s\n\n%s%s\n"+helpStyle.Render("Press Esc to exit."), m.keyword.View(), m.loading.View(), color.Gray.Text("Searching..."))
	case "items":
		return fmt.Sprintf("Please select a basic microcontroller model:\n\n%s\n\n%s\n"+helpStyle.Render("Up/Down to navigate, Enter to select, Esc to quit."), m.keyword.View(), m.items.View())
	}

	return ""
}

func (m McuCreateModel) onItemsRequest(keyword string) tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		items := []list.Item{
			item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
			item{title: "Nutella", desc: "It's good on toast"},
			item{title: "Bitter melon", desc: "It cools you down"},
			item{title: "Nice socks", desc: "And by that I mean socks without holes"},
			item{title: "Eight hours of sleep", desc: "I had this once"},
			item{title: "Cats", desc: "Usually"},
			item{title: "Plantasia, the album", desc: "My plants love it too"},
			item{title: "Pour over coffee", desc: "It takes forever to make though"},
			item{title: "VR", desc: "Virtual reality...what is there to say?"},
			item{title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
			item{title: "Linux", desc: "Pretty much the best OS"},
			item{title: "Business school", desc: "Just kidding"},
			item{title: "Pottery", desc: "Wet clay is a great feeling"},
			item{title: "Shampoo", desc: "Nothing like clean hair"},
			item{title: "Table tennis", desc: "It’s surprisingly exhausting"},
			item{title: "Milk crates", desc: "Great for packing in your extra stuff"},
			item{title: "Afternoon tea", desc: "Especially the tea sandwich part"},
			item{title: "Stickers", desc: "The thicker the vinyl the better"},
			item{title: "20° Weather", desc: "Celsius, not Fahrenheit"},
			item{title: "Warm light", desc: "Like around 2700 Kelvin"},
			item{title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
			item{title: "Gaffer’s tape", desc: "Basically sticky fabric"},
			item{title: "Terrycloth", desc: "In other words, towel fabric"},
		}
		return onItemsRequestMsg{true, "", items}
	})
}
