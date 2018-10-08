package main

import (
	"os/user"
	"flag"
	"fmt"
	"os"
	"log"
	"github.com/dansimone/safekubectl/pkg/safekubectl"
	"github.com/dansimone/safekubectl/pkg/ishell"
)

func main() {
	var clusterName, rootDir string

	flag.Usage = func() {
		c := safekubectl.GetHighlightColor()
		c.Printf("Safekubectl:")
		fmt.Printf(" The kubectl wrapper for safe execution of commands against your many Kubernetes clusters.\n\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Determine the user's home dir
	usr, err := user.Current()
	if err != nil {
		flag.Usage()
		log.Fatal(err)
	}
	flag.StringVar(&clusterName, "cluster", "", "Name of the cluster to connect to")
	flag.StringVar(&rootDir, "rootDir", fmt.Sprintf("%s/.safekubectl", usr.HomeDir), "Safekubectl home directory")
	flag.Parse()

	safekubectl := safekubectl.NewSafeKubectl(rootDir)
	clusterNames := safekubectl.ListClusters()
	if len(clusterNames) == 0 {
		fmt.Printf("No clusters found.  Please create at least one directory under %s containing a kubeconfig file.\n", rootDir)
		os.Exit(1)
	}

	// Prompt for cluster selection if cluster not specified
	if clusterName == "" {
		shell := ishell.New()
		for {
			choice := shell.MultiChoice(clusterNames, "Select the cluster to connect to:")
			if choice < 0 || choice >= len(clusterNames) {
				os.Exit(0)
			}
			clusterName = clusterNames[choice]

			err = safekubectl.ConnectToCluster(clusterName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

	} else {
		err = safekubectl.ConnectToCluster(clusterName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
