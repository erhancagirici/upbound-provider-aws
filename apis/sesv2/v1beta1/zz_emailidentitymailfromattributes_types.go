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

type EmailIdentityMailFromAttributesInitParameters struct {

	// The action to take if the required MX record isn't found when you send an email. Valid values: USE_DEFAULT_VALUE, REJECT_MESSAGE.
	BehaviorOnMxFailure *string `json:"behaviorOnMxFailure,omitempty" tf:"behavior_on_mx_failure,omitempty"`

	// The custom MAIL FROM domain that you want the verified identity to use. Required if behavior_on_mx_failure is REJECT_MESSAGE.
	MailFromDomain *string `json:"mailFromDomain,omitempty" tf:"mail_from_domain,omitempty"`
}

type EmailIdentityMailFromAttributesObservation struct {

	// The action to take if the required MX record isn't found when you send an email. Valid values: USE_DEFAULT_VALUE, REJECT_MESSAGE.
	BehaviorOnMxFailure *string `json:"behaviorOnMxFailure,omitempty" tf:"behavior_on_mx_failure,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The custom MAIL FROM domain that you want the verified identity to use. Required if behavior_on_mx_failure is REJECT_MESSAGE.
	MailFromDomain *string `json:"mailFromDomain,omitempty" tf:"mail_from_domain,omitempty"`
}

type EmailIdentityMailFromAttributesParameters struct {

	// The action to take if the required MX record isn't found when you send an email. Valid values: USE_DEFAULT_VALUE, REJECT_MESSAGE.
	// +kubebuilder:validation:Optional
	BehaviorOnMxFailure *string `json:"behaviorOnMxFailure,omitempty" tf:"behavior_on_mx_failure,omitempty"`

	// The custom MAIL FROM domain that you want the verified identity to use. Required if behavior_on_mx_failure is REJECT_MESSAGE.
	// +kubebuilder:validation:Optional
	MailFromDomain *string `json:"mailFromDomain,omitempty" tf:"mail_from_domain,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// EmailIdentityMailFromAttributesSpec defines the desired state of EmailIdentityMailFromAttributes
type EmailIdentityMailFromAttributesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailIdentityMailFromAttributesParameters `json:"forProvider"`
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
	InitProvider EmailIdentityMailFromAttributesInitParameters `json:"initProvider,omitempty"`
}

// EmailIdentityMailFromAttributesStatus defines the observed state of EmailIdentityMailFromAttributes.
type EmailIdentityMailFromAttributesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailIdentityMailFromAttributesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EmailIdentityMailFromAttributes is the Schema for the EmailIdentityMailFromAttributess API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EmailIdentityMailFromAttributes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmailIdentityMailFromAttributesSpec   `json:"spec"`
	Status            EmailIdentityMailFromAttributesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIdentityMailFromAttributesList contains a list of EmailIdentityMailFromAttributess
type EmailIdentityMailFromAttributesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailIdentityMailFromAttributes `json:"items"`
}

// Repository type metadata.
var (
	EmailIdentityMailFromAttributes_Kind             = "EmailIdentityMailFromAttributes"
	EmailIdentityMailFromAttributes_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailIdentityMailFromAttributes_Kind}.String()
	EmailIdentityMailFromAttributes_KindAPIVersion   = EmailIdentityMailFromAttributes_Kind + "." + CRDGroupVersion.String()
	EmailIdentityMailFromAttributes_GroupVersionKind = CRDGroupVersion.WithKind(EmailIdentityMailFromAttributes_Kind)
)

func init() {
	SchemeBuilder.Register(&EmailIdentityMailFromAttributes{}, &EmailIdentityMailFromAttributesList{})
}
