//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionRequired) DeepCopyInto(out *ActionRequired) {
	*out = *in
	if in.ActionRequiredCode != nil {
		in, out := &in.ActionRequiredCode, &out.ActionRequiredCode
		*out = new(string)
		**out = **in
	}
	if in.ActionRequiredInfo != nil {
		in, out := &in.ActionRequiredInfo, &out.ActionRequiredInfo
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionRequired.
func (in *ActionRequired) DeepCopy() *ActionRequired {
	if in == nil {
		return nil
	}
	out := new(ActionRequired)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityZone) DeepCopyInto(out *AvailabilityZone) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityZone.
func (in *AvailabilityZone) DeepCopy() *AvailabilityZone {
	if in == nil {
		return nil
	}
	out := new(AvailabilityZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Broker) DeepCopyInto(out *Broker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Broker.
func (in *Broker) DeepCopy() *Broker {
	if in == nil {
		return nil
	}
	out := new(Broker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Broker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerEngineType) DeepCopyInto(out *BrokerEngineType) {
	*out = *in
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerEngineType.
func (in *BrokerEngineType) DeepCopy() *BrokerEngineType {
	if in == nil {
		return nil
	}
	out := new(BrokerEngineType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerInstance) DeepCopyInto(out *BrokerInstance) {
	*out = *in
	if in.ConsoleURL != nil {
		in, out := &in.ConsoleURL, &out.ConsoleURL
		*out = new(string)
		**out = **in
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerInstance.
func (in *BrokerInstance) DeepCopy() *BrokerInstance {
	if in == nil {
		return nil
	}
	out := new(BrokerInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerInstanceOption) DeepCopyInto(out *BrokerInstanceOption) {
	*out = *in
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.HostInstanceType != nil {
		in, out := &in.HostInstanceType, &out.HostInstanceType
		*out = new(string)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.SupportedEngineVersions != nil {
		in, out := &in.SupportedEngineVersions, &out.SupportedEngineVersions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerInstanceOption.
func (in *BrokerInstanceOption) DeepCopy() *BrokerInstanceOption {
	if in == nil {
		return nil
	}
	out := new(BrokerInstanceOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerList) DeepCopyInto(out *BrokerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Broker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerList.
func (in *BrokerList) DeepCopy() *BrokerList {
	if in == nil {
		return nil
	}
	out := new(BrokerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BrokerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpec) DeepCopyInto(out *BrokerSpec) {
	*out = *in
	if in.AuthenticationStrategy != nil {
		in, out := &in.AuthenticationStrategy, &out.AuthenticationStrategy
		*out = new(string)
		**out = **in
	}
	if in.AutoMinorVersionUpgrade != nil {
		in, out := &in.AutoMinorVersionUpgrade, &out.AutoMinorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(ConfigurationID)
		(*in).DeepCopyInto(*out)
	}
	if in.CreatorRequestID != nil {
		in, out := &in.CreatorRequestID, &out.CreatorRequestID
		*out = new(string)
		**out = **in
	}
	if in.DeploymentMode != nil {
		in, out := &in.DeploymentMode, &out.DeploymentMode
		*out = new(string)
		**out = **in
	}
	if in.EncryptionOptions != nil {
		in, out := &in.EncryptionOptions, &out.EncryptionOptions
		*out = new(EncryptionOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.HostInstanceType != nil {
		in, out := &in.HostInstanceType, &out.HostInstanceType
		*out = new(string)
		**out = **in
	}
	if in.LDAPServerMetadata != nil {
		in, out := &in.LDAPServerMetadata, &out.LDAPServerMetadata
		*out = new(LDAPServerMetadataInput)
		(*in).DeepCopyInto(*out)
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = new(Logs)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceWindowStartTime != nil {
		in, out := &in.MaintenanceWindowStartTime, &out.MaintenanceWindowStartTime
		*out = new(WeeklyStartTime)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PubliclyAccessible != nil {
		in, out := &in.PubliclyAccessible, &out.PubliclyAccessible
		*out = new(bool)
		**out = **in
	}
	if in.SecurityGroupRefs != nil {
		in, out := &in.SecurityGroupRefs, &out.SecurityGroupRefs
		*out = make([]*corev1alpha1.AWSResourceReferenceWrapper, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.AWSResourceReferenceWrapper)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetRefs != nil {
		in, out := &in.SubnetRefs, &out.SubnetRefs
		*out = make([]*corev1alpha1.AWSResourceReferenceWrapper, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.AWSResourceReferenceWrapper)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]*User, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(User)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpec.
func (in *BrokerSpec) DeepCopy() *BrokerSpec {
	if in == nil {
		return nil
	}
	out := new(BrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerStatus) DeepCopyInto(out *BrokerStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.BrokerID != nil {
		in, out := &in.BrokerID, &out.BrokerID
		*out = new(string)
		**out = **in
	}
	if in.BrokerInstances != nil {
		in, out := &in.BrokerInstances, &out.BrokerInstances
		*out = make([]*BrokerInstance, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BrokerInstance)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.BrokerState != nil {
		in, out := &in.BrokerState, &out.BrokerState
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerStatus.
func (in *BrokerStatus) DeepCopy() *BrokerStatus {
	if in == nil {
		return nil
	}
	out := new(BrokerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSummary) DeepCopyInto(out *BrokerSummary) {
	*out = *in
	if in.BrokerARN != nil {
		in, out := &in.BrokerARN, &out.BrokerARN
		*out = new(string)
		**out = **in
	}
	if in.BrokerID != nil {
		in, out := &in.BrokerID, &out.BrokerID
		*out = new(string)
		**out = **in
	}
	if in.BrokerName != nil {
		in, out := &in.BrokerName, &out.BrokerName
		*out = new(string)
		**out = **in
	}
	if in.BrokerState != nil {
		in, out := &in.BrokerState, &out.BrokerState
		*out = new(string)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = (*in).DeepCopy()
	}
	if in.DeploymentMode != nil {
		in, out := &in.DeploymentMode, &out.DeploymentMode
		*out = new(string)
		**out = **in
	}
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.HostInstanceType != nil {
		in, out := &in.HostInstanceType, &out.HostInstanceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSummary.
func (in *BrokerSummary) DeepCopy() *BrokerSummary {
	if in == nil {
		return nil
	}
	out := new(BrokerSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationStrategy != nil {
		in, out := &in.AuthenticationStrategy, &out.AuthenticationStrategy
		*out = new(string)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = (*in).DeepCopy()
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationID) DeepCopyInto(out *ConfigurationID) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationID.
func (in *ConfigurationID) DeepCopy() *ConfigurationID {
	if in == nil {
		return nil
	}
	out := new(ConfigurationID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationRevision) DeepCopyInto(out *ConfigurationRevision) {
	*out = *in
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = (*in).DeepCopy()
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationRevision.
func (in *ConfigurationRevision) DeepCopy() *ConfigurationRevision {
	if in == nil {
		return nil
	}
	out := new(ConfigurationRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configurations) DeepCopyInto(out *Configurations) {
	*out = *in
	if in.Current != nil {
		in, out := &in.Current, &out.Current
		*out = new(ConfigurationID)
		(*in).DeepCopyInto(*out)
	}
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = make([]*ConfigurationID, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ConfigurationID)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Pending != nil {
		in, out := &in.Pending, &out.Pending
		*out = new(ConfigurationID)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configurations.
func (in *Configurations) DeepCopy() *Configurations {
	if in == nil {
		return nil
	}
	out := new(Configurations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionOptions) DeepCopyInto(out *EncryptionOptions) {
	*out = *in
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.UseAWSOwnedKey != nil {
		in, out := &in.UseAWSOwnedKey, &out.UseAWSOwnedKey
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionOptions.
func (in *EncryptionOptions) DeepCopy() *EncryptionOptions {
	if in == nil {
		return nil
	}
	out := new(EncryptionOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EngineVersion) DeepCopyInto(out *EngineVersion) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EngineVersion.
func (in *EngineVersion) DeepCopy() *EngineVersion {
	if in == nil {
		return nil
	}
	out := new(EngineVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LDAPServerMetadataInput) DeepCopyInto(out *LDAPServerMetadataInput) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RoleBase != nil {
		in, out := &in.RoleBase, &out.RoleBase
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchMatching != nil {
		in, out := &in.RoleSearchMatching, &out.RoleSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchSubtree != nil {
		in, out := &in.RoleSearchSubtree, &out.RoleSearchSubtree
		*out = new(bool)
		**out = **in
	}
	if in.ServiceAccountPassword != nil {
		in, out := &in.ServiceAccountPassword, &out.ServiceAccountPassword
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountUsername != nil {
		in, out := &in.ServiceAccountUsername, &out.ServiceAccountUsername
		*out = new(string)
		**out = **in
	}
	if in.UserBase != nil {
		in, out := &in.UserBase, &out.UserBase
		*out = new(string)
		**out = **in
	}
	if in.UserRoleName != nil {
		in, out := &in.UserRoleName, &out.UserRoleName
		*out = new(string)
		**out = **in
	}
	if in.UserSearchMatching != nil {
		in, out := &in.UserSearchMatching, &out.UserSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.UserSearchSubtree != nil {
		in, out := &in.UserSearchSubtree, &out.UserSearchSubtree
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LDAPServerMetadataInput.
func (in *LDAPServerMetadataInput) DeepCopy() *LDAPServerMetadataInput {
	if in == nil {
		return nil
	}
	out := new(LDAPServerMetadataInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LDAPServerMetadataOutput) DeepCopyInto(out *LDAPServerMetadataOutput) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RoleBase != nil {
		in, out := &in.RoleBase, &out.RoleBase
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchMatching != nil {
		in, out := &in.RoleSearchMatching, &out.RoleSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchSubtree != nil {
		in, out := &in.RoleSearchSubtree, &out.RoleSearchSubtree
		*out = new(bool)
		**out = **in
	}
	if in.ServiceAccountUsername != nil {
		in, out := &in.ServiceAccountUsername, &out.ServiceAccountUsername
		*out = new(string)
		**out = **in
	}
	if in.UserBase != nil {
		in, out := &in.UserBase, &out.UserBase
		*out = new(string)
		**out = **in
	}
	if in.UserRoleName != nil {
		in, out := &in.UserRoleName, &out.UserRoleName
		*out = new(string)
		**out = **in
	}
	if in.UserSearchMatching != nil {
		in, out := &in.UserSearchMatching, &out.UserSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.UserSearchSubtree != nil {
		in, out := &in.UserSearchSubtree, &out.UserSearchSubtree
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LDAPServerMetadataOutput.
func (in *LDAPServerMetadataOutput) DeepCopy() *LDAPServerMetadataOutput {
	if in == nil {
		return nil
	}
	out := new(LDAPServerMetadataOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logs) DeepCopyInto(out *Logs) {
	*out = *in
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(bool)
		**out = **in
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logs.
func (in *Logs) DeepCopy() *Logs {
	if in == nil {
		return nil
	}
	out := new(Logs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsSummary) DeepCopyInto(out *LogsSummary) {
	*out = *in
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(bool)
		**out = **in
	}
	if in.AuditLogGroup != nil {
		in, out := &in.AuditLogGroup, &out.AuditLogGroup
		*out = new(string)
		**out = **in
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(bool)
		**out = **in
	}
	if in.GeneralLogGroup != nil {
		in, out := &in.GeneralLogGroup, &out.GeneralLogGroup
		*out = new(string)
		**out = **in
	}
	if in.Pending != nil {
		in, out := &in.Pending, &out.Pending
		*out = new(PendingLogs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsSummary.
func (in *LogsSummary) DeepCopy() *LogsSummary {
	if in == nil {
		return nil
	}
	out := new(LogsSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingLogs) DeepCopyInto(out *PendingLogs) {
	*out = *in
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(bool)
		**out = **in
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingLogs.
func (in *PendingLogs) DeepCopy() *PendingLogs {
	if in == nil {
		return nil
	}
	out := new(PendingLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SanitizationWarning) DeepCopyInto(out *SanitizationWarning) {
	*out = *in
	if in.AttributeName != nil {
		in, out := &in.AttributeName, &out.AttributeName
		*out = new(string)
		**out = **in
	}
	if in.ElementName != nil {
		in, out := &in.ElementName, &out.ElementName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SanitizationWarning.
func (in *SanitizationWarning) DeepCopy() *SanitizationWarning {
	if in == nil {
		return nil
	}
	out := new(SanitizationWarning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	if in.ConsoleAccess != nil {
		in, out := &in.ConsoleAccess, &out.ConsoleAccess
		*out = new(bool)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(corev1alpha1.SecretKeyReference)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPendingChanges) DeepCopyInto(out *UserPendingChanges) {
	*out = *in
	if in.ConsoleAccess != nil {
		in, out := &in.ConsoleAccess, &out.ConsoleAccess
		*out = new(bool)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PendingChange != nil {
		in, out := &in.PendingChange, &out.PendingChange
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPendingChanges.
func (in *UserPendingChanges) DeepCopy() *UserPendingChanges {
	if in == nil {
		return nil
	}
	out := new(UserPendingChanges)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSummary) DeepCopyInto(out *UserSummary) {
	*out = *in
	if in.PendingChange != nil {
		in, out := &in.PendingChange, &out.PendingChange
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSummary.
func (in *UserSummary) DeepCopy() *UserSummary {
	if in == nil {
		return nil
	}
	out := new(UserSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyStartTime) DeepCopyInto(out *WeeklyStartTime) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.TimeOfDay != nil {
		in, out := &in.TimeOfDay, &out.TimeOfDay
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyStartTime.
func (in *WeeklyStartTime) DeepCopy() *WeeklyStartTime {
	if in == nil {
		return nil
	}
	out := new(WeeklyStartTime)
	in.DeepCopyInto(out)
	return out
}
