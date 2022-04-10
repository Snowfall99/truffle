package storage

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"truffle.io/model"
)

type Database struct {
	host     string
	port     int
	user     string
	password string
	name     string
	db       *gorm.DB
	log      *zap.Logger
}

type NewDatabaseOptions struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Log      *zap.Logger
}

func NewDatabase(opts *NewDatabaseOptions) *Database {
	if opts.Log == nil {
		opts.Log = zap.NewNop()
	}
	return &Database{
		host:     opts.Host,
		port:     opts.Port,
		user:     opts.User,
		password: opts.Password,
		name:     opts.Name,
		log:      opts.Log,
	}
}

func (d *Database) Connect() error {
	d.log.Info("Connecting to database", zap.String("url", d.createDataSourceName(false)))

	var err error
	d.db, err = gorm.Open(mysql.Open(d.createDataSourceName(true)), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db.AutoMigrate(&model.BFT{})

	return nil
}

func (d *Database) createDataSourceName(withPassword bool) string {
	password := d.password
	if !withPassword {
		password = "xxx"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.user, password, d.host, d.port, d.name)
}

func (d *Database) FindBFTs() []model.BFT {
	var bfts []model.BFT
	d.db.Find(&bfts)
	return bfts
}
