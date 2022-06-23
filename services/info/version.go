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
package info

import (
	"context"

	api "github.com/ehazlett/flow/api/services/info/v1"
	"github.com/ehazlett/flow/version"
)

func (s *service) Version(ctx context.Context, r *api.VersionRequest) (*api.VersionResponse, error) {
	return &api.VersionResponse{
		Name:          version.Name,
		Version:       version.Version,
		Build:         version.Build,
		Commit:        version.GitCommit,
		Authenticator: s.authenticator.Name(),
	}, nil
}