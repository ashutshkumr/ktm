package k8s

import (
	"context"
	"fmt"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type client struct {
	config    *rest.Config
	clientset *kubernetes.Clientset
	ctx       context.Context
}

func NewClient() (*client, error) {
	log.Println("Creating new K8S client ...")

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("could not build kube config from flag: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("could not create new clientset: %v", err)
	}

	log.Println("Successfully created K8s client !")
	return &client{clientset: clientset, config: config, ctx: context.TODO()}, nil
}
