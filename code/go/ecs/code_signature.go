// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

// These fields contain information about binary code signatures.
type CodeSignature struct {
	// Subject name of the code signer
	SubjectName string `ecs:"subject_name"`

	// Boolean to capture if the digital signature is verified against the
	// binary content.
	Valid bool `ecs:"valid"`

	// Boolean to capture if a signature is present.
	// This should only populated if the signature was checked.
	Exists bool `ecs:"exists"`

	// Stores the trust status of the certificate chain.
	Trusted bool `ecs:"trusted"`

	// Additional information about the certificate status.
	// This is useful for logging cryptographic errors with the certificate
	// validity or trust status.
	Status string `ecs:"status"`
}
