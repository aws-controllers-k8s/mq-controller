// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package broker

import (
	"testing"

	svcapitypes "github.com/aws-controllers-k8s/mq-controller/apis/v1alpha1"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestNewResourceDelta_EngineVersionPrefix(t *testing.T) {
	tests := []struct {
		name           string
		aEngineVersion *string
		bEngineVersion *string
		expectDelta    bool
	}{
		{
			name:           "a is prefix of b - no delta expected",
			aEngineVersion: aws.String("3.13"),
			bEngineVersion: aws.String("3.13.1"),
			expectDelta:    false,
		},
		{
			name:           "a is prefix of b with patch version - no delta expected",
			aEngineVersion: aws.String("5.18"),
			bEngineVersion: aws.String("5.18.4"),
			expectDelta:    false,
		},
		{
			name:           "exact match - no delta expected",
			aEngineVersion: aws.String("3.13.1"),
			bEngineVersion: aws.String("3.13.1"),
			expectDelta:    false,
		},
		{
			name:           "different patch versions - no delta expected",
			aEngineVersion: aws.String("3.13.2"),
			bEngineVersion: aws.String("3.13.1"),
			expectDelta:    true,
		},
		{
			name:           "different minor versions - delta expected",
			aEngineVersion: aws.String("3.13"),
			bEngineVersion: aws.String("3.14.1"),
			expectDelta:    true,
		},
		{
			name:           "a is longer than b - delta expected",
			aEngineVersion: aws.String("3.13.1"),
			bEngineVersion: aws.String("3.13"),
			expectDelta:    true,
		},
		{
			name:           "nil versions - no delta expected",
			aEngineVersion: nil,
			bEngineVersion: nil,
			expectDelta:    false,
		},
		{
			name:           "a nil, b not nil - delta expected",
			aEngineVersion: nil,
			bEngineVersion: aws.String("3.13.1"),
			expectDelta:    true,
		},
		{
			name:           "a not nil, b nil - delta expected",
			aEngineVersion: aws.String("3.13"),
			bEngineVersion: nil,
			expectDelta:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &resource{
				ko: &svcapitypes.Broker{
					Spec: svcapitypes.BrokerSpec{
						EngineVersion: tt.aEngineVersion,
					},
				},
			}
			b := &resource{
				ko: &svcapitypes.Broker{
					Spec: svcapitypes.BrokerSpec{
						EngineVersion: tt.bEngineVersion,
					},
				},
			}

			delta := newResourceDelta(a, b)
			hasDelta := delta.DifferentAt("Spec.EngineVersion")

			if tt.expectDelta && !hasDelta {
				t.Errorf("Expected delta for EngineVersion but none found")
			}
			if !tt.expectDelta && hasDelta {
				t.Errorf("Expected no delta for EngineVersion but found one")
			}
		})
	}
}
