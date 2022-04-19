package domain

import (
	"github.com/satori/go.uuid"
	"time"
)
type Base struct {
	ID 	uuid.UUID `gorm:"type:uuid;primaryKey;autoIncrement:false"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time `sql:"index"`
}

func (b *Base) BeforeCreate() (err error) {
	if b.ID.String() == "00000000-0000-0000-0000-000000000000" {
		b.ID = uuid.NewV4()
	}
	return
}
