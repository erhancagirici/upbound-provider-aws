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

type PublicKeyInitParameters struct {

	// An optional comment about the public key.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The name for the public key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PublicKeyObservation struct {

	// Internal value used by CloudFront to allow future updates to the public key configuration.
	CallerReference *string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`

	// An optional comment about the public key.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The current version of the public key. For example: E2QWRUHAPOMQZL.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The identifier for the public key. For example: K3D5EWEUDCCXON.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name for the public key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PublicKeyParameters struct {

	// An optional comment about the public key.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	// +kubebuilder:validation:Optional
	EncodedKeySecretRef v1.SecretKeySelector `json:"encodedKeySecretRef" tf:"-"`

	// The name for the public key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// PublicKeySpec defines the desired state of PublicKey
type PublicKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicKeyParameters `json:"forProvider"`
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
	InitProvider PublicKeyInitParameters `json:"initProvider,omitempty"`
}

// PublicKeyStatus defines the observed state of PublicKey.
type PublicKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PublicKey is the Schema for the PublicKeys API. Provides a CloudFront Public Key which you add to CloudFront to use with features like field-level encryption.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.encodedKeySecretRef)",message="spec.forProvider.encodedKeySecretRef is a required parameter"
	Spec   PublicKeySpec   `json:"spec"`
	Status PublicKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicKeyList contains a list of PublicKeys
type PublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicKey `json:"items"`
}

// Repository type metadata.
var (
	PublicKey_Kind             = "PublicKey"
	PublicKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicKey_Kind}.String()
	PublicKey_KindAPIVersion   = PublicKey_Kind + "." + CRDGroupVersion.String()
	PublicKey_GroupVersionKind = CRDGroupVersion.WithKind(PublicKey_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicKey{}, &PublicKeyList{})
}
