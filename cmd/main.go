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

package main

import (
	"fmt"
	"os"

	"github.com/geekros/geekros/cmd/command"
	"github.com/geekros/geekros/pkg/version"
	"github.com/spf13/cobra"
)

func main() {

	cmd := &cobra.Command{
		Use:   version.Name,
		Short: version.Describe,
		Long:  fmt.Sprintf("%s - %s %s (%s)", version.Name, version.Describe, version.Number, version.Site),
	}

	cmd.CompletionOptions.HiddenDefaultCmd = true

	cmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})

	cmd.AddCommand(command.Version())

	cmd.AddCommand(command.Server())

	cmd.AddCommand(command.Login())

	cmd.AddCommand(command.Mcu())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
