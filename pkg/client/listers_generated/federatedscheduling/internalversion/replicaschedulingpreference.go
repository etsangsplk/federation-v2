/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	federatedscheduling "github.com/kubernetes-sigs/federation-v2/pkg/apis/federatedscheduling"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReplicaSchedulingPreferenceLister helps list ReplicaSchedulingPreferences.
type ReplicaSchedulingPreferenceLister interface {
	// List lists all ReplicaSchedulingPreferences in the indexer.
	List(selector labels.Selector) (ret []*federatedscheduling.ReplicaSchedulingPreference, err error)
	// ReplicaSchedulingPreferences returns an object that can list and get ReplicaSchedulingPreferences.
	ReplicaSchedulingPreferences(namespace string) ReplicaSchedulingPreferenceNamespaceLister
	ReplicaSchedulingPreferenceListerExpansion
}

// replicaSchedulingPreferenceLister implements the ReplicaSchedulingPreferenceLister interface.
type replicaSchedulingPreferenceLister struct {
	indexer cache.Indexer
}

// NewReplicaSchedulingPreferenceLister returns a new ReplicaSchedulingPreferenceLister.
func NewReplicaSchedulingPreferenceLister(indexer cache.Indexer) ReplicaSchedulingPreferenceLister {
	return &replicaSchedulingPreferenceLister{indexer: indexer}
}

// List lists all ReplicaSchedulingPreferences in the indexer.
func (s *replicaSchedulingPreferenceLister) List(selector labels.Selector) (ret []*federatedscheduling.ReplicaSchedulingPreference, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*federatedscheduling.ReplicaSchedulingPreference))
	})
	return ret, err
}

// ReplicaSchedulingPreferences returns an object that can list and get ReplicaSchedulingPreferences.
func (s *replicaSchedulingPreferenceLister) ReplicaSchedulingPreferences(namespace string) ReplicaSchedulingPreferenceNamespaceLister {
	return replicaSchedulingPreferenceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReplicaSchedulingPreferenceNamespaceLister helps list and get ReplicaSchedulingPreferences.
type ReplicaSchedulingPreferenceNamespaceLister interface {
	// List lists all ReplicaSchedulingPreferences in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*federatedscheduling.ReplicaSchedulingPreference, err error)
	// Get retrieves the ReplicaSchedulingPreference from the indexer for a given namespace and name.
	Get(name string) (*federatedscheduling.ReplicaSchedulingPreference, error)
	ReplicaSchedulingPreferenceNamespaceListerExpansion
}

// replicaSchedulingPreferenceNamespaceLister implements the ReplicaSchedulingPreferenceNamespaceLister
// interface.
type replicaSchedulingPreferenceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReplicaSchedulingPreferences in the indexer for a given namespace.
func (s replicaSchedulingPreferenceNamespaceLister) List(selector labels.Selector) (ret []*federatedscheduling.ReplicaSchedulingPreference, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*federatedscheduling.ReplicaSchedulingPreference))
	})
	return ret, err
}

// Get retrieves the ReplicaSchedulingPreference from the indexer for a given namespace and name.
func (s replicaSchedulingPreferenceNamespaceLister) Get(name string) (*federatedscheduling.ReplicaSchedulingPreference, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(federatedscheduling.Resource("replicaschedulingpreference"), name)
	}
	return obj.(*federatedscheduling.ReplicaSchedulingPreference), nil
}
