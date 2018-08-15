package db

import (
	"errors"
	"strings"

	"github.com/ljesparis/micro-todo/services/tasks/models"
)

var Tasks []models.Task

func init() {
	Tasks = []models.Task{
		{ID: 1, Done: false, Name: "database configuration", Description: "this configuration will allow the app to open a database connection"},
		{ID: 2, Done: false, Name: "create login controller", Description: "this controller will allow users to login into the system"},
		{ID: 3, Done: false, Name: "create signup controller", Description: "this controller will allow users to create an account"},
		{ID: 4, Done: false, Name: "create reset password controller", Description: "this controller will allow users to reset their password"},
		{ID: 5, Done: false, Name: "create get user controller", Description: "this controller will allow the user to get its profile"},
		{ID: 6, Done: false, Name: "create follow controller", Description: "this controller will allow users to follow another user"},
		{ID: 7, Done: false, Name: "create pwinty client", Description: "this client will allow patopatilla system to create orders"},
	}
}

func GetTaskById(i int) (task *models.Task, err error) {
	for _, tmp := range Tasks {
		if tmp.ID == i {
			task = &tmp
			break
		}
	}

	if task == nil {
		err = errors.New("task does not exists")
	}

	return
}

func CreateTask(task *models.Task) {
	Tasks = append(Tasks, *task)
}

func UpdateTask(task *models.Task) error {
	j := 0
	for i, _ := range Tasks {
		if Tasks[i].ID == task.ID {
			if Tasks[i].Done != task.Done {
				Tasks[i].Done = task.Done
			}

			if strings.Compare(task.Description, Tasks[i].Description) != 0 {
				Tasks[i].Description = task.Description
			}

			if strings.Compare(task.Name, Tasks[i].Name) != 0 {
				Tasks[i].Name = task.Name
			}
		} else {
			j++
		}
	}

	if len(Tasks) == j {
		return errors.New("task does not exists")
	}

	return nil
}

func DeleteTask(id int) error {
	var tmpTasks []models.Task
	for _, tmp := range Tasks {
		if tmp.ID != id {
			tmpTasks = append(tmpTasks, tmp)
		}
	}

	if len(Tasks) == len(tmpTasks) {
		return errors.New("task does not exists")
	}

	Tasks = tmpTasks
	return nil
}
