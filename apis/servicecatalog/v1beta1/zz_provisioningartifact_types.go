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

type ProvisioningArtifactInitParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). The default value is en.
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Whether the product version is active. Inactive provisioning artifacts are invisible to end users. End users cannot launch or update a provisioned product from an inactive provisioning artifact. Default is true.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether AWS Service Catalog stops validating the specified provisioning artifact template even if it is invalid.
	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation,omitempty"`

	// Information set by the administrator to provide guidance to end users about which provisioning artifacts to use. Valid values are DEFAULT and DEPRECATED. The default is DEFAULT. Users are able to make updates to a provisioned product of a deprecated version but cannot launch new provisioned products using a deprecated version.
	Guidance *string `json:"guidance,omitempty" tf:"guidance,omitempty"`

	// Name of the provisioning artifact (for example, v1, v2beta). No spaces are allowed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Identifier of the product.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/servicecatalog/v1beta1.Product
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Reference to a Product in servicecatalog to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDRef *v1.Reference `json:"productIdRef,omitempty" tf:"-"`

	// Selector for a Product in servicecatalog to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDSelector *v1.Selector `json:"productIdSelector,omitempty" tf:"-"`

	// Template source as the physical ID of the resource that contains the template. Currently only supports CloudFormation stack ARN. Specify the physical ID as arn:[partition]:cloudformation:[region]:[account ID]:stack/[stack name]/[resource ID].
	TemplatePhysicalID *string `json:"templatePhysicalId,omitempty" tf:"template_physical_id,omitempty"`

	// Template source as URL of the CloudFormation template in Amazon S3.
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// Type of provisioning artifact. See AWS Docs for valid list of values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProvisioningArtifactObservation struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). The default value is en.
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Whether the product version is active. Inactive provisioning artifacts are invisible to end users. End users cannot launch or update a provisioned product from an inactive provisioning artifact. Default is true.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Time when the provisioning artifact was created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether AWS Service Catalog stops validating the specified provisioning artifact template even if it is invalid.
	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation,omitempty"`

	// Information set by the administrator to provide guidance to end users about which provisioning artifacts to use. Valid values are DEFAULT and DEPRECATED. The default is DEFAULT. Users are able to make updates to a provisioned product of a deprecated version but cannot launch new provisioned products using a deprecated version.
	Guidance *string `json:"guidance,omitempty" tf:"guidance,omitempty"`

	// Provisioning artifact identifier and product identifier separated by a colon.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the provisioning artifact (for example, v1, v2beta). No spaces are allowed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Identifier of the product.
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Provisioning artifact identifier.
	ProvisioningArtifactID *string `json:"provisioningArtifactId,omitempty" tf:"provisioning_artifact_id,omitempty"`

	// Template source as the physical ID of the resource that contains the template. Currently only supports CloudFormation stack ARN. Specify the physical ID as arn:[partition]:cloudformation:[region]:[account ID]:stack/[stack name]/[resource ID].
	TemplatePhysicalID *string `json:"templatePhysicalId,omitempty" tf:"template_physical_id,omitempty"`

	// Template source as URL of the CloudFormation template in Amazon S3.
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// Type of provisioning artifact. See AWS Docs for valid list of values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProvisioningArtifactParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). The default value is en.
	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Whether the product version is active. Inactive provisioning artifacts are invisible to end users. End users cannot launch or update a provisioned product from an inactive provisioning artifact. Default is true.
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether AWS Service Catalog stops validating the specified provisioning artifact template even if it is invalid.
	// +kubebuilder:validation:Optional
	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation,omitempty"`

	// Information set by the administrator to provide guidance to end users about which provisioning artifacts to use. Valid values are DEFAULT and DEPRECATED. The default is DEFAULT. Users are able to make updates to a provisioned product of a deprecated version but cannot launch new provisioned products using a deprecated version.
	// +kubebuilder:validation:Optional
	Guidance *string `json:"guidance,omitempty" tf:"guidance,omitempty"`

	// Name of the provisioning artifact (for example, v1, v2beta). No spaces are allowed.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Identifier of the product.
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

	// Template source as the physical ID of the resource that contains the template. Currently only supports CloudFormation stack ARN. Specify the physical ID as arn:[partition]:cloudformation:[region]:[account ID]:stack/[stack name]/[resource ID].
	// +kubebuilder:validation:Optional
	TemplatePhysicalID *string `json:"templatePhysicalId,omitempty" tf:"template_physical_id,omitempty"`

	// Template source as URL of the CloudFormation template in Amazon S3.
	// +kubebuilder:validation:Optional
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// Type of provisioning artifact. See AWS Docs for valid list of values.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ProvisioningArtifactSpec defines the desired state of ProvisioningArtifact
type ProvisioningArtifactSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProvisioningArtifactParameters `json:"forProvider"`
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
	InitProvider ProvisioningArtifactInitParameters `json:"initProvider,omitempty"`
}

// ProvisioningArtifactStatus defines the observed state of ProvisioningArtifact.
type ProvisioningArtifactStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProvisioningArtifactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProvisioningArtifact is the Schema for the ProvisioningArtifacts API. Manages a Service Catalog Provisioning Artifact
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ProvisioningArtifact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProvisioningArtifactSpec   `json:"spec"`
	Status            ProvisioningArtifactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProvisioningArtifactList contains a list of ProvisioningArtifacts
type ProvisioningArtifactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProvisioningArtifact `json:"items"`
}

// Repository type metadata.
var (
	ProvisioningArtifact_Kind             = "ProvisioningArtifact"
	ProvisioningArtifact_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProvisioningArtifact_Kind}.String()
	ProvisioningArtifact_KindAPIVersion   = ProvisioningArtifact_Kind + "." + CRDGroupVersion.String()
	ProvisioningArtifact_GroupVersionKind = CRDGroupVersion.WithKind(ProvisioningArtifact_Kind)
)

func init() {
	SchemeBuilder.Register(&ProvisioningArtifact{}, &ProvisioningArtifactList{})
}
