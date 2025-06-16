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

	"github.com/spf13/cobra"
)

func Mcu() *cobra.Command {

	command := &cobra.Command{
		Use:     "mcu",
		Short:   "Microcontroller unit management module",
		Long:    "Microcontroller unit management module",
		Example: "geekros mcu [init|install|remove|publish]",
		Run:     McuRun,
	}
	return command
}

func McuRun(cmd *cobra.Command, args []string) {
	log.Println(args)
}
