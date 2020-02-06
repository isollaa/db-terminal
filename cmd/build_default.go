// +build !mongo,!sql,!mysql,!postgres

package main

import (
	_ "github.com/isollaa/dbterm/register/mongo"
	_ "github.com/isollaa/dbterm/register/sql"
	_ "github.com/isollaa/dbterm/util/sql/mysql"
	_ "github.com/isollaa/dbterm/util/sql/postgres"
)
