package main

import (
	"log"
	"os"

	"github.com/symbiosis-cloud/symbiosis-go"
)

func main() {
	client, err := symbiosis.NewClientFromAPIKey(os.Getenv("SYMBIOSIS_API_KEY"))

	if err != nil {
		log.Fatalf("Error occurred: %s", err)
	}

	b, err := client.NodePool.Describe("2512cf91-73de-404b-b04e-4c49d57ec1f1")

	if err != nil {
		log.Fatalf("Error occurred: %s", err)
	}

	log.Println(b)

	// c, err := client.Cluster.Create(&symbiosis.ClusterInput{
	// 	Name:              "test-cluster-123",
	// 	Region:            "germany-1",
	// 	IsHighlyAvailable: true,
	// 	Nodes:             []symbiosis.ClusterNodeInput{},
	// 	KubeVersion:       "latest",
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// log.Printf("Cluster %s created", c.Name)
	// log.Printf("Highly available: %t", c.HighlyAvailable)
}
