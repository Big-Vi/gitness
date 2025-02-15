// Copyright 2023 Harness, Inc.
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

package connector

import (
	"github.com/harness/gitness/app/auth/authz"
	"github.com/harness/gitness/app/store"
	"github.com/harness/gitness/types/check"
)

type Controller struct {
	uidCheck       check.PathUID
	connectorStore store.ConnectorStore
	authorizer     authz.Authorizer
	spaceStore     store.SpaceStore
}

func NewController(
	uidCheck check.PathUID,
	authorizer authz.Authorizer,
	connectorStore store.ConnectorStore,
	spaceStore store.SpaceStore,
) *Controller {
	return &Controller{
		uidCheck:       uidCheck,
		connectorStore: connectorStore,
		authorizer:     authorizer,
		spaceStore:     spaceStore,
	}
}
