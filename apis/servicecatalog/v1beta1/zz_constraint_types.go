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

type ConstraintInitParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Description of the constraint.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Type of constraint. Valid values are LAUNCH, NOTIFICATION, RESOURCE_UPDATE, STACKSET, and TEMPLATE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConstraintObservation struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Description of the constraint.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Constraint identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Owner of the constraint.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Portfolio identifier.
	PortfolioID *string `json:"portfolioId,omitempty" tf:"portfolio_id,omitempty"`

	// Product identifier.
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Type of constraint. Valid values are LAUNCH, NOTIFICATION, RESOURCE_UPDATE, STACKSET, and TEMPLATE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConstraintParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Description of the constraint.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Constraint parameters in JSON format. The syntax depends on the constraint type. See details below.
	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Portfolio identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/servicecatalog/v1beta1.Portfolio
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PortfolioID *string `json:"portfolioId,omitempty" tf:"portfolio_id,omitempty"`

	// Reference to a Portfolio in servicecatalog to populate portfolioId.
	// +kubebuilder:validation:Optional
	PortfolioIDRef *v1.Reference `json:"portfolioIdRef,omitempty" tf:"-"`

	// Selector for a Portfolio in servicecatalog to populate portfolioId.
	// +kubebuilder:validation:Optional
	PortfolioIDSelector *v1.Selector `json:"portfolioIdSelector,omitempty" tf:"-"`

	// Product identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/servicecatalog/v1beta1.Product
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Reference to a Product in servicecatalog to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDRef *v1.Reference `json:"productIdRef,omitempty" tf:"-"`

	// Selector for a Product in servicecatalog to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDSelector *v1.Selector `json:"productIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Type of constraint. Valid values are LAUNCH, NOTIFICATION, RESOURCE_UPDATE, STACKSET, and TEMPLATE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ConstraintSpec defines the desired state of Constraint
type ConstraintSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConstraintParameters `json:"forProvider"`
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
	InitProvider ConstraintInitParameters `json:"initProvider,omitempty"`
}

// ConstraintStatus defines the observed state of Constraint.
type ConstraintStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConstraintObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Constraint is the Schema for the Constraints API. Manages a Service Catalog Constraint
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Constraint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.parameters) || has(self.initProvider.parameters)",message="parameters is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || has(self.initProvider.type)",message="type is a required parameter"
	Spec   ConstraintSpec   `json:"spec"`
	Status ConstraintStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConstraintList contains a list of Constraints
type ConstraintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Constraint `json:"items"`
}

// Repository type metadata.
var (
	Constraint_Kind             = "Constraint"
	Constraint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Constraint_Kind}.String()
	Constraint_KindAPIVersion   = Constraint_Kind + "." + CRDGroupVersion.String()
	Constraint_GroupVersionKind = CRDGroupVersion.WithKind(Constraint_Kind)
)

func init() {
	SchemeBuilder.Register(&Constraint{}, &ConstraintList{})
}
