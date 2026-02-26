package repository

import (
	"context"

	"github.com/sidz111/job-service/model"
	"gorm.io/gorm"
)

type JobRepository interface {
	Create(ctx context.Context, job *model.Job) error
	GetById(ctx context.Context, id int) (*model.Job, error)
	GetAll(ctx context.Context) ([]model.Job, error)
	Update(ctx context.Context, job *model.Job) error
	Delete(ctx context.Context, id int) error
}

type jobRepo struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepo{db: db}
}

func (r *jobRepo) Create(ctx context.Context, job *model.Job) error {
	err := r.db.WithContext(ctx).Create(job).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *jobRepo) GetById(ctx context.Context, id int) (*model.Job, error) {
	var job *model.Job
	err := r.db.WithContext(ctx).First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return job, err
}
func (r *jobRepo) GetAll(ctx context.Context) ([]model.Job, error) {
	var jobs []model.Job
	err := r.db.WithContext(ctx).Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
func (r *jobRepo) Update(ctx context.Context, job *model.Job) error {
	err := r.db.WithContext(ctx).Save(job).Error
	if err != nil {
		return nil
	}
	return nil
}
func (r *jobRepo) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&model.Job{}, id).Error; err != nil {
		return err
	}
	return nil
}
