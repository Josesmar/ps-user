package routes

import (
	"net/http"
	"ps-user/controllers"
)

//Routers users
var routesUsers = []Route{

	{
		URI:                  "/user",
		Method:               http.MethodPost,
		Function:             controllers.CreateUser,
		RequestAutentication: false,
	},
	{
		URI:                  "/user/{userID}",
		Method:               http.MethodGet,
		Function:             controllers.GetUser,
		RequestAutentication: false,
	},
}
