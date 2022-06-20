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

type DefinitionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IconID *string `json:"iconId,omitempty" tf:"icon_id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NumberOfItems *float64 `json:"numberOfItems,omitempty" tf:"number_of_items,omitempty"`

	SourceName *string `json:"sourceName,omitempty" tf:"source_name,omitempty"`

	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DefinitionParameters struct {
}

type ItemEntitlementObservation struct {
	Definition []DefinitionObservation `json:"definition,omitempty" tf:"definition,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ItemEntitlementParameters struct {

	// Catalog item id.
	// +kubebuilder:validation:Required
	CatalogItemID *string `json:"catalogItemId" tf:"catalog_item_id,omitempty"`

	// Project id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-vra/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

// ItemEntitlementSpec defines the desired state of ItemEntitlement
type ItemEntitlementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ItemEntitlementParameters `json:"forProvider"`
}

// ItemEntitlementStatus defines the observed state of ItemEntitlement.
type ItemEntitlementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ItemEntitlementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ItemEntitlement is the Schema for the ItemEntitlements API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vrajet}
type ItemEntitlement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ItemEntitlementSpec   `json:"spec"`
	Status            ItemEntitlementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ItemEntitlementList contains a list of ItemEntitlements
type ItemEntitlementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ItemEntitlement `json:"items"`
}

// Repository type metadata.
var (
	ItemEntitlement_Kind             = "ItemEntitlement"
	ItemEntitlement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ItemEntitlement_Kind}.String()
	ItemEntitlement_KindAPIVersion   = ItemEntitlement_Kind + "." + CRDGroupVersion.String()
	ItemEntitlement_GroupVersionKind = CRDGroupVersion.WithKind(ItemEntitlement_Kind)
)

func init() {
	SchemeBuilder.Register(&ItemEntitlement{}, &ItemEntitlementList{})
}
