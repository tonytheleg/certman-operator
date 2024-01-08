//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSPlatformSecrets) DeepCopyInto(out *AWSPlatformSecrets) {
	*out = *in
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSPlatformSecrets.
func (in *AWSPlatformSecrets) DeepCopy() *AWSPlatformSecrets {
	if in == nil {
		return nil
	}
	out := new(AWSPlatformSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzurePlatformSecrets) DeepCopyInto(out *AzurePlatformSecrets) {
	*out = *in
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzurePlatformSecrets.
func (in *AzurePlatformSecrets) DeepCopy() *AzurePlatformSecrets {
	if in == nil {
		return nil
	}
	out := new(AzurePlatformSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequest) DeepCopyInto(out *CertificateRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequest.
func (in *CertificateRequest) DeepCopy() *CertificateRequest {
	if in == nil {
		return nil
	}
	out := new(CertificateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestCondition) DeepCopyInto(out *CertificateRequestCondition) {
	*out = *in
	if in.LastProbeTime != nil {
		in, out := &in.LastProbeTime, &out.LastProbeTime
		*out = (*in).DeepCopy()
	}
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestCondition.
func (in *CertificateRequestCondition) DeepCopy() *CertificateRequestCondition {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestList) DeepCopyInto(out *CertificateRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestList.
func (in *CertificateRequestList) DeepCopy() *CertificateRequestList {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestSpec) DeepCopyInto(out *CertificateRequestSpec) {
	*out = *in
	out.CertificateSecret = in.CertificateSecret
	in.Platform.DeepCopyInto(&out.Platform)
	if in.DnsNames != nil {
		in, out := &in.DnsNames, &out.DnsNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestSpec.
func (in *CertificateRequestSpec) DeepCopy() *CertificateRequestSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestStatus) DeepCopyInto(out *CertificateRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateRequestCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestStatus.
func (in *CertificateRequestStatus) DeepCopy() *CertificateRequestStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPPlatformSecrets) DeepCopyInto(out *GCPPlatformSecrets) {
	*out = *in
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPPlatformSecrets.
func (in *GCPPlatformSecrets) DeepCopy() *GCPPlatformSecrets {
	if in == nil {
		return nil
	}
	out := new(GCPPlatformSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MockPlatformSecrets) DeepCopyInto(out *MockPlatformSecrets) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MockPlatformSecrets.
func (in *MockPlatformSecrets) DeepCopy() *MockPlatformSecrets {
	if in == nil {
		return nil
	}
	out := new(MockPlatformSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platform) DeepCopyInto(out *Platform) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSPlatformSecrets)
		**out = **in
	}
	if in.GCP != nil {
		in, out := &in.GCP, &out.GCP
		*out = new(GCPPlatformSecrets)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzurePlatformSecrets)
		**out = **in
	}
	if in.Mock != nil {
		in, out := &in.Mock, &out.Mock
		*out = new(MockPlatformSecrets)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platform.
func (in *Platform) DeepCopy() *Platform {
	if in == nil {
		return nil
	}
	out := new(Platform)
	in.DeepCopyInto(out)
	return out
}