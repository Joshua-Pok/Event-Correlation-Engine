/*
Copyright 2026.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type EventReference struct {
	Name      string      `json:"name"`
	EventType string      `json:"eventType"`
	Timestamp metav1.Time `json:"timestamp"`
}

// DpcEventGroupSpec defines the desired state of DpcEventGroup
type DpcEventGroupSpec struct {
	Node          string `json:"node,omitempty"`
	WindowSeconds int    `json:"windowSeconds"`
}

// DpcEventGroupStatus defines the observed state of DpcEventGroup.
type DpcEventGroupStatus struct {
	Phase      string           `json:"phase"`
	EventCount int              `json:"eventCount"`
	Events     []EventReference `json:"events,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DpcEventGroup is the Schema for the dpceventgroups API
type DpcEventGroup struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of DpcEventGroup
	// +required
	Spec DpcEventGroupSpec `json:"spec"`

	// status defines the observed state of DpcEventGroup
	// +optional
	Status DpcEventGroupStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// DpcEventGroupList contains a list of DpcEventGroup
type DpcEventGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []DpcEventGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypes(SchemeGroupVersion, &DpcEventGroup{}, &DpcEventGroupList{})
		return nil
	})
}
