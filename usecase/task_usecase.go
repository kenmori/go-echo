package usecase

import (
	"go-echo/model"
	"go-echo/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskByID(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdateAt:  v.UpdateAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskByID(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskByID(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdateAt:  task.UpdateAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	// 成功した場合はtaskが指し示す値が書き換わっているのでここで参照できる
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdateAt:  task.UpdateAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdateAt:  task.UpdateAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
