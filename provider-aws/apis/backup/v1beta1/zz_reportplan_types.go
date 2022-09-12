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

type ReportDeliveryChannelObservation struct {
}

type ReportDeliveryChannelParameters struct {

	// A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.
	// +kubebuilder:validation:Optional
	Formats []*string `json:"formats,omitempty" tf:"formats,omitempty"`

	// The unique name of the S3 bucket that receives your reports.
	// +kubebuilder:validation:Required
	S3BucketName *string `json:"s3BucketName" tf:"s3_bucket_name,omitempty"`

	// The prefix for where Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.
	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type ReportPlanObservation struct {

	// The ARN of the backup report plan.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date and time that a report plan is created, in Unix format and Coordinated Universal Time (UTC).
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The deployment status of a report plan. The statuses are: CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED.
	DeploymentStatus *string `json:"deploymentStatus,omitempty" tf:"deployment_status,omitempty"`

	// The id of the backup report plan.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ReportPlanParameters struct {

	// The description of the report plan with a maximum of 1,024 characters
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// An object that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports. Detailed below.
	// +kubebuilder:validation:Required
	ReportDeliveryChannel []ReportDeliveryChannelParameters `json:"reportDeliveryChannel" tf:"report_delivery_channel,omitempty"`

	// An object that identifies the report template for the report. Reports are built using a report template. Detailed below.
	// +kubebuilder:validation:Required
	ReportSetting []ReportSettingParameters `json:"reportSetting" tf:"report_setting,omitempty"`

	// Metadata that you can assign to help organize the report plans you create. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ReportSettingObservation struct {
}

type ReportSettingParameters struct {

	// Specifies the Amazon Resource Names (ARNs) of the frameworks a report covers.
	// +kubebuilder:validation:Optional
	FrameworkArns []*string `json:"frameworkArns,omitempty" tf:"framework_arns,omitempty"`

	// Specifies the number of frameworks a report covers.
	// +kubebuilder:validation:Optional
	NumberOfFrameworks *float64 `json:"numberOfFrameworks,omitempty" tf:"number_of_frameworks,omitempty"`

	// Identifies the report template for the report. Reports are built using a report template. The report templates are: RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT.
	// +kubebuilder:validation:Required
	ReportTemplate *string `json:"reportTemplate" tf:"report_template,omitempty"`
}

// ReportPlanSpec defines the desired state of ReportPlan
type ReportPlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReportPlanParameters `json:"forProvider"`
}

// ReportPlanStatus defines the observed state of ReportPlan.
type ReportPlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReportPlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReportPlan is the Schema for the ReportPlans API. Provides an AWS Backup Report Plan resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReportPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReportPlanSpec   `json:"spec"`
	Status            ReportPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReportPlanList contains a list of ReportPlans
type ReportPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReportPlan `json:"items"`
}

// Repository type metadata.
var (
	ReportPlan_Kind             = "ReportPlan"
	ReportPlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReportPlan_Kind}.String()
	ReportPlan_KindAPIVersion   = ReportPlan_Kind + "." + CRDGroupVersion.String()
	ReportPlan_GroupVersionKind = CRDGroupVersion.WithKind(ReportPlan_Kind)
)

func init() {
	SchemeBuilder.Register(&ReportPlan{}, &ReportPlanList{})
}
