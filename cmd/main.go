package main

import (
	"SeKai/internal/config"
	"SeKai/internal/server/http"
)

func main() {
	config.InitConfig()
	http.StartHTTP()
}
