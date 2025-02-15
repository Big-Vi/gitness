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

package job

import (
	"github.com/harness/gitness/app/store"
	"github.com/harness/gitness/lock"
	"github.com/harness/gitness/pubsub"
	"github.com/harness/gitness/types"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	ProvideExecutor,
	ProvideScheduler,
)

func ProvideExecutor(
	jobStore store.JobStore,
	pubsubService pubsub.PubSub,
) *Executor {
	return NewExecutor(
		jobStore,
		pubsubService,
	)
}

func ProvideScheduler(
	jobStore store.JobStore,
	executor *Executor,
	mutexManager lock.MutexManager,
	pubsubService pubsub.PubSub,
	config *types.Config,
) (*Scheduler, error) {
	return NewScheduler(
		jobStore,
		executor,
		mutexManager,
		pubsubService,
		config.InstanceID,
		config.BackgroundJobs.MaxRunning,
		config.BackgroundJobs.RetentionTime,
	)
}
