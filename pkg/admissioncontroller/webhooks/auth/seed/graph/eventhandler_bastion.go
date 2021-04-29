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

package graph

import (
	"time"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"

	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/cache"
)

func (g *graph) setupBastionWatch(informer cache.Informer) {
	informer.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			bastion, ok := obj.(*gardencorev1beta1.Bastion)
			if !ok {
				return
			}
			g.handleBastionCreateOrUpdate(bastion)
		},

		UpdateFunc: func(oldObj, newObj interface{}) {
			oldBastion, ok := oldObj.(*gardencorev1beta1.Bastion)
			if !ok {
				return
			}

			newBastion, ok := newObj.(*gardencorev1beta1.Bastion)
			if !ok {
				return
			}

			if oldBastion.Status.SeedName != newBastion.Status.SeedName {
				g.handleBastionCreateOrUpdate(newBastion)
			}
		},

		DeleteFunc: func(obj interface{}) {
			if tombstone, ok := obj.(toolscache.DeletedFinalStateUnknown); ok {
				obj = tombstone.Obj
			}
			bastion, ok := obj.(*gardencorev1beta1.Bastion)
			if !ok {
				return
			}
			g.handleBastionDelete(bastion)
		},
	})
}

func (g *graph) handleBastionCreateOrUpdate(bastion *gardencorev1beta1.Bastion) {
	start := time.Now()
	defer func() {
		metricUpdateDuration.WithLabelValues("Bastion", "CreateOrUpdate").Observe(time.Since(start).Seconds())
	}()
	g.lock.Lock()
	defer g.lock.Unlock()

	g.deleteAllOutgoingEdges(VertexTypeBastion, bastion.Namespace, bastion.Name, VertexTypeSeed)
	g.deleteVertex(VertexTypeBastion, bastion.Namespace, bastion.Name)

	bastionVertex := g.getOrCreateVertex(VertexTypeBastion, bastion.Namespace, bastion.Name)

	if bastion.Status.SeedName != "" {
		seedVertex := g.getOrCreateVertex(VertexTypeSeed, "", bastion.Status.SeedName)
		g.addEdge(bastionVertex, seedVertex)
	}
}

func (g *graph) handleBastionDelete(bastion *gardencorev1beta1.Bastion) {
	start := time.Now()
	defer func() {
		metricUpdateDuration.WithLabelValues("Bastion", "Delete").Observe(time.Since(start).Seconds())
	}()
	g.lock.Lock()
	defer g.lock.Unlock()

	g.deleteVertex(VertexTypeBastion, bastion.Namespace, bastion.Name)
}
