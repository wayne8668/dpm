package routers

var (
	optionsRouter = ControllerRouter{
		Route{
			Name:        "options_action", 
			Methods:     []string{"OPTIONS"},
			Pattern:     "/*action",
			HandlerFunc: CorsMW(),
		},
	}
)