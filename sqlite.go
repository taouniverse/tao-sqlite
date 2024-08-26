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
	"github.com/glebarez/sqlite"
	"github.com/taouniverse/tao"
	"gorm.io/gorm"

	// Load the required dependencies.
	// An error occurs when there was no package in the root directory.
	_ "github.com/glebarez/sqlite"
	_ "gorm.io/gorm"
)

/**
import _ "github.com/taouniverse/tao-sqlite"
*/

// S config of sqlite
var S = new(Config)

func init() {
	err := tao.Register(ConfigKey, S, setup)
	if err != nil {
		panic(err.Error())
	}
}

// DB orm client of sqlite
var DB *gorm.DB

// setup unit with the global config 'S'
// execute when init tao universe
func setup() (err error) {
	DB, err = gorm.Open(sqlite.Open(S.DB), &gorm.Config{})
	if err != nil {
		return tao.NewErrorWrapped("sqlite: fail to create gorm client", err)
	}
	return nil
}
