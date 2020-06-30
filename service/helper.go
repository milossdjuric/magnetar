package service

import (
	// "sort"
	// "strconv"
	"strings"
)

const (
	nodes     = "nodes"
	reserve   = "reserved"
	allocated = "allocated"
	mttl      = "mttl"
)

//keyspace for nodes
// nodes/nodeid
// example
// nodes/1234acvf

func nodeKey(nodeid string) string {
	return strings.Join([]string{nodes, nodeid}, "/")
}

func reservedKey(nodeid string) string {
	return strings.Join([]string{reserve, nodeid}, "/")
}

func allocatedKey(nodeid string) string {
	return strings.Join([]string{allocated, nodeid}, "/")
}

func isUsed(name string) bool {
	return strings.Contains(name, "/")
}
