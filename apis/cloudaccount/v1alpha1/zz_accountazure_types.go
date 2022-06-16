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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccountAzureLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type AccountAzureLinksParameters struct {
}

type AccountAzureObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Links []AccountAzureLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type AccountAzureParameters struct {

	// Azure Client Application ID.
	// +kubebuilder:validation:Required
	ApplicationID *string `json:"applicationId" tf:"application_id,omitempty"`

	// Azure Client Application Secret Key.
	// +kubebuilder:validation:Required
	ApplicationKeySecretRef v1.SecretKeySelector `json:"applicationKeySecretRef" tf:"-"`

	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of this resource instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The set of region ids that will be enabled for this cloud account.
	// +kubebuilder:validation:Required
	Regions []*string `json:"regions" tf:"regions,omitempty"`

	// Azure Subscription ID.
	// +kubebuilder:validation:Required
	SubscriptionID *string `json:"subscriptionId" tf:"subscription_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []AccountAzureTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// Azure Tenant ID.
	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

type AccountAzureTagsObservation struct {
}

type AccountAzureTagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// AccountAzureSpec defines the desired state of AccountAzure
type AccountAzureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountAzureParameters `json:"forProvider"`
}

// AccountAzureStatus defines the observed state of AccountAzure.
type AccountAzureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountAzureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountAzure is the Schema for the AccountAzures API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vrajet}
type AccountAzure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountAzureSpec   `json:"spec"`
	Status            AccountAzureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountAzureList contains a list of AccountAzures
type AccountAzureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountAzure `json:"items"`
}

// Repository type metadata.
var (
	AccountAzure_Kind             = "AccountAzure"
	AccountAzure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountAzure_Kind}.String()
	AccountAzure_KindAPIVersion   = AccountAzure_Kind + "." + CRDGroupVersion.String()
	AccountAzure_GroupVersionKind = CRDGroupVersion.WithKind(AccountAzure_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountAzure{}, &AccountAzureList{})
}
