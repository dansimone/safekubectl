# Safekubectl

Safekubectl is for anyone who frequently changes the Kubernetes cluster against which they run kubectl commands.
As many have experienced using kubectl directly, it's all too easy to accidentally point to the wrong cluster
and get yourself into trouble.  Safekubectl is a wrapper on top of kubectl (using a slightly modified version of
[ishell](https://github.com/abiosoft/ishell)), that helps you run commands against your many Kubernetes clusters,
with added protections that make it harder to make a mistake.

# Project Status

## Features

* Provides a prompt that lets you know without a doubt, at all times, which cluster you are pointing you.
* Maintains a command history per cluster.

## Open Issues

* After executing each command, the first character entered into stdin is lost and needs to be re-entered by the user.

# Installing

```
go get github.com/dansimone/safekubectl/...
```

# Usage

## Provide a Safekubectl Root directory

Before using Safekubectl, set up a root directory (the default value is $HOME/.safekubectl) containing one or more
directories, each containing a `kubeconfig` file for a different Kubernetes cluster.  For example:

```
cd $HOME/.safekubectl
find .

./cluster1/kubeconfig
./cluster2/kubeconfig
./cluster3/kubeconfig
./clusterA/kubeconfig
./myCluster/kubeconfig
```

## List Available Clusters

List the clusters available to Safekubectl with:

```
safekubectl -c list

Available Clusters:
cluster1
cluster2
cluster3
clusterA
myCluster
```

## Connect to a Cluster

Connect to a cluster (for example, "cluster1") with:

```
safekubectl -c connect -k cluster1
```

This creates a shell indicating the cluster name, within which you can interact with cluser using regular kubectl
commands:

```
cluster1> get nodes
NAME              STATUS    ROLES     AGE       VERSION
worker1           Ready     node      6d        v1.10.3
```

Log tailing (`logs -f <pod_name>`) and pod login (`exec -it <pod_name>`) are fully supported.

## Disconnect from the Cluster

Use Control-C to disconnect from the Safekubectl shell.

# Running Locally

## Clone

```
mkdir -p $GOPATH/src/github.com/dansimone
cd $GOPATH/src/github.com/dansimone
git clone https://github.com/dansimone/safekubectl
```

## Install

```
make go-install
```

## Run Unit Tests

```
make unit-test
```

# Contributions

Pull requests welcome!