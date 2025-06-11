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

package router

import (
	"github.com/geekros/geekros/pkg/server/handler"
	"github.com/geekros/geekros/pkg/server/handler/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func New() *Router {

	return &Router{}
}

func (r *Router) Init(mode string) *Router {

	gin.SetMode(mode)

	r.engine = gin.New()
	r.engine.Use(gin.Recovery())
	r.engine.Use(cors.Default())
	r.engine.Use(r.Authentication())

	return r
}

func (r *Router) Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()
	}
}

func (r *Router) InitFrontendHandler() *gin.Engine {

	return r.engine
}

func (r *Router) InitHandler() *gin.Engine {

	r.InitFrontendHandler()

	handlerGroup := r.engine.Group("handler")
	{
		handlerGroup.GET("/health", handler.Health)

		handlerGroup.GET("/auth/token", auth.AuthToken)
	}

	return r.engine
}
