# GEP 1323: Response Header Filter

* Issue: [#1323](https://github.com/kubernetes-sigs/gateway-api/issues/1323)
* Status: Standard

> **Note**: This GEP is exempt from the [Probationary Period][expprob] rules of
> our GEP overview as it existed before those rules did, and so it has been
> explicitly grandfathered in.

[expprob]:../overview.md#probationary-period

## TLDR
Similar to how we have `RequestHeaderModifier` in `HTTPRouteFilter`, which lets users modify request headers before the request is forwarded to a backend (or a group of backends), it’d be helpful to have a `ResponseHeaderModifier` field which would let users modify response headers before they are returned to the client.

## Goals
* Provide a way to modify HTTP response headers in a `HTTPRoute`.
* Reuse existing types as much as possible to reduce boilerplate code.

## Non Goals
* Provide a way to modify other parts of a HTTP response like status code.
* Add fields specifically for standard headers such as `Cookie`.

## Introduction
Currently, the `HTTPRouteFilter` API provides a way for request headers to be modified through the `RequestHeaderModifier` field of type `HTTPRequestHeaderModifier`. But, a similar API to modify response headers does not exist. This proposal intends to introduce a new field in `HTTPRouteFilter` named `ResponseHeaderModifier`.

## API
We could introduce a new API named `HTTPResponseHeaderModifier` which would look exactly like the existing `HTTPRequestHeaderModifier` API. But since HTTP headers have the same semantics for both requests and responses, it makes more sense to rename `HTTPRequestHeaderModifier` to `HTTPHeaderModifier` and use this for both `RequestHeaderModifier` and `ResponseHeaderModifier`.

```golang
// HTTPHeaderModifier defines a filter that modifies the headers of a HTTP
// request or response.
type HTTPHeaderModifier struct {
    // Set overwrites the request with the given header (name, value)
    // before the action.
    // +optional
    // +listType=map
    // +listMapKey=name
    // +kubebuilder:validation:MaxItems=16
    Set []HTTPHeader `json:"set,omitempty"`

    // Add adds the given header(s) (name, value) to the request
    // before the action. It appends to any existing values associated
    // with the header name.

    // +optional
    // +listType=map
    // +listMapKey=name
    // +kubebuilder:validation:MaxItems=16
    Add []HTTPHeader `json:"add,omitempty"`

    // Remove the given header(s) from the HTTP request before the action. The
    // value of Remove is a list of HTTP header names. Note that the header
    // names are case-insensitive (see
    // https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).
    // +optional
    // +kubebuilder:validation:MaxItems=16
    Remove []string `json:"remove,omitempty"`
}
```

Given the fact that this functionality is offered by only a few projects that currently implement Gateway API when using their own traffic routing CRDs, it’s better to support `ResponseHeaderModifier` as an _Extended_ feature, unlike `RequestHeaderModifier` which is a _Core_ feature. This will also not increase the difficulty of implementing Gateway API for any future ingress or service mesh.

This feature can be further extended via [Policy Attachment](../../reference/policy-attachment.md). The mechanism and use cases of this may be explored in a future GEP.

## Usage
Adding support for this unlocks a lot of real world use cases. Let’s review a couple of them:

* A team has a frontend web app, along with two different versions of their backends exposed as Kubernetes services. If, the frontend needs to know which backend it’s talking to, this can be easily achieved without any modifications to the application code.

```yaml
apiVersion: gateway.networking.k8s.io/v1beta1
kind: HTTPRoute
metadata:
  name: http-response-header
spec:
  hostnames:
    - response.header.example
  rules:
    - backendRefs:
      - name: example-svc-beta
        weight: 50
        port: 80
        # set a custom header for all responses being sent from the beta build of the backend server.
        filters:
           - type: ResponseHeaderModifier
             responseHeaderModifier:
              add:
                name: build
                value: beta
      - name: example-svc-stable
        weight: 50
        port: 80
```

* Cookies can be automatically injected into the response of services. This can enable services to identify users that were redirected to a certain backend.

```yaml
apiVersion: gateway.networking.k8s.io/v1beta1
kind: HTTPRoute
metadata:
  name: http-response-header
spec:
  hostnames:
    - response.header.example
  rules:
    # match against any requests that has the cookie set due to the below rule
    - matches:
      - headers:
        type: Exact
        name: Cookie
        value: user=insider
      backendRefs:
      - name: foo-svc
        port: 8080

    - filters:
      - type: ResponseHeaderModifier
        # set cookies for all requests being forwarded to this service
        responseHeaderModifier:
          set:
            name: Set-Cookie
            value: user=insider
      backendRefs:
      - name: example-svc
        weight: 1
        port: 80
```

> Note: Some projects like Envoy support interpolating a few predefined [variables into header values](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/headers#custom-request-response-headers). Similar functionality might be supported by other implementations but its unlikely to be portable and thus has been excluded from the API for the time being.

## Prior Art
A few projects that implement Gateway API already have support for similar functionality (in their custom CRDs), like:
* Istio’s `VirtualService`:

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews-route
spec:
  hosts:
  - reviews.prod.svc.cluster.local
  http:
  - headers:
      request:
        set:
          test: "true"
    route:
    - destination:
        host: reviews.prod.svc.cluster.local
        subset: v2
      weight: 25
    - destination:
        host: reviews.prod.svc.cluster.local
        subset: v1
      headers:
        response:
          remove:
          - foo
      weight: 75
```

* Contour’s `HTTPProxy`:

```yaml
apiVersion: projectcontour.io/v1
kind: HTTPProxy
metadata:
  name: basic
spec:
  virtualhost:
    fqdn: foo-basic.bar.com
  routes:
    - conditions:
      - prefix: /
      services:
        - name: s1
          port: 80
      responseHeadersPolicy:
        set:
          name: test
          value: true
```

* Ingress NGINX:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nginx-headers
  annotations:
    nginx.ingress.kubernetes.io/configuration-snippet: |
      add_header ingress nginx;
spec:
  ingressClassName: nginx
  rules:
  - host: custom.configuration.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: http-svc
            port:
              number: 80
```

## External Links
* [Contour `HeaderValue`](https://projectcontour.io/docs/v1.22.0/config/api/#projectcontour.io/v1.HeaderValue)
* [Istio `Headers`](https://istio.io/latest/docs/reference/config/networking/virtual-service/#Headers)
* [Ingress NGINX Configuration Snippets](https://kubernetes.github.io/ingress-nginx/examples/customization/configuration-snippets/)
* [NGINX `add_header` directive](https://nginx.org/en/docs/http/ngx_http_headers_module.html#add_header)

