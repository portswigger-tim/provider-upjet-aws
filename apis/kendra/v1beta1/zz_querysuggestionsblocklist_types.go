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

type QuerySuggestionsBlockListObservation struct {

	// ARN of the block list.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique indentifier of the block list.
	QuerySuggestionsBlockListID *string `json:"querySuggestionsBlockListId,omitempty" tf:"query_suggestions_block_list_id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type QuerySuggestionsBlockListParameters struct {

	// The description for a block list.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the index for a block list.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kendra/v1beta1.Index
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IndexID *string `json:"indexId,omitempty" tf:"index_id,omitempty"`

	// Reference to a Index in kendra to populate indexId.
	// +kubebuilder:validation:Optional
	IndexIDRef *v1.Reference `json:"indexIdRef,omitempty" tf:"-"`

	// Selector for a Index in kendra to populate indexId.
	// +kubebuilder:validation:Optional
	IndexIDSelector *v1.Selector `json:"indexIdSelector,omitempty" tf:"-"`

	// The name for the block list.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The IAM (Identity and Access Management) role used to access the block list text file in S3.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The S3 path where your block list text file sits in S3. Detailed below.
	// +kubebuilder:validation:Required
	SourceS3Path []SourceS3PathParameters `json:"sourceS3Path" tf:"source_s3_path,omitempty"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SourceS3PathObservation struct {
}

type SourceS3PathParameters struct {

	// The name of the S3 bucket that contains the file.
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

	// The name of the file.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

// QuerySuggestionsBlockListSpec defines the desired state of QuerySuggestionsBlockList
type QuerySuggestionsBlockListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QuerySuggestionsBlockListParameters `json:"forProvider"`
}

// QuerySuggestionsBlockListStatus defines the observed state of QuerySuggestionsBlockList.
type QuerySuggestionsBlockListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QuerySuggestionsBlockListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QuerySuggestionsBlockList is the Schema for the QuerySuggestionsBlockLists API. Upbound official provider resource for managing an aws kendra block list used for query suggestions for an index
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type QuerySuggestionsBlockList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QuerySuggestionsBlockListSpec   `json:"spec"`
	Status            QuerySuggestionsBlockListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QuerySuggestionsBlockListList contains a list of QuerySuggestionsBlockLists
type QuerySuggestionsBlockListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuerySuggestionsBlockList `json:"items"`
}

// Repository type metadata.
var (
	QuerySuggestionsBlockList_Kind             = "QuerySuggestionsBlockList"
	QuerySuggestionsBlockList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QuerySuggestionsBlockList_Kind}.String()
	QuerySuggestionsBlockList_KindAPIVersion   = QuerySuggestionsBlockList_Kind + "." + CRDGroupVersion.String()
	QuerySuggestionsBlockList_GroupVersionKind = CRDGroupVersion.WithKind(QuerySuggestionsBlockList_Kind)
)

func init() {
	SchemeBuilder.Register(&QuerySuggestionsBlockList{}, &QuerySuggestionsBlockListList{})
}
