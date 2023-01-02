package configuration

import (
	"time"

	l "github.com/SbFinanceFbd/golib/logger"
	"github.com/SbFinanceFbd/golib/sql"
)

var Db sql.Database

func Init() (err error) {
	l.Debug.Printf("::Init:: Trying to connect DB\n")
	Db.Ip = AppConfig.MustGetString("psql.db.host")
	Db.Port = AppConfig.MustGetString("psql.db.port")
	Db.Type = AppConfig.MustGetString("psql.db.type")
	Db.Schema = AppConfig.MustGetString("psql.db.dbname")
	Db.Username = AppConfig.MustGetString("psql.db.user")
	Db.Password = AppConfig.MustGetString("psql.db.password")

	err = Db.Connect()
	Db.ConnPtr.SetConnMaxLifetime(time.Second * time.Duration(AppConfig.MustGetInt("psql.conn.max.life.time")))

	if err != nil {
		l.Debug.Printf("::Init:: db connect fail\n")
		return
	}
	l.Debug.Printf("::Init:: DB Connected successfully\n")
	return
}
