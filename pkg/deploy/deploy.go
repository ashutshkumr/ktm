package deploy

import (
	"fmt"

	"github.com/ashutshkumr/ktm/pkg/topology"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Deploy(t *topology.Topology) error {
	config, err := clientcmd.BuildConfigFromFlags("", "~/.kube/config")
	if err != nil {
		return fmt.Errorf("could not build kube config from flag: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("could not get new clientset: %v", err)
	}

	fmt.Println(clientset)
	// clientset.CoreV1().Pods()
	return nil
}
