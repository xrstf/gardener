/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	core "github.com/gardener/gardener/pkg/apis/core"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BastionLister helps list Bastions.
// All objects returned here must be treated as read-only.
type BastionLister interface {
	// List lists all Bastions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*core.Bastion, err error)
	// Bastions returns an object that can list and get Bastions.
	Bastions(namespace string) BastionNamespaceLister
	BastionListerExpansion
}

// bastionLister implements the BastionLister interface.
type bastionLister struct {
	indexer cache.Indexer
}

// NewBastionLister returns a new BastionLister.
func NewBastionLister(indexer cache.Indexer) BastionLister {
	return &bastionLister{indexer: indexer}
}

// List lists all Bastions in the indexer.
func (s *bastionLister) List(selector labels.Selector) (ret []*core.Bastion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*core.Bastion))
	})
	return ret, err
}

// Bastions returns an object that can list and get Bastions.
func (s *bastionLister) Bastions(namespace string) BastionNamespaceLister {
	return bastionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BastionNamespaceLister helps list and get Bastions.
// All objects returned here must be treated as read-only.
type BastionNamespaceLister interface {
	// List lists all Bastions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*core.Bastion, err error)
	// Get retrieves the Bastion from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*core.Bastion, error)
	BastionNamespaceListerExpansion
}

// bastionNamespaceLister implements the BastionNamespaceLister
// interface.
type bastionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Bastions in the indexer for a given namespace.
func (s bastionNamespaceLister) List(selector labels.Selector) (ret []*core.Bastion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*core.Bastion))
	})
	return ret, err
}

// Get retrieves the Bastion from the indexer for a given namespace and name.
func (s bastionNamespaceLister) Get(name string) (*core.Bastion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(core.Resource("bastion"), name)
	}
	return obj.(*core.Bastion), nil
}
