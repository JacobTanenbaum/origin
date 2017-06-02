// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	api "github.com/openshift/origin/pkg/sdn/api"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_ClusterNetwork_To_api_ClusterNetwork,
		Convert_api_ClusterNetwork_To_v1_ClusterNetwork,
		Convert_v1_ClusterNetworkEntry_To_api_ClusterNetworkEntry,
		Convert_api_ClusterNetworkEntry_To_v1_ClusterNetworkEntry,
		Convert_v1_ClusterNetworkList_To_api_ClusterNetworkList,
		Convert_api_ClusterNetworkList_To_v1_ClusterNetworkList,
		Convert_v1_EgressNetworkPolicy_To_api_EgressNetworkPolicy,
		Convert_api_EgressNetworkPolicy_To_v1_EgressNetworkPolicy,
		Convert_v1_EgressNetworkPolicyList_To_api_EgressNetworkPolicyList,
		Convert_api_EgressNetworkPolicyList_To_v1_EgressNetworkPolicyList,
		Convert_v1_EgressNetworkPolicyPeer_To_api_EgressNetworkPolicyPeer,
		Convert_api_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer,
		Convert_v1_EgressNetworkPolicyRule_To_api_EgressNetworkPolicyRule,
		Convert_api_EgressNetworkPolicyRule_To_v1_EgressNetworkPolicyRule,
		Convert_v1_EgressNetworkPolicySpec_To_api_EgressNetworkPolicySpec,
		Convert_api_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec,
		Convert_v1_HostSubnet_To_api_HostSubnet,
		Convert_api_HostSubnet_To_v1_HostSubnet,
		Convert_v1_HostSubnetList_To_api_HostSubnetList,
		Convert_api_HostSubnetList_To_v1_HostSubnetList,
		Convert_v1_NetNamespace_To_api_NetNamespace,
		Convert_api_NetNamespace_To_v1_NetNamespace,
		Convert_v1_NetNamespaceList_To_api_NetNamespaceList,
		Convert_api_NetNamespaceList_To_v1_NetNamespaceList,
	)
}

func autoConvert_v1_ClusterNetwork_To_api_ClusterNetwork(in *ClusterNetwork, out *api.ClusterNetwork, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Network = in.Network
	out.ClusterDef = *(*[]api.ClusterNetworkEntry)(unsafe.Pointer(&in.ClusterDef))
	out.HostSubnetLength = in.HostSubnetLength
	out.ServiceNetwork = in.ServiceNetwork
	out.PluginName = in.PluginName
	return nil
}

func Convert_v1_ClusterNetwork_To_api_ClusterNetwork(in *ClusterNetwork, out *api.ClusterNetwork, s conversion.Scope) error {
	return autoConvert_v1_ClusterNetwork_To_api_ClusterNetwork(in, out, s)
}

func autoConvert_api_ClusterNetwork_To_v1_ClusterNetwork(in *api.ClusterNetwork, out *ClusterNetwork, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Network = in.Network
	if in.ClusterDef == nil {
		out.ClusterDef = make([]ClusterNetworkEntry, 0)
	} else {
		out.ClusterDef = *(*[]ClusterNetworkEntry)(unsafe.Pointer(&in.ClusterDef))
	}
	out.HostSubnetLength = in.HostSubnetLength
	out.ServiceNetwork = in.ServiceNetwork
	out.PluginName = in.PluginName
	return nil
}

func Convert_api_ClusterNetwork_To_v1_ClusterNetwork(in *api.ClusterNetwork, out *ClusterNetwork, s conversion.Scope) error {
	return autoConvert_api_ClusterNetwork_To_v1_ClusterNetwork(in, out, s)
}

func autoConvert_v1_ClusterNetworkEntry_To_api_ClusterNetworkEntry(in *ClusterNetworkEntry, out *api.ClusterNetworkEntry, s conversion.Scope) error {
	out.ClusterNetworkCIDR = in.ClusterNetworkCIDR
	return nil
}

func Convert_v1_ClusterNetworkEntry_To_api_ClusterNetworkEntry(in *ClusterNetworkEntry, out *api.ClusterNetworkEntry, s conversion.Scope) error {
	return autoConvert_v1_ClusterNetworkEntry_To_api_ClusterNetworkEntry(in, out, s)
}

func autoConvert_api_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(in *api.ClusterNetworkEntry, out *ClusterNetworkEntry, s conversion.Scope) error {
	out.ClusterNetworkCIDR = in.ClusterNetworkCIDR
	return nil
}

func Convert_api_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(in *api.ClusterNetworkEntry, out *ClusterNetworkEntry, s conversion.Scope) error {
	return autoConvert_api_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(in, out, s)
}

func autoConvert_v1_ClusterNetworkList_To_api_ClusterNetworkList(in *ClusterNetworkList, out *api.ClusterNetworkList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]api.ClusterNetwork)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1_ClusterNetworkList_To_api_ClusterNetworkList(in *ClusterNetworkList, out *api.ClusterNetworkList, s conversion.Scope) error {
	return autoConvert_v1_ClusterNetworkList_To_api_ClusterNetworkList(in, out, s)
}

func autoConvert_api_ClusterNetworkList_To_v1_ClusterNetworkList(in *api.ClusterNetworkList, out *ClusterNetworkList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]ClusterNetwork, 0)
	} else {
		out.Items = *(*[]ClusterNetwork)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_api_ClusterNetworkList_To_v1_ClusterNetworkList(in *api.ClusterNetworkList, out *ClusterNetworkList, s conversion.Scope) error {
	return autoConvert_api_ClusterNetworkList_To_v1_ClusterNetworkList(in, out, s)
}

func autoConvert_v1_EgressNetworkPolicy_To_api_EgressNetworkPolicy(in *EgressNetworkPolicy, out *api.EgressNetworkPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_EgressNetworkPolicySpec_To_api_EgressNetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_EgressNetworkPolicy_To_api_EgressNetworkPolicy(in *EgressNetworkPolicy, out *api.EgressNetworkPolicy, s conversion.Scope) error {
	return autoConvert_v1_EgressNetworkPolicy_To_api_EgressNetworkPolicy(in, out, s)
}

func autoConvert_api_EgressNetworkPolicy_To_v1_EgressNetworkPolicy(in *api.EgressNetworkPolicy, out *EgressNetworkPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_api_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_EgressNetworkPolicy_To_v1_EgressNetworkPolicy(in *api.EgressNetworkPolicy, out *EgressNetworkPolicy, s conversion.Scope) error {
	return autoConvert_api_EgressNetworkPolicy_To_v1_EgressNetworkPolicy(in, out, s)
}

func autoConvert_v1_EgressNetworkPolicyList_To_api_EgressNetworkPolicyList(in *EgressNetworkPolicyList, out *api.EgressNetworkPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]api.EgressNetworkPolicy)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1_EgressNetworkPolicyList_To_api_EgressNetworkPolicyList(in *EgressNetworkPolicyList, out *api.EgressNetworkPolicyList, s conversion.Scope) error {
	return autoConvert_v1_EgressNetworkPolicyList_To_api_EgressNetworkPolicyList(in, out, s)
}

func autoConvert_api_EgressNetworkPolicyList_To_v1_EgressNetworkPolicyList(in *api.EgressNetworkPolicyList, out *EgressNetworkPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]EgressNetworkPolicy, 0)
	} else {
		out.Items = *(*[]EgressNetworkPolicy)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_api_EgressNetworkPolicyList_To_v1_EgressNetworkPolicyList(in *api.EgressNetworkPolicyList, out *EgressNetworkPolicyList, s conversion.Scope) error {
	return autoConvert_api_EgressNetworkPolicyList_To_v1_EgressNetworkPolicyList(in, out, s)
}

func autoConvert_v1_EgressNetworkPolicyPeer_To_api_EgressNetworkPolicyPeer(in *EgressNetworkPolicyPeer, out *api.EgressNetworkPolicyPeer, s conversion.Scope) error {
	out.CIDRSelector = in.CIDRSelector
	out.DNSName = in.DNSName
	return nil
}

func Convert_v1_EgressNetworkPolicyPeer_To_api_EgressNetworkPolicyPeer(in *EgressNetworkPolicyPeer, out *api.EgressNetworkPolicyPeer, s conversion.Scope) error {
	return autoConvert_v1_EgressNetworkPolicyPeer_To_api_EgressNetworkPolicyPeer(in, out, s)
}

func autoConvert_api_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer(in *api.EgressNetworkPolicyPeer, out *EgressNetworkPolicyPeer, s conversion.Scope) error {
	out.CIDRSelector = in.CIDRSelector
	out.DNSName = in.DNSName
	return nil
}

func Convert_api_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer(in *api.EgressNetworkPolicyPeer, out *EgressNetworkPolicyPeer, s conversion.Scope) error {
	return autoConvert_api_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer(in, out, s)
}

func autoConvert_v1_EgressNetworkPolicyRule_To_api_EgressNetworkPolicyRule(in *EgressNetworkPolicyRule, out *api.EgressNetworkPolicyRule, s conversion.Scope) error {
	out.Type = api.EgressNetworkPolicyRuleType(in.Type)
	if err := Convert_v1_EgressNetworkPolicyPeer_To_api_EgressNetworkPolicyPeer(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_EgressNetworkPolicyRule_To_api_EgressNetworkPolicyRule(in *EgressNetworkPolicyRule, out *api.EgressNetworkPolicyRule, s conversion.Scope) error {
	return autoConvert_v1_EgressNetworkPolicyRule_To_api_EgressNetworkPolicyRule(in, out, s)
}

func autoConvert_api_EgressNetworkPolicyRule_To_v1_EgressNetworkPolicyRule(in *api.EgressNetworkPolicyRule, out *EgressNetworkPolicyRule, s conversion.Scope) error {
	out.Type = EgressNetworkPolicyRuleType(in.Type)
	if err := Convert_api_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_EgressNetworkPolicyRule_To_v1_EgressNetworkPolicyRule(in *api.EgressNetworkPolicyRule, out *EgressNetworkPolicyRule, s conversion.Scope) error {
	return autoConvert_api_EgressNetworkPolicyRule_To_v1_EgressNetworkPolicyRule(in, out, s)
}

func autoConvert_v1_EgressNetworkPolicySpec_To_api_EgressNetworkPolicySpec(in *EgressNetworkPolicySpec, out *api.EgressNetworkPolicySpec, s conversion.Scope) error {
	out.Egress = *(*[]api.EgressNetworkPolicyRule)(unsafe.Pointer(&in.Egress))
	return nil
}

func Convert_v1_EgressNetworkPolicySpec_To_api_EgressNetworkPolicySpec(in *EgressNetworkPolicySpec, out *api.EgressNetworkPolicySpec, s conversion.Scope) error {
	return autoConvert_v1_EgressNetworkPolicySpec_To_api_EgressNetworkPolicySpec(in, out, s)
}

func autoConvert_api_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec(in *api.EgressNetworkPolicySpec, out *EgressNetworkPolicySpec, s conversion.Scope) error {
	if in.Egress == nil {
		out.Egress = make([]EgressNetworkPolicyRule, 0)
	} else {
		out.Egress = *(*[]EgressNetworkPolicyRule)(unsafe.Pointer(&in.Egress))
	}
	return nil
}

func Convert_api_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec(in *api.EgressNetworkPolicySpec, out *EgressNetworkPolicySpec, s conversion.Scope) error {
	return autoConvert_api_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec(in, out, s)
}

func autoConvert_v1_HostSubnet_To_api_HostSubnet(in *HostSubnet, out *api.HostSubnet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Host = in.Host
	out.HostIP = in.HostIP
	out.Subnet = in.Subnet
	return nil
}

func Convert_v1_HostSubnet_To_api_HostSubnet(in *HostSubnet, out *api.HostSubnet, s conversion.Scope) error {
	return autoConvert_v1_HostSubnet_To_api_HostSubnet(in, out, s)
}

func autoConvert_api_HostSubnet_To_v1_HostSubnet(in *api.HostSubnet, out *HostSubnet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Host = in.Host
	out.HostIP = in.HostIP
	out.Subnet = in.Subnet
	return nil
}

func Convert_api_HostSubnet_To_v1_HostSubnet(in *api.HostSubnet, out *HostSubnet, s conversion.Scope) error {
	return autoConvert_api_HostSubnet_To_v1_HostSubnet(in, out, s)
}

func autoConvert_v1_HostSubnetList_To_api_HostSubnetList(in *HostSubnetList, out *api.HostSubnetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]api.HostSubnet)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1_HostSubnetList_To_api_HostSubnetList(in *HostSubnetList, out *api.HostSubnetList, s conversion.Scope) error {
	return autoConvert_v1_HostSubnetList_To_api_HostSubnetList(in, out, s)
}

func autoConvert_api_HostSubnetList_To_v1_HostSubnetList(in *api.HostSubnetList, out *HostSubnetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]HostSubnet, 0)
	} else {
		out.Items = *(*[]HostSubnet)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_api_HostSubnetList_To_v1_HostSubnetList(in *api.HostSubnetList, out *HostSubnetList, s conversion.Scope) error {
	return autoConvert_api_HostSubnetList_To_v1_HostSubnetList(in, out, s)
}

func autoConvert_v1_NetNamespace_To_api_NetNamespace(in *NetNamespace, out *api.NetNamespace, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.NetName = in.NetName
	out.NetID = in.NetID
	return nil
}

func Convert_v1_NetNamespace_To_api_NetNamespace(in *NetNamespace, out *api.NetNamespace, s conversion.Scope) error {
	return autoConvert_v1_NetNamespace_To_api_NetNamespace(in, out, s)
}

func autoConvert_api_NetNamespace_To_v1_NetNamespace(in *api.NetNamespace, out *NetNamespace, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.NetName = in.NetName
	out.NetID = in.NetID
	return nil
}

func Convert_api_NetNamespace_To_v1_NetNamespace(in *api.NetNamespace, out *NetNamespace, s conversion.Scope) error {
	return autoConvert_api_NetNamespace_To_v1_NetNamespace(in, out, s)
}

func autoConvert_v1_NetNamespaceList_To_api_NetNamespaceList(in *NetNamespaceList, out *api.NetNamespaceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]api.NetNamespace)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1_NetNamespaceList_To_api_NetNamespaceList(in *NetNamespaceList, out *api.NetNamespaceList, s conversion.Scope) error {
	return autoConvert_v1_NetNamespaceList_To_api_NetNamespaceList(in, out, s)
}

func autoConvert_api_NetNamespaceList_To_v1_NetNamespaceList(in *api.NetNamespaceList, out *NetNamespaceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]NetNamespace, 0)
	} else {
		out.Items = *(*[]NetNamespace)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_api_NetNamespaceList_To_v1_NetNamespaceList(in *api.NetNamespaceList, out *NetNamespaceList, s conversion.Scope) error {
	return autoConvert_api_NetNamespaceList_To_v1_NetNamespaceList(in, out, s)
}
