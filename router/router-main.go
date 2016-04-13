package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func main()  {
  r := httprouter.New()
  r.GET("/", PageHandler)

  r.GET("/posts", PageHandler)
  r.POST("/posts", PageHandler)

  r.GET("/posts/:id", PageHandler)
  r.PUT("/posts/:id", PageHandler)
  r.GET("/posts/:id/edit", PageHandler)

  fmt.Println("starting server on 8080")
  http.ListenAndServe(":8080", r)
}

func PageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
  fmt.Fprintln(w, "home", p)
}
