// Copyright 2020 Google LLC
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

// Package authentication provides functionality for authentication.
package authentication

import (
	"encoding/json"
	"net/http"
)

// Middleware works as a wrapper around http.HandlerFunc.
// It verifies that the caller provided a valid api secret when requesting a authenticated endpoint.
// This method is expecting the requests to contain a header with the format:
//		"Authorization: value".
func Middleware(next http.HandlerFunc, apiKey string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			if authorizationHeader == apiKey {
				next(w, req)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode("Invalid Authorization header")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("An Authorization header is required")
		}
	})
}
