package main

import (
	"os/user"
	"flag"
	"fmt"
	"os"
	"log"
	"github.com/fatih/color"
	"github.com/dansimone/safekubectl/pkg/safekubectl"
)

var bold = color.New(color.Bold)

const ListCommand = "list"
const ConnectCommand = "connect"
const HelpCommand = "help"

func main() {
	var command, clusterName, rootDir string

	flag.Usage = func() {
		c := color.New(color.Bold).Add(color.Underline)
		c.Printf("Safekubectl:")
		c = color.New(color.Bold)
		c.Printf(" The kubectl wrapper for safe execution of commands against your many Kubernetes clusters.\n\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Determine the user's home dir
	usr, err := user.Current()
	if err != nil {
		flag.Usage()
		log.Fatal(err)
	}
	flag.StringVar(&command, "c", "", fmt.Sprintf("Command to run, either \"%s\" or \"%s\"", ListCommand, ConnectCommand))
	flag.StringVar(&clusterName, "k", "", "Name of the cluster to connect to")
	flag.StringVar(&rootDir, "rootDir", fmt.Sprintf("%s/.safekubectl", usr.HomeDir), "Safekubectl home directory")
	flag.Parse()

	if command == "" {
		flag.Usage()
		os.Exit(1)
	}

	safekubectl := safekubectl.NewSafeKubectl(rootDir)
	if command == ListCommand {
		clusterNames := safekubectl.ListClusters()
		if len(clusterNames) == 0 {
			c := color.New(color.Bold)
			c.Printf("No clusters found.  Please create at least one directory under %s containing a kubeconfig file.\n", rootDir)
		} else {
			c := color.New(color.Bold).Add(color.Underline)
			c.Println("Available Clusters:")
			for _, clusterName := range clusterNames {
				fmt.Println(clusterName)
			}
		}
	} else if command == ConnectCommand {
		if clusterName == "" {
			flag.Usage()
			os.Exit(1)
		}
		err = safekubectl.ConnectToCluster(clusterName)
		if err != nil {
			c := color.New(color.Bold)
			c.Println(err)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
