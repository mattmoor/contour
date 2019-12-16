# Header Rewrite For HTTPProxy

**Status**: _Draft_

This document specifies a design for supporting request / response header manipulation in the HTTPProxy CRD.

## Goals

Adding and removing headers from requests and responses at two levels:
1. Per-route (pre-split),
1. Per-cluster (post-split).

## Non Goals

- Support for projecting dynamic values into header values.

## Background

There are a number of use cases where having this raw capability is useful, and many are called out in #70.

## High-Level Design

Add a new type: `HeaderRewritePolicy` that captures adding and removing headers.

```Go
// HeaderRewritePolicy defines alterations to the headers being sent to or returned from a service.
type HeaderRewritePolicy struct {
	// Add holds the header key/value pairs to add to those sent to or returned from a service.
	// +optional
	Add []HeaderAddition `json:"add,omitempty"`
	// Remove lists the header keys to remove from those sent to or returned from a service.
	// +optional
	Remove []string `json:"remove,omitempty"`
}

// HeaderAddition defines a header key/value pair to be added to those sent to or return from a service.
type HeaderAddition struct {
	// Name is the header key to add.
	Name string `json:"name"`
	// Value is the header value to add.
	Value string `json:"value"`
}
```

This will be added in two flavors (request and response) to both Route (pre-split) and Service (post-split).

```Go
	// RequestHeaders defines how to add or remove headers from requests routed to this service.
	// +optional
	RequestHeaders *HeaderRewritePolicy `json:"requestHeadersPolicy,omitempty"`
	// ResponseHeaders defines how to add or remove headers from responses returned from this service.
	// +optional
	ResponseHeaders *HeaderRewritePolicy `json:"responseHeadersPolicy,omitempty"`
```

## Detailed Design

For the most part, these fields will be directly translated to the following fields in the respective Envoy proto:
 - `RequestHeadersToAdd`
 - `RequestHeadersToRemove`
 - `ResponseHeadersToAdd`
 - `ResponseHeadersToRemove`

There are two notable exceptions to this translation:
1. `Host` header manipulations must be done via a separate Envoy directive.
1. [`%`-encoded variables](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/headers#custom-request-response-headers) are not supported, so literal `%`'s must be escaped in header values.


## Alternatives Considered

There were a few alternatives discussed around API shapes for general header manipulation in #70, but this shape was chosen to address the various scenarios raised there (e.g. request/response).

There was some discussion of API shapes for `Host` rewriting in #1982.


## Security Considerations

None at this time.