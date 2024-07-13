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

type ClientDataInitParameters struct {

	// A user-defined comment about the disk upload.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The time that the disk upload ends.
	UploadEnd *string `json:"uploadEnd,omitempty" tf:"upload_end,omitempty"`

	// The size of the uploaded disk image, in GiB.
	UploadSize *float64 `json:"uploadSize,omitempty" tf:"upload_size,omitempty"`

	// The time that the disk upload starts.
	UploadStart *string `json:"uploadStart,omitempty" tf:"upload_start,omitempty"`
}

type ClientDataObservation struct {

	// A user-defined comment about the disk upload.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The time that the disk upload ends.
	UploadEnd *string `json:"uploadEnd,omitempty" tf:"upload_end,omitempty"`

	// The size of the uploaded disk image, in GiB.
	UploadSize *float64 `json:"uploadSize,omitempty" tf:"upload_size,omitempty"`

	// The time that the disk upload starts.
	UploadStart *string `json:"uploadStart,omitempty" tf:"upload_start,omitempty"`
}

type ClientDataParameters struct {

	// A user-defined comment about the disk upload.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The time that the disk upload ends.
	// +kubebuilder:validation:Optional
	UploadEnd *string `json:"uploadEnd,omitempty" tf:"upload_end,omitempty"`

	// The size of the uploaded disk image, in GiB.
	// +kubebuilder:validation:Optional
	UploadSize *float64 `json:"uploadSize,omitempty" tf:"upload_size,omitempty"`

	// The time that the disk upload starts.
	// +kubebuilder:validation:Optional
	UploadStart *string `json:"uploadStart,omitempty" tf:"upload_start,omitempty"`
}

type DiskContainerInitParameters struct {

	// The description of the disk image being imported.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The format of the disk image being imported. One of VHD or VMDK.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The URL to the Amazon S3-based disk image being imported. It can either be a https URL (https://..) or an Amazon S3 URL (s3://..). One of url or user_bucket must be set.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The Amazon S3 bucket for the disk image. One of url or user_bucket must be set. Detailed below.
	UserBucket *UserBucketInitParameters `json:"userBucket,omitempty" tf:"user_bucket,omitempty"`
}

type DiskContainerObservation struct {

	// The description of the disk image being imported.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The format of the disk image being imported. One of VHD or VMDK.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The URL to the Amazon S3-based disk image being imported. It can either be a https URL (https://..) or an Amazon S3 URL (s3://..). One of url or user_bucket must be set.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The Amazon S3 bucket for the disk image. One of url or user_bucket must be set. Detailed below.
	UserBucket *UserBucketObservation `json:"userBucket,omitempty" tf:"user_bucket,omitempty"`
}

type DiskContainerParameters struct {

	// The description of the disk image being imported.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The format of the disk image being imported. One of VHD or VMDK.
	// +kubebuilder:validation:Optional
	Format *string `json:"format" tf:"format,omitempty"`

	// The URL to the Amazon S3-based disk image being imported. It can either be a https URL (https://..) or an Amazon S3 URL (s3://..). One of url or user_bucket must be set.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The Amazon S3 bucket for the disk image. One of url or user_bucket must be set. Detailed below.
	// +kubebuilder:validation:Optional
	UserBucket *UserBucketParameters `json:"userBucket,omitempty" tf:"user_bucket,omitempty"`
}

type EBSSnapshotImportInitParameters struct {

	// The client-specific data. Detailed below.
	ClientData *ClientDataInitParameters `json:"clientData,omitempty" tf:"client_data,omitempty"`

	// The description string for the import snapshot task.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Information about the disk container. Detailed below.
	DiskContainer *DiskContainerInitParameters `json:"diskContainer,omitempty" tf:"disk_container,omitempty"`

	// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore *bool `json:"permanentRestore,omitempty" tf:"permanent_restore,omitempty"`

	// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: vmimport
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// The name of the storage tier. Valid values are archive and standard. Default value is standard.
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays *float64 `json:"temporaryRestoreDays,omitempty" tf:"temporary_restore_days,omitempty"`
}

type EBSSnapshotImportObservation struct {

	// Amazon Resource Name (ARN) of the EBS Snapshot.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The client-specific data. Detailed below.
	ClientData *ClientDataObservation `json:"clientData,omitempty" tf:"client_data,omitempty"`

	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyID *string `json:"dataEncryptionKeyId,omitempty" tf:"data_encryption_key_id,omitempty"`

	// The description string for the import snapshot task.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Information about the disk container. Detailed below.
	DiskContainer *DiskContainerObservation `json:"diskContainer,omitempty" tf:"disk_container,omitempty"`

	// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The snapshot ID (e.g., snap-59fcb34e).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Amazon Resource Name (ARN) of the EBS Snapshot.
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// Value from an Amazon-maintained list (amazon, aws-marketplace, microsoft) of snapshot owners.
	OwnerAlias *string `json:"ownerAlias,omitempty" tf:"owner_alias,omitempty"`

	// The AWS account ID of the EBS snapshot owner.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore *bool `json:"permanentRestore,omitempty" tf:"permanent_restore,omitempty"`

	// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: vmimport
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// The name of the storage tier. Valid values are archive and standard. Default value is standard.
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays *float64 `json:"temporaryRestoreDays,omitempty" tf:"temporary_restore_days,omitempty"`

	// The snapshot ID (e.g., snap-59fcb34e).
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// The size of the drive in GiBs.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
}

type EBSSnapshotImportParameters struct {

	// The client-specific data. Detailed below.
	// +kubebuilder:validation:Optional
	ClientData *ClientDataParameters `json:"clientData,omitempty" tf:"client_data,omitempty"`

	// The description string for the import snapshot task.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Information about the disk container. Detailed below.
	// +kubebuilder:validation:Optional
	DiskContainer *DiskContainerParameters `json:"diskContainer,omitempty" tf:"disk_container,omitempty"`

	// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Indicates whether to permanently restore an archived snapshot.
	// +kubebuilder:validation:Optional
	PermanentRestore *bool `json:"permanentRestore,omitempty" tf:"permanent_restore,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: vmimport
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// The name of the storage tier. Valid values are archive and standard. Default value is standard.
	// +kubebuilder:validation:Optional
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	// +kubebuilder:validation:Optional
	TemporaryRestoreDays *float64 `json:"temporaryRestoreDays,omitempty" tf:"temporary_restore_days,omitempty"`
}

type UserBucketInitParameters struct {

	// The name of the Amazon S3 bucket where the disk image is located.
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// The file name of the disk image.
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`
}

type UserBucketObservation struct {

	// The name of the Amazon S3 bucket where the disk image is located.
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// The file name of the disk image.
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`
}

type UserBucketParameters struct {

	// The name of the Amazon S3 bucket where the disk image is located.
	// +kubebuilder:validation:Optional
	S3Bucket *string `json:"s3Bucket" tf:"s3_bucket,omitempty"`

	// The file name of the disk image.
	// +kubebuilder:validation:Optional
	S3Key *string `json:"s3Key" tf:"s3_key,omitempty"`
}

// EBSSnapshotImportSpec defines the desired state of EBSSnapshotImport
type EBSSnapshotImportSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EBSSnapshotImportParameters `json:"forProvider"`
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
	InitProvider EBSSnapshotImportInitParameters `json:"initProvider,omitempty"`
}

// EBSSnapshotImportStatus defines the observed state of EBSSnapshotImport.
type EBSSnapshotImportStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EBSSnapshotImportObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// EBSSnapshotImport is the Schema for the EBSSnapshotImports API. Provides an elastic block storage snapshot import resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EBSSnapshotImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.diskContainer) || (has(self.initProvider) && has(self.initProvider.diskContainer))",message="spec.forProvider.diskContainer is a required parameter"
	Spec   EBSSnapshotImportSpec   `json:"spec"`
	Status EBSSnapshotImportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EBSSnapshotImportList contains a list of EBSSnapshotImports
type EBSSnapshotImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EBSSnapshotImport `json:"items"`
}

// Repository type metadata.
var (
	EBSSnapshotImport_Kind             = "EBSSnapshotImport"
	EBSSnapshotImport_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EBSSnapshotImport_Kind}.String()
	EBSSnapshotImport_KindAPIVersion   = EBSSnapshotImport_Kind + "." + CRDGroupVersion.String()
	EBSSnapshotImport_GroupVersionKind = CRDGroupVersion.WithKind(EBSSnapshotImport_Kind)
)

func init() {
	SchemeBuilder.Register(&EBSSnapshotImport{}, &EBSSnapshotImportList{})
}