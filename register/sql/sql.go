package sql

import (
	"github.com/isollaa/dbterm/plugin/disk"
	"github.com/isollaa/dbterm/plugin/list"
	"github.com/isollaa/dbterm/plugin/ping"
	"github.com/isollaa/dbterm/registry"
)

func init() {
	registry.RegisterDriver(disk.SQL)
	registry.RegisterDriver(list.SQL)
	registry.RegisterDriver(ping.SQL)
}
