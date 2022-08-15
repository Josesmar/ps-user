package routes

import (
	"net/http"
	controllers "ps-user/src/adapter/api"
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
	{
		URI:                  "/users/credentials/",
		Method:               http.MethodGet,
		Function:             controllers.ValidCredentials,
		RequestAutentication: false,
	},
	{
		URI:                  "/users/{userID}",
		Method:               http.MethodDelete,
		Function:             controllers.DeleteUser,
		RequestAutentication: false,
	},
	{
		URI:                  "/users/{userID}",
		Method:               http.MethodPatch,
		Function:             controllers.UpdateUser,
		RequestAutentication: false,
	},
}
