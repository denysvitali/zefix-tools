package zefix

import (
	"errors"
	"gorm.io/gorm"
)

func (c *Client) FindCompany(name string) (*Company, error) {
	var company Company
	tx := c.db.First(&company, "legal_name = ? OR name = ?", name, name)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		c.logger.Errorf("unable to find company: %v", tx.Error)
		return nil, tx.Error
	}
	return &company, nil
}
