//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package main

import (
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	go proxy("github.com:443")

	testcases := []struct {
		name  string
		title string
	}{
		{"Github", "Github"},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get("http://localhost:8443")
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("status code error | expected %d | got %d", http.StatusOK, resp.StatusCode)
			}
		})
	}

}
