/*
Copyright 2025.

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
    corev1 "k8s.io/api/core/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CompilerEnvironmentSpec defines the desired state of CompilerEnvironment
type CompilerEnvironmentSpec struct {
    // +kubebuilder:validation:Required
    // 工具链容器镜像 (例如: "gcc:11.3", "python:3.9-slim")
    Image string `json:"image"`

    // 传递给容器的命令 (覆盖 Dockerfile 的 CMD)
    Command []string `json:"command,omitempty"`

    // 传递给容器的参数
    Args []string `json:"args,omitempty"`

    // +kubebuilder:validation:Required
    // 资源请求和限制
    Resources corev1.ResourceRequirements `json:"resources"`

    // 存储卷大小请求 (例如: "5Gi")
    StorageRequest string `json:"storageRequest"`

    // 环境变量
    Env []corev1.EnvVar `json:"env,omitempty"`
}

// CompilerEnvironmentStatus defines the observed state of CompilerEnvironment.
type CompilerEnvironmentStatus struct {
    // 环境的状态 (Creating, Running, Failed, Terminating)
    Phase string `json:"phase,omitempty"`

    // 所创建 Pod 的名称
    PodName string `json:"podName,omitempty"`

    // 所创建 PVC 的名称
    PVCName string `json:"pvcName,omitempty"`

    // 如果环境是服务型的，可以在这里记录访问地址
    AccessURL string `json:"accessURL,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CompilerEnvironment is the Schema for the compilerenvironments API
type CompilerEnvironment struct {
    metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	//metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec defines the desired state of CompilerEnvironment
	// +required
	//Spec CompilerEnvironmentSpec `json:"spec"`
    Spec CompilerEnvironmentSpec `json:"spec,omitempty"`

	// status defines the observed state of CompilerEnvironment
	// +optional
	//Status CompilerEnvironmentStatus `json:"status,omitempty,omitzero"`
    Status CompilerEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CompilerEnvironmentList contains a list of CompilerEnvironment
type CompilerEnvironmentList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []CompilerEnvironment `json:"items"`
}

func init() {
    SchemeBuilder.Register(&CompilerEnvironment{}, &CompilerEnvironmentList{})
}
