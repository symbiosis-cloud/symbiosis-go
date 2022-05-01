package main

import (
	"github.com/symbiosis-cloud/symbiosis-go"
	"log"
	"os"
)

func main() {
	client, err := symbiosis.NewClient(os.Getenv("SYMBIOSIS_API_KEY"))

	if err != nil {
		log.Fatalf("Error occurred: %s", err)
	}

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
