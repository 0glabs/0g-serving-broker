package db

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Task struct {
	ID                  *uuid.UUID            `gorm:"type:char(36);primaryKey" json:"id" readonly:"true"`
	CreatedAt           *time.Time            `json:"createdAt" readonly:"true" gen:"-"`
	UpdatedAt           *time.Time            `json:"updatedAt" readonly:"true" gen:"-"`
	CustomerAddress     string                `gorm:"type:text;not null"`
	PreTrainedModelHash string                `gorm:"type:text;not null"`
	DatasetHash         string                `gorm:"type:text;not null"`
	TrainingParams      string                `gorm:"type:text;not null"`
	OutputRootHash      string                `gorm:"type:text;"`
	IsTurbo             bool                  `gorm:"type:bool;not null;default:false"`
	Progress            string                `gorm:"type:varchar(255);not null;default:InProgress"`
	DeletedAt           soft_delete.DeletedAt `gorm:"softDelete:nano;not null;default:0;index:deleted_name"`
}

// BeforeCreate hook for generating a UUID
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == nil {
		id := uuid.New()
		t.ID = &id
	}
	return
}

func (d *DB) Migrate() error {
	d.db.Set("gorm:table_options", "ENGINE=InnoDB")

	m := gormigrate.New(d.db, &gormigrate.Options{UseTransaction: false}, []*gormigrate.Migration{
		{
			ID: "create-task",
			Migrate: func(tx *gorm.DB) error {

				return tx.AutoMigrate(&Task{})
			},
		},
	})

	return errors.Wrap(m.Migrate(), "migrate database")
}
