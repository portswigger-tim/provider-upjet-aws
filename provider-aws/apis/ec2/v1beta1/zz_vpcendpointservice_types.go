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

type PrivateDNSNameConfigurationObservation struct {

	// Name of the record subdomain the service provider needs to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Endpoint service verification type, for example TXT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Value the service provider adds to the private DNS name domain record before verification.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PrivateDNSNameConfigurationParameters struct {
}

type VPCEndpointServiceObservation struct {

	// The ARNs of one or more principals allowed to discover the endpoint service.
	AllowedPrincipals []*string `json:"allowedPrincipals,omitempty" tf:"allowed_principals,omitempty"`

	// The Amazon Resource Name (ARN) of the VPC endpoint service.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A set of Availability Zones in which the service is available.
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// A set of DNS names for the service.
	BaseEndpointDNSNames []*string `json:"baseEndpointDnsNames,omitempty" tf:"base_endpoint_dns_names,omitempty"`

	// The ID of the VPC endpoint service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether or not the service manages its VPC endpoints - true or false.
	ManagesVPCEndpoints *bool `json:"managesVpcEndpoints,omitempty" tf:"manages_vpc_endpoints,omitempty"`

	// List of objects containing information about the endpoint service private DNS name configuration.
	PrivateDNSNameConfiguration []PrivateDNSNameConfigurationObservation `json:"privateDnsNameConfiguration,omitempty" tf:"private_dns_name_configuration,omitempty"`

	// The service name.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The service type, Gateway or Interface.
	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCEndpointServiceParameters struct {

	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - true or false.
	// +kubebuilder:validation:Required
	AcceptanceRequired *bool `json:"acceptanceRequired" tf:"acceptance_required,omitempty"`

	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	// +kubebuilder:validation:Optional
	GatewayLoadBalancerArns []*string `json:"gatewayLoadBalancerArns,omitempty" tf:"gateway_load_balancer_arns,omitempty"`

	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerArns []*string `json:"networkLoadBalancerArns,omitempty" tf:"network_load_balancer_arns,omitempty"`

	// The private DNS name for the service.
	// +kubebuilder:validation:Optional
	PrivateDNSName *string `json:"privateDnsName,omitempty" tf:"private_dns_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCEndpointServiceSpec defines the desired state of VPCEndpointService
type VPCEndpointServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointServiceParameters `json:"forProvider"`
}

// VPCEndpointServiceStatus defines the observed state of VPCEndpointService.
type VPCEndpointServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointService is the Schema for the VPCEndpointServices API. Provides a VPC Endpoint Service resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpointService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointServiceSpec   `json:"spec"`
	Status            VPCEndpointServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointServiceList contains a list of VPCEndpointServices
type VPCEndpointServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointService `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointService_Kind             = "VPCEndpointService"
	VPCEndpointService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpointService_Kind}.String()
	VPCEndpointService_KindAPIVersion   = VPCEndpointService_Kind + "." + CRDGroupVersion.String()
	VPCEndpointService_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpointService_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointService{}, &VPCEndpointServiceList{})
}
