package k8s

import (
	"fmt"
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) CreateNamespace(name string) error {
	if name == "" {
		return fmt.Errorf("namespace not provided")
	}

	log.Printf("Creating namespace %s ...\n", name)
	_, err := c.clientset.CoreV1().Namespaces().Create(
		c.ctx, &apiv1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:   name,
				Labels: map[string]string{"owner": "ktm"},
			},
		},
		metav1.CreateOptions{},
	)

	if err != nil {
		return fmt.Errorf("could not create namespace %s: %v", name, err)
	}
	log.Println("Successfully created namespace !")

	return nil
}

func (c *client) DeleteNamespace(name string) error {
	if name == "" {
		return fmt.Errorf("namespace not provided")
	}

	log.Printf("Deleting namespace %s ...\n", name)
	err := c.clientset.CoreV1().Namespaces().Delete(
		c.ctx, name,
		metav1.DeleteOptions{},
	)

	if err != nil {
		return fmt.Errorf("could not delete namespace %s: %v", name, err)
	}
	log.Println("Successfully deleted namespace !")

	return nil
}
