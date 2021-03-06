/*
Copyright 2020 AccuKnox.

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

type SelectorType struct {
	MatchNames  map[string]string `json:"matchNames,omitempty"`
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// +kubebuilder:validation:Pattern=^\/([A-z0-9-_.]+\/)*([A-z0-9-_.]+)$
type MatchPathType string

// +kubebuilder:validation:Pattern=^\/([A-z0-9-_.]+\/)*([A-z0-9-_.]+)+\/$
type MatchDirectoryType string

type MatchSourceType struct {
	Path      MatchPathType      `json:"path,omitempty"`
	Directory MatchDirectoryType `json:"dir,omitempty"`
}

type ProcessPathType struct {
	Path MatchPathType `json:"path"`

	// +kubebuilder:validation:Optional
	OwnerOnly bool `json:"ownerOnly,omitempty"`

	// +kubebuilder:validation:optional
	FromSource []MatchSourceType `json:"fromSource,omitempty"`
}

type ProcessDirectoryType struct {
	Directory MatchDirectoryType `json:"dir"`

	// +kubebuilder:validation:Optional
	Recursive bool `json:"recursive,omitempty"`
	// +kubebuilder:validation:Optional
	OwnerOnly bool `json:"ownerOnly,omitempty"`

	// +kubebuilder:validation:optional
	FromSource []MatchSourceType `json:"fromSource,omitempty"`
}

type ProcessPatternType struct {
	Pattern string `json:"pattern"`

	// +kubebuilder:validation:Optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// +kubebuilder:validation:Optional
	OwnerOnly bool `json:"ownerOnly,omitempty"`

	// +kubebuilder:validation:optional
	FromSource []MatchSourceType `json:"fromSource,omitempty"`
}

type ProcessType struct {
	MatchPaths       []ProcessPathType      `json:"matchPaths,omitempty"`
	MatchDirectories []ProcessDirectoryType `json:"matchDirectories,omitempty"`
	MatchPatterns    []ProcessPatternType   `json:"matchPatterns,omitempty"`
}

type FilePathType struct {
	Path MatchPathType `json:"path"`

	// +kubebuilder:validation:Optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// +kubebuilder:validation:Optional
	OwnerOnly bool `json:"ownerOnly,omitempty"`

	// +kubebuilder:validation:optional
	FromSource []MatchSourceType `json:"fromSource,omitempty"`
}

type FileDirectoryType struct {
	Directory MatchDirectoryType `json:"dir"`

	// +kubebuilder:validation:Optional
	Recursive bool `json:"recursive,omitempty"`
	// +kubebuilder:validation:Optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// +kubebuilder:validation:Optional
	OwnerOnly bool `json:"ownerOnly,omitempty"`

	// +kubebuilder:validation:optional
	FromSource []MatchSourceType `json:"fromSource,omitempty"`
}

type FilePatternType struct {
	Pattern string `json:"pattern"`

	// +kubebuilder:validation:Optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// +kubebuilder:validation:Optional
	OwnerOnly bool `json:"ownerOnly,omitempty"`

	// +kubebuilder:validation:optional
	FromSource []MatchSourceType `json:"fromSource,omitempty"`
}

type FileType struct {
	MatchPaths       []FilePathType      `json:"matchPaths,omitempty"`
	MatchDirectories []FileDirectoryType `json:"matchDirectories,omitempty"`
	MatchPatterns    []FilePatternType   `json:"matchPatterns,omitempty"`
}

// +kubebuilder:validation:Enum=TCP;tcp;UDP;udp;ICMP;icmp
type MatchProtocolType string

type NetworkProtocolType struct {
	Protocol MatchProtocolType `json:"protocol"`
	IPv4     bool              `json:"ipv4,omitempty"`
	IPv6     bool              `json:"ipv6,omitempty"`
}

type NetworkProcFileType struct {
	MatchPaths       []string `json:"matchPaths,omitempty"`
	MatchDirectories []string `json:"matchDirectories,omitempty"`
}

type NetworkSourceType struct {
	Process NetworkProcFileType `json:"process,omitempty"`
	File    NetworkProcFileType `json:"file,omitempty"`
}

type NetworkType struct {
	MatchProtocols []NetworkProtocolType `json:"matchProtocols,omitempty"`
	MatchSources   []NetworkSourceType   `json:"matchSources,omitempty"`
}

// +kubebuilder:validation:Enum=chown;dac_override;dac_read_search;fowner;fsetid;kill;setgid;setuid;setpcap;linux_immutable;net_bind_service;net_broadcast;net_admin;net_raw;ipc_lock;ipc_owner;sys_module;sys_rawio;sys_chroot;sys_ptrace;sys_pacct;sys_admin;sys_boot;sys_nice;sys_resource;sys_time;sys_tty_config;mknod;lease;audit_write;audit_control;setfcap;mac_override;mac_admin
type MatchCapabilitiesType string

// +kubebuilder:validation:Enum=TODO
type MatchOperationsType string

type CapabilitiesType struct {
	MatchCapabilities []MatchCapabilitiesType `json:"matchCapabilities,omitempty"`
	MatchOperations   []MatchOperationsType   `json:"matchOperations,omitempty"`
}

// +kubebuilder:validation:Enum=Allow;allow;Block;block;Audit;audit
type ActionType string

// KubeArmorPolicySpec defines the desired state of KubeArmorPolicy
type KubeArmorPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Selector SelectorType `json:"selector"`

	Process      ProcessType      `json:"process,omitempty"`
	File         FileType         `json:"file,omitempty"`
	Network      NetworkType      `json:"network,omitempty"`
	Capabilities CapabilitiesType `json:"capabilities,omitempty"`

	Action ActionType `json:"action"`
}

// KubeArmorPolicyStatus defines the observed state of KubeArmorPolicy
type KubeArmorPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// KubeArmorPolicy is the Schema for the kubearmorpolicies API
type KubeArmorPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeArmorPolicySpec   `json:"spec,omitempty"`
	Status KubeArmorPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeArmorPolicyList contains a list of KubeArmorPolicy
type KubeArmorPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeArmorPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeArmorPolicy{}, &KubeArmorPolicyList{})
}
