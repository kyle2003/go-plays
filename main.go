package main

import (
	"net/http"

	"github.com/go-plays/modules"
)

func main() {
	//
	var artzp modules.Theme
	artzp.Name = "artzp"

	modules.GetTopicData(artzp.Name)
	http.HandleFunc("artzp/", artzp.Handler)
	http.ListenAndServe("localhost:8866", nil)
}
