/*
Copyright 2021.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AvailableHostSpec defines the desired state of AvailableHost
type AvailableHostSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ID              string `json:"id,omitempty"`
	Zone            string `json:"zone,omitempty"`
	Address         string `json:"address,omitempty"`
	InternalAddress string `json:"internalAddress,omitempty"`
	User            string `json:"user,omitempty"`
	Port            int    `json:"port,omitempty"`
	Password        string `json:"password,omitempty"`
	PrivateKey      string `json:"privateKey,omitempty"`
	ARCH            string `json:"arch,omitempty"`
	OSName          string `json:"osName,omitempty"`
	CPU             int    `json:"cpu,omitempty"`
	Memory          int    `json:"memory,omitempty"`
	Storage         int    `json:"storage,omitempty"`
}

// AvailableHostStatus defines the observed state of AvailableHost
type AvailableHostStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AvailableHost is the Schema for the availablehosts API
type AvailableHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AvailableHostSpec   `json:"spec,omitempty"`
	Status AvailableHostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AvailableHostList contains a list of AvailableHost
type AvailableHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AvailableHost `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AvailableHost{}, &AvailableHostList{})
}

type HostCfg struct {
	Name            string `yaml:"name,omitempty" json:"name,omitempty"`
	Address         string `yaml:"address,omitempty" json:"address,omitempty"`
	InternalAddress string `yaml:"internalAddress,omitempty" json:"internalAddress,omitempty"`
	Port            int    `yaml:"port,omitempty" json:"port,omitempty"`
	User            string `yaml:"user,omitempty" json:"user,omitempty"`
	Password        string `yaml:"password,omitempty" json:"password,omitempty"`
	PrivateKey      string `yaml:"privateKey,omitempty" json:"privateKey,omitempty"`
	PrivateKeyPath  string `yaml:"privateKeyPath,omitempty" json:"privateKeyPath,omitempty"`
	Arch            string `yaml:"arch,omitempty" json:"arch,omitempty"`

	Labels   map[string]string `yaml:"labels,omitempty" json:"labels,omitempty"`
	ID       string            `yaml:"id,omitempty" json:"id,omitempty"`
	Index    int               `json:"-"`
	IsEtcd   bool              `json:"-"`
	IsMaster bool              `json:"-"`
	IsWorker bool              `json:"-"`
}

type HostsAction struct {
	Hosts  []HostCfg `json:"hosts,omitempty"`
	Action int       `json:"action,omitempty"`
}
