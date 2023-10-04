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

type DataSetInitParameters struct {

	// The type of asset that is added to a data set. Valid values are: S3_SNAPSHOT, REDSHIFT_DATA_SHARE, and API_GATEWAY_API.
	AssetType *string `json:"assetType,omitempty" tf:"asset_type,omitempty"`

	// A description for the data set.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the data set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DataSetObservation struct {

	// The Amazon Resource Name of this data set.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The type of asset that is added to a data set. Valid values are: S3_SNAPSHOT, REDSHIFT_DATA_SHARE, and API_GATEWAY_API.
	AssetType *string `json:"assetType,omitempty" tf:"asset_type,omitempty"`

	// A description for the data set.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Id of the data set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the data set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DataSetParameters struct {

	// The type of asset that is added to a data set. Valid values are: S3_SNAPSHOT, REDSHIFT_DATA_SHARE, and API_GATEWAY_API.
	// +kubebuilder:validation:Optional
	AssetType *string `json:"assetType,omitempty" tf:"asset_type,omitempty"`

	// A description for the data set.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the data set.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DataSetSpec defines the desired state of DataSet
type DataSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetParameters `json:"forProvider"`
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
	InitProvider DataSetInitParameters `json:"initProvider,omitempty"`
}

// DataSetStatus defines the observed state of DataSet.
type DataSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSet is the Schema for the DataSets API. Provides a DataExchange DataSet
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DataSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.assetType) || has(self.initProvider.assetType)",message="assetType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || has(self.initProvider.description)",message="description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   DataSetSpec   `json:"spec"`
	Status DataSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetList contains a list of DataSets
type DataSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSet `json:"items"`
}

// Repository type metadata.
var (
	DataSet_Kind             = "DataSet"
	DataSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSet_Kind}.String()
	DataSet_KindAPIVersion   = DataSet_Kind + "." + CRDGroupVersion.String()
	DataSet_GroupVersionKind = CRDGroupVersion.WithKind(DataSet_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSet{}, &DataSetList{})
}
