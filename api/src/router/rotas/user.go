package rotas

import "net/http"

//Routers users
var routesUsers = []Route{
	{
		URI:                  "/users",
		Method:               http.MethodGet,
		Function:             controllers.GetAll,
		RequestAutentication: false,
	},
}
