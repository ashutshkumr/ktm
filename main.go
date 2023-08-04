package main

import (
	"flag"
	"log"

	"github.com/ashutshkumr/ktm/pkg/topology"
)

func main() {
	printUsage := true
	topoPtr := flag.String("topo", "", "Create new topology from YAML file")
	cleanupPtr := flag.Bool("cleanup", false, "Remove any active topology")
	flag.Parse()

	if *cleanupPtr {
		printUsage = false

		if err := topology.Cleanup(); err != nil {
			log.Fatalf("Failed cleaning up: %v\n", err)
		}
	}

	if *topoPtr != "" {
		printUsage = false

		if err := topology.NewFromFile(*topoPtr); err != nil {
			log.Fatalf("Failed creating new topology: %v\n", err)
		}
	}

	if printUsage {
		flag.Usage()
	}
}
