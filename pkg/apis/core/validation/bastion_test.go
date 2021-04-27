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

package validation_test

import (
	"github.com/gardener/gardener/pkg/apis/core"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	. "github.com/gardener/gardener/pkg/apis/core/validation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("validation", func() {
	var bastion *core.Bastion

	BeforeEach(func() {
		bastion = &core.Bastion{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "example-bastion",
				Namespace: "garden",
			},
			Spec: core.BastionSpec{
				ShootRef: core.BastionShootReference{
					Name: "example-shoot",
				},
				SSHPublicKey: "ssh-rsa 1234",
				Ingress: []core.BastionIngressPolicy{{
					IPBlock: networkingv1.IPBlock{
						CIDR: "1.2.3.4/8",
					},
				}},
			},
		}
	})

	Describe("#ValidateBastion", func() {
		It("should not return any errors", func() {
			errorList := ValidateBastion(bastion)

			Expect(errorList).To(HaveLen(0))
		})

		It("should forbid Bastion resources with empty metadata", func() {
			bastion.ObjectMeta = metav1.ObjectMeta{}

			errorList := ValidateBastion(bastion)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("metadata.name"),
			})), PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("metadata.namespace"),
			}))))
		})

		It("should forbid Bastion specification with empty SSH key", func() {
			bastion.Spec.SSHPublicKey = ""

			errorList := ValidateBastion(bastion)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.sshPublicKey"),
			}))))
		})

		It("should forbid Bastion specification with empty Shoot ref", func() {
			bastion.Spec.ShootRef.Name = ""

			errorList := ValidateBastion(bastion)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.shootRef.name"),
			}))))
		})

		It("should forbid Bastion specification with empty ingress", func() {
			bastion.Spec.Ingress = []core.BastionIngressPolicy{}

			errorList := ValidateBastion(bastion)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.ingress"),
			}))))
		})

		It("should forbid Bastion specification with invalid ingress", func() {
			bastion.Spec.Ingress[0].IPBlock.CIDR = "not-a-cidr"

			errorList := ValidateBastion(bastion)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.ingress"),
			}))))
		})

		It("should forbid changing Shoot ref", func() {
			newBastion := prepareBastionForUpdate(bastion)
			newBastion.Spec.ShootRef.Name = "another-shoot"

			errorList := ValidateBastionUpdate(newBastion, bastion)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.shootRef.name"),
			}))))
		})

		It("should forbid changing SSH key", func() {
			newBastion := prepareBastionForUpdate(bastion)
			newBastion.Spec.SSHPublicKey = "ssh-rsa another-key"

			errorList := ValidateBastionUpdate(newBastion, bastion)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.sshPublicKey"),
			}))))
		})
	})
})

func prepareBastionForUpdate(obj *core.Bastion) *core.Bastion {
	newObj := obj.DeepCopy()
	newObj.ResourceVersion = "1"
	return newObj
}
