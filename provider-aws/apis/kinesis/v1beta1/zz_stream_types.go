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

type StreamModeDetailsObservation struct {
}

type StreamModeDetailsParameters struct {

	// Specifies the capacity mode of the stream. Must be either PROVISIONED or ON_DEMAND.
	// +kubebuilder:validation:Required
	StreamMode *string `json:"streamMode" tf:"stream_mode,omitempty"`
}

type StreamObservation struct {

	// The unique Stream id
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type StreamParameters struct {

	// The Amazon Resource Name  specifying the Stream
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The encryption type to use. The only acceptable values are NONE or KMS. The default value is NONE.
	// +kubebuilder:validation:Optional
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is false.
	// +kubebuilder:validation:Optional
	EnforceConsumerDeletion *bool `json:"enforceConsumerDeletion,omitempty" tf:"enforce_consumer_deletion,omitempty"`

	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias alias/aws/kinesis.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
	// +kubebuilder:validation:Optional
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	ShardCount *float64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`

	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See Monitoring with CloudWatch for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	// +kubebuilder:validation:Optional
	ShardLevelMetrics []*string `json:"shardLevelMetrics,omitempty" tf:"shard_level_metrics,omitempty"`

	// Indicates the capacity mode of the data stream. Detailed below.
	// +kubebuilder:validation:Optional
	StreamModeDetails []StreamModeDetailsParameters `json:"streamModeDetails,omitempty" tf:"stream_mode_details,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StreamSpec defines the desired state of Stream
type StreamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamParameters `json:"forProvider"`
}

// StreamStatus defines the observed state of Stream.
type StreamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stream is the Schema for the Streams API. Provides a AWS Kinesis Stream
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Stream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamSpec   `json:"spec"`
	Status            StreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamList contains a list of Streams
type StreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stream `json:"items"`
}

// Repository type metadata.
var (
	Stream_Kind             = "Stream"
	Stream_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stream_Kind}.String()
	Stream_KindAPIVersion   = Stream_Kind + "." + CRDGroupVersion.String()
	Stream_GroupVersionKind = CRDGroupVersion.WithKind(Stream_Kind)
)

func init() {
	SchemeBuilder.Register(&Stream{}, &StreamList{})
}
