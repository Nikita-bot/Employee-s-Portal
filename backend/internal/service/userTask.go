package service

import (
	"portal/internal/entity"
	"portal/internal/repository"

	"go.uber.org/zap"
)

type (
	UserTaskService interface {
		TaskForUser(userId int) ([]entity.UserTask, error)
		TaskByUser(userId int) ([]entity.UserTask, error)
		CreateTask(uc entity.UserTaskCreate) error
		GetTaskByID(id int) (entity.UserTask, error)
		ExecTask(id int, date string) error
		ChangeExecutor(id, executor int) error
		DeleteTask(id int) error
	}
	userTaskService struct {
		r repository.UserTaskRepository
		l *zap.Logger
	}
)

func NewUserTaskService(r repository.UserTaskRepository, l *zap.Logger) UserTaskService {
	return &userTaskService{
		r: r,
		l: l,
	}
}

func (u userTaskService) TaskForUser(userId int) ([]entity.UserTask, error) {
	u.l.Debug("IN USER TASK SERVICE :: GET TASK CREATED FOR USER")

	ut, err := u.r.TaskForUser(userId)
	if err != nil {
		return nil, err
	}

	return ut, nil
}

func (u userTaskService) TaskByUser(userId int) ([]entity.UserTask, error) {
	u.l.Debug("IN USER TASK SERVICE :: GET TASK CREATED BY USER")

	ut, err := u.r.TaskByUser(userId)
	if err != nil {
		return nil, err
	}

	return ut, nil
}

func (u userTaskService) CreateTask(uc entity.UserTaskCreate) error {
	u.l.Debug("IN USER TASK SERVICE :: CREATE TASK")

	var ut []entity.UserCountTask
	ut, err := u.r.GetUserAndCountTasksByTaskID(uc.Task)
	if err != nil {
		return err
	}

	uc.Executor = findMinCountTask(ut)

	err = u.r.CreateTask(uc)
	if err != nil {
		return err
	}

	return nil
}

func findMinCountTask(ut []entity.UserCountTask) int {
	min := 9999999
	id := 0
	for _, u := range ut {
		if u.Count < min {
			min = u.Count
			id = u.UserID
		}
	}
	return id
}

func (u userTaskService) GetTaskByID(id int) (entity.UserTask, error) {
	u.l.Debug("IN USER TASK SERVICE :: GET TASK BY ID")

	ut, err := u.r.GetTaskByID(id)
	if err != nil {
		return entity.UserTask{}, err
	}

	return ut, nil

}

func (u userTaskService) ExecTask(id int, date string) error {
	u.l.Debug("IN USER TASK SERVICE :: EXEC TASK")

	err := u.r.UpdateTask(id, date)
	if err != nil {
		return err
	}

	return nil
}

func (u userTaskService) DeleteTask(id int) error {
	u.l.Debug("IN USER TASK SERVICE :: EXEC TASK")

	err := u.r.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
}

func (u userTaskService) ChangeExecutor(id, executor int) error {
	u.l.Debug("IN USER TASK SERVICE :: CHANGE EXECUTOR FOR TASK")

	err := u.r.UpdateExecutor(id, executor)
	if err != nil {
		return err
	}

	return nil
}
