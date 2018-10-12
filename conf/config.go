package conf

import (
	"pandora/database"

	"github.com/go-ini/ini"
)

var GlobalDb = database.SqliteObj{}
var Setup, _ = ini.Load("setup.ini")
