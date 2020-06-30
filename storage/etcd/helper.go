package etcd

import (
	// "sort"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	nodes     = "nodes"
	reserve   = "reserved"
	topology  = "topology"
	allocated = "allocated"
	mttl      = "mttl"
	query     = "labels/"
	labels    = "labels"

	B  = "b"
	KB = "kb"
	MB = "mb"
	GB = "gb"
	TB = "tb"
	BV = 1
)

var maper = map[string]int64{
	B:  BV,
	KB: BV << 10,
	MB: BV << 20,
	GB: BV << 30,
	TB: BV << 40,
}

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

func reservedTTLKey(key string) string {
	return strings.Join([]string{key, mttl}, "/")
}

func reserveKeyProblem(regionid, clusterid, nodeid string) string {
	return strings.Join([]string{regionid, clusterid, nodeid}, ".")
}

func allocatedKey(nodeid string) string {
	return strings.Join([]string{allocated, nodeid}, "/")
}

func isUsed(name string) bool {
	return strings.Contains(name, "/")
}

// regionid.clusterid.nodeid -> [regionid, clusterid, nodeid]
func split(id string) []string {
	return strings.Split(id, ".")
}

func isHostInfo(name string) bool {
	return strings.Contains(name, "host")
}

func labelsKey(key string) string {
	return strings.ReplaceAll(key, nodes, labels)
}

func extractNodeID(key string) string {
	return strings.ReplaceAll(key, "labels/", "")
}

func tobytes(n int64, unit string) (int64, error) {
	if val, ok := maper[unit]; ok {
		return n * val, nil
	}
	return 0, errors.New(fmt.Sprintf("%s not valid unit.valid units are b,kb,mb,gb,tb", unit))
}

func bytestostring(n int64, unit string) (string, error) {
	b, err := tobytes(n, unit)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(b, 10), nil

}

func tolabels(l string) map[string]string {
	rez := map[string]string{}
	for _, item := range strings.Split(l, ",") {
		parts := strings.Split(item, ":")
		rez[parts[0]] = parts[1]
	}
	return rez
}
