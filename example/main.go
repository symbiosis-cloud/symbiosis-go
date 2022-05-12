package main

import (
	"github.com/symbiosis-cloud/symbiosis-go"
	"log"
	"os"
)

func main() {
	client, err := symbiosis.NewClientFromAPIKey(os.Getenv("SYMBIOSIS_API_KEY"))

	if err != nil {
		log.Fatalf("Error occurred: %s", err)
	}

	c, err := client.Cluster.Create(&symbiosis.NewCluster{
		Name: "hello-world-test-golang",
		Nodes: []symbiosis.ClusterNodeInput{
			symbiosis.ClusterNodeInput{
				Quantity: 1,
				NodeType: "general-int-1",
			},
		},
		KubeVersion: "1.23.5",
		Region:      "germany-1",
		Configuration: symbiosis.ClusterConfigurationInput{
			EnableCsiDriver:    true,
			EnableNginxIngress: false,
		},
	})

	if err != nil {
		log.Fatalf("Call failed: %s", err)
	}

	log.Printf("Cluster created: %s", c.Name)

	clusters, err := client.Cluster.List(10, 0)

	if err != nil {
		log.Fatalf("Call failed: %s", err)
	}

	for _, cluster := range clusters.Clusters {
		nodes, err := client.Cluster.ListNodes(cluster.Name)

		if err != nil {
			log.Fatalf("Call failed: %s", err)
		}

		for _, node := range nodes.Nodes {
			log.Printf("Node: %s (state: %s)", node.Name, node.State)
		}

	}
}
