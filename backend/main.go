package main

import (
	//	api_v1 "github.com/dqx0/GoHalves/go/api/v1"

	"github.com/dqx0/GoHalves/go/db"
	"github.com/dqx0/GoHalves/go/handler"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/dqx0/GoHalves/go/router"
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/dqx0/GoHalves/go/validator"
)

func main() {
	db := db.NewDB()
	br := repository.NewBaseRepository(db)
	bv := validator.NewBaseValidator()
	bu := usecase.NewBaseUsecase(br, bv)
	bh := handler.NewBaseHandler(bu)
	r := router.NewRouter(bh)
	r.Run()
}
