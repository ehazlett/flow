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
package flow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateToken returns a sha256 of the specified data
func GenerateToken(v ...string) string {
	data := fmt.Sprintf("%s", time.Now())
	for _, x := range v {
		data += x
	}
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}