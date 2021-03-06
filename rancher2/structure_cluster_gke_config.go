package rancher2

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Flatteners

func flattenClusterGKEConfig(in *GoogleKubernetesEngineConfig) ([]interface{}, error) {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}, nil
	}

	if len(in.ClusterIpv4Cidr) > 0 {
		obj["cluster_ipv4_cidr"] = in.ClusterIpv4Cidr
	}

	if len(in.Credential) > 0 {
		obj["credential"] = in.Credential
	}

	if len(in.Description) > 0 {
		obj["description"] = in.Description
	}

	obj["enable_alpha_feature"] = in.EnableAlphaFeature

	if in.EnableHTTPLoadBalancing != nil {
		obj["enable_http_load_balancing"] = *in.EnableHTTPLoadBalancing
	}

	if in.EnableHorizontalPodAutoscaling != nil {
		obj["enable_horizontal_pod_autoscaling"] = *in.EnableHorizontalPodAutoscaling
	}

	obj["enable_kubernetes_dashboard"] = in.EnableKubernetesDashboard
	obj["enable_legacy_abac"] = in.EnableLegacyAbac
	obj["enable_master_authorized_network"] = in.EnableMasterAuthorizedNetwork

	if in.EnableNetworkPolicyConfig != nil {
		obj["enable_network_policy_config"] = *in.EnableNetworkPolicyConfig
	}

	obj["enable_nodepool_autoscaling"] = in.EnableNodepoolAutoscaling
	obj["enable_private_endpoint"] = in.EnablePrivateEndpoint
	obj["enable_private_nodes"] = in.EnablePrivateNodes

	if in.EnableStackdriverLogging != nil {
		obj["enable_stackdriver_logging"] = *in.EnableStackdriverLogging
	}

	if in.EnableHorizontalPodAutoscaling != nil {
		obj["enable_stackdriver_monitoring"] = *in.EnableStackdriverMonitoring
	}

	if len(in.IPPolicyClusterIpv4CidrBlock) > 0 {
		obj["ip_policy_cluster_ipv4_cidr_block"] = in.IPPolicyClusterIpv4CidrBlock
	}

	if len(in.IPPolicyClusterSecondaryRangeName) > 0 {
		obj["ip_policy_cluster_secondary_range_name"] = in.IPPolicyClusterSecondaryRangeName
	}

	obj["ip_policy_create_subnetwork"] = in.IPPolicyCreateSubnetwork

	if len(in.IPPolicyNodeIpv4CidrBlock) > 0 {
		obj["ip_policy_node_ipv4_cidr_block"] = in.IPPolicyNodeIpv4CidrBlock
	}

	if len(in.IPPolicyServicesIpv4CidrBlock) > 0 {
		obj["ip_policy_services_ipv4_cidr_block"] = in.IPPolicyServicesIpv4CidrBlock
	}

	if len(in.IPPolicyServicesSecondaryRangeName) > 0 {
		obj["ip_policy_services_secondary_range_name"] = in.IPPolicyServicesSecondaryRangeName
	}

	if len(in.IPPolicySubnetworkName) > 0 {
		obj["ip_policy_subnetwork_name"] = in.IPPolicySubnetworkName
	}

	obj["issue_client_certificate"] = in.IssueClientCertificate
	obj["kubernetes_dashboard"] = in.KubernetesDashboard

	if len(in.Locations) > 0 {
		obj["locations"] = toArrayInterface(in.Locations)
	}

	if len(in.MaintenanceWindow) > 0 {
		obj["maintenance_window"] = in.MaintenanceWindow
	}

	if len(in.MasterAuthorizedNetworkCidrBlocks) > 0 {
		obj["master_authorized_network_cidr_blocks"] = toArrayInterface(in.MasterAuthorizedNetworkCidrBlocks)
	}

	if len(in.MasterIpv4CidrBlock) > 0 {
		obj["master_ipv4_cidr_block"] = in.MasterIpv4CidrBlock
	}

	if len(in.MasterVersion) > 0 {
		obj["master_version"] = in.MasterVersion
	}

	if len(in.Network) > 0 {
		obj["network"] = in.Network
	}

	if len(in.ProjectID) > 0 {
		obj["project_id"] = in.ProjectID
	}

	if len(in.ResourceLabels) > 0 {
		obj["resource_labels"] = toMapInterface(in.ResourceLabels)
	}

	if len(in.SubNetwork) > 0 {
		obj["sub_network"] = in.SubNetwork
	}

	obj["use_ip_aliases"] = in.UseIPAliases

	if len(in.Zone) > 0 {
		obj["zone"] = in.Zone
	}

	if len(in.Region) > 0 {
		obj["region"] = in.Region
	}

	if in.DefaultMaxPodsConstraint > 0 {
		obj["default_max_pods_constraint"] = int(in.DefaultMaxPodsConstraint)
	}

	var nodePoolObjs []interface{}

	for _, nodePoolIn := range in.NodePools {
		var nodePool GoogleKubernetesEngineNodePool
		if err := json.Unmarshal([]byte(nodePoolIn), &nodePool); err != nil {
			return nil, err
		}

		nodePoolObj := flattenClusterBaseNodePool(nodePool.BaseNodePool)

		if nodePool.DiskSizeGb > 0 {
			nodePoolObj["disk_size_gb"] = int(nodePool.DiskSizeGb)
		}

		if len(nodePool.DiskType) > 0 {
			nodePoolObj["disk_type"] = nodePool.DiskType
		}

		nodePoolObj["enable_auto_repair"] = nodePool.EnableAutoRepair
		nodePoolObj["enable_auto_upgrade"] = nodePool.EnableAutoUpgrade

		if len(nodePool.ImageType) > 0 {
			nodePoolObj["image_type"] = nodePool.ImageType
		}

		if nodePool.LocalSsdCount > 0 {
			nodePoolObj["local_ssd_count"] = int(nodePool.LocalSsdCount)
		}

		if len(nodePool.MachineType) > 0 {
			nodePoolObj["machine_type"] = nodePool.MachineType
		}

		if nodePool.MaximumNodeCount > 0 {
			nodePoolObj["max_node_count"] = int(nodePool.MaximumNodeCount)
		}

		if nodePool.MinimumNodeCount > 0 {
			nodePoolObj["min_node_count"] = int(nodePool.MinimumNodeCount)
		}

		if len(nodePool.MinimumCpuPlatform) > 0 {
			nodePoolObj["min_cpu_platform"] = nodePool.MinimumCpuPlatform
		}

		if len(nodePool.OauthScopes) > 0 {
			nodePoolObj["oauth_scopes"] = toArrayInterface(nodePool.OauthScopes)
		}

		nodePoolObj["preemptible"] = nodePool.Preemptible

		if len(nodePool.ServiceAccount) > 0 {
			nodePoolObj["service_account"] = nodePool.ServiceAccount
		}

		if len(nodePool.Version) > 0 {
			nodePoolObj["version"] = nodePool.Version
		}

		nodePoolObjs = append(nodePoolObjs, nodePoolObj)
	}

	// when rancher returns details of a cluster that hasn't been migrated to new state model, we fallback to old
	// fields to extract existing node pool details
	if len(nodePoolObjs) == 0 {
		nodePoolObj := make(map[string]interface{})

		nodePoolObj["add_default_label"] = false
		nodePoolObj["add_default_taint"] = false

		if in.DiskSizeGb > 0 {
			nodePoolObj["disk_size_gb"] = int(in.DiskSizeGb)
		}

		if len(in.DiskType) > 0 {
			nodePoolObj["disk_type"] = in.DiskType
		}

		nodePoolObj["enable_auto_repair"] = in.EnableAutoRepair
		nodePoolObj["enable_auto_upgrade"] = in.EnableAutoUpgrade

		if len(in.ImageType) > 0 {
			nodePoolObj["image_type"] = in.ImageType
		}

		if in.LocalSsdCount > 0 {
			nodePoolObj["local_ssd_count"] = int(in.LocalSsdCount)
		}

		if len(in.MachineType) > 0 {
			nodePoolObj["machine_type"] = in.MachineType
		}

		if in.MaxNodeCount > 0 {
			nodePoolObj["max_node_count"] = int(in.MaxNodeCount)
		}

		if in.MinNodeCount > 0 {
			nodePoolObj["min_node_count"] = int(in.MinNodeCount)
		}

		if len(in.MinCpuPlatform) > 0 {
			nodePoolObj["min_cpu_platform"] = in.MinCpuPlatform
		}

		if len(in.NodePool) > 0 {
			nodePoolObj["name"] = in.NodePool
		}

		if len(in.NodeVersion) > 0 {
			nodePoolObj["version"] = in.NodeVersion
		}

		if len(in.OauthScopes) > 0 {
			nodePoolObj["oauth_scopes"] = toArrayInterface(in.OauthScopes)
		}

		nodePoolObj["preemptible"] = in.Preemptible

		if len(in.ServiceAccount) > 0 {
			nodePoolObj["service_account"] = in.ServiceAccount
		}

		nodePoolObjs = append(nodePoolObjs, nodePoolObj)
	}

	obj["node_pools"] = nodePoolObjs
	return []interface{}{obj}, nil
}

func flattenClusterGKENodePoolLegacyTaints(in []string) []interface{} {
	var taints []interface{}

	for _, part := range in {
		taint := make(map[string]interface{})
		ekv := strings.Split(part, ":")
		if len(ekv) == 2 {
			taint["effect"] = ekv[0]
			kv := strings.Split(ekv[1], "=")
			if len(kv) == 2 {
				taint["key"] = kv[0]
				taint["value"] = kv[1]
			}
		}

		taints = append(taints, taint)
	}

	return taints
}

// Expanders

func expandClusterGKEConfig(obj *GoogleKubernetesEngineConfig, p []interface{}, name string) (*GoogleKubernetesEngineConfig, error) {
	if len(p) == 0 || p[0] == nil {
		return obj, nil
	}
	in := p[0].(map[string]interface{})

	obj.DisplayName = name
	obj.Name = name

	if v, ok := in["cluster_ipv4_cidr"].(string); ok && len(v) > 0 {
		obj.ClusterIpv4Cidr = v
	}

	if v, ok := in["credential"].(string); ok && len(v) > 0 {
		obj.Credential = v
	}

	if v, ok := in["description"].(string); ok && len(v) > 0 {
		obj.Description = v
	}

	if v, ok := in["enable_alpha_feature"].(bool); ok {
		obj.EnableAlphaFeature = v
	}

	if v, ok := in["enable_http_load_balancing"].(bool); ok {
		obj.EnableHTTPLoadBalancing = &v
	}

	if v, ok := in["enable_horizontal_pod_autoscaling"].(bool); ok {
		obj.EnableHorizontalPodAutoscaling = &v
	}

	if v, ok := in["enable_kubernetes_dashboard"].(bool); ok {
		obj.EnableKubernetesDashboard = v
	}

	if v, ok := in["enable_legacy_abac"].(bool); ok {
		obj.EnableLegacyAbac = v
	}

	if v, ok := in["enable_master_authorized_network"].(bool); ok {
		obj.EnableMasterAuthorizedNetwork = v
	}

	if v, ok := in["enable_network_policy_config"].(bool); ok {
		obj.EnableNetworkPolicyConfig = &v
	}

	if v, ok := in["enable_nodepool_autoscaling"].(bool); ok {
		obj.EnableNodepoolAutoscaling = v
	}

	if v, ok := in["enable_private_endpoint"].(bool); ok {
		obj.EnablePrivateEndpoint = v
	}

	if v, ok := in["enable_private_nodes"].(bool); ok {
		obj.EnablePrivateNodes = v
	}

	if v, ok := in["enable_stackdriver_logging"].(bool); ok {
		obj.EnableStackdriverLogging = &v
	}

	if v, ok := in["enable_stackdriver_monitoring"].(bool); ok {
		obj.EnableStackdriverMonitoring = &v
	}

	if v, ok := in["ip_policy_cluster_ipv4_cidr_block"].(string); ok && len(v) > 0 {
		obj.IPPolicyClusterIpv4CidrBlock = v
	}

	if v, ok := in["ip_policy_cluster_secondary_range_name"].(string); ok && len(v) > 0 {
		obj.IPPolicyClusterSecondaryRangeName = v
	}

	if v, ok := in["ip_policy_create_subnetwork"].(bool); ok {
		obj.IPPolicyCreateSubnetwork = v
	}

	if v, ok := in["ip_policy_node_ipv4_cidr_block"].(string); ok && len(v) > 0 {
		obj.IPPolicyNodeIpv4CidrBlock = v
	}

	if v, ok := in["ip_policy_services_ipv4_cidr_block"].(string); ok && len(v) > 0 {
		obj.IPPolicyServicesIpv4CidrBlock = v
	}

	if v, ok := in["ip_policy_services_secondary_range_name"].(string); ok && len(v) > 0 {
		obj.IPPolicyServicesSecondaryRangeName = v
	}

	if v, ok := in["ip_policy_subnetwork_name"].(string); ok && len(v) > 0 {
		obj.IPPolicySubnetworkName = v
	}

	if v, ok := in["issue_client_certificate"].(bool); ok {
		obj.IssueClientCertificate = v
	}

	if v, ok := in["kubernetes_dashboard"].(bool); ok {
		obj.KubernetesDashboard = v
	}

	if v, ok := in["locations"].([]interface{}); ok && len(v) > 0 {
		obj.Locations = toArrayString(v)
	}

	if v, ok := in["maintenance_window"].(string); ok && len(v) > 0 {
		obj.MaintenanceWindow = v
	}

	if v, ok := in["master_authorized_network_cidr_blocks"].([]interface{}); ok && len(v) > 0 {
		obj.MasterAuthorizedNetworkCidrBlocks = toArrayString(v)
	}

	if v, ok := in["master_ipv4_cidr_block"].(string); ok && len(v) > 0 {
		obj.MasterIpv4CidrBlock = v
	}

	if v, ok := in["master_version"].(string); ok && len(v) > 0 {
		obj.MasterVersion = v
	}

	if v, ok := in["network"].(string); ok && len(v) > 0 {
		obj.Network = v
	}

	if v, ok := in["node_count"].(int); ok && v > 0 {
		obj.NodeCount = int64(v)
	}

	if v, ok := in["project_id"].(string); ok && len(v) > 0 {
		obj.ProjectID = v
	}

	if v, ok := in["resource_labels"].(map[string]interface{}); ok && len(v) > 0 {
		obj.ResourceLabels = toMapString(v)
	}

	if v, ok := in["sub_network"].(string); ok && len(v) > 0 {
		obj.SubNetwork = v
	}

	if v, ok := in["use_ip_aliases"].(bool); ok {
		obj.UseIPAliases = v
	}

	if v, ok := in["zone"].(string); ok && len(v) > 0 {
		obj.Zone = v
	}

	if v, ok := in["region"].(string); ok && len(v) > 0 {
		obj.Region = v
	}

	if v, ok := in["default_max_pods_constraint"].(int); ok && v > 0 {
		obj.DefaultMaxPodsConstraint = int64(v)
	}

	var nodePoolObjs []string

	if vs, ok := in["node_pools"]; ok {
		if nodePoolIns, ok := vs.([]interface{}); !ok {
			return nil, errors.New("unexpected content in 'node_pools'")
		} else if len(nodePoolIns) > 0 {
			for index, v := range nodePoolIns {
				if nodePoolIn, ok := v.(map[string]interface{}); ok {
					nodePoolObj, err := expandClusterGKENodePool(nodePoolIn, false)
					if err != nil {
						return nil, err
					}

					nodePoolObjs = append(nodePoolObjs, nodePoolObj)
				} else {
					return nil, fmt.Errorf("unexpected content in node pool with index %d", index)
				}
			}
		}
	}

	if len(nodePoolObjs) == 0 {
		if nodePoolObj, err := expandClusterGKENodePool(in, true); err == nil {
			nodePoolObjs = append(nodePoolObjs, nodePoolObj)
		} else {
			return nil, err
		}

	}

	obj.NodePools = nodePoolObjs

	return obj, nil
}

func expandClusterGKENodePool(in map[string]interface{}, legacy bool) (string, error) {
	var versionField string
	var bnp BaseNodePool

	if legacy {
		versionField = "node_version"

		bnp = BaseNodePool{}

		if v, ok := in["node_pool"].(string); ok && len(v) > 0 {
			bnp.Name = v
		} else {
			bnp.Name = "default-0"
		}

		if v, ok := in["labels"].(map[string]interface{}); ok && len(v) > 0 {
			bnp.Labels = toMapString(v)
		}

		bnp.Taints = expandClusterGKENodePoolLegacyTaints(in)

	} else {
		versionField = "version"

		var err error
		bnp, err = expandClusterBaseNodePool(in)
		if err != nil {
			return "", err
		}

	}

	nodePoolObj := GoogleKubernetesEngineNodePool{
		BaseNodePool: bnp,
	}

	if v, ok := in["disk_size_gb"].(int); ok && v > 0 {
		nodePoolObj.DiskSizeGb = int64(v)
	}

	if v, ok := in["disk_type"].(string); ok && len(v) > 0 {
		nodePoolObj.DiskType = v
	}

	if v, ok := in["enable_auto_repair"].(bool); ok {
		nodePoolObj.EnableAutoRepair = v
	}

	if v, ok := in["enable_auto_upgrade"].(bool); ok {
		nodePoolObj.EnableAutoUpgrade = v
	}

	if v, ok := in["image_type"].(string); ok && len(v) > 0 {
		nodePoolObj.ImageType = v
	}

	if v, ok := in["local_ssd_count"].(int); ok && v > 0 {
		nodePoolObj.LocalSsdCount = int64(v)
	}

	if v, ok := in["machine_type"].(string); ok && len(v) > 0 {
		nodePoolObj.MachineType = v
	}

	if v, ok := in["max_node_count"].(int); ok && v > 0 {
		nodePoolObj.MaximumNodeCount = int64(v)
	}

	if v, ok := in["min_node_count"].(int); ok && v > 0 {
		nodePoolObj.MinimumNodeCount = int64(v)
	}

	if v, ok := in["min_cpu_platform"].(string); ok && len(v) > 0 {
		nodePoolObj.MinimumCpuPlatform = v
	}

	if v, ok := in[versionField].(string); ok && len(v) > 0 {
		nodePoolObj.Version = v
	}

	if v, ok := in["oauth_scopes"].([]interface{}); ok && len(v) > 0 {
		nodePoolObj.OauthScopes = toArrayString(v)
	}

	if v, ok := in["preemptible"].(bool); ok {
		nodePoolObj.Preemptible = v
	}

	if v, ok := in["service_account"].(string); ok && len(v) > 0 {
		nodePoolObj.ServiceAccount = v
	}

	bs, err := json.Marshal(nodePoolObj)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func expandClusterGKENodePoolLegacyTaints(in map[string]interface{}) []K8sTaint {
	var taints []K8sTaint

	if v, ok := in["taints"].([]interface{}); ok && len(v) > 0 {
		legacyTaints := toArrayString(v)

		for _, part := range legacyTaints {
			taint := K8sTaint{}
			ekv := strings.Split(part, ":")
			if len(ekv) == 2 {
				taint.Effect = ekv[0]
				kv := strings.Split(ekv[1], "=")
				if len(kv) == 2 {
					taint.Key = kv[0]
					taint.Value = kv[1]
				}
			}

			taints = append(taints, taint)
		}
	}

	return taints
}
