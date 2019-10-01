package main

import (
        "net/http"
       )

func fivexx(w http.ResponseWriter, r *http.Request) {

  w.WriteHeader(http.StatusServiceUnavailable)
  w.Write([]byte("Nope"))

}

func main(){

  http.HandleFunc("/", fivexx)
  http.ListenAndServe(":8080", nil)

}
