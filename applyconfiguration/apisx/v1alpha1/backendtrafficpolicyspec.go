/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "sigs.k8s.io/gateway-api/applyconfiguration/apis/v1"
	v1alpha2 "sigs.k8s.io/gateway-api/applyconfiguration/apis/v1alpha2"
)

// BackendTrafficPolicySpecApplyConfiguration represents a declarative configuration of the BackendTrafficPolicySpec type for use
// with apply.
type BackendTrafficPolicySpecApplyConfiguration struct {
	TargetRefs         []v1alpha2.LocalPolicyTargetReferenceApplyConfiguration `json:"targetRefs,omitempty"`
	RetryConstraint    *RetryConstraintApplyConfiguration                      `json:"retryConstraint,omitempty"`
	SessionPersistence *v1.SessionPersistenceApplyConfiguration                `json:"sessionPersistence,omitempty"`
}

// BackendTrafficPolicySpecApplyConfiguration constructs a declarative configuration of the BackendTrafficPolicySpec type for use with
// apply.
func BackendTrafficPolicySpec() *BackendTrafficPolicySpecApplyConfiguration {
	return &BackendTrafficPolicySpecApplyConfiguration{}
}

// WithTargetRefs adds the given value to the TargetRefs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TargetRefs field.
func (b *BackendTrafficPolicySpecApplyConfiguration) WithTargetRefs(values ...*v1alpha2.LocalPolicyTargetReferenceApplyConfiguration) *BackendTrafficPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTargetRefs")
		}
		b.TargetRefs = append(b.TargetRefs, *values[i])
	}
	return b
}

// WithRetryConstraint sets the RetryConstraint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RetryConstraint field is set to the value of the last call.
func (b *BackendTrafficPolicySpecApplyConfiguration) WithRetryConstraint(value *RetryConstraintApplyConfiguration) *BackendTrafficPolicySpecApplyConfiguration {
	b.RetryConstraint = value
	return b
}

// WithSessionPersistence sets the SessionPersistence field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SessionPersistence field is set to the value of the last call.
func (b *BackendTrafficPolicySpecApplyConfiguration) WithSessionPersistence(value *v1.SessionPersistenceApplyConfiguration) *BackendTrafficPolicySpecApplyConfiguration {
	b.SessionPersistence = value
	return b
}
