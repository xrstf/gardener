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

package validation

import (
	"net"

	"github.com/gardener/gardener/pkg/apis/core"
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateBastion validates a Bastion object.
func ValidateBastion(bastion *core.Bastion) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, apivalidation.ValidateObjectMeta(&bastion.ObjectMeta, true, apivalidation.NameIsDNSLabel, field.NewPath("metadata"))...)
	allErrs = append(allErrs, ValidateBastionSpec(&bastion.Spec, field.NewPath("spec"))...)

	return allErrs
}

// ValidateBastionUpdate validates a Bastion object before an update.
func ValidateBastionUpdate(newBastion, oldBastion *core.Bastion) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, apivalidation.ValidateObjectMetaUpdate(&newBastion.ObjectMeta, &oldBastion.ObjectMeta, field.NewPath("metadata"))...)
	allErrs = append(allErrs, ValidateBastionSpecUpdate(&newBastion.Spec, &oldBastion.Spec, field.NewPath("spec"))...)
	allErrs = append(allErrs, ValidateBastion(newBastion)...)

	return allErrs
}

// ValidateBastionSpec validates the specification of a Bastion object.
func ValidateBastionSpec(spec *core.BastionSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(spec.ShootRef.Name) == 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("shootRef.name"), spec.ShootRef.Name, "shoot reference must not be empty"))
	}
	if len(spec.SSHPublicKey) == 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("sshPublicKey"), spec.SSHPublicKey, "sshPublicKey must not be empty"))
	}
	// TODO: validate SSH key thoroughly?

	if len(spec.Ingress) == 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("ingress"), spec.Ingress, "ingress must not be empty"))
	}

	for _, block := range spec.Ingress {
		if len(block.IPBlock.CIDR) == 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("ingress"), block.IPBlock.CIDR, "CIDR must not be empty"))
		} else if _, _, err := net.ParseCIDR(block.IPBlock.CIDR); err != nil {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("ingress"), block.IPBlock.CIDR, "invalid CIDR"))
		}
	}

	return allErrs
}

// ValidateBastionSpecUpdate validates the specification of a Bastion object.
func ValidateBastionSpecUpdate(newSpec, oldSpec *core.BastionSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, apivalidation.ValidateImmutableField(newSpec.ShootRef.Name, oldSpec.ShootRef.Name, fldPath.Child("shootRef.name"))...)
	allErrs = append(allErrs, apivalidation.ValidateImmutableField(newSpec.SSHPublicKey, oldSpec.SSHPublicKey, fldPath.Child("sshPublicKey"))...)

	return allErrs
}

// ValidateBastionStatusUpdate validates the status field of a Bastion object.
func ValidateBastionStatusUpdate(newBastion, oldBastion *core.Bastion) field.ErrorList {
	allErrs := field.ErrorList{}

	return allErrs
}
