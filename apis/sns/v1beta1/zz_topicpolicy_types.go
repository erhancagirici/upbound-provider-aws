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

type TopicPolicyInitParameters struct {

	// The fully-formed AWS policy as JSON.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type TopicPolicyObservation struct {

	// The ARN of the SNS topic
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS Account ID of the SNS topic owner
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The fully-formed AWS policy as JSON.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type TopicPolicyParameters struct {

	// The ARN of the SNS topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Reference to a Topic in sns to populate arn.
	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate arn.
	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// The fully-formed AWS policy as JSON.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// TopicPolicySpec defines the desired state of TopicPolicy
type TopicPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicPolicyParameters `json:"forProvider"`
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
	InitProvider TopicPolicyInitParameters `json:"initProvider,omitempty"`
}

// TopicPolicyStatus defines the observed state of TopicPolicy.
type TopicPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TopicPolicy is the Schema for the TopicPolicys API. Provides an SNS topic policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TopicPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || has(self.initProvider.policy)",message="policy is a required parameter"
	Spec   TopicPolicySpec   `json:"spec"`
	Status TopicPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicPolicyList contains a list of TopicPolicys
type TopicPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicPolicy `json:"items"`
}

// Repository type metadata.
var (
	TopicPolicy_Kind             = "TopicPolicy"
	TopicPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicPolicy_Kind}.String()
	TopicPolicy_KindAPIVersion   = TopicPolicy_Kind + "." + CRDGroupVersion.String()
	TopicPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TopicPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicPolicy{}, &TopicPolicyList{})
}
