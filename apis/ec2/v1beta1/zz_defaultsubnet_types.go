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

type DefaultSubnetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// , cidr_block and vpc_id arguments become computed attributes
	AvailabilityZoneID *string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id,omitempty"`

	// The IPv4 CIDR block assigned to the subnet
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	ExistingDefaultSubnet *bool `json:"existingDefaultSubnet,omitempty" tf:"existing_default_subnet,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPv6CidrBlockAssociationID *string `json:"ipv6CidrBlockAssociationId,omitempty" tf:"ipv6_cidr_block_association_id,omitempty"`

	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the VPC the subnet is in
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DefaultSubnetParameters struct {

	// +kubebuilder:validation:Optional
	AssignIPv6AddressOnCreation *bool `json:"assignIpv6AddressOnCreation,omitempty" tf:"assign_ipv6_address_on_creation,omitempty"`

	// is required
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// +kubebuilder:validation:Optional
	EnableDns64 *bool `json:"enableDns64,omitempty" tf:"enable_dns64,omitempty"`

	// +kubebuilder:validation:Optional
	EnableResourceNameDNSARecordOnLaunch *bool `json:"enableResourceNameDnsARecordOnLaunch,omitempty" tf:"enable_resource_name_dns_a_record_on_launch,omitempty"`

	// +kubebuilder:validation:Optional
	EnableResourceNameDNSAaaaRecordOnLaunch *bool `json:"enableResourceNameDnsAaaaRecordOnLaunch,omitempty" tf:"enable_resource_name_dns_aaaa_record_on_launch,omitempty"`

	// Whether destroying the resource deletes the default subnet. Default: false
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The IPv4 CIDR block assigned to the subnet
	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Native *bool `json:"ipv6Native,omitempty" tf:"ipv6_native,omitempty"`

	// +kubebuilder:validation:Optional
	MapCustomerOwnedIPOnLaunch *bool `json:"mapCustomerOwnedIpOnLaunch,omitempty" tf:"map_customer_owned_ip_on_launch,omitempty"`

	// is true
	// +kubebuilder:validation:Optional
	MapPublicIPOnLaunch *bool `json:"mapPublicIpOnLaunch,omitempty" tf:"map_public_ip_on_launch,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateDNSHostnameTypeOnLaunch *string `json:"privateDnsHostnameTypeOnLaunch,omitempty" tf:"private_dns_hostname_type_on_launch,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DefaultSubnetSpec defines the desired state of DefaultSubnet
type DefaultSubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultSubnetParameters `json:"forProvider"`
}

// DefaultSubnetStatus defines the observed state of DefaultSubnet.
type DefaultSubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultSubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSubnet is the Schema for the DefaultSubnets API. Manage a default subnet resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DefaultSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSubnetSpec   `json:"spec"`
	Status            DefaultSubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSubnetList contains a list of DefaultSubnets
type DefaultSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSubnet `json:"items"`
}

// Repository type metadata.
var (
	DefaultSubnet_Kind             = "DefaultSubnet"
	DefaultSubnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultSubnet_Kind}.String()
	DefaultSubnet_KindAPIVersion   = DefaultSubnet_Kind + "." + CRDGroupVersion.String()
	DefaultSubnet_GroupVersionKind = CRDGroupVersion.WithKind(DefaultSubnet_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultSubnet{}, &DefaultSubnetList{})
}