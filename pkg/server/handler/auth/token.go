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

package auth

import (
	"github.com/geekros/geekros/pkg/config"
	"github.com/geekros/geekros/pkg/structs"
	"github.com/geekros/geekros/pkg/utils"
	"github.com/gin-gonic/gin"
)

type responseAuthToken struct {
	Token      string `json:"token"`
	Expiration int    `json:"expiration"`
}

func AuthToken(c *gin.Context) {

	responseData := responseAuthToken{}

	role := c.DefaultQuery("role", "")
	if role == "" {
		utils.Error(c, structs.EmptyData{})
	}

	data := map[string]interface{}{
		"role": role,
	}

	token, err := utils.GenerateToken(config.Get.Auth.Secret, data, config.Get.Auth.Expiration)
	if err != nil {
		utils.Error(c, structs.EmptyData{})
	}

	responseData.Token = token
	responseData.Expiration = config.Get.Auth.Expiration

	utils.Success(c, responseData)
}
