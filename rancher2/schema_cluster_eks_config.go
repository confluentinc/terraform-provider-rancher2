package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	clusterEKSKind   = "eks"
	clusterDriverEKS = "amazonElasticContainerService"
)

//Types

type AmazonElasticContainerServiceConfig struct {
	AccessKey               string            `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`
	Annotations             map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	DisplayName             string            `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	DriverName              string            `json:"driverName" yaml:"driverName"`
	KeyPairName             string            `json:"keyPairName,omitempty" yaml:"keyPairName,omitempty"`
	KubernetesVersion       string            `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
	ManageOwnSecurityGroups *bool             `json:"manageOwnSecurityGroups,omitempty" yaml:"manageOwnSecurityGroups,omitempty"`
	NodeSecurityGroups      []string          `json:"nodeSecurityGroups,omitempty" yaml:"nodeSecurityGroups,omitempty"`
	Region                  string            `json:"region,omitempty" yaml:"region,omitempty"`
	SecretKey               string            `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`
	SecurityGroups          []string          `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`
	ServiceRole             string            `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`
	SessionToken            string            `json:"sessionToken,omitempty" yaml:"sessionToken,omitempty"`
	Subnets                 []string          `json:"subnets,omitempty" yaml:"subnets,omitempty"`
	VirtualNetwork          string            `json:"virtualNetwork,omitempty" yaml:"virtualNetwork,omitempty"`

	// list of json objects. Each object matches the schema defined by `AmazonElasticContainerWorkerPool`
	WorkerPools []string `json:"workerPools,omitempty" yaml:"workerPools,omitempty"`
}

type AmazonElasticContainerWorkerPool struct {
	BaseNodePool `json:",inline" yaml:",inline"`

	AMI                         string   `json:"ami,omitempty" yaml:"ami,omitempty"`
	AssociateWorkerNodePublicIP *bool    `json:"associateWorkerNodePublicIp,omitempty" yaml:"associateWorkerNodePublicIp,omitempty"`
	CreatePoolPerSubnet         bool     `json:"createPoolPerSubnet" yaml:"createPoolPerSubnet"`
	DesiredNodes                int64    `json:"desiredNodes,omitempty" yaml:"desiredNodes,omitempty"`
	EBSEncryption               bool     `json:"ebsEncryption,omitempty" yaml:"ebsEncryption,omitempty"`
	InstanceType                string   `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	MaximumNodes                int64    `json:"maximumNodes,omitempty" yaml:"maximumNodes,omitempty"`
	MinimumNodes                int64    `json:"minimumNodes" yaml:"minimumNodes"`
	NodeVolumeSize              int64    `json:"nodeVolumeSize,omitempty" yaml:"nodeVolumeSize,omitempty"`
	PlacementGroup              string   `json:"placementGroup,omitempty" yaml:"placementGroup,omitempty"`
	UserData                    string   `json:"userData,omitempty" yaml:"userData,omitempty"`
	Subnets                     []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}

//Schemas

func clusterEKSConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"access_key": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "The AWS Client ID to use",
		},
		"kubernetes_version": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The kubernetes master version",
		},
		"secret_key": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "The AWS Client Secret associated with the Client ID",
		},
		"key_pair_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Allow user to specify key name to use",
		},
		"manage_own_security_groups": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "When true, do not create or edit security groups, only assign them",
		},
		"node_security_groups": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of security groups to assign to the worker nodes",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"region": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "us-west-2",
			Description: "The AWS Region to create the EKS cluster in",
		},
		"security_groups": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of security groups to use for the cluster",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"service_role": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The service role to use to perform the cluster operations in AWS",
		},
		"session_token": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "A session token to use with the client key and secret if applicable",
		},
		"subnets": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of subnets in the virtual network to use for the master",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"virtual_network": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the virtual network to use",
		},
		"worker_pools": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "List of worker pools",
			MinItems:    1,
			Elem: &schema.Resource{
				Schema: newNodePoolSchema(map[string]*schema.Schema{
					"ami": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "A custom AMI ID to use for the worker nodes instead of the default",
					},
					"associate_worker_node_public_ip": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     true,
						Description: "Associate public ip EKS worker nodes",
					},
					"create_pool_per_subnet": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     true,
						Description: "Whether to create a pool per subnet or one pool for all subnets",
					},
					"desired_nodes": {
						Type:        schema.TypeInt,
						Optional:    true,
						Default:     3,
						Description: "The desired number of worker nodes",
					},
					"ebs_encryption": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Enable EBS encryption for EKS worker nodes",
					},
					"instance_type": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "t2.medium",
						Description: "The type of machine to use for worker nodes",
					},
					"placement_group": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "",
						Description: "The name of an existing cluster placement group into which you want to launch your instances",
					},
					"maximum_nodes": {
						Type:        schema.TypeInt,
						Optional:    true,
						Default:     3,
						Description: "The maximum number of worker nodes",
					},
					"minimum_nodes": {
						Type:        schema.TypeInt,
						Optional:    true,
						Default:     1,
						Description: "The minimum number of worker nodes",
					},
					"node_volume_size": {
						Type:        schema.TypeInt,
						Optional:    true,
						Default:     20,
						Description: "The volume size for each node",
					},
					"user_data": {
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
						Description: "Pass user-data to the nodes to perform automated configuration tasks",
					},
					"subnets": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "List of worker subnets in the virtual network to use",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				}),
			},
		},
	}

	return s
}
