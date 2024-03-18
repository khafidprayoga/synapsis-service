package commonConn

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func PostgresConnect() (*gorm.DB, error) {
	var (
		host     = os.Getenv("PGHOST")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDATABASE")
	)

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s sslmode=require`, host, user, password, dbname)
	gConf := &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		TranslateError:                           true,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}

	pConf := postgres.New(postgres.Config{
		DSN: dsn,
	})

	db, errConnect := gorm.Open(pConf, gConf)
	if errConnect != nil {
		return nil, errConnect
	}

	return db, nil
}
