// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumNodeLister helps list CiliumNodes.
type CiliumNodeLister interface {
	// List lists all CiliumNodes in the indexer.
	List(selector labels.Selector) (ret []*v2.CiliumNode, err error)
	// CiliumNodes returns an object that can list and get CiliumNodes.
	CiliumNodes(namespace string) CiliumNodeNamespaceLister
	CiliumNodeListerExpansion
}

// ciliumNodeLister implements the CiliumNodeLister interface.
type ciliumNodeLister struct {
	indexer cache.Indexer
}

// NewCiliumNodeLister returns a new CiliumNodeLister.
func NewCiliumNodeLister(indexer cache.Indexer) CiliumNodeLister {
	return &ciliumNodeLister{indexer: indexer}
}

// List lists all CiliumNodes in the indexer.
func (s *ciliumNodeLister) List(selector labels.Selector) (ret []*v2.CiliumNode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumNode))
	})
	return ret, err
}

// CiliumNodes returns an object that can list and get CiliumNodes.
func (s *ciliumNodeLister) CiliumNodes(namespace string) CiliumNodeNamespaceLister {
	return ciliumNodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CiliumNodeNamespaceLister helps list and get CiliumNodes.
type CiliumNodeNamespaceLister interface {
	// List lists all CiliumNodes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.CiliumNode, err error)
	// Get retrieves the CiliumNode from the indexer for a given namespace and name.
	Get(name string) (*v2.CiliumNode, error)
	CiliumNodeNamespaceListerExpansion
}

// ciliumNodeNamespaceLister implements the CiliumNodeNamespaceLister
// interface.
type ciliumNodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CiliumNodes in the indexer for a given namespace.
func (s ciliumNodeNamespaceLister) List(selector labels.Selector) (ret []*v2.CiliumNode, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumNode))
	})
	return ret, err
}

// Get retrieves the CiliumNode from the indexer for a given namespace and name.
func (s ciliumNodeNamespaceLister) Get(name string) (*v2.CiliumNode, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("ciliumnode"), name)
	}
	return obj.(*v2.CiliumNode), nil
}
