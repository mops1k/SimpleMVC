package service

import (
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mssql"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
    dialect string
    url     string
    orm     *gorm.DB
    logger  *log.Logger
}

func (d *Database) SetLogger(logger *log.Logger) {
    d.logger = logger
}

func (d *Database) SetDialect(dialect string) {
    d.dialect = dialect
}

func (d *Database) SetUrl(url string) {
    d.url = url
}

func (d *Database) Connect() {
    var err error
    d.orm, err = gorm.Open(d.dialect, d.url)
    if err != nil {
        Container.GetLogger().Database.Panic(err)
    }

    if d.logger != nil {
        d.orm.SetLogger(d.logger)
    }
    d.orm.Debug().Set("gorm:auto_preload", true)
}

func (d *Database) Close() {
    err := d.orm.Close()
    if err != nil {
        Container.GetLogger().Database.Panic(err)
    }
}

func (d *Database) GetORM() *gorm.DB {
    return d.orm
}
