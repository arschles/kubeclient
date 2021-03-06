/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package v0

import "github.com/jchauncey/kubeclient/api/unversioned"

// Policy contains a single ABAC policy rule
type Policy struct {
	unversioned.TypeMeta `json:",inline"`

	// User is the username this rule applies to.
	// Either user or group is required to match the request.
	// "*" matches all users.
	User string `json:"user,omitempty"`

	// Group is the group this rule applies to.
	// Either user or group is required to match the request.
	// "*" matches all groups.
	Group string `json:"group,omitempty"`

	// Readonly matches readonly requests when true, and all requests when false
	Readonly bool `json:"readonly,omitempty"`

	// Resource is the name of a resource
	// "*" matches all resources
	Resource string `json:"resource,omitempty"`

	// Namespace is the name of a namespace
	// "*" matches all namespaces (including unnamespaced requests)
	Namespace string `json:"namespace,omitempty"`
}
