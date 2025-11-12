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
		ExecTask(id, status int, date string) error
		ChangeExecutor(id, executor int) error
		DeleteTask(id int) error
	}
	userTaskService struct {
		r  repository.UserTaskRepository
		l  *zap.Logger
		ts TaskService
	}
)

func NewUserTaskService(r repository.UserTaskRepository, l *zap.Logger, ts TaskService) UserTaskService {
	return &userTaskService{
		r:  r,
		l:  l,
		ts: ts,
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

	u.l.Debug("IN USER TASK SERVICE", zap.Int("branch_id", uc.BranchID))

	var isSupp bool

	taskType, err := u.ts.GetTaskByID(uc.Task)
	if err != nil {
		u.l.Error("IN USER TASK SERVICE", zap.Error(err))
		return err
	}

	isSupp = taskType.Type == "support"
	u.l.Debug("IN CREATE TASK SERVICE :: ", zap.Bool("Is Support", isSupp))

	if uc.Executor != 0 {
		u.l.Info("IN USER TASK SERVICE :: CREATE TASK WITH EXECUTOR", zap.Int("Executor", uc.Executor))
		err = u.r.CreateTask(uc)
		if err != nil {
			return err
		}
		return nil
	}

	u.l.Info("IN USER TASK SERVICE :: CREATE TASK WITHOUT EXECUTOR")

	role, _ := u.r.GetTaskRoles(uc.Task)

	ut, err := u.r.GetUserAndCountTasksByRoleID(role, uc.BranchID, uc.Initiator, isSupp)
	if err != nil {
		return err
	}

	if isSupp {
		uc.Executor = findMinCountTask(ut)
		err = u.r.CreateTask(uc)
		if err != nil {
			return err
		}
		return nil
	}

	executors := make(map[int]entity.UserCountTask)

	for _, user := range ut {

		ex, ok := executors[user.UserDep]
		if !ok {
			executors[user.UserDep] = user
			continue
		}

		if user.Count < ex.Count {
			executors[user.UserDep] = user
		}

		// u.l.Debug("IN SERVICE", zap.Int("Iteration", i+1), zap.Any("Executors:", executors))
	}

	u.l.Debug("IN USER TASK SERVICE", zap.Any("Executors:", executors))

	for _, user := range executors {
		u.l.Debug("IN SERVICE", zap.Any("Executor for Mass:", user))
		uc.Executor = user.UserID
		err = u.r.CreateTask(uc)
		if err != nil {
			return err
		}
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

func (u userTaskService) ExecTask(id, status int, date string) error {
	u.l.Debug("IN USER TASK SERVICE :: EXEC TASK")

	err := u.r.UpdateTask(id, status, date)
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
