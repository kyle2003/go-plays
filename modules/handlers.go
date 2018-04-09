package modules

import (
	"log"
	"net/http"
)

func (th *Theme) Handler(res http.ResponseWriter, req *http.Request) {
	data, err := ParseTheme(th.Name)

	if err != nil {
		log.Fatal("xx")
		return
	}

	log.Println(data)

	//	for _, t := range data {
	//
	//	}
	return

}
