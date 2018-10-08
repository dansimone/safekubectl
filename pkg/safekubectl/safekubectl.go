package safekubectl

import (
	"io/ioutil"
	"fmt"
	"errors"
	"github.com/dansimone/safekubectl/pkg/ishell"
)

type Safekubectl struct {
	clusterRootDir string
}

func NewSafeKubectl(clusterRootDir string) *Safekubectl {
	safekubectl := &Safekubectl{
		clusterRootDir: clusterRootDir,
	}
	return safekubectl
}

// Returns the names of the available clusters found.
func (s *Safekubectl) ListClusters() ([]string) {
	clusterNames := []string{}
	files, err := ioutil.ReadDir(s.clusterRootDir)
	if err != nil {
		return clusterNames
	}
	for _, file := range files {
		k := fmt.Sprintf("%s/%s/kubeconfig", s.clusterRootDir, file.Name())
		if fileExists(k) {
			clusterNames = append(clusterNames, file.Name())
		}
	}
	return clusterNames
}

// Create a shell connection to a specific cluster.
func (s *Safekubectl) ConnectToCluster(clusterName string) (error) {
	// Ensure the cluster provided is in the list of available clusters
	clusters := s.ListClusters()
	found := false
	for _, cluster := range clusters {
		if cluster == clusterName {
			found = true
		}
	}
	if !found {
		return errors.New(fmt.Sprintf("Cluster %s not found\n", clusterName))
	}

	// Launch iShell
	shell := ishell.New()
	shell.KubeConfig(fmt.Sprintf("%s/%s/kubeconfig", s.clusterRootDir, clusterName))
	shell.SetPrompt(fmt.Sprintf("\033[1m%s> \033[m", clusterName))
	shell.SetHistoryPath(fmt.Sprintf("%s/%s/.safekubectl_history", s.clusterRootDir, clusterName))
	shell.Run()
	return nil
}
