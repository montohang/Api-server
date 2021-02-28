package domain

import (
	"api_server/config"
	"api_server/domain/auth"
	"api_server/domain/user"
	"context"

	"github.com/gorilla/mux"
)


type AppRoute struct {
	Router *mux.Router
	Ctx context.Context
}

func Init() {
	route := AppRoute{}
	route.Router = config.MuxRouter()
	route.Ctx = context.Background()
	route.InitRoute()
	route.Serve()
}

func (route *AppRoute) Serve()  {
	config.ListenServe(route.Router)
}

func (route *AppRoute) InitRoute()  {
	env := config.NewEnv()
	db := config.InitDB(env)

	userRepo := user.NewUserRepoImple(db)
	userUsecase := user.NewUserUsecase(userRepo)
	userRoute := user.UserController{Usecase: userUsecase}
	userRoute.UserRoute(route.Router)

	authRepo := auth.NewAuthRepo(db)
	authUsecase := auth.NewAuthUsecase(authRepo)
	authRoute := auth.AuthController{Usecase: authUsecase}
	authRoute.Auth(route.Router)
}