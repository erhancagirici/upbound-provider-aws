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

type LayerVersionInitParameters struct {

	// List of Architectures this layer is compatible with. Currently x86_64 and arm64 can be specified.
	CompatibleArchitectures []*string `json:"compatibleArchitectures,omitempty" tf:"compatible_architectures,omitempty"`

	// List of Runtimes this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes []*string `json:"compatibleRuntimes,omitempty" tf:"compatible_runtimes,omitempty"`

	// Description of what your Lambda Layer does.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// prefixed options cannot be used.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// Unique name for your Lambda Layer
	LayerName *string `json:"layerName,omitempty" tf:"layer_name,omitempty"`

	// License info for your Lambda Layer. See License Info.
	LicenseInfo *string `json:"licenseInfo,omitempty" tf:"license_info,omitempty"`

	// S3 bucket location containing the function's deployment package. Conflicts with filename. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// S3 key of an object containing the function's deployment package. Conflicts with filename.
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`

	// Object version containing the function's deployment package. Conflicts with filename.
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`

	// Whether to retain the old version of a previously deployed Lambda Layer. Default is false. When this is not set to true, changing any of compatible_architectures, compatible_runtimes, description, filename, layer_name, license_info, s3_bucket, s3_key, s3_object_version, or source_code_hash forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either filename or s3_key. The usual way to set this is ${filebase64sha256("file.11.12 or later) or ${base64sha256(file("file.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
	SourceCodeHash *string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`
}

type LayerVersionObservation struct {

	// ARN of the Lambda Layer with version.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// List of Architectures this layer is compatible with. Currently x86_64 and arm64 can be specified.
	CompatibleArchitectures []*string `json:"compatibleArchitectures,omitempty" tf:"compatible_architectures,omitempty"`

	// List of Runtimes this layer is compatible with. Up to 5 runtimes can be specified.
	CompatibleRuntimes []*string `json:"compatibleRuntimes,omitempty" tf:"compatible_runtimes,omitempty"`

	// Date this resource was created.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// Description of what your Lambda Layer does.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// prefixed options cannot be used.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the Lambda Layer without version.
	LayerArn *string `json:"layerArn,omitempty" tf:"layer_arn,omitempty"`

	// Unique name for your Lambda Layer
	LayerName *string `json:"layerName,omitempty" tf:"layer_name,omitempty"`

	// License info for your Lambda Layer. See License Info.
	LicenseInfo *string `json:"licenseInfo,omitempty" tf:"license_info,omitempty"`

	// S3 bucket location containing the function's deployment package. Conflicts with filename. This bucket must reside in the same AWS region where you are creating the Lambda function.
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// S3 key of an object containing the function's deployment package. Conflicts with filename.
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`

	// Object version containing the function's deployment package. Conflicts with filename.
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`

	// ARN of a signing job.
	SigningJobArn *string `json:"signingJobArn,omitempty" tf:"signing_job_arn,omitempty"`

	// ARN for a signing profile version.
	SigningProfileVersionArn *string `json:"signingProfileVersionArn,omitempty" tf:"signing_profile_version_arn,omitempty"`

	// Whether to retain the old version of a previously deployed Lambda Layer. Default is false. When this is not set to true, changing any of compatible_architectures, compatible_runtimes, description, filename, layer_name, license_info, s3_bucket, s3_key, s3_object_version, or source_code_hash forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either filename or s3_key. The usual way to set this is ${filebase64sha256("file.11.12 or later) or ${base64sha256(file("file.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
	SourceCodeHash *string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`

	// Size in bytes of the function .zip file.
	SourceCodeSize *int64 `json:"sourceCodeSize,omitempty" tf:"source_code_size,omitempty"`

	// Lambda Layer version.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type LayerVersionParameters struct {

	// List of Architectures this layer is compatible with. Currently x86_64 and arm64 can be specified.
	// +kubebuilder:validation:Optional
	CompatibleArchitectures []*string `json:"compatibleArchitectures,omitempty" tf:"compatible_architectures,omitempty"`

	// List of Runtimes this layer is compatible with. Up to 5 runtimes can be specified.
	// +kubebuilder:validation:Optional
	CompatibleRuntimes []*string `json:"compatibleRuntimes,omitempty" tf:"compatible_runtimes,omitempty"`

	// Description of what your Lambda Layer does.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// prefixed options cannot be used.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// Unique name for your Lambda Layer
	// +kubebuilder:validation:Optional
	LayerName *string `json:"layerName,omitempty" tf:"layer_name,omitempty"`

	// License info for your Lambda Layer. See License Info.
	// +kubebuilder:validation:Optional
	LicenseInfo *string `json:"licenseInfo,omitempty" tf:"license_info,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// S3 bucket location containing the function's deployment package. Conflicts with filename. This bucket must reside in the same AWS region where you are creating the Lambda function.
	// +kubebuilder:validation:Optional
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// S3 key of an object containing the function's deployment package. Conflicts with filename.
	// +kubebuilder:validation:Optional
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`

	// Object version containing the function's deployment package. Conflicts with filename.
	// +kubebuilder:validation:Optional
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`

	// Whether to retain the old version of a previously deployed Lambda Layer. Default is false. When this is not set to true, changing any of compatible_architectures, compatible_runtimes, description, filename, layer_name, license_info, s3_bucket, s3_key, s3_object_version, or source_code_hash forces deletion of the existing layer version and creation of a new layer version.
	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either filename or s3_key. The usual way to set this is ${filebase64sha256("file.11.12 or later) or ${base64sha256(file("file.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
	// +kubebuilder:validation:Optional
	SourceCodeHash *string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`
}

// LayerVersionSpec defines the desired state of LayerVersion
type LayerVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LayerVersionParameters `json:"forProvider"`
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
	InitProvider LayerVersionInitParameters `json:"initProvider,omitempty"`
}

// LayerVersionStatus defines the observed state of LayerVersion.
type LayerVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LayerVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LayerVersion is the Schema for the LayerVersions API. Provides a Lambda Layer Version resource. Lambda Layers allow you to reuse shared bits of code across multiple lambda functions.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LayerVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.layerName) || has(self.initProvider.layerName)",message="layerName is a required parameter"
	Spec   LayerVersionSpec   `json:"spec"`
	Status LayerVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LayerVersionList contains a list of LayerVersions
type LayerVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LayerVersion `json:"items"`
}

// Repository type metadata.
var (
	LayerVersion_Kind             = "LayerVersion"
	LayerVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LayerVersion_Kind}.String()
	LayerVersion_KindAPIVersion   = LayerVersion_Kind + "." + CRDGroupVersion.String()
	LayerVersion_GroupVersionKind = CRDGroupVersion.WithKind(LayerVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&LayerVersion{}, &LayerVersionList{})
}
