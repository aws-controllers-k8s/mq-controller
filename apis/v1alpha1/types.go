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
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// Action required for a broker.
type ActionRequired struct {
	ActionRequiredCode *string `json:"actionRequiredCode,omitempty"`
	ActionRequiredInfo *string `json:"actionRequiredInfo,omitempty"`
}

// Name of the availability zone.
type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

// Types of broker engines.
type BrokerEngineType struct {
	// The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.
	EngineType *string `json:"engineType,omitempty"`
}

// Returns information about all brokers.
type BrokerInstance struct {
	ConsoleURL *string   `json:"consoleURL,omitempty"`
	Endpoints  []*string `json:"endpoints,omitempty"`
	IPAddress  *string   `json:"ipAddress,omitempty"`
}

// Option for host instance type.
type BrokerInstanceOption struct {
	// The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.
	EngineType       *string `json:"engineType,omitempty"`
	HostInstanceType *string `json:"hostInstanceType,omitempty"`
	// The broker's storage type.
	//
	// EFS is not supported for RabbitMQ engine type.
	StorageType             *string   `json:"storageType,omitempty"`
	SupportedEngineVersions []*string `json:"supportedEngineVersions,omitempty"`
}

// Returns information about all brokers.
type BrokerSummary struct {
	BrokerARN  *string `json:"brokerARN,omitempty"`
	BrokerID   *string `json:"brokerID,omitempty"`
	BrokerName *string `json:"brokerName,omitempty"`
	// The broker's status.
	BrokerState *string      `json:"brokerState,omitempty"`
	Created     *metav1.Time `json:"created,omitempty"`
	// The broker's deployment mode.
	DeploymentMode *string `json:"deploymentMode,omitempty"`
	// The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.
	EngineType       *string `json:"engineType,omitempty"`
	HostInstanceType *string `json:"hostInstanceType,omitempty"`
}

// Returns information about all configurations.
type Configuration struct {
	ARN *string `json:"arn,omitempty"`
	// Optional. The authentication strategy used to secure the broker. The default
	// is SIMPLE.
	AuthenticationStrategy *string      `json:"authenticationStrategy,omitempty"`
	Created                *metav1.Time `json:"created,omitempty"`
	Description            *string      `json:"description,omitempty"`
	// The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.
	EngineType    *string            `json:"engineType,omitempty"`
	EngineVersion *string            `json:"engineVersion,omitempty"`
	ID            *string            `json:"id,omitempty"`
	Name          *string            `json:"name,omitempty"`
	Tags          map[string]*string `json:"tags,omitempty"`
}

// A list of information about the configuration.
type ConfigurationID struct {
	ID       *string `json:"id,omitempty"`
	Revision *int64  `json:"revision,omitempty"`
}

// Returns information about the specified configuration revision.
type ConfigurationRevision struct {
	Created     *metav1.Time `json:"created,omitempty"`
	Description *string      `json:"description,omitempty"`
	Revision    *int64       `json:"revision,omitempty"`
}

// Broker configuration information
type Configurations struct {
	// A list of information about the configuration.
	Current *ConfigurationID   `json:"current,omitempty"`
	History []*ConfigurationID `json:"history,omitempty"`
	// A list of information about the configuration.
	Pending *ConfigurationID `json:"pending,omitempty"`
}

// Specifies a broker in a data replication pair.
type DataReplicationCounterpart struct {
	BrokerID *string `json:"brokerID,omitempty"`
	Region   *string `json:"region,omitempty"`
}

// The replication details of the data replication-enabled broker. Only returned
// if dataReplicationMode or pendingDataReplicationMode is set to CRDR.
type DataReplicationMetadataOutput struct {
	// Specifies a broker in a data replication pair.
	DataReplicationCounterpart *DataReplicationCounterpart `json:"dataReplicationCounterpart,omitempty"`
	DataReplicationRole        *string                     `json:"dataReplicationRole,omitempty"`
}

// Encryption options for the broker.
type EncryptionOptions struct {
	KMSKeyID       *string `json:"kmsKeyID,omitempty"`
	UseAWSOwnedKey *bool   `json:"useAWSOwnedKey,omitempty"`
}

// Id of the engine version.
type EngineVersion struct {
	Name *string `json:"name,omitempty"`
}

// Optional. The metadata of the LDAP server used to authenticate and authorize
// connections to the broker.
//
// Does not apply to RabbitMQ brokers.
type LDAPServerMetadataInput struct {
	Hosts                  []*string `json:"hosts,omitempty"`
	RoleBase               *string   `json:"roleBase,omitempty"`
	RoleName               *string   `json:"roleName,omitempty"`
	RoleSearchMatching     *string   `json:"roleSearchMatching,omitempty"`
	RoleSearchSubtree      *bool     `json:"roleSearchSubtree,omitempty"`
	ServiceAccountPassword *string   `json:"serviceAccountPassword,omitempty"`
	ServiceAccountUsername *string   `json:"serviceAccountUsername,omitempty"`
	UserBase               *string   `json:"userBase,omitempty"`
	UserRoleName           *string   `json:"userRoleName,omitempty"`
	UserSearchMatching     *string   `json:"userSearchMatching,omitempty"`
	UserSearchSubtree      *bool     `json:"userSearchSubtree,omitempty"`
}

// Optional. The metadata of the LDAP server used to authenticate and authorize
// connections to the broker.
type LDAPServerMetadataOutput struct {
	Hosts                  []*string `json:"hosts,omitempty"`
	RoleBase               *string   `json:"roleBase,omitempty"`
	RoleName               *string   `json:"roleName,omitempty"`
	RoleSearchMatching     *string   `json:"roleSearchMatching,omitempty"`
	RoleSearchSubtree      *bool     `json:"roleSearchSubtree,omitempty"`
	ServiceAccountUsername *string   `json:"serviceAccountUsername,omitempty"`
	UserBase               *string   `json:"userBase,omitempty"`
	UserRoleName           *string   `json:"userRoleName,omitempty"`
	UserSearchMatching     *string   `json:"userSearchMatching,omitempty"`
	UserSearchSubtree      *bool     `json:"userSearchSubtree,omitempty"`
}

// The list of information about logs to be enabled for the specified broker.
type Logs struct {
	Audit   *bool `json:"audit,omitempty"`
	General *bool `json:"general,omitempty"`
}

// The list of information about logs currently enabled and pending to be deployed
// for the specified broker.
type LogsSummary struct {
	Audit           *bool   `json:"audit,omitempty"`
	AuditLogGroup   *string `json:"auditLogGroup,omitempty"`
	General         *bool   `json:"general,omitempty"`
	GeneralLogGroup *string `json:"generalLogGroup,omitempty"`
	// The list of information about logs to be enabled for the specified broker.
	Pending *PendingLogs `json:"pending,omitempty"`
}

// The list of information about logs to be enabled for the specified broker.
type PendingLogs struct {
	Audit   *bool `json:"audit,omitempty"`
	General *bool `json:"general,omitempty"`
}

// Returns information about the configuration element or attribute that was
// sanitized in the configuration.
type SanitizationWarning struct {
	AttributeName *string `json:"attributeName,omitempty"`
	ElementName   *string `json:"elementName,omitempty"`
}

// A user associated with the broker. For Amazon MQ for RabbitMQ brokers, one
// and only one administrative user is accepted and created when a broker is
// first provisioned. All subsequent broker users are created by making RabbitMQ
// API calls directly to brokers or via the RabbitMQ web console.
type User struct {
	ConsoleAccess *bool                           `json:"consoleAccess,omitempty"`
	Groups        []*string                       `json:"groups,omitempty"`
	Password      *ackv1alpha1.SecretKeyReference `json:"password,omitempty"`
	Username      *string                         `json:"username,omitempty"`
}

// Returns information about the status of the changes pending for the ActiveMQ
// user.
type UserPendingChanges struct {
	ConsoleAccess *bool     `json:"consoleAccess,omitempty"`
	Groups        []*string `json:"groups,omitempty"`
	// The type of change pending for the ActiveMQ user.
	PendingChange *string `json:"pendingChange,omitempty"`
}

// Returns a list of all broker users. Does not apply to RabbitMQ brokers.
type UserSummary struct {
	// The type of change pending for the ActiveMQ user.
	PendingChange *string `json:"pendingChange,omitempty"`
	Username      *string `json:"username,omitempty"`
}

// The scheduled time period relative to UTC during which Amazon MQ begins to
// apply pending updates or patches to the broker.
type WeeklyStartTime struct {
	DayOfWeek *string `json:"dayOfWeek,omitempty"`
	TimeOfDay *string `json:"timeOfDay,omitempty"`
	TimeZone  *string `json:"timeZone,omitempty"`
}
