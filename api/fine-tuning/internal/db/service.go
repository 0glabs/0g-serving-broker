package db

import (
	"github.com/0glabs/0g-serving-broker/fine-tuning/schema"
	"github.com/google/uuid"
)

func (d *DB) AddTask(task *schema.Task) error {
	ret := d.db.Create(&task)
	return ret.Error
}

func (d *DB) GetTask(id *uuid.UUID) (schema.Task, error) {
	svc := schema.Task{}
	ret := d.db.Where(&schema.Task{ID: id}).First(&svc)
	return svc, ret.Error
}

func (d *DB) ListTask(userAddress string, latest bool) ([]schema.Task, error) {
	var tasks []schema.Task
	query := d.db.Where(&schema.Task{UserAddress: userAddress})
	if latest {
		query = query.Order("created_at DESC").Limit(1)
	}
	ret := query.Find(&tasks)
	return tasks, ret.Error
}

func (d *DB) UpdateTask(id *uuid.UUID, new schema.Task) error {
	ret := d.db.Where(&schema.Task{ID: id}).Updates(new)
	return ret.Error
}

func (d *DB) InProgressTaskCount() (int64, error) {
	var count int64
	d.db.Model(&schema.Task{}).Where("Progress <> ?", schema.ProgressStateFinished.String()).Count(&count)
	return count, nil
}

func (d *DB) GetDeliveredTasks() ([]schema.Task, error) {
	var filteredTasks []schema.Task
	d.db.Where(&schema.Task{Progress: schema.ProgressStateDelivered.String()}).Find(&filteredTasks)
	return filteredTasks, nil
}
