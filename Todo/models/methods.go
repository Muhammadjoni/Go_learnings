package Models

import (
	"gin-todo-prc/Config"
)

func CreateATodo(todo *Todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}

	return nil
}
