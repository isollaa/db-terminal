package mongo

import (
	"github.com/isollaa/dbterm/plugin/disk"
	"github.com/isollaa/dbterm/plugin/info"
	"github.com/isollaa/dbterm/plugin/list"
	"github.com/isollaa/dbterm/plugin/ping"
	"github.com/isollaa/dbterm/registry"
)

func init() {
	registry.RegisterDriver(disk.Mongo)
	registry.RegisterDriver(info.Mongo)
	registry.RegisterDriver(list.Mongo)
	registry.RegisterDriver(ping.Mongo)
}
