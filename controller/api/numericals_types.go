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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NumericalsSpec defines the desired state of Numericals
type NumericalsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Expression is an math expression field of Numericals, which the user wants to solve  numericals_types.go to remove/update
    Expression string `json:"expression,omitempty"`
}

// NumericalsStatus defines the observed state of Numericals
type NumericalsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Answer is the solution to the expression
	Answer  string  'json: "answer"'
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Numericals is the Schema for the numericals API
type Numericals struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NumericalsSpec   `json:"spec,omitempty"`
	Status NumericalsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NumericalsList contains a list of Numericals
type NumericalsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Numericals `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Numericals{}, &NumericalsList{})
}
