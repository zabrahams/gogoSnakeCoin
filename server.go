package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Server struct {
	txions []*Transaction
}

func (s *Server) Start(port string) {
	http.HandleFunc("/txions", s.createTxionHandler)
	http.ListenAndServe(":"+port, nil)
}

func (s *Server) createTxionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newTxion := Transaction{}
		defer r.Body.Close()

		rawJson, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("error reading request body: %+v\n", err.Error())
			http.Error(w, "500 problem reading json body", 500)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(rawJson, &newTxion)
		if err != nil {
			fmt.Printf("error processing json: %+v\n", err.Error())
			http.Error(w, "400 bad request - invalid json body", 400)
			return
		}

		err = newTxion.Validate()
		if err != nil {
			fmt.Println("error validating json")
			http.Error(w, "400 bad request - invalid json body", 400)
			return
		}

		s.txions = append(s.txions, &newTxion)
		fmt.Println("New Transaction Added:")
		fmt.Print(newTxion.String())
		w.WriteHeader(http.StatusCreated)
	default:
		w.Write([]byte("404 page not found"))
		w.WriteHeader(http.StatusNotFound)
	}
}
