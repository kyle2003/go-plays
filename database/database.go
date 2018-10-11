package database

import (
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type SqliteObj struct {
	DbName string
	db     *gorm.DB
}

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
}

func (s *SqliteObj) Get() *gorm.DB {
	if s.db == nil {
		s.Connect()
	}

	return s.db
}

func (s *SqliteObj) Connect() error {
	if s.DbName == "" {
		cfg, err := ini.Load("conf/setup.ini")
		if err != nil {
			logrus.Fatalln(err)
		}
		dbPath := "db/"
		if cfg.Section("db").Haskey("db_path") {
			dbPath = cfg.Section("db").Key("db_path").String()
		}

		if cfg.Section("db").Haskey("db_name") {
			s.DbName = dbPath + cfg.Section("db").Key("db_name").String()
		} else {
			s.DbName = dbPath + "test.db"
		}
	}

	db, err := gorm.Open("sqlite3", s.DbName)
	if err != nil {
		logrus.Fatalln(err)
		return nil
	} else {
		s.db = db
		return nil
	}
}
