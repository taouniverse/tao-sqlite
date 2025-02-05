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
	"github.com/stretchr/testify/assert"
	"github.com/taouniverse/tao"
	"testing"
)

func TestTao(t *testing.T) {
	err := tao.DevelopMode()
	assert.Nil(t, err)

	db, err := DB.DB()
	assert.Nil(t, err)

	err = db.Ping()
	assert.Nil(t, err)

	err = tao.Run(nil, nil)
	assert.Nil(t, err)
}
