package user

import handler "funding/src/app/handlers"

 
func init() {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
}
