/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	gentype "k8s.io/client-go/gentype"
	v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
	apisv1beta1 "sigs.k8s.io/gateway-api/applyconfiguration/apis/v1beta1"
	typedapisv1beta1 "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/typed/apis/v1beta1"
)

// fakeHTTPRoutes implements HTTPRouteInterface
type fakeHTTPRoutes struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.HTTPRoute, *v1beta1.HTTPRouteList, *apisv1beta1.HTTPRouteApplyConfiguration]
	Fake *FakeGatewayV1beta1
}

func newFakeHTTPRoutes(fake *FakeGatewayV1beta1, namespace string) typedapisv1beta1.HTTPRouteInterface {
	return &fakeHTTPRoutes{
		gentype.NewFakeClientWithListAndApply[*v1beta1.HTTPRoute, *v1beta1.HTTPRouteList, *apisv1beta1.HTTPRouteApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("httproutes"),
			v1beta1.SchemeGroupVersion.WithKind("HTTPRoute"),
			func() *v1beta1.HTTPRoute { return &v1beta1.HTTPRoute{} },
			func() *v1beta1.HTTPRouteList { return &v1beta1.HTTPRouteList{} },
			func(dst, src *v1beta1.HTTPRouteList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.HTTPRouteList) []*v1beta1.HTTPRoute { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.HTTPRouteList, items []*v1beta1.HTTPRoute) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
