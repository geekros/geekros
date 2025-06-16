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
	"fmt"
	"log"

	"github.com/geekros/geekros/pkg/version"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func Version() *cobra.Command {

	command := &cobra.Command{
		Use:     "version",
		Short:   "Get version number",
		Long:    "Get version number",
		Example: "geekros version",
		Run:     versionRun,
	}

	return command
}

func versionRun(cmd *cobra.Command, args []string) {
	log.Println(color.Gray.Text(fmt.Sprintf("%s %s (%s)", version.Name, version.Number, version.Site)))
}
