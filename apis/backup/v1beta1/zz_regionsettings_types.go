// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegionSettingsInitParameters struct {

	// A map of services along with the management preferences for the Region. For more information, see the AWS Documentation.
	//
	// WARNING: All parameters are required to be given: EFS, DynamoDB
	ResourceTypeManagementPreference map[string]*bool `json:"resourceTypeManagementPreference,omitempty" tf:"resource_type_management_preference,omitempty"`

	// A map of services along with the opt-in preferences for the Region.
	//
	// WARNING: All parameters are required to be given: EFS, DynamoDB, EBS, EC2, FSx, S3, Aurora, RDS, Storage Gateway, VirtualMachine
	ResourceTypeOptInPreference map[string]*bool `json:"resourceTypeOptInPreference,omitempty" tf:"resource_type_opt_in_preference,omitempty"`
}

type RegionSettingsObservation struct {

	// The AWS region.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of services along with the management preferences for the Region. For more information, see the AWS Documentation.
	//
	// WARNING: All parameters are required to be given: EFS, DynamoDB
	ResourceTypeManagementPreference map[string]*bool `json:"resourceTypeManagementPreference,omitempty" tf:"resource_type_management_preference,omitempty"`

	// A map of services along with the opt-in preferences for the Region.
	//
	// WARNING: All parameters are required to be given: EFS, DynamoDB, EBS, EC2, FSx, S3, Aurora, RDS, Storage Gateway, VirtualMachine
	ResourceTypeOptInPreference map[string]*bool `json:"resourceTypeOptInPreference,omitempty" tf:"resource_type_opt_in_preference,omitempty"`
}

type RegionSettingsParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of services along with the management preferences for the Region. For more information, see the AWS Documentation.
	//
	// WARNING: All parameters are required to be given: EFS, DynamoDB
	// +kubebuilder:validation:Optional
	ResourceTypeManagementPreference map[string]*bool `json:"resourceTypeManagementPreference,omitempty" tf:"resource_type_management_preference,omitempty"`

	// A map of services along with the opt-in preferences for the Region.
	//
	// WARNING: All parameters are required to be given: EFS, DynamoDB, EBS, EC2, FSx, S3, Aurora, RDS, Storage Gateway, VirtualMachine
	// +kubebuilder:validation:Optional
	ResourceTypeOptInPreference map[string]*bool `json:"resourceTypeOptInPreference,omitempty" tf:"resource_type_opt_in_preference,omitempty"`
}

// RegionSettingsSpec defines the desired state of RegionSettings
type RegionSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionSettingsParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RegionSettingsInitParameters `json:"initProvider,omitempty"`
}

// RegionSettingsStatus defines the observed state of RegionSettings.
type RegionSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionSettings is the Schema for the RegionSettingss API. Provides an AWS Backup Region Settings resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RegionSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceTypeOptInPreference) || has(self.initProvider.resourceTypeOptInPreference)",message="resourceTypeOptInPreference is a required parameter"
	Spec   RegionSettingsSpec   `json:"spec"`
	Status RegionSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionSettingsList contains a list of RegionSettingss
type RegionSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionSettings `json:"items"`
}

// Repository type metadata.
var (
	RegionSettings_Kind             = "RegionSettings"
	RegionSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionSettings_Kind}.String()
	RegionSettings_KindAPIVersion   = RegionSettings_Kind + "." + CRDGroupVersion.String()
	RegionSettings_GroupVersionKind = CRDGroupVersion.WithKind(RegionSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionSettings{}, &RegionSettingsList{})
}
