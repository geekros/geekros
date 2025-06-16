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
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/geekros/geekros/pkg/cline"
	"github.com/spf13/cobra"
)

func Login() *cobra.Command {

	command := &cobra.Command{
		Use:     "login",
		Short:   "Authorize your account",
		Long:    "Authorize your account",
		Example: "geekros login [phone]",
		Run:     LoginRun,
	}
	return command
}

func LoginRun(cmd *cobra.Command, args []string) {

	if _, err := tea.NewProgram(cline.InitModel()).Run(); err != nil {
		log.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
