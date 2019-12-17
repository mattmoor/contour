// +build !ignore_autogenerated

/*
Copyright 2019 VMware

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateDelegation) DeepCopyInto(out *CertificateDelegation) {
	*out = *in
	if in.TargetNamespaces != nil {
		in, out := &in.TargetNamespaces, &out.TargetNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateDelegation.
func (in *CertificateDelegation) DeepCopy() *CertificateDelegation {
	if in == nil {
		return nil
	}
	out := new(CertificateDelegation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.Header != nil {
		in, out := &in.Header, &out.Header
		*out = new(HeaderCondition)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPHealthCheckPolicy) DeepCopyInto(out *HTTPHealthCheckPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHealthCheckPolicy.
func (in *HTTPHealthCheckPolicy) DeepCopy() *HTTPHealthCheckPolicy {
	if in == nil {
		return nil
	}
	out := new(HTTPHealthCheckPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProxy) DeepCopyInto(out *HTTPProxy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProxy.
func (in *HTTPProxy) DeepCopy() *HTTPProxy {
	if in == nil {
		return nil
	}
	out := new(HTTPProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPProxy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProxyList) DeepCopyInto(out *HTTPProxyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HTTPProxy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProxyList.
func (in *HTTPProxyList) DeepCopy() *HTTPProxyList {
	if in == nil {
		return nil
	}
	out := new(HTTPProxyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPProxyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProxySpec) DeepCopyInto(out *HTTPProxySpec) {
	*out = *in
	if in.VirtualHost != nil {
		in, out := &in.VirtualHost, &out.VirtualHost
		*out = new(VirtualHost)
		(*in).DeepCopyInto(*out)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TCPProxy != nil {
		in, out := &in.TCPProxy, &out.TCPProxy
		*out = new(TCPProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.Includes != nil {
		in, out := &in.Includes, &out.Includes
		*out = make([]Include, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProxySpec.
func (in *HTTPProxySpec) DeepCopy() *HTTPProxySpec {
	if in == nil {
		return nil
	}
	out := new(HTTPProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderCondition) DeepCopyInto(out *HeaderCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderCondition.
func (in *HeaderCondition) DeepCopy() *HeaderCondition {
	if in == nil {
		return nil
	}
	out := new(HeaderCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderPolicy) DeepCopyInto(out *HeaderPolicy) {
	*out = *in
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make([]HeaderValue, len(*in))
		copy(*out, *in)
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderPolicy.
func (in *HeaderPolicy) DeepCopy() *HeaderPolicy {
	if in == nil {
		return nil
	}
	out := new(HeaderPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderValue) DeepCopyInto(out *HeaderValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderValue.
func (in *HeaderValue) DeepCopy() *HeaderValue {
	if in == nil {
		return nil
	}
	out := new(HeaderValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Include) DeepCopyInto(out *Include) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Include.
func (in *Include) DeepCopy() *Include {
	if in == nil {
		return nil
	}
	out := new(Include)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerPolicy) DeepCopyInto(out *LoadBalancerPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerPolicy.
func (in *LoadBalancerPolicy) DeepCopy() *LoadBalancerPolicy {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PathRewritePolicy) DeepCopyInto(out *PathRewritePolicy) {
	*out = *in
	if in.ReplacePrefix != nil {
		in, out := &in.ReplacePrefix, &out.ReplacePrefix
		*out = make([]ReplacePrefix, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PathRewritePolicy.
func (in *PathRewritePolicy) DeepCopy() *PathRewritePolicy {
	if in == nil {
		return nil
	}
	out := new(PathRewritePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplacePrefix) DeepCopyInto(out *ReplacePrefix) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplacePrefix.
func (in *ReplacePrefix) DeepCopy() *ReplacePrefix {
	if in == nil {
		return nil
	}
	out := new(ReplacePrefix)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryPolicy) DeepCopyInto(out *RetryPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryPolicy.
func (in *RetryPolicy) DeepCopy() *RetryPolicy {
	if in == nil {
		return nil
	}
	out := new(RetryPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeoutPolicy != nil {
		in, out := &in.TimeoutPolicy, &out.TimeoutPolicy
		*out = new(TimeoutPolicy)
		**out = **in
	}
	if in.RetryPolicy != nil {
		in, out := &in.RetryPolicy, &out.RetryPolicy
		*out = new(RetryPolicy)
		**out = **in
	}
	if in.HealthCheckPolicy != nil {
		in, out := &in.HealthCheckPolicy, &out.HealthCheckPolicy
		*out = new(HTTPHealthCheckPolicy)
		**out = **in
	}
	if in.LoadBalancerPolicy != nil {
		in, out := &in.LoadBalancerPolicy, &out.LoadBalancerPolicy
		*out = new(LoadBalancerPolicy)
		**out = **in
	}
	if in.PathRewrite != nil {
		in, out := &in.PathRewrite, &out.PathRewrite
		*out = new(PathRewritePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestHeadersPolicy != nil {
		in, out := &in.RequestHeadersPolicy, &out.RequestHeadersPolicy
		*out = new(HeaderPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.ResponseHeadersPolicy != nil {
		in, out := &in.ResponseHeadersPolicy, &out.ResponseHeadersPolicy
		*out = new(HeaderPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.UpstreamValidation != nil {
		in, out := &in.UpstreamValidation, &out.UpstreamValidation
		*out = new(UpstreamValidation)
		**out = **in
	}
	if in.RequestHeadersPolicy != nil {
		in, out := &in.RequestHeadersPolicy, &out.RequestHeadersPolicy
		*out = new(HeaderPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.ResponseHeadersPolicy != nil {
		in, out := &in.ResponseHeadersPolicy, &out.ResponseHeadersPolicy
		*out = new(HeaderPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPProxy) DeepCopyInto(out *TCPProxy) {
	*out = *in
	if in.LoadBalancerPolicy != nil {
		in, out := &in.LoadBalancerPolicy, &out.LoadBalancerPolicy
		*out = new(LoadBalancerPolicy)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = new(TCPProxyInclude)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPProxy.
func (in *TCPProxy) DeepCopy() *TCPProxy {
	if in == nil {
		return nil
	}
	out := new(TCPProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPProxyInclude) DeepCopyInto(out *TCPProxyInclude) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPProxyInclude.
func (in *TCPProxyInclude) DeepCopy() *TCPProxyInclude {
	if in == nil {
		return nil
	}
	out := new(TCPProxyInclude)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLS) DeepCopyInto(out *TLS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLS.
func (in *TLS) DeepCopy() *TLS {
	if in == nil {
		return nil
	}
	out := new(TLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSCertificateDelegation) DeepCopyInto(out *TLSCertificateDelegation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSCertificateDelegation.
func (in *TLSCertificateDelegation) DeepCopy() *TLSCertificateDelegation {
	if in == nil {
		return nil
	}
	out := new(TLSCertificateDelegation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TLSCertificateDelegation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSCertificateDelegationList) DeepCopyInto(out *TLSCertificateDelegationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TLSCertificateDelegation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSCertificateDelegationList.
func (in *TLSCertificateDelegationList) DeepCopy() *TLSCertificateDelegationList {
	if in == nil {
		return nil
	}
	out := new(TLSCertificateDelegationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TLSCertificateDelegationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSCertificateDelegationSpec) DeepCopyInto(out *TLSCertificateDelegationSpec) {
	*out = *in
	if in.Delegations != nil {
		in, out := &in.Delegations, &out.Delegations
		*out = make([]CertificateDelegation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSCertificateDelegationSpec.
func (in *TLSCertificateDelegationSpec) DeepCopy() *TLSCertificateDelegationSpec {
	if in == nil {
		return nil
	}
	out := new(TLSCertificateDelegationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeoutPolicy) DeepCopyInto(out *TimeoutPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutPolicy.
func (in *TimeoutPolicy) DeepCopy() *TimeoutPolicy {
	if in == nil {
		return nil
	}
	out := new(TimeoutPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamValidation) DeepCopyInto(out *UpstreamValidation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamValidation.
func (in *UpstreamValidation) DeepCopy() *UpstreamValidation {
	if in == nil {
		return nil
	}
	out := new(UpstreamValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualHost) DeepCopyInto(out *VirtualHost) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHost.
func (in *VirtualHost) DeepCopy() *VirtualHost {
	if in == nil {
		return nil
	}
	out := new(VirtualHost)
	in.DeepCopyInto(out)
	return out
}
