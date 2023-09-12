package campaign

import (
	handler "funding/src/app/handlers"
	"funding/src/app/user"
)

func init() {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
}
