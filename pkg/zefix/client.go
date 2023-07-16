package zefix

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func New(dsn string) (*Client, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	c := Client{
		db: db,
	}
	err = c.initModels()
	return &c, err
}

func (c *Client) SetLogger(logger *logrus.Logger) {
	if logger != nil {
		c.logger = logger
	}
}

func (c *Client) initModels() error {
	err := c.db.AutoMigrate(&Company{})
	if err != nil {
		return err
	}
	return nil
}
