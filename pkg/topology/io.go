package topology

import (
	"fmt"
	"log"
	"os"

	"github.com/ashutshkumr/ktm/pkg/api"
	"gopkg.in/yaml.v3"
)

func NewFromFile(fileName string) error {
	log.Printf("Reading topology file %s ...\n", fileName)
	b, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("could not read topology file: %v", err)
	}

	t := api.Topology{}
	log.Printf("Parsing topology ...\n%s\n", b)
	if err := yaml.Unmarshal(b, &t); err != nil {
		return fmt.Errorf("could not unmarshal topology file: %v", err)
	}
	log.Println("Successfully parsed topology !")

	return New(&t)
}
