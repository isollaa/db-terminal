// +build sql postgres

package cmd

import (
	_ "github.com/isollaa/dbterm/register/sql"
	_ "github.com/isollaa/dbterm/util/sql/postgres"
)
