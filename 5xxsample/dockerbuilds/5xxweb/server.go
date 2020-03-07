package main

import (
        "net/http"
        "log"
       )

func fivexx(w http.ResponseWriter, r *http.Request) {

  log.Print("5xx endpoint hit")
  w.WriteHeader(http.StatusGatewayTimeout)
  w.Write([]byte("500x error!......."))

}

func healthz(w http.ResponseWriter, r *http.Request) {

  w.Write([]byte("Good"))

}

func main(){

  http.HandleFunc("/", fivexx)
  http.HandleFunc("/healthz", healthz)
  http.ListenAndServe(":80", nil)

}
