// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/types/commonpb/v3/kube.proto

package commonv3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Taint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The taint key to be applied to a node.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The taint value corresponding to the taint key.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Required. The effect of the taint on pods
	// that do not tolerate the taint.
	// Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	Effect string `protobuf:"bytes,3,opt,name=effect,proto3" json:"effect,omitempty"`
	// TimeAdded represents the time at which the taint was added.
	// It is only written for NoExecute taints.
	TimeAdded *Time `protobuf:"bytes,4,opt,name=timeAdded,proto3" json:"timeAdded,omitempty"`
}

func (x *Taint) Reset() {
	*x = Taint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_commonpb_v3_kube_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Taint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Taint) ProtoMessage() {}

func (x *Taint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_commonpb_v3_kube_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Taint.ProtoReflect.Descriptor instead.
func (*Taint) Descriptor() ([]byte, []int) {
	return file_proto_types_commonpb_v3_kube_proto_rawDescGZIP(), []int{0}
}

func (x *Taint) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Taint) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Taint) GetEffect() string {
	if x != nil {
		return x.Effect
	}
	return ""
}

func (x *Taint) GetTimeAdded() *Time {
	if x != nil {
		return x.TimeAdded
	}
	return nil
}

type Time struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive. This field may be limited in precision depending on context.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Time) Reset() {
	*x = Time{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_commonpb_v3_kube_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Time) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Time) ProtoMessage() {}

func (x *Time) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_commonpb_v3_kube_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Time.ProtoReflect.Descriptor instead.
func (*Time) Descriptor() ([]byte, []int) {
	return file_proto_types_commonpb_v3_kube_proto_rawDescGZIP(), []int{1}
}

func (x *Time) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Time) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type NodeCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of node condition.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Last time we got an update on a given condition.
	LastHeartbeatTime *Time `protobuf:"bytes,3,opt,name=lastHeartbeatTime,proto3" json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	LastTransitionTime *Time `protobuf:"bytes,4,opt,name=lastTransitionTime,proto3" json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	Reason string `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NodeCondition) Reset() {
	*x = NodeCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_commonpb_v3_kube_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeCondition) ProtoMessage() {}

func (x *NodeCondition) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_commonpb_v3_kube_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeCondition.ProtoReflect.Descriptor instead.
func (*NodeCondition) Descriptor() ([]byte, []int) {
	return file_proto_types_commonpb_v3_kube_proto_rawDescGZIP(), []int{2}
}

func (x *NodeCondition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NodeCondition) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NodeCondition) GetLastHeartbeatTime() *Time {
	if x != nil {
		return x.LastHeartbeatTime
	}
	return nil
}

func (x *NodeCondition) GetLastTransitionTime() *Time {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

func (x *NodeCondition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *NodeCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NodeSystemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MachineID reported by the node. For unique machine identification
	// in the cluster this field is preferred. Learn more from man(5)
	// machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
	MachineID string `protobuf:"bytes,1,opt,name=machineID,proto3" json:"machineID,omitempty"`
	// SystemUUID reported by the node. For unique machine identification
	// MachineID is preferred. This field is specific to Red Hat hosts
	// https://access.redhat.com/documentation/en-us/red_hat_subscription_management/1/html/rhsm/uuid
	SystemUUID string `protobuf:"bytes,2,opt,name=systemUUID,proto3" json:"systemUUID,omitempty"`
	// Boot ID reported by the node.
	BootID string `protobuf:"bytes,3,opt,name=bootID,proto3" json:"bootID,omitempty"`
	// Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	KernelVersion string `protobuf:"bytes,4,opt,name=kernelVersion,proto3" json:"kernelVersion,omitempty"`
	// OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	OsImage string `protobuf:"bytes,5,opt,name=osImage,proto3" json:"osImage,omitempty"`
	// ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	ContainerRuntimeVersion string `protobuf:"bytes,6,opt,name=containerRuntimeVersion,proto3" json:"containerRuntimeVersion,omitempty"`
	// Kubelet Version reported by the node.
	KubeletVersion string `protobuf:"bytes,7,opt,name=kubeletVersion,proto3" json:"kubeletVersion,omitempty"`
	// KubeProxy Version reported by the node.
	KubeProxyVersion string `protobuf:"bytes,8,opt,name=kubeProxyVersion,proto3" json:"kubeProxyVersion,omitempty"`
	// The Operating System reported by the node
	OperatingSystem string `protobuf:"bytes,9,opt,name=operatingSystem,proto3" json:"operatingSystem,omitempty"`
	// The Architecture reported by the node
	Architecture string `protobuf:"bytes,10,opt,name=architecture,proto3" json:"architecture,omitempty"`
}

func (x *NodeSystemInfo) Reset() {
	*x = NodeSystemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_commonpb_v3_kube_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeSystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeSystemInfo) ProtoMessage() {}

func (x *NodeSystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_commonpb_v3_kube_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeSystemInfo.ProtoReflect.Descriptor instead.
func (*NodeSystemInfo) Descriptor() ([]byte, []int) {
	return file_proto_types_commonpb_v3_kube_proto_rawDescGZIP(), []int{3}
}

func (x *NodeSystemInfo) GetMachineID() string {
	if x != nil {
		return x.MachineID
	}
	return ""
}

func (x *NodeSystemInfo) GetSystemUUID() string {
	if x != nil {
		return x.SystemUUID
	}
	return ""
}

func (x *NodeSystemInfo) GetBootID() string {
	if x != nil {
		return x.BootID
	}
	return ""
}

func (x *NodeSystemInfo) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *NodeSystemInfo) GetOsImage() string {
	if x != nil {
		return x.OsImage
	}
	return ""
}

func (x *NodeSystemInfo) GetContainerRuntimeVersion() string {
	if x != nil {
		return x.ContainerRuntimeVersion
	}
	return ""
}

func (x *NodeSystemInfo) GetKubeletVersion() string {
	if x != nil {
		return x.KubeletVersion
	}
	return ""
}

func (x *NodeSystemInfo) GetKubeProxyVersion() string {
	if x != nil {
		return x.KubeProxyVersion
	}
	return ""
}

func (x *NodeSystemInfo) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *NodeSystemInfo) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

var File_proto_types_commonpb_v3_kube_proto protoreflect.FileDescriptor

var file_proto_types_commonpb_v3_kube_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaa, 0x01, 0x0a, 0x05, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x41, 0x64, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x41, 0x64, 0x64, 0x65, 0x64, 0x3a, 0x22, 0x92, 0x41, 0x1f, 0x0a, 0x1d, 0x2a,
	0x05, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x32, 0x14, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x04,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x3a, 0x11, 0x92, 0x41, 0x0e, 0x0a, 0x0c, 0x2a, 0x04, 0x54, 0x69, 0x6d,
	0x65, 0x32, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xd6, 0x02, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0x92, 0x41, 0x1e, 0x2a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x32, 0x16, 0x54, 0x79, 0x70, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x6e, 0x6f, 0x64, 0x65,
	0x20, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4d, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x24, 0x92, 0x41, 0x21, 0x0a,
	0x1f, 0x2a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x32, 0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xab, 0x03, 0x0a, 0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x55, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x6b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6b, 0x75, 0x62,
	0x65, 0x6c, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x6b,
	0x75, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x3a, 0x27, 0x92, 0x41, 0x24, 0x0a, 0x22, 0x2a, 0x0e, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x10, 0x4e, 0x6f,
	0x64, 0x65, 0x20, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x20, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x8b,
	0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33,
	0x42, 0x09, 0x4b, 0x75, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x76, 0x33, 0xa2, 0x02, 0x04, 0x52, 0x44, 0x54, 0x43, 0xaa, 0x02, 0x19, 0x52,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x19, 0x52, 0x61, 0x66, 0x61, 0x79,
	0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x25, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76,
	0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x33,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x52,
	0x61, 0x66, 0x61, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_commonpb_v3_kube_proto_rawDescOnce sync.Once
	file_proto_types_commonpb_v3_kube_proto_rawDescData = file_proto_types_commonpb_v3_kube_proto_rawDesc
)

func file_proto_types_commonpb_v3_kube_proto_rawDescGZIP() []byte {
	file_proto_types_commonpb_v3_kube_proto_rawDescOnce.Do(func() {
		file_proto_types_commonpb_v3_kube_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_commonpb_v3_kube_proto_rawDescData)
	})
	return file_proto_types_commonpb_v3_kube_proto_rawDescData
}

var file_proto_types_commonpb_v3_kube_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_types_commonpb_v3_kube_proto_goTypes = []interface{}{
	(*Taint)(nil),          // 0: rafay.dev.types.common.v3.Taint
	(*Time)(nil),           // 1: rafay.dev.types.common.v3.Time
	(*NodeCondition)(nil),  // 2: rafay.dev.types.common.v3.NodeCondition
	(*NodeSystemInfo)(nil), // 3: rafay.dev.types.common.v3.NodeSystemInfo
}
var file_proto_types_commonpb_v3_kube_proto_depIdxs = []int32{
	1, // 0: rafay.dev.types.common.v3.Taint.timeAdded:type_name -> rafay.dev.types.common.v3.Time
	1, // 1: rafay.dev.types.common.v3.NodeCondition.lastHeartbeatTime:type_name -> rafay.dev.types.common.v3.Time
	1, // 2: rafay.dev.types.common.v3.NodeCondition.lastTransitionTime:type_name -> rafay.dev.types.common.v3.Time
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_types_commonpb_v3_kube_proto_init() }
func file_proto_types_commonpb_v3_kube_proto_init() {
	if File_proto_types_commonpb_v3_kube_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_commonpb_v3_kube_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Taint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_commonpb_v3_kube_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Time); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_commonpb_v3_kube_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeCondition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_commonpb_v3_kube_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeSystemInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_types_commonpb_v3_kube_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_commonpb_v3_kube_proto_goTypes,
		DependencyIndexes: file_proto_types_commonpb_v3_kube_proto_depIdxs,
		MessageInfos:      file_proto_types_commonpb_v3_kube_proto_msgTypes,
	}.Build()
	File_proto_types_commonpb_v3_kube_proto = out.File
	file_proto_types_commonpb_v3_kube_proto_rawDesc = nil
	file_proto_types_commonpb_v3_kube_proto_goTypes = nil
	file_proto_types_commonpb_v3_kube_proto_depIdxs = nil
}
