package webservice

import (
	"net/http"
	"typosquatch/webservice/handlers"
)

func Run(port string) {
	http.HandleFunc("/scan/", handlers.Scan)
	http.HandleFunc("/result/", handlers.Result)
	http.ListenAndServe(":"+port, nil)
}
