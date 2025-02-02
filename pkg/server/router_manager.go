/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"github.com/apache/dubbo-go-pixiu/pkg/logger"
	"github.com/apache/dubbo-go-pixiu/pkg/model"
)

type (
	RouterListener interface {
		OnAddRouter(r *model.Router)
		OnDeleteRouter(r *model.Router)
	}

	RouterManager struct {
		rls []RouterListener
	}
)

func CreateDefaultRouterManager(server *Server, bs *model.Bootstrap) *RouterManager {
	rm := &RouterManager{}
	return rm
}

func (rm *RouterManager) AddRouterListener(l RouterListener) {
	rm.rls = append(rm.rls, l)
}

func (rm *RouterManager) AddRouter(r *model.Router) {
	logger.Infof("add router: %v", r)
	for _, l := range rm.rls {
		l.OnAddRouter(r)
	}
}

func (rm *RouterManager) DeleteRouter(r *model.Router) {
	logger.Infof("del router: %v", r)
	for _, l := range rm.rls {
		l.OnDeleteRouter(r)
	}
}
