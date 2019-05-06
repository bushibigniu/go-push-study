package main 

import "net/http"

func callback(w http.ResponseWriter, r *http.Request) {
	//
	w.Write([]byte("welcome"))
}

func main() {
	
	//lu you
	http.HandleFunc("/ws", callback)

	//server
	http.ListenAndServe("0.0.0.0:9999", nil)
}
