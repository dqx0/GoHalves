package di

/*
package di

import (
	"github.com/yourapp/db"
	"github.com/yourapp/repository"
	"github.com/yourapp/usecase"
	"github.com/yourapp/validator"
	"github.com/yourapp/controller"
)

type Container struct {
	TaskController *controller.TaskController
	// 他のコントローラーもここに追加
}

func NewContainer() *Container {
	dbInstance := db.NewDB()
	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(dbInstance)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	return &Container{
		TaskController: taskController,
		// 他のコントローラーの初期化もここで行う
	}
}
*/
