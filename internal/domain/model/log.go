package model

import (
	"time"

	"gorm.io/datatypes"
)

//MetaField metadata field model
type Log struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Url       string         `gorm:"column:url;" json:"url,omitempty"`
	Response  datatypes.JSON `gorm:"column:response;" json:"response,omitempty"`
	CreatedAt time.Time      `gorm:"column:createdAt" json:"created_at,omitempty"`
}

func (Log) TableName() string {
	return "log"
}
