// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mail

import (
	"testing"

	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	assert := assert.New(t)
	id, err := NewID()
	assert.NoError(err)
	assert.Len(id, 46)
	assert.Equal(byte(multihash.ID), id[0])
}
