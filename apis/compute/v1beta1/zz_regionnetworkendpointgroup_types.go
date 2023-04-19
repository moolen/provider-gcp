/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AppEngineObservation struct {

	// Optional serving service.
	// The service name must be 1-63 characters long, and comply with RFC1035.
	// Example value: "default", "my-service".
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// A template to parse service and version fields from a request URL.
	// URL mask allows for routing to multiple App Engine services without
	// having to create multiple Network Endpoint Groups and backend services.
	// For example, the request URLs "foo1-dot-appname.appspot.com/v1" and
	// "foo1-dot-appname.appspot.com/v2" can be backed by the same Serverless NEG with
	// URL mask "-dot-appname.appspot.com/". The URL mask will parse
	// them to { service = "foo1", version = "v1" } and { service = "foo1", version = "v2" } respectively.
	URLMask *string `json:"urlMask,omitempty" tf:"url_mask,omitempty"`

	// Optional serving version.
	// The version must be 1-63 characters long, and comply with RFC1035.
	// Example value: "v1", "v2".
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AppEngineParameters struct {

	// Optional serving service.
	// The service name must be 1-63 characters long, and comply with RFC1035.
	// Example value: "default", "my-service".
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// A template to parse service and version fields from a request URL.
	// URL mask allows for routing to multiple App Engine services without
	// having to create multiple Network Endpoint Groups and backend services.
	// For example, the request URLs "foo1-dot-appname.appspot.com/v1" and
	// "foo1-dot-appname.appspot.com/v2" can be backed by the same Serverless NEG with
	// URL mask "-dot-appname.appspot.com/". The URL mask will parse
	// them to { service = "foo1", version = "v1" } and { service = "foo1", version = "v2" } respectively.
	// +kubebuilder:validation:Optional
	URLMask *string `json:"urlMask,omitempty" tf:"url_mask,omitempty"`

	// Optional serving version.
	// The version must be 1-63 characters long, and comply with RFC1035.
	// Example value: "v1", "v2".
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type CloudFunctionObservation struct {

	// A user-defined name of the Cloud Function.
	// The function name is case-sensitive and must be 1-63 characters long.
	// Example value: "func1".
	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	// A template to parse function field from a request URL. URL mask allows
	// for routing to multiple Cloud Functions without having to create
	// multiple Network Endpoint Groups and backend services.
	// For example, request URLs "mydomain.com/function1" and "mydomain.com/function2"
	// can be backed by the same Serverless NEG with URL mask "/". The URL mask
	// will parse them to { function = "function1" } and { function = "function2" } respectively.
	URLMask *string `json:"urlMask,omitempty" tf:"url_mask,omitempty"`
}

type CloudFunctionParameters struct {

	// A user-defined name of the Cloud Function.
	// The function name is case-sensitive and must be 1-63 characters long.
	// Example value: "func1".
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudfunctions/v1beta1.Function
	// +kubebuilder:validation:Optional
	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	// Reference to a Function in cloudfunctions to populate function.
	// +kubebuilder:validation:Optional
	FunctionRef *v1.Reference `json:"functionRef,omitempty" tf:"-"`

	// Selector for a Function in cloudfunctions to populate function.
	// +kubebuilder:validation:Optional
	FunctionSelector *v1.Selector `json:"functionSelector,omitempty" tf:"-"`

	// A template to parse function field from a request URL. URL mask allows
	// for routing to multiple Cloud Functions without having to create
	// multiple Network Endpoint Groups and backend services.
	// For example, request URLs "mydomain.com/function1" and "mydomain.com/function2"
	// can be backed by the same Serverless NEG with URL mask "/". The URL mask
	// will parse them to { function = "function1" } and { function = "function2" } respectively.
	// +kubebuilder:validation:Optional
	URLMask *string `json:"urlMask,omitempty" tf:"url_mask,omitempty"`
}

type CloudRunObservation struct {

	// Cloud Run service is the main resource of Cloud Run.
	// The service must be 1-63 characters long, and comply with RFC1035.
	// Example value: "run-service".
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// Cloud Run tag represents the "named-revision" to provide
	// additional fine-grained traffic routing information.
	// The tag must be 1-63 characters long, and comply with RFC1035.
	// Example value: "revision-0010".
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// A template to parse service and tag fields from a request URL.
	// URL mask allows for routing to multiple Run services without having
	// to create multiple network endpoint groups and backend services.
	// For example, request URLs "foo1.domain.com/bar1" and "foo1.domain.com/bar2"
	// an be backed by the same Serverless Network Endpoint Group (NEG) with
	// URL mask ".domain.com/". The URL mask will parse them to { service="bar1", tag="foo1" }
	// and { service="bar2", tag="foo2" } respectively.
	URLMask *string `json:"urlMask,omitempty" tf:"url_mask,omitempty"`
}

type CloudRunParameters struct {

	// Cloud Run service is the main resource of Cloud Run.
	// The service must be 1-63 characters long, and comply with RFC1035.
	// Example value: "run-service".
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudrun/v1beta1.Service
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// Reference to a Service in cloudrun to populate service.
	// +kubebuilder:validation:Optional
	ServiceRef *v1.Reference `json:"serviceRef,omitempty" tf:"-"`

	// Selector for a Service in cloudrun to populate service.
	// +kubebuilder:validation:Optional
	ServiceSelector *v1.Selector `json:"serviceSelector,omitempty" tf:"-"`

	// Cloud Run tag represents the "named-revision" to provide
	// additional fine-grained traffic routing information.
	// The tag must be 1-63 characters long, and comply with RFC1035.
	// Example value: "revision-0010".
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// A template to parse service and tag fields from a request URL.
	// URL mask allows for routing to multiple Run services without having
	// to create multiple network endpoint groups and backend services.
	// For example, request URLs "foo1.domain.com/bar1" and "foo1.domain.com/bar2"
	// an be backed by the same Serverless Network Endpoint Group (NEG) with
	// URL mask ".domain.com/". The URL mask will parse them to { service="bar1", tag="foo1" }
	// and { service="bar2", tag="foo2" } respectively.
	// +kubebuilder:validation:Optional
	URLMask *string `json:"urlMask,omitempty" tf:"url_mask,omitempty"`
}

type RegionNetworkEndpointGroupObservation struct {

	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	// Structure is documented below.
	AppEngine []AppEngineObservation `json:"appEngine,omitempty" tf:"app_engine,omitempty"`

	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	// Structure is documented below.
	CloudFunction []CloudFunctionObservation `json:"cloudFunction,omitempty" tf:"cloud_function,omitempty"`

	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	// Structure is documented below.
	CloudRun []CloudRunObservation `json:"cloudRun,omitempty" tf:"cloud_run,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// This field is only used for PSC.
	// The URL of the network to which all network endpoints in the NEG belong. Uses
	// "default" project network if unspecified.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
	// Default value is SERVERLESS.
	// Possible values are SERVERLESS and PRIVATE_SERVICE_CONNECT.
	NetworkEndpointType *string `json:"networkEndpointType,omitempty" tf:"network_endpoint_type,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The target service url used to set up private service connection to
	// a Google API or a PSC Producer Service Attachment.
	PscTargetService *string `json:"pscTargetService,omitempty" tf:"psc_target_service,omitempty"`

	// A reference to the region where the Serverless NEGs Reside.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// This field is only used for PSC.
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
}

type RegionNetworkEndpointGroupParameters struct {

	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AppEngine []AppEngineParameters `json:"appEngine,omitempty" tf:"app_engine,omitempty"`

	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CloudFunction []CloudFunctionParameters `json:"cloudFunction,omitempty" tf:"cloud_function,omitempty"`

	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CloudRun []CloudRunParameters `json:"cloudRun,omitempty" tf:"cloud_run,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// This field is only used for PSC.
	// The URL of the network to which all network endpoints in the NEG belong. Uses
	// "default" project network if unspecified.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
	// Default value is SERVERLESS.
	// Possible values are SERVERLESS and PRIVATE_SERVICE_CONNECT.
	// +kubebuilder:validation:Optional
	NetworkEndpointType *string `json:"networkEndpointType,omitempty" tf:"network_endpoint_type,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The target service url used to set up private service connection to
	// a Google API or a PSC Producer Service Attachment.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.ServiceAttachment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	PscTargetService *string `json:"pscTargetService,omitempty" tf:"psc_target_service,omitempty"`

	// Reference to a ServiceAttachment in compute to populate pscTargetService.
	// +kubebuilder:validation:Optional
	PscTargetServiceRef *v1.Reference `json:"pscTargetServiceRef,omitempty" tf:"-"`

	// Selector for a ServiceAttachment in compute to populate pscTargetService.
	// +kubebuilder:validation:Optional
	PscTargetServiceSelector *v1.Selector `json:"pscTargetServiceSelector,omitempty" tf:"-"`

	// A reference to the region where the Serverless NEGs Reside.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// This field is only used for PSC.
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Reference to a Subnetwork in compute to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork in compute to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

// RegionNetworkEndpointGroupSpec defines the desired state of RegionNetworkEndpointGroup
type RegionNetworkEndpointGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionNetworkEndpointGroupParameters `json:"forProvider"`
}

// RegionNetworkEndpointGroupStatus defines the observed state of RegionNetworkEndpointGroup.
type RegionNetworkEndpointGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionNetworkEndpointGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionNetworkEndpointGroup is the Schema for the RegionNetworkEndpointGroups API. A regional NEG that can support Serverless Products.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegionNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionNetworkEndpointGroupSpec   `json:"spec"`
	Status            RegionNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionNetworkEndpointGroupList contains a list of RegionNetworkEndpointGroups
type RegionNetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionNetworkEndpointGroup `json:"items"`
}

// Repository type metadata.
var (
	RegionNetworkEndpointGroup_Kind             = "RegionNetworkEndpointGroup"
	RegionNetworkEndpointGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionNetworkEndpointGroup_Kind}.String()
	RegionNetworkEndpointGroup_KindAPIVersion   = RegionNetworkEndpointGroup_Kind + "." + CRDGroupVersion.String()
	RegionNetworkEndpointGroup_GroupVersionKind = CRDGroupVersion.WithKind(RegionNetworkEndpointGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionNetworkEndpointGroup{}, &RegionNetworkEndpointGroupList{})
}
