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

package bastion

import (
	"context"
	"fmt"

	"github.com/gardener/gardener/pkg/api"
	"github.com/gardener/gardener/pkg/apis/core"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/pkg/apis/core/validation"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

type bastionStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

// Strategy defines the storage strategy for Bastions.
var Strategy = bastionStrategy{api.Scheme, names.SimpleNameGenerator}

func (bastionStrategy) NamespaceScoped() bool {
	return true
}

func (bastionStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	bastion := obj.(*core.Bastion)

	bastion.Generation = 1
	bastion.Status = core.BastionStatus{}
}

func (bastionStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	newBastion := obj.(*core.Bastion)
	oldBastion := old.(*core.Bastion)
	newBastion.Status = oldBastion.Status

	if mustIncreaseGeneration(oldBastion, newBastion) {
		newBastion.Generation = oldBastion.Generation + 1
	}
}

func mustIncreaseGeneration(oldBastion, newBastion *core.Bastion) bool {
	// The Bastion specification changes.
	if !apiequality.Semantic.DeepEqual(oldBastion.Spec, newBastion.Spec) {
		return true
	}

	// The deletion timestamp was set.
	if oldBastion.DeletionTimestamp == nil && newBastion.DeletionTimestamp != nil {
		return true
	}

	if kutil.HasMetaDataAnnotation(&newBastion.ObjectMeta, v1beta1constants.GardenerOperation, v1beta1constants.GardenerOperationReconcile) {
		return true
	}

	return false
}

func (bastionStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	bastion := obj.(*core.Bastion)
	return validation.ValidateBastion(bastion)
}

func (bastionStrategy) Canonicalize(obj runtime.Object) {
}

func (bastionStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (bastionStrategy) ValidateUpdate(ctx context.Context, newObj, oldObj runtime.Object) field.ErrorList {
	oldBastion, newBastion := oldObj.(*core.Bastion), newObj.(*core.Bastion)
	return validation.ValidateBastionUpdate(newBastion, oldBastion)
}

func (bastionStrategy) AllowUnconditionalUpdate() bool {
	return false
}

type bastionStatusStrategy struct {
	bastionStrategy
}

// StatusStrategy defines the storage strategy for the status subresource of Bastions.
var StatusStrategy = bastionStatusStrategy{Strategy}

func (bastionStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	newBastion := obj.(*core.Bastion)
	oldBastion := old.(*core.Bastion)
	newBastion.Spec = oldBastion.Spec
}

func (bastionStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return validation.ValidateBastionStatusUpdate(obj.(*core.Bastion), old.(*core.Bastion))
}

// ToSelectableFields returns a field set that represents the object
// TODO: fields are not labels, and the validation rules for them do not apply.
func ToSelectableFields(bastion *core.Bastion) fields.Set {
	// The purpose of allocation with a given number of elements is to reduce
	// amount of allocations needed to create the fields.Set. If you add any
	// field here or the number of object-meta related fields changes, this should
	// be adjusted.
	bastionSpecificFieldsSet := make(fields.Set, 3)
	bastionSpecificFieldsSet[core.BastionSeedName] = getSeedName(bastion)
	return generic.AddObjectMetaFieldsSet(bastionSpecificFieldsSet, &bastion.ObjectMeta, true)
}

// GetAttrs returns labels and fields of a given object for filtering purposes.
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	bastion, ok := obj.(*core.Bastion)
	if !ok {
		return nil, nil, fmt.Errorf("not a bastion")
	}
	return labels.Set(bastion.ObjectMeta.Labels), ToSelectableFields(bastion), nil
}

// MatchBastion returns a generic matcher for a given label and field selector.
func MatchBastion(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:       label,
		Field:       field,
		GetAttrs:    GetAttrs,
		IndexFields: []string{core.BastionSeedName},
	}
}

// SeedNameTriggerFunc returns spec.seedName of given Bastion.
func SeedNameTriggerFunc(obj runtime.Object) string {
	bastion, ok := obj.(*core.Bastion)
	if !ok {
		return ""
	}

	return getSeedName(bastion)
}

func getSeedName(bastion *core.Bastion) string {
	return bastion.Status.SeedName
}
