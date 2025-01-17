/*
Copyright 2024.

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
	"github.com/google/go-tpm-tools/launcher/spec"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConfidentialSpaceSpec defines the desired state of ConfidentialSpace
type ConfidentialSpaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	spec.LaunchSpec
}

// ConfidentialSpaceStatus defines the observed state of ConfidentialSpace
type ConfidentialSpaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// TODO(alexmwu): add more status fields
	// InstanceStatus is the GCE status of the CS instance.
	InstanceStatus        string
	InstanceLastStartTime string
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConfidentialSpace is the Schema for the confidentialspaces API
type ConfidentialSpace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfidentialSpaceSpec   `json:"spec,omitempty"`
	Status ConfidentialSpaceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConfidentialSpaceList contains a list of ConfidentialSpace
type ConfidentialSpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfidentialSpace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfidentialSpace{}, &ConfidentialSpaceList{})
}
