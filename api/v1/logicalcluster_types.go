/*
Copyright 2023 Nguyen Thanh Nguyen.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Cluster Member State contains state of cluster member
type ClusterMemberStatus struct {
	// Cluster Name
	ClusterName string `json:"clustername,omitempty"`
	// Status of Cluster
	Ready bool `json:"ready,omitempty"`
	// Failure Message
	// +optional
	FailureMessage string `json:"failureMessage,omitempty"`
	// Failure Reason
	// +optional
	FailureReason string `json:"failureReason,omitempty"`
	// Logical Cluster Conditions
	// +optional
	Conditions ConditionType `json:"conditions,omitempty"`
	// Registration Status of Member Cluster
	Registration bool `json:"registration,omitempty"`
}

// LogicalClusterSpec defines the desired state of LogicalCluster
type LogicalClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of LogicalCluster. Edit logicalcluster_types.go to remove/update
	Clusters []ClusterMember `json:"clusters,omitempty"`
}
type ClusterMemberSpec struct {
	// Cluster Catalog
	//+optional
	ClusterCatalog string `json:"clustercatalog,omitempty"`
	// Cluster Detail Spec
	//+optional
	ClusterSpec ClusterSpec `json:"spec,omitempty"`
}
type ClusterMember struct {
	Name string `json:"name,omitempty"`
	// ClusterCatalog    string `json:"clustercatalog,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Cluster member of logical cluster. Each cluster member associate with a physical cluster (CAPI)
	//+optional
	ClusterMemberSpec ClusterMemberSpec `json:"spec,omitempty"`
	// ClusterRef is a reference to a L-KaaS cluster that holds the details cluster
	// +optional
	ClusterRef *corev1.ObjectReference `json:"clusterref,omitempty"`
}

// LogicalClusterStatus defines the observed state of LogicalCluster
type LogicalClusterStatus struct {

	// Ready state of Logical cluster
	Ready bool `json:"ready,omitempty"`
	// Status of cluster
	Status string `json:"status,omitempty"`
	// State of Each Cluster Member
	ClusterMemberStates []ClusterMemberStatus `json:"clusterMemberStates,omitempty"`
	// Failure Message
	// +optional
	FailureMessage string `json:"failureMessage,omitempty"`
	// Failure Reason
	// +optional
	FailureReason string `json:"failureReason,omitempty"`
	// Logical Cluster Conditions
	// +optional
	Conditions ConditionType `json:"conditions,omitempty"`
	// InfrastructureReady is the state of the infrastructure provider.
	// +optional
	InfrastructureReady bool `json:"infrastructureReady,omitempty"`
	// Logical Cluster Phase
	// +optional
	Phase ClusterPhase `json:"phase,omitempty"`
	// Registration Status of Logical Cluster
	Registration bool `json:"registration,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="Time duration since creation of Cluster"
// +kubebuilder:printcolumn:name="InfrastructureReady",type="string",JSONPath=".status.infrastructureReady",description="infrastructureReady"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.ready",description="Ready"
// +kubebuilder:printcolumn:name="FailureReason",type="string",JSONPath=".status.failureReason",description="failureReason"
// +kubebuilder:printcolumn:name="FailureMessage",type="string",JSONPath=".status.failureMessage",description="failureMessage"
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase",description="Phase"
// +kubebuilder:printcolumn:name="Registration",type="string",JSONPath=".status.registration",description="Registration"

// LogicalCluster is the Schema for the logicalclusters API
type LogicalCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LogicalClusterSpec   `json:"spec,omitempty"`
	Status LogicalClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LogicalClusterList contains a list of LogicalCluster
type LogicalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogicalCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LogicalCluster{}, &LogicalClusterList{})
}

// SetTypedPhase sets the Phase field to the string representation of ClusterPhase.
func (c *LogicalClusterStatus) SetTypedPhase(p ClusterPhase) {
	c.Phase = p
}
