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

type RevocationRecordObservation struct {
	RevocationEffectiveFrom *string `json:"revocationEffectiveFrom,omitempty" tf:"revocation_effective_from,omitempty"`

	RevokedAt *string `json:"revokedAt,omitempty" tf:"revoked_at,omitempty"`

	RevokedBy *string `json:"revokedBy,omitempty" tf:"revoked_by,omitempty"`
}

type RevocationRecordParameters struct {
}

type SignatureValidityPeriodObservation struct {
}

type SignatureValidityPeriodParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type SigningProfileObservation struct {

	// The Amazon Resource Name (ARN) for the signing profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A human-readable name for the signing platform associated with the signing profile.
	PlatformDisplayName *string `json:"platformDisplayName,omitempty" tf:"platform_display_name,omitempty"`

	// Revocation information for a signing profile.
	RevocationRecord []RevocationRecordObservation `json:"revocationRecord,omitempty" tf:"revocation_record,omitempty"`

	// The status of the target signing profile.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The current version of the signing profile.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The signing profile ARN, including the profile version.
	VersionArn *string `json:"versionArn,omitempty" tf:"version_arn,omitempty"`
}

type SigningProfileParameters struct {

	// The ID of the platform that is used by the target signing profile.
	// +kubebuilder:validation:Required
	PlatformID *string `json:"platformId" tf:"platform_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The validity period for a signing job.
	// +kubebuilder:validation:Optional
	SignatureValidityPeriod []SignatureValidityPeriodParameters `json:"signatureValidityPeriod,omitempty" tf:"signature_validity_period,omitempty"`

	// A list of tags associated with the signing profile. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SigningProfileSpec defines the desired state of SigningProfile
type SigningProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SigningProfileParameters `json:"forProvider"`
}

// SigningProfileStatus defines the observed state of SigningProfile.
type SigningProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SigningProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SigningProfile is the Schema for the SigningProfiles API. Creates a Signer Signing Profile.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SigningProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SigningProfileSpec   `json:"spec"`
	Status            SigningProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SigningProfileList contains a list of SigningProfiles
type SigningProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SigningProfile `json:"items"`
}

// Repository type metadata.
var (
	SigningProfile_Kind             = "SigningProfile"
	SigningProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SigningProfile_Kind}.String()
	SigningProfile_KindAPIVersion   = SigningProfile_Kind + "." + CRDGroupVersion.String()
	SigningProfile_GroupVersionKind = CRDGroupVersion.WithKind(SigningProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&SigningProfile{}, &SigningProfileList{})
}
