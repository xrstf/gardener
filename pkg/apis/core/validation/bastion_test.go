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

// package validation_test

// import (
// 	"github.com/gardener/gardener/pkg/apis/core"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/apimachinery/pkg/util/validation/field"

// 	. "github.com/gardener/gardener/pkg/apis/core/validation"
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	. "github.com/onsi/gomega/gstruct"
// )

// var _ = Describe("validation", func() {
// 	var bastion *core.Bastion

// 	BeforeEach(func() {
// 		bastion = &core.Bastion{
// 			ObjectMeta: metav1.ObjectMeta{
// 				Name:      "example-backup-entry",
// 				Namespace: "garden",
// 				Annotations: map[string]string{
// 					core.BastionForceDeletion: "true",
// 				},
// 			},
// 			Spec: core.BastionSpec{
// 				BucketName: "some-bucket",
// 			},
// 		}
// 	})

// 	Describe("#ValidateBastion", func() {
// 		It("should not return any errors", func() {
// 			errorList := ValidateBastion(bastion)

// 			Expect(errorList).To(HaveLen(0))
// 		})

// 		It("should forbid Bastion resources with empty metadata", func() {
// 			bastion.ObjectMeta = metav1.ObjectMeta{}

// 			errorList := ValidateBastion(bastion)

// 			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
// 				"Type":  Equal(field.ErrorTypeRequired),
// 				"Field": Equal("metadata.name"),
// 			})),
// 				PointTo(MatchFields(IgnoreExtras, Fields{
// 					"Type":  Equal(field.ErrorTypeRequired),
// 					"Field": Equal("metadata.namespace"),
// 				}))))
// 		})

// 		It("should forbid Bastion specification with empty or invalid keys", func() {
// 			bastion.Spec.BucketName = ""

// 			errorList := ValidateBastion(bastion)

// 			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
// 				"Type":  Equal(field.ErrorTypeInvalid),
// 				"Field": Equal("spec.bucketName"),
// 			}))))
// 		})

// 		It("should forbid updating some keys", func() {
// 			newBastion := prepareBastionForUpdate(bastion)
// 			newBastion.Spec.BucketName = "another-bucketName"

// 			errorList := ValidateBastionUpdate(newBastion, bastion)

// 			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
// 				"Type":  Equal(field.ErrorTypeInvalid),
// 				"Field": Equal("spec.bucketName"),
// 			}))))
// 		})
// 	})
// })

// func prepareBastionForUpdate(obj *core.Bastion) *core.Bastion {
// 	newObj := obj.DeepCopy()
// 	newObj.ResourceVersion = "1"
// 	return newObj
// }
