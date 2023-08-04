package topology

import (
	"fmt"
	"log"

	"github.com/ashutshkumr/ktm/pkg/api"
	"github.com/ashutshkumr/ktm/pkg/k8s"
)

func New(t *api.Topology) error {
	if t == nil {
		return fmt.Errorf("no topology provided")
	}

	log.Printf("Creating topology %s ...\n", t.Name)
	k8sClient, err := k8s.NewClient()
	if err != nil {
		return fmt.Errorf("could not create k8s client: %v", err)
	}

	if err := k8sClient.CreateNamespace(t.Name); err != nil {
		return err
	}

	for _, n := range t.Nodes {
		if err := k8sClient.CreateNode(t.Name, &n); err != nil {
			return err
		}
	}

	log.Println("Successfully created topology !")
	return nil
}

func Cleanup() error {
	log.Println("Cleaning up any active topologies ...")
	// k8sClient, err := k8s.NewClient()
	// if err != nil {
	// 	return fmt.Errorf("could not create k8s client: %v", err)
	// }

	// if err := k8sClient.CreateNamespace(t.Name); err != nil {
	// 	return err
	// }

	// for _, n := range t.Nodes {
	// 	if err := k8sClient.CreateNode(t.Name, &n); err != nil {
	// 		return err
	// 	}
	// }

	log.Println("Successfully created topology !")
	return nil
}
