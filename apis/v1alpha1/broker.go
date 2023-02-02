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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BrokerSpec defines the desired state of Broker.
type BrokerSpec struct {
	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty"`
	// +kubebuilder:validation:Required
	AutoMinorVersionUpgrade *bool            `json:"autoMinorVersionUpgrade"`
	Configuration           *ConfigurationID `json:"configuration,omitempty"`
	CreatorRequestID        *string          `json:"creatorRequestID,omitempty"`
	// +kubebuilder:validation:Required
	DeploymentMode    *string            `json:"deploymentMode"`
	EncryptionOptions *EncryptionOptions `json:"encryptionOptions,omitempty"`
	// +kubebuilder:validation:Required
	EngineType *string `json:"engineType"`
	// +kubebuilder:validation:Required
	EngineVersion *string `json:"engineVersion"`
	// +kubebuilder:validation:Required
	HostInstanceType           *string                  `json:"hostInstanceType"`
	LDAPServerMetadata         *LDAPServerMetadataInput `json:"ldapServerMetadata,omitempty"`
	Logs                       *Logs                    `json:"logs,omitempty"`
	MaintenanceWindowStartTime *WeeklyStartTime         `json:"maintenanceWindowStartTime,omitempty"`
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// +kubebuilder:validation:Required
	PubliclyAccessible *bool                                      `json:"publiclyAccessible"`
	SecurityGroupRefs  []*ackv1alpha1.AWSResourceReferenceWrapper `json:"securityGroupRefs,omitempty"`
	SecurityGroups     []*string                                  `json:"securityGroups,omitempty"`
	StorageType        *string                                    `json:"storageType,omitempty"`
	SubnetIDs          []*string                                  `json:"subnetIDs,omitempty"`
	SubnetRefs         []*ackv1alpha1.AWSResourceReferenceWrapper `json:"subnetRefs,omitempty"`
	Tags               map[string]*string                         `json:"tags,omitempty"`
	// +kubebuilder:validation:Required
	Users []*User `json:"users"`
}

// BrokerStatus defines the observed state of Broker
type BrokerStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// +kubebuilder:validation:Optional
	BrokerID *string `json:"brokerID,omitempty"`
	// +kubebuilder:validation:Optional
	BrokerInstances []*BrokerInstance `json:"brokerInstances,omitempty"`
	// +kubebuilder:validation:Optional
	BrokerState *string `json:"brokerState,omitempty"`
}

// Broker is the Schema for the Brokers API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Broker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BrokerSpec   `json:"spec,omitempty"`
	Status            BrokerStatus `json:"status,omitempty"`
}

// BrokerList contains a list of Broker
// +kubebuilder:object:root=true
type BrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Broker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Broker{}, &BrokerList{})
}
