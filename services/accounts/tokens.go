// Copyright 2022 Evan Hazlett
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
package accounts

import (
	"context"

	api "github.com/ehazlett/flow/api/services/accounts/v1"
)

func (s *service) GenerateServiceToken(ctx context.Context, req *api.GenerateServiceTokenRequest) (*api.GenerateServiceTokenResponse, error) {
	st, err := s.authenticator.GenerateServiceToken(ctx, req.Description, req.TTL)
	if err != nil {
		return nil, err
	}
	return &api.GenerateServiceTokenResponse{
		ServiceToken: st,
	}, nil
}
