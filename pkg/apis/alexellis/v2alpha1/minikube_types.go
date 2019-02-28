package v2alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MinikubeSpec defines the desired state of Minikube
// +k8s:openapi-gen=true
type MinikubeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	MemoryMB    int    `json:"memoryMB"`
	CPUCount    int    `json:"cpuCount"`
	ClusterName string `json:"clusterName"`
}

// MinikubeStatus defines the observed state of Minikube
// +k8s:openapi-gen=true
type MinikubeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Minikube is the Schema for the minikubes API
// +k8s:openapi-gen=true
type Minikube struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MinikubeSpec   `json:"spec,omitempty"`
	Status MinikubeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MinikubeList contains a list of Minikube
type MinikubeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Minikube `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Minikube{}, &MinikubeList{})
}
