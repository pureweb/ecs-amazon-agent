// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
package mux

const (
	// AnythingRegEx is a regex pattern that matches anything.
	AnythingRegEx = ".*"

	// AnythingButSlashRegEx is a regex pattern that matches any string without slash.
	AnythingButSlashRegEx = "[^/]*"

	// AnythingButEmptyRegEx is a regex pattern that matches anything but an empty string.
	AnythingButEmptyRegEx = ".+"
)

// ConstructMuxVar constructs the mux var that is used in the gorilla/mux styled
// path, example: {id}, {id:[0-9]+}.
func ConstructMuxVar(name string, pattern string) string {
	if pattern == "" {
		return "{" + name + "}"
	}

	return "{" + name + ":" + pattern + "}"
}
