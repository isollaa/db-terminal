package register

import (
	"github.com/isollaa/dbterm/plugin/disk"
	"github.com/isollaa/dbterm/plugin/list"
	"github.com/isollaa/dbterm/plugin/ping"
	"github.com/isollaa/dbterm/registry"
)

func init() {
	registry.RegisterDriver(disk.Sql)
	registry.RegisterDriver(list.Sql)
	registry.RegisterDriver(ping.Sql)
}
