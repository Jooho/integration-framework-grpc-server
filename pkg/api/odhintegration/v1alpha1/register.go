/*
Copyright 2022.

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

// Package v1alpha1 contains API Schema definitions for the odh v1alpha1 API group
//+kubebuilder:object:generate=true
//+groupName=odh.redhat.com
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "odh.redhat.com"
const GroupVersion = "v1alpha1"

var (
	// GroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: GroupVersion}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	// SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// Install is a function which adds this version to a scheme
	// Install = SchemeBuilder.AddToScheme
	Install = func(scheme *runtime.Scheme) error {
		if err := SchemeBuilder.AddToScheme(scheme); err != nil {
			return err
		}
		ParameterCodec = runtime.NewParameterCodec(scheme)
		return nil
	}

	ParameterCodec runtime.ParameterCodec
	// AddToScheme exists solely to keep the old generators creating valid code
	// DEPRECATED
	AddToScheme = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&ODHIntegration{},
		&ODHIntegrationList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
