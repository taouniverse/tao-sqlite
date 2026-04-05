// Copyright 2021-2026 huija
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

// InstanceConfig 单实例配置
type InstanceConfig struct {
	DB string `json:"db" yaml:"db"`
}

// Config 总配置，实现 tao.MultiConfig 接口
type Config struct {
	tao.BaseMultiConfig[InstanceConfig]
	RunAfters []string `json:"run_after,omitempty" yaml:"run_after,omitempty"`
}

var defaultInstance = &InstanceConfig{
	DB: "sqlite.db",
}

// Name of Config
func (s *Config) Name() string {
	return ConfigKey
}

// ValidSelf with some default values
func (s *Config) ValidSelf() {
	for name, instance := range s.Instances {
		if instance.DB == "" {
			instance.DB = defaultInstance.DB
		}
		s.Instances[name] = instance
	}
	if s.RunAfters == nil {
		s.RunAfters = []string{}
	}
}

// ToTask transform itself to Task
func (s *Config) ToTask() tao.Task {
	return tao.NewTask(
		ConfigKey,
		func(ctx context.Context, param tao.Parameter) (tao.Parameter, error) {
			select {
			case <-ctx.Done():
				return param, tao.NewError(tao.ContextCanceled, "%s: context has been canceled", ConfigKey)
			default:
			}
			for name := range s.Instances {
				db, err := Factory.Get(name)
				if err != nil {
					return param, err
				}
				sqlDB, err := db.DB()
				if err != nil {
					return param, err
				}
				if err := sqlDB.Ping(); err != nil {
					return param, err
				}
			}
			return param, nil
		})
}

// RunAfter defines pre task names
func (s *Config) RunAfter() []string {
	return s.RunAfters
}
