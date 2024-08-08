package main

import (
	lib "github.com/redhat-best-practices-for-k8s/graphsolver-lib"
	l2lib "github.com/redhat-best-practices-for-k8s/l2discovery-exports"
)

type testGraph struct {
	// list of cluster interfaces indexed with a simple integer (X) for readability in the graph
	ifList []*l2lib.PtpIf
	// LANs identified in the graph
	lans *[][]int
	// List of port receiving PTP frames (assuming valid GM signal received)
	ptpInterfaces []*l2lib.PtpIf
}

// list of cluster interfaces indexed with a simple integer (X) for readability in the graph
func (config testGraph) GetPtpIfList() []*l2lib.PtpIf {
	return config.ifList
}

// LANs identified in the graph
func (config testGraph) GetLANs() *[][]int {
	return config.lans
}

// List of port receiving PTP frames (assuming valid GM signal received)
func (config testGraph) GetPortsGettingPTP() []*l2lib.PtpIf {
	return config.ptpInterfaces
}

// Runs Solver to find optimal configurations
func main() {
	const (
		// problem/scenario name
		findOCProblemName = "OC"

		// unique id for each tag, e.g. solution role
		tagSlave       = 0
		tagGrandmaster = 1
	)

	if1 := l2lib.PtpIf{IfClusterIndex: l2lib.IfClusterIndex{InterfaceName: "ens3f0", NodeName: "node1"}, Iface: l2lib.Iface{IfMac: l2lib.Mac{Data: "52:55:00:81:c2:62"}, IfPci: l2lib.PCIAddress{Device: "00:03", Function: "0"}}}
	if2 := l2lib.PtpIf{IfClusterIndex: l2lib.IfClusterIndex{InterfaceName: "ens3f0", NodeName: "node2"}, Iface: l2lib.Iface{IfMac: l2lib.Mac{Data: "52:55:00:81:c2:63"}, IfPci: l2lib.PCIAddress{Device: "00:03", Function: "0"}}}
	lans := [][]int{{0, 1}}
	aGraph := testGraph{ifList: []*l2lib.PtpIf{&if1, &if2}, lans: &lans, ptpInterfaces: nil}

	// initialize L2 config in solver
	lib.GlobalConfig.SetL2Config(&aGraph)

	// Initializing problems
	lib.GlobalConfig.InitProblem(
		findOCProblemName,
		[][][]int{
			{{int(lib.StepNil), 0, 0}},         // step1
			{{int(lib.StepSameLan2), 2, 0, 1}}, // step2
		},
		[]int{tagSlave: 0, tagGrandmaster: 1},
	)

	// Run solver for problem
	lib.GlobalConfig.Run(findOCProblemName)

	// print first solution
	lib.GlobalConfig.PrintAllSolutions()
}
