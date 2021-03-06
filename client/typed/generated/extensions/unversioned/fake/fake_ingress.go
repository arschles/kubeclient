/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package fake

import (
	api "github.com/jchauncey/kubeclient/api"
	extensions "github.com/jchauncey/kubeclient/apis/extensions"
	core "github.com/jchauncey/kubeclient/client/testing/core"
	labels "github.com/jchauncey/kubeclient/labels"
	watch "github.com/jchauncey/kubeclient/watch"
)

// FakeIngresses implements IngressInterface
type FakeIngresses struct {
	Fake *FakeExtensions
	ns   string
}

func (c *FakeIngresses) Create(ingress *extensions.Ingress) (result *extensions.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(core.NewCreateAction("ingresses", c.ns, ingress), &extensions.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.Ingress), err
}

func (c *FakeIngresses) Update(ingress *extensions.Ingress) (result *extensions.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(core.NewUpdateAction("ingresses", c.ns, ingress), &extensions.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.Ingress), err
}

func (c *FakeIngresses) UpdateStatus(ingress *extensions.Ingress) (*extensions.Ingress, error) {
	obj, err := c.Fake.
		Invokes(core.NewUpdateSubresourceAction("ingresses", "status", c.ns, ingress), &extensions.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.Ingress), err
}

func (c *FakeIngresses) Delete(name string, options *api.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewDeleteAction("ingresses", c.ns, name), &extensions.Ingress{})

	return err
}

func (c *FakeIngresses) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	action := core.NewDeleteCollectionAction("events", c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &extensions.IngressList{})
	return err
}

func (c *FakeIngresses) Get(name string) (result *extensions.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(core.NewGetAction("ingresses", c.ns, name), &extensions.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.Ingress), err
}

func (c *FakeIngresses) List(opts api.ListOptions) (result *extensions.IngressList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewListAction("ingresses", c.ns, opts), &extensions.IngressList{})

	if obj == nil {
		return nil, err
	}

	label := opts.LabelSelector
	if label == nil {
		label = labels.Everything()
	}
	list := &extensions.IngressList{}
	for _, item := range obj.(*extensions.IngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingresses.
func (c *FakeIngresses) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewWatchAction("ingresses", c.ns, opts))

}
