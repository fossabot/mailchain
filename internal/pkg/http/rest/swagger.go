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

// Package rest Mailchain API
//
// All the information needed to talk to the API.
//
// To raise see anything wrong? Raise an [issue](https://github.com/mailchain/mailchain/issues)
//
//     Schemes: https
//     BasePath: /api
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package rest

import (
	// blank imports to support doc creations
	_ "github.com/mailchain/mailchain/internal/pkg/http/rest/addresses"
	_ "github.com/mailchain/mailchain/internal/pkg/http/rest/ethereum/address/messages"
	_ "github.com/mailchain/mailchain/internal/pkg/http/rest/ethereum/address/publickey"
	_ "github.com/mailchain/mailchain/internal/pkg/http/rest/ethereum/messages/send"
	_ "github.com/mailchain/mailchain/internal/pkg/http/rest/messages/read"
	_ "github.com/mailchain/mailchain/internal/pkg/http/rest/spec"
)
