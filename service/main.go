// Copyright 2021 Authors of KubeArmor
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/kubearmor/KVMService/service/core"
	kg "github.com/kubearmor/KVMService/service/log"
)

// GitCommit string passed from govvv
var GitCommit string

// GitBranch string passed from govvv
var GitBranch string

// BuildDate string passed from govvv
var BuildDate string

// Version string passed from govvv
var Version string

func printBuildDetails() {
	kg.Printf("BUILD-INFO: commit:%s, branch: %s, date: %s, version: %s",
		GitCommit, GitBranch, BuildDate, Version)
}

func main() {
	printBuildDetails()
	if os.Geteuid() != 0 {
		kg.Printf("Need to have root privileges to run %s\n", os.Args[0])
		return
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		kg.Err(err.Error())
		return
	}

	if err := os.Chdir(dir); err != nil {
		kg.Err(err.Error())
		return
	}

	/*
		// options (string)
		clusterPtr := flag.String("cluster", "", "cluster name")
		gRPCPtr := flag.String("gRPC", "32767", "gRPC port number")
		logPathPtr := flag.String("logPath", "none", "log file path, {path|stdout|none}")
		logFilterPtr := flag.String("logFilter", "policy", "Filter for what kinds of alerts and logs to receive, {policy|system|all}")

		// options (boolean)
		enableEnforcerPerPodPtr := flag.Bool("enableEnforcerPerPod", false, "enabling the enforcer per pod")
		//enableExternalWorkloadPolicyPtr := true
	*/
	portPtr := flag.Int("port", 0, "Cluster Port")
	ipAddressPtr := flag.String("ipAddress", "", "Cluster Address")

	flag.Parse()

	// == //

	core.KVMSDaemon(*portPtr, *ipAddressPtr)

	// == //
}