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

type SubnetCidrReservationInitParameters struct {

	// The CIDR block for the reservation.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// A brief description of the reservation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The type of reservation to create. Valid values: explicit, prefix
	ReservationType *string `json:"reservationType,omitempty" tf:"reservation_type,omitempty"`

	// The ID of the subnet to create the reservation for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type SubnetCidrReservationObservation struct {

	// The CIDR block for the reservation.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// A brief description of the reservation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the CIDR reservation.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the AWS account that owns this CIDR reservation.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// The type of reservation to create. Valid values: explicit, prefix
	ReservationType *string `json:"reservationType,omitempty" tf:"reservation_type,omitempty"`

	// The ID of the subnet to create the reservation for.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SubnetCidrReservationParameters struct {

	// The CIDR block for the reservation.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// A brief description of the reservation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The type of reservation to create. Valid values: explicit, prefix
	// +kubebuilder:validation:Optional
	ReservationType *string `json:"reservationType,omitempty" tf:"reservation_type,omitempty"`

	// The ID of the subnet to create the reservation for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// SubnetCidrReservationSpec defines the desired state of SubnetCidrReservation
type SubnetCidrReservationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetCidrReservationParameters `json:"forProvider"`
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
	InitProvider SubnetCidrReservationInitParameters `json:"initProvider,omitempty"`
}

// SubnetCidrReservationStatus defines the observed state of SubnetCidrReservation.
type SubnetCidrReservationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetCidrReservationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubnetCidrReservation is the Schema for the SubnetCidrReservations API. Provides a subnet CIDR reservation resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SubnetCidrReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidrBlock) || (has(self.initProvider) && has(self.initProvider.cidrBlock))",message="spec.forProvider.cidrBlock is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.reservationType) || (has(self.initProvider) && has(self.initProvider.reservationType))",message="spec.forProvider.reservationType is a required parameter"
	Spec   SubnetCidrReservationSpec   `json:"spec"`
	Status SubnetCidrReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetCidrReservationList contains a list of SubnetCidrReservations
type SubnetCidrReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetCidrReservation `json:"items"`
}

// Repository type metadata.
var (
	SubnetCidrReservation_Kind             = "SubnetCidrReservation"
	SubnetCidrReservation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetCidrReservation_Kind}.String()
	SubnetCidrReservation_KindAPIVersion   = SubnetCidrReservation_Kind + "." + CRDGroupVersion.String()
	SubnetCidrReservation_GroupVersionKind = CRDGroupVersion.WithKind(SubnetCidrReservation_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetCidrReservation{}, &SubnetCidrReservationList{})
}
