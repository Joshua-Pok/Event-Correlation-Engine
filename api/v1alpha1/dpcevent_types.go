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

type DPCSource struct { //needs to be exported
	Outcome string `json:"outcome,omitempty"`
	OCP     string `json:"ocp,omitempty"`
	IDRAC   string `json:"idrac,omitempty"`
}

// DpcEventSpec defines the desired state of DpcEvent
type DpcEventSpec struct {
	Source      *DPCSource `json:"soure,omitempty"`
	Component   string     `json:"component,omitempty"`
	ComponentID int32      `json:"componentid,omitempty"`
	Severity    string     `json:"severity,omitempty"`
}

// DpcEventStatus defines the observed state of DpcEvent.
type DpcEventStatus struct {
	Phase           string             `json:"phase,omitempty"`
	Reason          string             `json:"reason,omitempty"`
	AcknowledgeTime metav1.Time        `json:"acknowledgetime,omitempty"` //need to use metav1.Time for kubernetes controller
	Callhomed       bool               `json:"callhomed,omitempty"`
	RoutedTo        string             `json:"routedto,omitempty"`
	Conditions      []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DpcEvent is the Schema for the dpcevents API
type DpcEvent struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of DpcEvent
	// +required
	Spec DpcEventSpec `json:"spec"`

	// status defines the observed state of DpcEvent
	// +optional
	Status DpcEventStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// DpcEventList contains a list of DpcEvent
type DpcEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []DpcEvent `json:"items"`
}

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypes(SchemeGroupVersion, &DpcEvent{}, &DpcEventList{})
		return nil
	})
}
