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

type DedicatedIPPoolInitParameters struct {

	// IP pool scaling mode. Valid values: STANDARD, MANAGED. If omitted, the AWS API will default to a standard pool.
	ScalingMode *string `json:"scalingMode,omitempty" tf:"scaling_mode,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DedicatedIPPoolObservation struct {

	// ARN of the Dedicated IP Pool.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP pool scaling mode. Valid values: STANDARD, MANAGED. If omitted, the AWS API will default to a standard pool.
	ScalingMode *string `json:"scalingMode,omitempty" tf:"scaling_mode,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DedicatedIPPoolParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// IP pool scaling mode. Valid values: STANDARD, MANAGED. If omitted, the AWS API will default to a standard pool.
	// +kubebuilder:validation:Optional
	ScalingMode *string `json:"scalingMode,omitempty" tf:"scaling_mode,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DedicatedIPPoolSpec defines the desired state of DedicatedIPPool
type DedicatedIPPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedIPPoolParameters `json:"forProvider"`
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
	InitProvider DedicatedIPPoolInitParameters `json:"initProvider,omitempty"`
}

// DedicatedIPPoolStatus defines the observed state of DedicatedIPPool.
type DedicatedIPPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedIPPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedIPPool is the Schema for the DedicatedIPPools API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DedicatedIPPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedIPPoolSpec   `json:"spec"`
	Status            DedicatedIPPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedIPPoolList contains a list of DedicatedIPPools
type DedicatedIPPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedIPPool `json:"items"`
}

// Repository type metadata.
var (
	DedicatedIPPool_Kind             = "DedicatedIPPool"
	DedicatedIPPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedIPPool_Kind}.String()
	DedicatedIPPool_KindAPIVersion   = DedicatedIPPool_Kind + "." + CRDGroupVersion.String()
	DedicatedIPPool_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedIPPool_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedIPPool{}, &DedicatedIPPoolList{})
}
