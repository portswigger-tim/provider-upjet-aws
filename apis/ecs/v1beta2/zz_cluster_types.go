// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterInitParameters struct {

	// The execute command configuration for the cluster. Detailed below.
	Configuration *ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Configures a default Service Connect namespace. Detailed below.
	ServiceConnectDefaults *ServiceConnectDefaultsInitParameters `json:"serviceConnectDefaults,omitempty" tf:"service_connect_defaults,omitempty"`

	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
	Setting []SettingInitParameters `json:"setting,omitempty" tf:"setting,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClusterObservation struct {

	// ARN that identifies the cluster.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The execute command configuration for the cluster. Detailed below.
	Configuration *ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// ARN that identifies the cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configures a default Service Connect namespace. Detailed below.
	ServiceConnectDefaults *ServiceConnectDefaultsObservation `json:"serviceConnectDefaults,omitempty" tf:"service_connect_defaults,omitempty"`

	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
	Setting []SettingObservation `json:"setting,omitempty" tf:"setting,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterParameters struct {

	// The execute command configuration for the cluster. Detailed below.
	// +kubebuilder:validation:Optional
	Configuration *ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configures a default Service Connect namespace. Detailed below.
	// +kubebuilder:validation:Optional
	ServiceConnectDefaults *ServiceConnectDefaultsParameters `json:"serviceConnectDefaults,omitempty" tf:"service_connect_defaults,omitempty"`

	// Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
	// +kubebuilder:validation:Optional
	Setting []SettingParameters `json:"setting,omitempty" tf:"setting,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigurationInitParameters struct {

	// The details of the execute command configuration. Detailed below.
	ExecuteCommandConfiguration *ExecuteCommandConfigurationInitParameters `json:"executeCommandConfiguration,omitempty" tf:"execute_command_configuration,omitempty"`
}

type ConfigurationObservation struct {

	// The details of the execute command configuration. Detailed below.
	ExecuteCommandConfiguration *ExecuteCommandConfigurationObservation `json:"executeCommandConfiguration,omitempty" tf:"execute_command_configuration,omitempty"`
}

type ConfigurationParameters struct {

	// The details of the execute command configuration. Detailed below.
	// +kubebuilder:validation:Optional
	ExecuteCommandConfiguration *ExecuteCommandConfigurationParameters `json:"executeCommandConfiguration,omitempty" tf:"execute_command_configuration,omitempty"`
}

type ExecuteCommandConfigurationInitParameters struct {

	// The AWS Key Management Service key ID to encrypt the data between the local client and the container.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The log configuration for the results of the execute command actions Required when logging is OVERRIDE. Detailed below.
	LogConfiguration *LogConfigurationInitParameters `json:"logConfiguration,omitempty" tf:"log_configuration,omitempty"`

	// The log setting to use for redirecting logs for your execute command results. Valid values are NONE, DEFAULT, and OVERRIDE.
	Logging *string `json:"logging,omitempty" tf:"logging,omitempty"`
}

type ExecuteCommandConfigurationObservation struct {

	// The AWS Key Management Service key ID to encrypt the data between the local client and the container.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The log configuration for the results of the execute command actions Required when logging is OVERRIDE. Detailed below.
	LogConfiguration *LogConfigurationObservation `json:"logConfiguration,omitempty" tf:"log_configuration,omitempty"`

	// The log setting to use for redirecting logs for your execute command results. Valid values are NONE, DEFAULT, and OVERRIDE.
	Logging *string `json:"logging,omitempty" tf:"logging,omitempty"`
}

type ExecuteCommandConfigurationParameters struct {

	// The AWS Key Management Service key ID to encrypt the data between the local client and the container.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The log configuration for the results of the execute command actions Required when logging is OVERRIDE. Detailed below.
	// +kubebuilder:validation:Optional
	LogConfiguration *LogConfigurationParameters `json:"logConfiguration,omitempty" tf:"log_configuration,omitempty"`

	// The log setting to use for redirecting logs for your execute command results. Valid values are NONE, DEFAULT, and OVERRIDE.
	// +kubebuilder:validation:Optional
	Logging *string `json:"logging,omitempty" tf:"logging,omitempty"`
}

type LogConfigurationInitParameters struct {

	// Whether or not to enable encryption on the CloudWatch logs. If not specified, encryption will be disabled.
	CloudWatchEncryptionEnabled *bool `json:"cloudWatchEncryptionEnabled,omitempty" tf:"cloud_watch_encryption_enabled,omitempty"`

	// The name of the CloudWatch log group to send logs to.
	CloudWatchLogGroupName *string `json:"cloudWatchLogGroupName,omitempty" tf:"cloud_watch_log_group_name,omitempty"`

	// Whether or not to enable encryption on the logs sent to S3. If not specified, encryption will be disabled.
	S3BucketEncryptionEnabled *bool `json:"s3BucketEncryptionEnabled,omitempty" tf:"s3_bucket_encryption_enabled,omitempty"`

	// The name of the S3 bucket to send logs to.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// An optional folder in the S3 bucket to place logs in.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type LogConfigurationObservation struct {

	// Whether or not to enable encryption on the CloudWatch logs. If not specified, encryption will be disabled.
	CloudWatchEncryptionEnabled *bool `json:"cloudWatchEncryptionEnabled,omitempty" tf:"cloud_watch_encryption_enabled,omitempty"`

	// The name of the CloudWatch log group to send logs to.
	CloudWatchLogGroupName *string `json:"cloudWatchLogGroupName,omitempty" tf:"cloud_watch_log_group_name,omitempty"`

	// Whether or not to enable encryption on the logs sent to S3. If not specified, encryption will be disabled.
	S3BucketEncryptionEnabled *bool `json:"s3BucketEncryptionEnabled,omitempty" tf:"s3_bucket_encryption_enabled,omitempty"`

	// The name of the S3 bucket to send logs to.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// An optional folder in the S3 bucket to place logs in.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type LogConfigurationParameters struct {

	// Whether or not to enable encryption on the CloudWatch logs. If not specified, encryption will be disabled.
	// +kubebuilder:validation:Optional
	CloudWatchEncryptionEnabled *bool `json:"cloudWatchEncryptionEnabled,omitempty" tf:"cloud_watch_encryption_enabled,omitempty"`

	// The name of the CloudWatch log group to send logs to.
	// +kubebuilder:validation:Optional
	CloudWatchLogGroupName *string `json:"cloudWatchLogGroupName,omitempty" tf:"cloud_watch_log_group_name,omitempty"`

	// Whether or not to enable encryption on the logs sent to S3. If not specified, encryption will be disabled.
	// +kubebuilder:validation:Optional
	S3BucketEncryptionEnabled *bool `json:"s3BucketEncryptionEnabled,omitempty" tf:"s3_bucket_encryption_enabled,omitempty"`

	// The name of the S3 bucket to send logs to.
	// +kubebuilder:validation:Optional
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// An optional folder in the S3 bucket to place logs in.
	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type ServiceConnectDefaultsInitParameters struct {

	// The ARN of the aws_service_discovery_http_namespace that's used when you create a service and don't specify a Service Connect configuration.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type ServiceConnectDefaultsObservation struct {

	// The ARN of the aws_service_discovery_http_namespace that's used when you create a service and don't specify a Service Connect configuration.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type ServiceConnectDefaultsParameters struct {

	// The ARN of the aws_service_discovery_http_namespace that's used when you create a service and don't specify a Service Connect configuration.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type SettingInitParameters struct {

	// Name of the setting to manage. Valid values: containerInsights.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value to assign to the setting. Valid values are enabled and disabled.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SettingObservation struct {

	// Name of the setting to manage. Valid values: containerInsights.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value to assign to the setting. Valid values are enabled and disabled.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SettingParameters struct {

	// Name of the setting to manage. Valid values: containerInsights.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value to assign to the setting. Valid values are enabled and disabled.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Cluster is the Schema for the Clusters API. Provides an ECS cluster.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}