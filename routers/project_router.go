package routers

import (
	"dpm/api"
)

var projectRouter = ControllerRouter{
	Route{
		Name:        "projects_index",
		Methods:     []string{"GET"},
		Pattern:     prefixion + "/projects",
		HandlerFunc: api.IndexProject,
	},
	Route{
		Name:        "projects_create",
		Methods:     []string{"POST"},
		Pattern:     prefixion + "/projects",
		HandlerFunc: api.CreateProject,
	},
	Route{
		Name:        "projects_delete",
		Methods:     []string{"DELETE"},
		Pattern:     prefixion + "/projects",
		HandlerFunc: api.DeleteProject,
	},
	Route{
		Name:        "projects_update",
		Methods:     []string{"PUT"},
		Pattern:     prefixion + "/projects/{id}",
		HandlerFunc: api.UpdateProject,
	},
}
