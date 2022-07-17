package routes

import (
	"net/http"
	"ps-user/controllers"
)

//Routers users
var routesUsers = []Route{

	{
		URI:                  "/users",
		Method:               http.MethodPost,
		Function:             controllers.CreateUser,
		RequestAutentication: false,
	},
	{
		URI:                  "/users/{userID}",
		Method:               http.MethodGet,
		Function:             controllers.GetUser,
		RequestAutentication: false,
	},
	{
		URI:                  "/users",
		Method:               http.MethodGet,
		Function:             controllers.GetAllUser,
		RequestAutentication: false,
	},
}
