package modules

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ThemeHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	theme := vars["theme"]

	data, err := ParseTheme(theme)

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

func TopicHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	topic := vars["topic"]

	data, err := ParseTopic(topic)

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
