package server

import (
	"net/http"
)

// RoutAndHandler is a struct to combine rout and handler
type RoutAndHandler struct {
	Rout string
	Func func(http.ResponseWriter, *http.Request)
}

// StartServer is use apis to create a api server
func StartServer(port string, apis []RoutAndHandler) {
	for _, api := range apis {
		http.HandleFunc(api.Rout, api.Func)
	}
	http.ListenAndServe(":"+port, nil)
}
