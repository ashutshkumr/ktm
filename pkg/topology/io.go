package topology

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func NewFromFile(fileName string) error {
	log.Printf("Reading topology file %s ...\n", fileName)
	b, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("could not read topology file: %v", err)
	}

	t := Topology{}
	log.Printf("Parsing topology ...\n%s\n", b)
	if err := yaml.Unmarshal(b, &t); err != nil {
		return fmt.Errorf("could not unmarshal topology file: %v", err)
	}

	return New(&t)
}
