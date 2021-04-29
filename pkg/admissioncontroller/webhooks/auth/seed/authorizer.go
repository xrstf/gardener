// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package seed

import (
	"context"
	"fmt"

	"github.com/gardener/gardener/pkg/admissioncontroller/seedidentity"
	"github.com/gardener/gardener/pkg/admissioncontroller/webhooks/auth/seed/graph"
	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	seedmanagementv1alpha1 "github.com/gardener/gardener/pkg/apis/seedmanagement/v1alpha1"
	"github.com/gardener/gardener/pkg/utils"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	auth "k8s.io/apiserver/pkg/authorization/authorizer"
)

// AuthorizerName is the name of this authorizer.
const AuthorizerName = "seedauthorizer"

// NewAuthorizer returns a new authorizer for requests from gardenlets. It never has an opinion on the request.
func NewAuthorizer(logger logr.Logger, graph graph.Interface) *authorizer {
	return &authorizer{
		logger: logger,
		graph:  graph,
	}
}

type authorizer struct {
	logger logr.Logger
	graph  graph.Interface
}

var _ = auth.Authorizer(&authorizer{})

var (
	// Only take v1beta1 for the core.gardener.cloud API group because the Authorize function only checks the resource
	// group and the resource (but it ignores the version).
	backupBucketResource           = gardencorev1beta1.Resource("backupbuckets")
	backupEntryResource            = gardencorev1beta1.Resource("backupentries")
	bastionResource                = gardencorev1beta1.Resource("bastions")
	cloudProfileResource           = gardencorev1beta1.Resource("cloudprofiles")
	configMapResource              = corev1.Resource("configmaps")
	controllerInstallationResource = gardencorev1beta1.Resource("controllerinstallations")
	managedSeedResource            = seedmanagementv1alpha1.Resource("managedseeds")
	namespaceResource              = corev1.Resource("namespaces")
	projectResource                = gardencorev1beta1.Resource("projects")
	secretBindingResource          = gardencorev1beta1.Resource("secretbindings")
	shootStateResource             = gardencorev1alpha1.Resource("shootstates")
)

// TODO: Revisit all `DecisionNoOpinion` later. Today we cannot deny the request for backwards compatibility
// because older Gardenlet versions might not be compatible at the time this authorization plugin is enabled.
// With `DecisionNoOpinion`, RBAC will be respected in the authorization chain afterwards.
func (a *authorizer) Authorize(_ context.Context, attrs auth.Attributes) (auth.Decision, string, error) {
	seedName, isSeed := seedidentity.FromUserInfoInterface(attrs.GetUser())
	if !isSeed {
		return auth.DecisionNoOpinion, "", nil
	}

	if attrs.IsResourceRequest() {
		requestResource := schema.GroupResource{Group: attrs.GetAPIGroup(), Resource: attrs.GetResource()}
		switch requestResource {
		case backupBucketResource:
			return a.authorize(seedName, graph.VertexTypeBackupBucket, attrs,
				[]string{"update", "patch", "delete"},
				[]string{"create", "get", "list", "watch"},
				[]string{"status"},
			)
		case backupEntryResource:
			return a.authorize(seedName, graph.VertexTypeBackupEntry, attrs,
				[]string{"update", "patch"},
				[]string{"create", "get", "list", "watch"},
				[]string{"status"},
			)
		case bastionResource:
			return a.authorize(seedName, graph.VertexTypeBastion, attrs,
				[]string{"update", "patch"},
				[]string{"create", "get", "list", "watch"},
				[]string{"status"},
			)
		case cloudProfileResource:
			return a.authorizeRead(seedName, graph.VertexTypeCloudProfile, attrs)
		case configMapResource:
			return a.authorizeRead(seedName, graph.VertexTypeConfigMap, attrs)
		case controllerInstallationResource:
			return a.authorize(seedName, graph.VertexTypeControllerInstallation, attrs,
				[]string{"update", "patch"},
				[]string{"get", "list", "watch"},
				[]string{"status"},
			)
		case managedSeedResource:
			return a.authorize(seedName, graph.VertexTypeManagedSeed, attrs,
				[]string{"update", "patch"},
				[]string{"get", "list", "watch"},
				[]string{"status"},
			)
		case namespaceResource:
			return a.authorizeRead(seedName, graph.VertexTypeNamespace, attrs)
		case projectResource:
			return a.authorizeRead(seedName, graph.VertexTypeProject, attrs)
		case secretBindingResource:
			return a.authorizeRead(seedName, graph.VertexTypeSecretBinding, attrs)
		case shootStateResource:
			return a.authorize(seedName, graph.VertexTypeShootState, attrs,
				[]string{"get", "update", "patch"},
				[]string{"create"},
				nil,
			)
		}
	}

	return auth.DecisionNoOpinion, "", nil
}

func (a *authorizer) authorizeRead(seedName string, fromType graph.VertexType, attrs auth.Attributes) (auth.Decision, string, error) {
	return a.authorize(seedName, fromType, attrs, []string{"get"}, nil, nil)
}

func (a *authorizer) authorize(
	seedName string,
	fromType graph.VertexType,
	attrs auth.Attributes,
	allowedVerbs []string,
	alwaysAllowedVerbs []string,
	allowedSubresources []string,
) (
	auth.Decision,
	string,
	error,
) {
	if ok, reason := a.checkSubresource(seedName, attrs, allowedSubresources...); !ok {
		return auth.DecisionNoOpinion, reason, nil
	}

	// When a new object is created then it doesn't yet exist in the graph, so usually such requests are always allowed
	// as the 'create case' is typically handled in the SeedRestriction admission handler. Similarly, resources for
	// which the gardenlet has a controller need to be listed/watched, so those verbs would also be allowed here.
	if utils.ValueExists(attrs.GetVerb(), alwaysAllowedVerbs) {
		return auth.DecisionAllow, "", nil
	}

	if ok, reason := a.checkVerb(seedName, attrs, append(alwaysAllowedVerbs, allowedVerbs...)...); !ok {
		return auth.DecisionNoOpinion, reason, nil
	}

	return a.hasPathFrom(seedName, fromType, attrs)
}

func (a *authorizer) hasPathFrom(seedName string, fromType graph.VertexType, attrs auth.Attributes) (auth.Decision, string, error) {
	if len(attrs.GetName()) == 0 {
		a.logger.Info(fmt.Sprintf("SEED DENY: '%s' %#v", seedName, attrs))
		return auth.DecisionNoOpinion, "No Object name found", nil
	}

	// Allow request if seed name is not known because a target seed cannot be used to find a path.
	if seedName == "" {
		return auth.DecisionAllow, "", nil
	}

	// If the request is made for a namespace then the attributes.Namespace field is not empty. It contains the name of
	// the namespace.
	namespace := attrs.GetNamespace()
	if fromType == graph.VertexTypeNamespace {
		namespace = ""
	}

	if !a.graph.HasPathFrom(fromType, namespace, attrs.GetName(), graph.VertexTypeSeed, "", seedName) {
		a.logger.Info(fmt.Sprintf("SEED DENY: '%s' %#v", seedName, attrs))
		return auth.DecisionNoOpinion, fmt.Sprintf("no relationship found between seed '%s' and this object", seedName), nil
	}

	return auth.DecisionAllow, "", nil
}

func (a *authorizer) checkVerb(seedName string, attrs auth.Attributes, allowedVerbs ...string) (bool, string) {
	if !utils.ValueExists(attrs.GetVerb(), allowedVerbs) {
		a.logger.Info(fmt.Sprintf("SEED DENY: '%s' %#v", seedName, attrs))
		return false, fmt.Sprintf("only the following verbs are allowed for this resource type: %+v", allowedVerbs)
	}

	return true, ""
}

func (a *authorizer) checkSubresource(seedName string, attrs auth.Attributes, allowedSubresources ...string) (bool, string) {
	if subresource := attrs.GetSubresource(); len(subresource) > 0 && !utils.ValueExists(attrs.GetSubresource(), allowedSubresources) {
		a.logger.Info(fmt.Sprintf("SEED DENY: '%s' %#v", seedName, attrs))
		return false, fmt.Sprintf("only the following subresources are allowed for this resource type: %+v", allowedSubresources)
	}

	return true, ""
}
