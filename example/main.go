package main

import (
	"github.com/symbiosis-cloud/symbiosis-go"
	"log"
	"os"
)

func main() {
	client, err := symbiosis.NewClient("https://api.symbiosis.host", os.Getenv("SYMBIOSIS_API_KEY"))

	if err != nil {
		log.Fatalf("Error occurred: %s", err)
	}

	result, err := client.ListClusters(10)

	if err != nil {
		log.Fatalf("Call failed: %s", err)
	}

	for _, cluster := range result.Clusters {
		describeCluster, err := client.DescribeCluster(cluster.Name)

		if err != nil {
			log.Fatalf("Describe failed: %s", err)
		}

		poolCount := 0
		for _, pool := range describeCluster.NodePools {
			poolCount += pool.DesiredQuantity
		}

		log.Printf("Cluster name: %s, Node pools: %d", describeCluster.Name, len(describeCluster.NodePools))
		log.Printf("Total nodes in cluster: %d", poolCount)

		node, err := client.DescribeNode(describeCluster.Nodes[0].Name)

		if err != nil {
			log.Fatalf("Describe failed: %s", err)
		}

		log.Printf("Node name: %s", node.Name)

		err = node.Recycle()

		if err != nil {
			log.Fatalf("Recycle failed (%s): %s", node.Name, err)
		}

	}
}
