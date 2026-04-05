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
	"github.com/glebarez/sqlite"
	"github.com/taouniverse/tao"
	"gorm.io/gorm"

	// sqlite driver for gorm
	_ "github.com/glebarez/sqlite"
	// gorm package
	_ "gorm.io/gorm"
)

/**
import _ "github.com/taouniverse/tao-sqlite"
*/

// S is the global config instance for tao-sqlite
var S = &Config{}

// Factory is the global factory instance for managing gorm.DB
var Factory *tao.BaseFactory[*gorm.DB]

func init() {
	var err error
	Factory, err = tao.Register(ConfigKey, S, NewSQLite)
	if err != nil {
		panic(err.Error())
	}
}

// NewSQLite creates a new SQLite client for factory pattern
func NewSQLite(name string, config InstanceConfig) (*gorm.DB, func() error, error) {
	db, err := gorm.Open(sqlite.Open(config.DB), &gorm.Config{})
	if err != nil {
		return nil, nil, tao.NewErrorWrapped("sqlite: fail to create gorm client", err)
	}

	closer := func() error {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return db, closer, nil
}

// DB returns the default gorm DB instance
func DB() (*gorm.DB, error) {
	return Factory.Get(S.GetDefaultInstanceName())
}

// GetDB returns the gorm DB instance by name
func GetDB(name string) (*gorm.DB, error) {
	return Factory.Get(name)
}
