package webservice

import (
	"net/http"
	"typosquatch/webservice/handlers"
)

func Run(port string) {
	handler := handlers.NewHandler()
	http.HandleFunc("/scan/", handler.Scan)
	http.HandleFunc("/result/", handler.Result)
	http.ListenAndServe(":"+port, nil)
}
