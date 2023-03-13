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

type BucketIntelligentTieringConfigurationFilterObservation struct {
}

type BucketIntelligentTieringConfigurationFilterParameters struct {

	// Object key name prefix that identifies the subset of objects to which the configuration applies.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BucketIntelligentTieringConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketIntelligentTieringConfigurationParameters struct {

	// Name of the bucket this intelligent tiering configuration is associated with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Bucket filter. The configuration only includes objects that meet the filter's criteria (documented below).
	// +kubebuilder:validation:Optional
	Filter []BucketIntelligentTieringConfigurationFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique name used to identify the S3 Intelligent-Tiering configuration for the bucket.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the status of the configuration. Valid values: Enabled, Disabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// S3 Intelligent-Tiering storage class tiers of the configuration (documented below).
	// +kubebuilder:validation:Required
	Tiering []TieringParameters `json:"tiering" tf:"tiering,omitempty"`
}

type TieringObservation struct {
}

type TieringParameters struct {

	// S3 Intelligent-Tiering access tier. Valid values: ARCHIVE_ACCESS, DEEP_ARCHIVE_ACCESS.
	// +kubebuilder:validation:Required
	AccessTier *string `json:"accessTier" tf:"access_tier,omitempty"`

	// Number of consecutive days of no access after which an object will be eligible to be transitioned to the corresponding tier.
	// +kubebuilder:validation:Required
	Days *float64 `json:"days" tf:"days,omitempty"`
}

// BucketIntelligentTieringConfigurationSpec defines the desired state of BucketIntelligentTieringConfiguration
type BucketIntelligentTieringConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketIntelligentTieringConfigurationParameters `json:"forProvider"`
}

// BucketIntelligentTieringConfigurationStatus defines the observed state of BucketIntelligentTieringConfiguration.
type BucketIntelligentTieringConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketIntelligentTieringConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketIntelligentTieringConfiguration is the Schema for the BucketIntelligentTieringConfigurations API. Provides an S3 Intelligent-Tiering configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketIntelligentTieringConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketIntelligentTieringConfigurationSpec   `json:"spec"`
	Status            BucketIntelligentTieringConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketIntelligentTieringConfigurationList contains a list of BucketIntelligentTieringConfigurations
type BucketIntelligentTieringConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketIntelligentTieringConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketIntelligentTieringConfiguration_Kind             = "BucketIntelligentTieringConfiguration"
	BucketIntelligentTieringConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketIntelligentTieringConfiguration_Kind}.String()
	BucketIntelligentTieringConfiguration_KindAPIVersion   = BucketIntelligentTieringConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketIntelligentTieringConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketIntelligentTieringConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketIntelligentTieringConfiguration{}, &BucketIntelligentTieringConfigurationList{})
}
