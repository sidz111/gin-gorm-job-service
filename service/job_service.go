package service

import (
	"context"
	"fmt"

	"github.com/sidz111/job-service/model"
	"github.com/sidz111/job-service/repository"
)

type JobService interface {
	Create(ctx context.Context, job *model.Job) error
	GetById(ctx context.Context, id int) (*model.Job, error)
	GetAll(ctx context.Context) ([]model.Job, error)
	Update(ctx context.Context, job *model.Job) error
	Delete(ctx context.Context, id int) error
}

type jobService struct {
	repo repository.JobRepository
}

func NewJobService(repo repository.JobRepository) JobService {
	return &jobService{repo: repo}
}

func (s *jobService) Create(ctx context.Context, job *model.Job) error {
	if err := JobValidation(job); err != nil {
		return err
	}
	return s.repo.Create(ctx, job)
}
func (s *jobService) GetById(ctx context.Context, id int) (*model.Job, error) {
	if err := id <= 0; err != false {
		return nil, fmt.Errorf("Id must be positive")
	}
	return s.repo.GetById(ctx, id)
}
func (s *jobService) GetAll(ctx context.Context) ([]model.Job, error) {
	return s.repo.GetAll(ctx)
}
func (s *jobService) Update(ctx context.Context, job *model.Job) error {
	if err := JobValidation(job); err != nil {
		return err
	}
	return s.repo.Update(ctx, job)
}
func (s *jobService) Delete(ctx context.Context, id int) error {
	if err := id <= 0; err != false {
		return fmt.Errorf("Id must be positive number")
	}
	return s.repo.Delete(ctx, id)
}
func JobValidation(job *model.Job) error {
	if job.Title == "" {
		return fmt.Errorf("Title Required")
	}
	if job.Salary <= 0 {
		return fmt.Errorf("Salary Should be positive number")
	}
	return nil
}
