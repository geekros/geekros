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

package command

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/geekros/geekros/pkg/cline"
	"github.com/spf13/cobra"
)

func Mcu() *cobra.Command {

	command := &cobra.Command{
		Use:     "mcu",
		Short:   "Microcontroller unit management module",
		Long:    "Microcontroller unit management module",
		Example: "geekros mcu [create|install|uninstall|publish]",
		Args:    cobra.MaximumNArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"create", "install", "uninstall", "publish"}, cobra.ShellCompDirectiveNoFileComp
		},
		Run: McuRun,
	}
	return command
}

func McuRun(cmd *cobra.Command, args []string) {

	if len(args) > 0 {
		if args[0] == "create" {
			if _, err := tea.NewProgram(cline.InitMcuCreateModel()).Run(); err != nil {
				os.Exit(1)
			}
		}
		if args[0] == "install" {

		}
		if args[0] == "uninstall" {

		}
		if args[0] == "publish" {

		}
	}
}
