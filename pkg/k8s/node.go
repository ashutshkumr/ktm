package k8s

import (
	"fmt"
	"log"

	"github.com/ashutshkumr/ktm/pkg/api"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) CreateNode(ns string, n *api.Node) error {
	if n == nil {
		return fmt.Errorf("node not provided")
	}

	log.Printf("Creating node %s ...\n", n.Name)
	_, err := c.clientset.CoreV1().Pods(ns).Create(
		c.ctx,
		&apiv1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      n.Name,
				Namespace: ns,
				Labels:    map[string]string{"owner": "ktm"},
			},
			Spec: apiv1.PodSpec{
				Containers: []apiv1.Container{
					{
						Name:            n.Name + "-ctr",
						Image:           n.Image,
						Command:         n.Cmd,
						ImagePullPolicy: apiv1.PullIfNotPresent,
					},
				},
				RestartPolicy: apiv1.RestartPolicyAlways,
			},
		},
		metav1.CreateOptions{},
	)

	if err != nil {
		return fmt.Errorf("could not create node %s: %v", n.Name, err)
	}
	log.Println("Successfully created node !")

	return nil
}
