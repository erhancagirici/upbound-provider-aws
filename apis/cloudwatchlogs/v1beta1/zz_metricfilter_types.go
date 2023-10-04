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

type MetricFilterInitParameters struct {

	// A block defining collection of information needed to define how metric data gets emitted. See below.
	MetricTransformation []MetricTransformationInitParameters `json:"metricTransformation,omitempty" tf:"metric_transformation,omitempty"`

	// A valid CloudWatch Logs filter pattern
	// for extracting metric data out of ingested log events.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type MetricFilterObservation struct {

	// The name of the metric filter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the log group to associate the metric filter with.
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// A block defining collection of information needed to define how metric data gets emitted. See below.
	MetricTransformation []MetricTransformationObservation `json:"metricTransformation,omitempty" tf:"metric_transformation,omitempty"`

	// A valid CloudWatch Logs filter pattern
	// for extracting metric data out of ingested log events.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type MetricFilterParameters struct {

	// The name of the log group to associate the metric filter with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1.Group
	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// Reference to a Group in cloudwatchlogs to populate logGroupName.
	// +kubebuilder:validation:Optional
	LogGroupNameRef *v1.Reference `json:"logGroupNameRef,omitempty" tf:"-"`

	// Selector for a Group in cloudwatchlogs to populate logGroupName.
	// +kubebuilder:validation:Optional
	LogGroupNameSelector *v1.Selector `json:"logGroupNameSelector,omitempty" tf:"-"`

	// A block defining collection of information needed to define how metric data gets emitted. See below.
	// +kubebuilder:validation:Optional
	MetricTransformation []MetricTransformationParameters `json:"metricTransformation,omitempty" tf:"metric_transformation,omitempty"`

	// A valid CloudWatch Logs filter pattern
	// for extracting metric data out of ingested log events.
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type MetricTransformationInitParameters struct {

	// The value to emit when a filter pattern does not match a log event. Conflicts with dimensions.
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// Map of fields to use as dimensions for the metric. Up to 3 dimensions are allowed. Conflicts with default_value.
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// The name of the CloudWatch metric to which the monitored log information should be published (e.g., ErrorCount)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The destination namespace of the CloudWatch metric.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The unit to assign to the metric. If you omit this, the unit is set as None.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// What to publish to the metric. For example, if you're counting the occurrences of a particular term like "Error", the value will be "1" for each occurrence. If you're counting the bytes transferred the published value will be the value in the log event.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MetricTransformationObservation struct {

	// The value to emit when a filter pattern does not match a log event. Conflicts with dimensions.
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// Map of fields to use as dimensions for the metric. Up to 3 dimensions are allowed. Conflicts with default_value.
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// The name of the CloudWatch metric to which the monitored log information should be published (e.g., ErrorCount)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The destination namespace of the CloudWatch metric.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The unit to assign to the metric. If you omit this, the unit is set as None.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// What to publish to the metric. For example, if you're counting the occurrences of a particular term like "Error", the value will be "1" for each occurrence. If you're counting the bytes transferred the published value will be the value in the log event.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MetricTransformationParameters struct {

	// The value to emit when a filter pattern does not match a log event. Conflicts with dimensions.
	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// Map of fields to use as dimensions for the metric. Up to 3 dimensions are allowed. Conflicts with default_value.
	// +kubebuilder:validation:Optional
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// The name of the CloudWatch metric to which the monitored log information should be published (e.g., ErrorCount)
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The destination namespace of the CloudWatch metric.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// The unit to assign to the metric. If you omit this, the unit is set as None.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// What to publish to the metric. For example, if you're counting the occurrences of a particular term like "Error", the value will be "1" for each occurrence. If you're counting the bytes transferred the published value will be the value in the log event.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// MetricFilterSpec defines the desired state of MetricFilter
type MetricFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricFilterParameters `json:"forProvider"`
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
	InitProvider MetricFilterInitParameters `json:"initProvider,omitempty"`
}

// MetricFilterStatus defines the observed state of MetricFilter.
type MetricFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MetricFilter is the Schema for the MetricFilters API. Provides a CloudWatch Log Metric Filter resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MetricFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metricTransformation) || has(self.initProvider.metricTransformation)",message="metricTransformation is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pattern) || has(self.initProvider.pattern)",message="pattern is a required parameter"
	Spec   MetricFilterSpec   `json:"spec"`
	Status MetricFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricFilterList contains a list of MetricFilters
type MetricFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricFilter `json:"items"`
}

// Repository type metadata.
var (
	MetricFilter_Kind             = "MetricFilter"
	MetricFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetricFilter_Kind}.String()
	MetricFilter_KindAPIVersion   = MetricFilter_Kind + "." + CRDGroupVersion.String()
	MetricFilter_GroupVersionKind = CRDGroupVersion.WithKind(MetricFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&MetricFilter{}, &MetricFilterList{})
}
