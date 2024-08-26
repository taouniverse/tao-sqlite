// Copyright 2024 huija
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlite

import (
	"context"
	"github.com/taouniverse/tao"
)

// ConfigKey for this repo
const ConfigKey = "sqlite"

// Config implements tao.Config
// declare the configuration you want & define some default values
type Config struct {
	DB        string   `json:"db"`
	RunAfters []string `json:"run_after,omitempty"`
}

var defaultSqlite = &Config{
	DB:        "sqlite.db",
	RunAfters: []string{},
}

// Name of Config
func (s *Config) Name() string {
	return ConfigKey
}

// ValidSelf with some default values
func (s *Config) ValidSelf() {
	if s.DB == "" {
		s.DB = defaultSqlite.DB
	}

	if s.RunAfters == nil {
		s.RunAfters = defaultSqlite.RunAfters
	}
}

// ToTask transform itself to Task
func (s *Config) ToTask() tao.Task {
	return tao.NewTask(
		ConfigKey,
		func(ctx context.Context, param tao.Parameter) (tao.Parameter, error) {
			// non-block check
			select {
			case <-ctx.Done():
				return param, tao.NewError(tao.ContextCanceled, "%s: context has been canceled", ConfigKey)
			default:
			}
			// JOB code run after RunAfters, you can just do nothing here
			db, err := DB.DB()
			if err != nil {
				return param, err
			}

			err = db.Ping()
			return param, err
		})
}

// RunAfter defines pre task names
func (s *Config) RunAfter() []string {
	return s.RunAfters
}
