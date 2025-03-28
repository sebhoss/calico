// Copyright (c) 2025 Tigera, Inc. All rights reserved.

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

package types

import (
	"encoding/json"
	"unique"

	"github.com/sirupsen/logrus"

	"github.com/projectcalico/calico/goldmane/proto"
)

// FlowKey is a unique key for a flow. It matches the protobuf API exactly. Unfortunately,
// we cannot use the protobuf API structures as map keys due to private fields that are inserted
// by the protobuf Go code generation. So we need a copy of the struct here.
//
// This struct should be an exact copy of the proto.FlowKey structure, but without the private fields.
type FlowKey struct {
	// SourceName is the name of the source for this Flow. It represents one or more
	// source pods that share a GenerateName.
	SourceName string `protobuf:"bytes,1,opt,name=source_name,json=sourceName,proto3" json:"source_name,omitempty"`
	// SourceNamespace is the namespace of the source pods for this flow.
	SourceNamespace string `protobuf:"bytes,2,opt,name=source_namespace,json=sourceNamespace,proto3" json:"source_namespace,omitempty"`
	// SourceType is the type of the source, used to contextualize the source
	// name and namespace fields.
	SourceType proto.EndpointType `protobuf:"varint,3,opt,name=source_type,json=sourceType,proto3,enum=goldmane.EndpointType" json:"source_type,omitempty"`
	// DestName is the name of the destination for this Flow. It represents one or more
	// destination pods that share a GenerateName.
	DestName string `protobuf:"bytes,4,opt,name=dest_name,json=destName,proto3" json:"dest_name,omitempty"`
	// DestNamespace is the namespace of the destination pods for this flow.
	DestNamespace string `protobuf:"bytes,5,opt,name=dest_namespace,json=destNamespace,proto3" json:"dest_namespace,omitempty"`
	// DestType is the type of the destination, used to contextualize the dest
	// name and namespace fields.
	DestType proto.EndpointType `protobuf:"varint,6,opt,name=dest_type,json=destType,proto3,enum=goldmane.EndpointType" json:"dest_type,omitempty"`
	// DestPort is the destination port on the specified protocol accessed by this flow.
	DestPort int64 `protobuf:"varint,7,opt,name=dest_port,json=destPort,proto3" json:"dest_port,omitempty"`
	// DestServiceName is the name of the destination service, if any.
	DestServiceName string `protobuf:"bytes,8,opt,name=dest_service_name,json=destServiceName,proto3" json:"dest_service_name,omitempty"`
	// DestServiceNamespace is the namespace of the destination service, if any.
	DestServiceNamespace string `protobuf:"bytes,9,opt,name=dest_service_namespace,json=destServiceNamespace,proto3" json:"dest_service_namespace,omitempty"`
	// DestServicePortName is the name of the port on the destination service, if any.
	DestServicePortName string `protobuf:"bytes,10,opt,name=dest_service_port_name,json=destServicePortName,proto3" json:"dest_service_port_name,omitempty"`
	// DestServicePort is the port number on the destination service.
	DestServicePort int64 `protobuf:"varint,11,opt,name=dest_service_port,json=destServicePort,proto3" json:"dest_service_port,omitempty"`
	// Proto is the L4 protocol for this flow. For example, TCP, UDP, SCTP, ICMP.
	Proto string `protobuf:"bytes,12,opt,name=proto,proto3" json:"proto,omitempty"`
	// Reporter is either "src" or "dst", depending on whether this flow was generated
	// at the initiating or terminating end of the connection attempt.
	Reporter proto.Reporter `protobuf:"varint,13,opt,name=reporter,proto3,enum=goldmane.Reporter" json:"reporter,omitempty"`
	// Action is the ultimate action taken on the flow.
	Action proto.Action `protobuf:"varint,14,opt,name=action,proto3,enum=goldmane.Action" json:"action,omitempty"`
	// Policies includes an entry for each policy rule that took an action on the connections
	// aggregated into this flow.
	Policies unique.Handle[PolicyTrace] `protobuf:"bytes,15,opt,name=policies,proto3" json:"policies,omitempty"`
}

// This struct should be an exact copy of the proto.Flow structure, but without the private fields.
type Flow struct {
	// Key includes the identifying fields for this flow.
	Key *FlowKey `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// StartTime is the start time for this flow. It is represented as the number of
	// seconds since the UNIX epoch.
	StartTime int64 `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// EndTime is the end time for this flow. It is always at least one aggregation
	// interval after the start time.
	EndTime int64 `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// SourceLabels contains the intersection of labels that appear on all source
	// pods that contributed to this flow.
	SourceLabels []string `protobuf:"bytes,4,rep,name=source_labels,json=sourceLabels,proto3" json:"source_labels,omitempty"`
	// SourceLabels contains the intersection of labels that appear on all destination
	// pods that contributed to this flow.
	DestLabels []string `protobuf:"bytes,5,rep,name=dest_labels,json=destLabels,proto3" json:"dest_labels,omitempty"`
	// Statistics.
	PacketsIn  int64 `protobuf:"varint,6,opt,name=packets_in,json=packetsIn,proto3" json:"packets_in,omitempty"`
	PacketsOut int64 `protobuf:"varint,7,opt,name=packets_out,json=packetsOut,proto3" json:"packets_out,omitempty"`
	BytesIn    int64 `protobuf:"varint,8,opt,name=bytes_in,json=bytesIn,proto3" json:"bytes_in,omitempty"`
	BytesOut   int64 `protobuf:"varint,9,opt,name=bytes_out,json=bytesOut,proto3" json:"bytes_out,omitempty"`
	// NumConnectionsStarted tracks the total number of new connections recorded for this Flow. It counts each
	// connection attempt that matches the FlowKey that was made between this Flow's StartTime and EndTime.
	NumConnectionsStarted int64 `protobuf:"varint,10,opt,name=num_connections_started,json=numConnectionsStarted,proto3" json:"num_connections_started,omitempty"`
	// NumConnectionsCompleted tracks the total number of completed TCP connections recorded for this Flow. It counts each
	// connection that matches the FlowKey that was completed between this Flow's StartTime and EndTime.
	NumConnectionsCompleted int64 `protobuf:"varint,11,opt,name=num_connections_completed,json=numConnectionsCompleted,proto3" json:"num_connections_completed,omitempty"`
	// NumConnectionsLive tracks the total number of still active connections recorded for this Flow. It counts each
	// connection that matches the FlowKey that was active at this Flow's EndTime.
	NumConnectionsLive int64 `protobuf:"varint,12,opt,name=num_connections_live,json=numConnectionsLive,proto3" json:"num_connections_live,omitempty"`
}

type PolicyTrace struct {
	// EnforcedPolicies shows the active dataplane policy rules traversed by this Flow.
	EnforcedPolicies string `protobuf:"bytes,1,rep,name=enforced_policies,json=enforcedPolicies,proto3" json:"enforced_policies,omitempty"`
	// PendingPolicies shows the set of staged policy rules traversed by this Flow.
	PendingPolicies string `protobuf:"bytes,2,rep,name=pending_policies,json=pendingPolicies,proto3" json:"pending_policies,omitempty"`
}

type PolicyHit struct {
	Kind proto.PolicyKind `protobuf:"varint,1,opt,name=kind,proto3,enum=goldmane.PolicyKind" json:"kind,omitempty"`
	// Namespace is the Kubernetes namespace of the Policy, if namespaced. It is empty for global /
	// cluster-scoped policy kinds.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name is the Name of the policy object.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Tier is the Tier of the policy object.
	Tier string `protobuf:"bytes,4,opt,name=tier,proto3" json:"tier,omitempty"`
	// Action is the action taken by this policy rule.
	Action proto.Action `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	// PolicyIndex is the order of the Policy among all policies traversed.
	PolicyIndex int64 `protobuf:"varint,6,opt,name=policy_index,json=policyIndex,proto3" json:"policy_index,omitempty"`
	// RuleIndex is the order of the Rule within the Policy rules.
	RuleIndex int64 `protobuf:"varint,7,opt,name=rule_index,json=ruleIndex,proto3" json:"rule_index,omitempty"`
}

func ProtoToPolicyHit(p *proto.PolicyHit) *PolicyHit {
	return &PolicyHit{
		Kind:        p.Kind,
		Namespace:   p.Namespace,
		Name:        p.Name,
		Tier:        p.Tier,
		Action:      p.Action,
		PolicyIndex: p.PolicyIndex,
		RuleIndex:   p.RuleIndex,
	}
}

func PolicyHitToProto(p *PolicyHit) *proto.PolicyHit {
	return &proto.PolicyHit{
		Kind:        p.Kind,
		Namespace:   p.Namespace,
		Name:        p.Name,
		Tier:        p.Tier,
		Action:      p.Action,
		PolicyIndex: p.PolicyIndex,
		RuleIndex:   p.RuleIndex,
	}
}

func ProtoToFlow(p *proto.Flow) *Flow {
	return &Flow{
		Key:                     ProtoToFlowKey(p.Key),
		StartTime:               p.StartTime,
		EndTime:                 p.EndTime,
		SourceLabels:            p.SourceLabels,
		DestLabels:              p.DestLabels,
		PacketsIn:               p.PacketsIn,
		PacketsOut:              p.PacketsOut,
		BytesIn:                 p.BytesIn,
		BytesOut:                p.BytesOut,
		NumConnectionsStarted:   p.NumConnectionsStarted,
		NumConnectionsCompleted: p.NumConnectionsCompleted,
		NumConnectionsLive:      p.NumConnectionsLive,
	}
}

func ProtoToFlowKey(p *proto.FlowKey) *FlowKey {
	if p == nil {
		return nil
	}
	return &FlowKey{
		SourceName:           p.SourceName,
		SourceNamespace:      p.SourceNamespace,
		SourceType:           p.SourceType,
		DestName:             p.DestName,
		DestNamespace:        p.DestNamespace,
		DestType:             p.DestType,
		DestPort:             p.DestPort,
		DestServiceName:      p.DestServiceName,
		DestServiceNamespace: p.DestServiceNamespace,
		DestServicePortName:  p.DestServicePortName,
		DestServicePort:      p.DestServicePort,
		Proto:                p.Proto,
		Reporter:             p.Reporter,
		Action:               p.Action,
		Policies:             ProtoToFlowLogPolicy(p.Policies),
	}
}

func ProtoToFlowLogPolicy(p *proto.PolicyTrace) unique.Handle[PolicyTrace] {
	var ep, pp string
	if p != nil {
		if len(p.EnforcedPolicies) > 0 {
			epb, err := json.Marshal(p.EnforcedPolicies)
			if err != nil {
				logrus.WithError(err).Fatal("Failed to marshal enforced policies")
			}
			ep = string(epb)
		}

		if len(p.PendingPolicies) > 0 {
			ppb, err := json.Marshal(p.PendingPolicies)
			if err != nil {
				logrus.WithError(err).Fatal("Failed to marshal pending policies")
			}
			pp = string(ppb)
		}
	}
	flp := PolicyTrace{
		EnforcedPolicies: ep,
		PendingPolicies:  pp,
	}
	return unique.Make(flp)
}

func FlowToProto(f *Flow) *proto.Flow {
	return &proto.Flow{
		Key:                     FlowKeyToProto(f.Key),
		StartTime:               f.StartTime,
		EndTime:                 f.EndTime,
		SourceLabels:            f.SourceLabels,
		DestLabels:              f.DestLabels,
		PacketsIn:               f.PacketsIn,
		PacketsOut:              f.PacketsOut,
		BytesIn:                 f.BytesIn,
		BytesOut:                f.BytesOut,
		NumConnectionsStarted:   f.NumConnectionsStarted,
		NumConnectionsCompleted: f.NumConnectionsCompleted,
		NumConnectionsLive:      f.NumConnectionsLive,
	}
}

func FlowKeyToProto(f *FlowKey) *proto.FlowKey {
	if f == nil {
		return nil
	}
	return &proto.FlowKey{
		SourceName:           f.SourceName,
		SourceNamespace:      f.SourceNamespace,
		SourceType:           f.SourceType,
		DestName:             f.DestName,
		DestNamespace:        f.DestNamespace,
		DestType:             f.DestType,
		DestPort:             f.DestPort,
		DestServiceName:      f.DestServiceName,
		DestServiceNamespace: f.DestServiceNamespace,
		DestServicePortName:  f.DestServicePortName,
		DestServicePort:      f.DestServicePort,
		Proto:                f.Proto,
		Reporter:             f.Reporter,
		Action:               f.Action,
		Policies:             FlowLogPolicyToProto(f.Policies),
	}
}

func FlowLogPolicyToProto(h unique.Handle[PolicyTrace]) *proto.PolicyTrace {
	f := h.Value()
	var eps, pps []*proto.PolicyHit
	if f.EnforcedPolicies != "" {
		err := json.Unmarshal([]byte(f.EnforcedPolicies), &eps)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to unmarshal enforced policies")
		}
	}
	if f.PendingPolicies != "" {
		err := json.Unmarshal([]byte(f.PendingPolicies), &pps)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to unmarshal pending policies")
		}
	}
	return &proto.PolicyTrace{
		EnforcedPolicies: eps,
		PendingPolicies:  pps,
	}
}
