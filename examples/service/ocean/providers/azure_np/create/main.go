package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as account and credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	// Create a new instance of the service's client with a Session.
	// Optional spotinst.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// service specific configuration.
	svc := ocean.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new cluster.
	out, err := svc.CloudProviderAzureNP().CreateCluster(ctx, &azure_np.CreateClusterInput{
		Cluster: &azure_np.Cluster{
			Name:                spotinst.String("foo"),
			ControllerClusterID: spotinst.String("foo"),
			AKS: &azure_np.AKS{
				ClusterName:                     spotinst.String("foo"),
				ResourceGroupName:               spotinst.String("foo"),
				Region:                          spotinst.String("eastus"),
				InfrastructureResourceGroupName: spotinst.String("foo"),
			},
			AutoScaler: &azure_np.AutoScaler{
				IsEnabled: spotinst.Bool(true),
				ResourceLimits: &azure_np.ResourceLimits{
					MaxVCPU:      spotinst.Int(20000),
					MaxMemoryGib: spotinst.Int(100000),
				},
				Down: &azure_np.Down{
					MaxScaleDownPercentage: spotinst.Int(10),
				},
				Headroom: &azure_np.Headroom{
					Automatic: &azure_np.Automatic{
						IsEnabled:  spotinst.Bool(true),
						Percentage: spotinst.Int(5),
					},
				},
			},
			Scheduling: &azure_np.Scheduling{
				ShutdownHours: &azure_np.ShutdownHours{
					TimeWindows: []string{
						"Sat:08:00-Sun:08:00",
						"Mon:08:00-Tue:08:00",
					},
					IsEnabled: spotinst.Bool(true),
				},
			},
			Health: &azure_np.Health{
				GracePeriod: spotinst.Int(600),
			},
			VirtualNodeGroupTemplate: &azure_np.VirtualNodeGroupTemplate{
				AvailabilityZones: []string{"1", "2"},
				NodePoolProperties: &azure_np.NodePoolProperties{
					MaxPodsPerNode:     spotinst.Int(110),
					EnableNodePublicIP: spotinst.Bool(false),
					OsDiskSizeGB:       spotinst.Int(128),
					OsDiskType:         spotinst.String("Managed"),
					OsType:             spotinst.String("Windows"),
					OsSKU:              spotinst.String("Windows2019"),
					KubernetesVersion:  spotinst.String("1.26"),
					PodSubnetIDs:       []string{"/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"},
					VnetSubnetIDs:      []string{"/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"},
				},
				NodeCountLimits: &azure_np.NodeCountLimits{
					MinCount: spotinst.Int(0),
					MaxCount: spotinst.Int(1000),
				},
				Strategy: &azure_np.Strategy{
					SpotPercentage: spotinst.Int(100),
					FallbackToOD:   spotinst.Bool(true),
				},
				AutoScale: &azure_np.AutoScale{
					Headrooms: []*azure_np.Headrooms{
						{
							CpuPerUnit:    spotinst.Int(10),
							MemoryPerUnit: spotinst.Int(30),
							GpuPerUnit:    spotinst.Int(5),
							NumberOfUnits: spotinst.Int(2),
						},
					},
				},
				Tags: &map[string]string{
					"key1":  "creator",
					"test":  "ocean",
					"test3": "ocean-aks",
				},

				Labels: &map[string]string{
					"creator": "automation",
					"test":    "aks",
					"test1":   "aks1",
				},

				Taints: []*azure_np.Taint{
					{
						Key:    spotinst.String("test"),
						Value:  spotinst.String("veryMuch"),
						Effect: spotinst.String("NoSchedule"),
					},
				},
				VmSizes: &azure_np.VmSizes{
					Filters: &azure_np.Filters{
						MinVcpu:      spotinst.Int(2),
						MaxVcpu:      spotinst.Int(16),
						MinMemoryGiB: spotinst.Float64(8),
						MaxMemoryGiB: spotinst.Float64(16),
						ExcludeSeries: []string{
							"Bs",
							"Da v4",
						},
						Architectures: []string{
							"X86_64",
						},
						Series: []string{
							"D v3",
							"F",
							"E v4",
						},
						AcceleratedNetworking: spotinst.String("Enabled"),
						DiskPerformance:       spotinst.String("Premium"),
						MinGpu:                spotinst.Float64(1),
						MaxGpu:                spotinst.Float64(2),
						MinNICs:               spotinst.Int(1),
						VmTypes: []string{
							"generalPurpose",
							"GPU",
						},
						MinDisk: spotinst.Int(2),
						GpuTypes: []string{
							"nvidia-tesla-t4",
							"nvidia-tesla-a100",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create cluster: %v", err)
	}

	// Output.
	if out.Cluster != nil {
		log.Printf("Cluster %q: %s",
			spotinst.StringValue(out.Cluster.ID),
			stringutil.Stringify(out.Cluster))
	}
}
