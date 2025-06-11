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

	"github.com/geekros/geekros/pkg/config"
	"github.com/geekros/geekros/pkg/server"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func Server() *cobra.Command {

	command := &cobra.Command{
		Use:     "server",
		Short:   "Start server module",
		Long:    "Start server module",
		Example: "geekros server",
		Run:     serverRun,
	}
	return command
}

func serverRun(cmd *cobra.Command, args []string) {

	log.Println(color.Gray.Text("start server..."))

	config.Get = config.New().LoadConfig()

	server.Get = server.New()
	server.Get.Start(func() {
		log.Println(color.Gray.Text("started server"))
	}, func() {
		log.Println(color.Gray.Text("exited server"))
	})
}
