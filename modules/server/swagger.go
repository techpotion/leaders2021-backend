package server

import "net/http"

// serving swagger.json file for opening it in swagger-ui
func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "gen/swagger/api.swagger.json")
}
