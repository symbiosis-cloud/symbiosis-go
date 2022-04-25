package main

import (
	"fmt"
	"github.com/symbiosis-cloud/symbiosis-go"
	"os"
)

func main() {
	client, err := symbiosis.NewClient("https://api.symbiosis.host", os.Getenv("SYMBIOSIS_API_KEY"))

	if err != nil {
		fmt.Errorf("Error occurred: %s", err)
	}

	result, err := client.ListClusters(10)

	if err != nil {
		fmt.Errorf("Error occurred: %s", err)
	}

	for _, cluster := range result.Clusters {
		fmt.Println(cluster)
	}
}
