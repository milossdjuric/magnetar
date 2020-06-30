package influx

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	TIMES     = "times"
	TAGS      = "tags"
	POINTS    = "points"
	DB_EXISTS = "SHOW DATABASES"
)

func createDB(id string) string {
	return fmt.Sprintf("CREATE DATABASE %s", id)
}

func createRetentionPolicy(id, policy string) string {
	return fmt.Sprintf("CREATE RETENTION POLICY \"%s\" ON \"%s\" DURATION %s REPLICATION 1 DEFAULT", toPolicy(id), id, policy)
}

func queryDB(query string) string {
	return fmt.Sprintf("SELECT * FROM %s", query)
}

func toTime(timestamp string) time.Time {
	i, _ := strconv.ParseInt(timestamp, 10, 64)
	return time.Unix(i, 0)
}

func toID(name string) string {
	if strings.Contains(name, "/") {
		p := strings.Split(name, "/")
		return p[len(p)-1]
	}
	return name
}

func toPolicy(id string) string {
	return strings.Join([]string{id, "policy"}, "_")
}

func toAddress(address string) string {
	if strings.Contains(address, "http://") || strings.Contains(address, "https://") {
		return address
	}
	return strings.Join([]string{"http://", address}, "")
}
