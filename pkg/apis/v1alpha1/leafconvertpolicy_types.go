package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope="Cluster"
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LeafPodConvertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec LeafPodConvertPolicySpec `json:"spec"`

	// +optional
	Status LeafPodConvertPolicyStatus `json:"status"`
}

type LeafPodConvertPolicySpec struct {
	// +optional
	Image string `json:"image,omitempty"`

	// +optional
	HostAliases []corev1.HostAlias `json:"hostAliases,omitempty"`
}

type LeafPodConvertPolicyStatus struct {
	// +optional
	CUEScript string `json:"cue,omitempty"`
}
