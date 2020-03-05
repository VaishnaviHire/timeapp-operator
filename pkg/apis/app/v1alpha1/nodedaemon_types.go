package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodeDaemonSpec defines the desired state of NodeDaemon
type NodeDaemonSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	Label string `json:"label"`
	Image string `json:"image"`
	//image string
}

// NodeDaemonStatus defines the observed state of NodeDaemon
type NodeDaemonStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	Count int32 `json:"count"`

}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeDaemon is the Schema for the nodedaemons API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nodedaemons,scope=Namespaced
type NodeDaemon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeDaemonSpec   `json:"spec,omitempty"`
	Status NodeDaemonStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeDaemonList contains a list of NodeDaemon
type NodeDaemonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeDaemon `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeDaemon{}, &NodeDaemonList{})
}
