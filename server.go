package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Server struct {
	Node Node
}

func (s *Server) Start(port string) {
	http.HandleFunc("/txions", s.txionHandler)
	http.HandleFunc("/mine", s.mineHandler)
	http.HandleFunc("/blockchain", s.blockchainHandler)
	http.ListenAndServe(":"+port, nil)
}

func (s *Server) txionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newTxion := Transaction{}
		defer r.Body.Close()

		rawJson, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("error reading request body: %+v\n", err.Error())
			http.Error(w, "500 problem reading json body", 500)
			return
		}

		err = json.Unmarshal(rawJson, &newTxion)
		if err != nil {
			fmt.Printf("error processing json: %+v\n", err.Error())
			http.Error(w, "400 bad request - invalid json body", http.StatusBadRequest)
			return
		}

		err = newTxion.Validate()
		if err != nil {
			fmt.Println("error validating json")
			http.Error(w, "400 bad request - invalid json body", http.StatusBadRequest)
			return
		}

		s.Node.Transactions = append(s.Node.Transactions, newTxion)
		fmt.Println("{{{--- New Transaction Added ---}}}")
		fmt.Print(newTxion.String() + "\n")
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}

func (s *Server) mineHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := s.Node.Mine(); err != nil {
			w.Write([]byte("mining failure"))
		}
		fmt.Println("{{{--- New Block Mined ---}}}")
		fmt.Print(s.Node.Blockchain[len(s.Node.Blockchain)-1])
		return
	default:
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
}

func (s *Server) blockchainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		jsonChain, err := json.Marshal(s.Node.Blockchain)
		if err != nil {
			http.Error(w, "502 problem rendering blockchain", http.StatusInternalServerError)
			return
		}
		w.Write(jsonChain)
		return
	default:
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
}
